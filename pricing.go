package pitOrgan

import (
	"bufio"
	"encoding/json"
	"golang.org/x/xerrors"
	"strings"
	"sync"
	"time"
)

// Receivers

type ReceiverPricing struct {
	AccountID  string
	Connection *Connection
}

type ReceiverPricingStream struct {
	AccountID  string
	Connection *Connection
}

// Params

type GetPricingParams struct {
	Instruments            []string
	Since                  time.Time
	IncludeUnitsAvailable  JsonBool
	IncludeHomeConversions JsonBool
}

type GetPricingStreamParams struct {
	BufferSize  int
	Instruments []string
}

// Schemas

type GetPricingSchema struct {
	Prices          []PriceDefinition           `json:"prices"`
	HomeConversions []HomeConversionsDefinition `json:"homeConversions"`
	Time            DateTimeDefinition          `json:"time"`
}

// Types

type PriceChannels struct {
	Price     <-chan *PriceDefinition
	Error     <-chan error
	close     chan<- struct{}
	closeWait *sync.WaitGroup
}

func (r *ReceiverAccountID) Pricing() *ReceiverPricing {
	return &ReceiverPricing{
		AccountID:  r.AccountID,
		Connection: r.Connection,
	}
}

// GET /v3/accounts/{accountID}/pricing
func (r *ReceiverPricing) Get(params *GetPricingParams) (*GetPricingSchema, error) {
	resp, err := r.Connection.request(
		&requestParams{
			method:   "GET",
			endPoint: "/v3/accounts/" + r.AccountID + "/pricing",
			headers: []header{
				{key: "Accept-Datetime-Format", value: "RFC3339"},
			},
			queries: func() []query {
				q := make([]query, 0, 4)
				q = append(q, query{key: "instruments", value: strings.Join(params.Instruments, ",")})
				if b, err := params.IncludeUnitsAvailable.string(); err == nil {
					q = append(q, query{key: "includeUnitsAvailable", value: b})
				}
				if b, err := params.IncludeHomeConversions.string(); err == nil {
					q = append(q, query{key: "includeHomeConversions", value: b})
				}
				if !params.Since.IsZero() {
					q = append(q, query{key: "since", value: params.Since.Format(time.RFC3339Nano)})
				}
				return q
			}(),
		},
	)
	if err != nil {
		return nil, xerrors.Errorf("Get pricing canceled: %w", err)
	}
	defer resp.Body.Close()

	var data interface{}
	switch resp.StatusCode {
	case 200:
		data = new(GetPricingSchema)
	}

	data, err = parseResponse(resp, data)
	if err != nil {
		return nil, xerrors.Errorf("Get pricing failed: %w", err)
	}
	return data.(*GetPricingSchema), nil
}

func (r *ReceiverPricing) Stream() *ReceiverPricingStream {
	return &ReceiverPricingStream{
		AccountID:  r.AccountID,
		Connection: r.Connection,
	}
}

// GET /v3/accounts/{accountID}/pricing/stream
func (r *ReceiverPricingStream) Get(params *GetPricingStreamParams) (*PriceChannels, error) {
	resp, err := r.Connection.stream(
		&requestParams{
			method:   "GET",
			endPoint: "/v3/accounts/" + r.AccountID + "/pricing/stream",
			headers: []header{
				{key: "Accept-Datetime-Format", value: "RFC3339"},
			},
			queries: []query{
				{key: "instruments", value: strings.Join(params.Instruments, ",")},
			},
		},
	)
	if err != nil {
		return nil, xerrors.Errorf("Get pricing stream canceled: %w", err)
	}

	if resp.StatusCode != 200 {
		defer resp.Body.Close()
		var err interface{}
		_, err = parseResponse(resp, err)
		return nil, xerrors.Errorf("Get pricing stream failed: %w", err)
	}

	closeWait := new(sync.WaitGroup)

	priceCh := make(chan *PriceDefinition, params.BufferSize)
	errorCh := make(chan error, 2)
	closeCh := make(chan struct{})

	// closeChがcloseされたらstreamを終了するgoroutine
	// 下記の readre.ReadBytes はデータを受信しないと処理が進まないため
	// streamの途中で途切れると永遠に待ち続けてしまうので終了用goroutineを用意。
	go func() {
		defer func() {
			resp.Body.Close()
			closeWait.Done()
		}()
		closeWait.Add(1)
		_, _ = <-closeCh
	}()

	// 受信したデータ(JSON)を構造体にしてreaderCh channelに送信するgoroutine
	readerCh := make(chan *PriceDefinition, params.BufferSize)
	go func() {
		defer func() {
			close(readerCh)
			closeWait.Done()
		}()
		closeWait.Add(1)

		reader := bufio.NewReader(resp.Body)
		for {
			line, err := reader.ReadBytes('\n')
			if err != nil {
				select {
				case _, _ = <-closeCh:
				default:
					errorCh <- xerrors.Errorf("Read response stream failed: %w", err)
				}
				return
			}

			data := new(PriceDefinition)
			err = json.Unmarshal(line, data)
			if err != nil {
				errorCh <- xerrors.Errorf("Unmarshal response stream failed: %w", err)
				return
			}

			readerCh <- data
		}
	}()

	// readeerCh channleから受信した構造体をユーザーに渡すgoroutine
	// heartbeat(5秒間隔)が途切れた場合を検知するためにタイムアウトも管理する。
	go func() {
		defer func() {
			close(priceCh)
			closeWait.Done()
		}()
		closeWait.Add(1)

		timeout := time.NewTimer(0)
		received := true

		for {
			select {
			case _, _ = <-closeCh:
				return
			case data := <-readerCh:
				received = true
				if data.Type == "HEARTBEAT" {
					continue
				}
				priceCh <- data
			case <-timeout.C:
				timeout.Reset(r.Connection.Timeout)
				if !received {
					var err error = &StreamHeartbeatBroken{ErrorMessage: "Heartbeat was broken"}
					errorCh <- xerrors.Errorf("Get pricing stream heartbeat was broken: %w", err)
					return
				}
				received = false
			}
		}
	}()

	return &PriceChannels{
		Price:     priceCh,
		Error:     errorCh,
		close:     closeCh,
		closeWait: closeWait,
	}, nil
}

func (ch *PriceChannels) Close() {
	close(ch.close)
}
