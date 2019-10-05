package pitOrgan

import (
	"bufio"
	"context"
	"encoding/json"
	"golang.org/x/xerrors"
	"strconv"
	"strings"
	"sync"
	"time"
)

/* Receivers */

type ReceiverPricing struct {
	AccountID  string
	Connection *Connection
}

func (r *ReceiverAccountID) Pricing() *ReceiverPricing {
	return &ReceiverPricing{
		AccountID:  r.AccountID,
		Connection: r.Connection,
	}
}

type ReceiverPricingStream struct {
	AccountID  string
	Connection *Connection
}

func (r *ReceiverPricing) Stream() *ReceiverPricingStream {
	return &ReceiverPricingStream{
		AccountID:  r.AccountID,
		Connection: r.Connection,
	}
}

/* Params */

type GetPricingParams struct {
	Instruments            []string
	Since                  time.Time
	IncludeUnitsAvailable  *bool
	IncludeHomeConversions *bool
}

type GetPricingStreamParams struct {
	BufferSize  int
	Instruments []string
}

/* Schemas */

type GetPricingSchema struct {
	Prices          []*PriceDefinition           `json:"prices,omitempty"`
	HomeConversions []*HomeConversionsDefinition `json:"homeConversions,omitempty"`
	Time            DateTimeDefinition           `json:"time,omitempty"`
}

/* Streams */

type PriceChannels struct {
	PriceCh   <-chan *PriceDefinition
	lastError error
	errorCh   <-chan error
	close     context.CancelFunc
	closeWait *sync.WaitGroup
}

/* API */

// GET /v3/accounts/{accountID}/pricing
func (r *ReceiverPricing) Get(ctx context.Context, params *GetPricingParams) (*GetPricingSchema, error) {
	resp, err := r.Connection.request(
		ctx,
		&requestParams{
			method:   "GET",
			endPoint: "/v3/accounts/" + r.AccountID + "/pricing",
			headers: []header{
				{key: "Accept-Datetime-Format", value: "RFC3339"},
			},
			queries: func() []query {
				q := make([]query, 0, 4)
				q = append(q, query{key: "instruments", value: strings.Join(params.Instruments, ",")})
				if params.IncludeUnitsAvailable != nil {
					q = append(q, query{key: "includeUnitsAvailable", value: strconv.FormatBool(*params.IncludeUnitsAvailable)})
				}
				if params.IncludeHomeConversions != nil {
					q = append(q, query{key: "includeHomeConversions", value: strconv.FormatBool(*params.IncludeHomeConversions)})
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

	data, err = parseResponse(resp, data, r.Connection.strict)
	if err != nil {
		return nil, xerrors.Errorf("Get pricing failed: %w", err)
	}
	return data.(*GetPricingSchema), nil
}

// GET /v3/accounts/{accountID}/pricing/stream
func (r *ReceiverPricingStream) Get(ctx context.Context, params *GetPricingStreamParams) (*PriceChannels, error) {
	childCtx, cancel := context.WithCancel(ctx)

	resp, err := r.Connection.stream(
		childCtx,
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
		defer func() {
			resp.Body.Close()
			cancel()
		}()
		var err interface{}
		_, err = parseResponse(resp, err, r.Connection.strict)
		return nil, xerrors.Errorf("Get pricing stream failed: %w", err)
	}

	closeWait := new(sync.WaitGroup)

	priceCh := make(chan *PriceDefinition, params.BufferSize)
	errorCh := make(chan error, 3)

	// closeChがcloseされたらstreamを終了するgoroutine
	// 下記の readre.ReadBytes はデータを受信しないと処理が進まないため
	// streamの途中で途切れると永遠に待ち続けてしまうので終了用goroutineを用意。
	go func() {
		defer func() {
			resp.Body.Close()
			cancel()
			closeWait.Done()
		}()
		closeWait.Add(1)
		<-childCtx.Done()
	}()

	// 受信したデータ(JSON)を構造体にしてreaderCh channelに送信するgoroutine
	readerCh := make(chan *PriceDefinition, params.BufferSize)
	go func() {
		defer func() {
			close(readerCh)
			cancel()
			closeWait.Done()
		}()
		closeWait.Add(1)

		reader := bufio.NewReader(resp.Body)
		for {
			line, err := reader.ReadBytes('\n')
			if err != nil {
				select {
				case <-childCtx.Done():
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
			cancel()
			closeWait.Done()
		}()
		closeWait.Add(1)

		timeout := time.NewTimer(0)
		received := true

		for {
			select {
			case <-childCtx.Done():
				return
			case data := <-readerCh:
				received = true
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
		PriceCh:   priceCh,
		lastError: nil,
		errorCh:   errorCh,
		close:     cancel,
		closeWait: closeWait,
	}, nil
}

/* Utils */

func (ch *PriceChannels) Close() {
	ch.close()
	ch.closeWait.Wait()
}

func (ch *PriceChannels) Err() error {
	if ch.lastError == nil {
		select {
		case ch.lastError = <-ch.errorCh:
		default:
		}
	}
	return ch.lastError
}
