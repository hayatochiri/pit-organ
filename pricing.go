package pitOrgan

import (
	"bufio"
	"encoding/json"
	"golang.org/x/xerrors"
	"net/http"
	"strings"
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
	Price <-chan *PriceDefinition
	Error <-chan error
	close chan<- interface{}
	resp  *http.Response // for test
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

	// TODO: Catch Response code 401, 404, 405

	priceCh := make(chan *PriceDefinition, params.BufferSize)
	errorCh := make(chan error, 1)
	closeCh := make(chan interface{})
	go func() {
		defer func() {
			resp.Body.Close()
			close(priceCh)
		}()

		reader := bufio.NewReader(resp.Body)

		for {
			select {
			case _, ok := <-closeCh:
				if ok == false {
					return
				}
			default:
			}

			line, err := reader.ReadBytes('\n')
			if err != nil {
				errorCh <- xerrors.Errorf("TODO Error Message: %w", err)
				return
			}

			data := new(PriceDefinition)
			err = json.Unmarshal(line, data)
			if err != nil {
				errorCh <- xerrors.Errorf("TODO Error Message: %w", err)
				return
			}

			if data.Type == "HEARTBEAT" {
				continue
			}

			priceCh <- data
		}
	}()

	return &PriceChannels{
		Price: priceCh,
		Error: errorCh,
		close: closeCh,
		resp:  resp,
	}, nil
}

func (ch *PriceChannels) Close() {
	close(ch.close)
}
