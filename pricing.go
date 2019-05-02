package pitOrgan

import (
	"bufio"
	"encoding/json"
	"golang.org/x/xerrors"
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
	close chan<- struct{}
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

	priceCh := make(chan *PriceDefinition, params.BufferSize)
	errorCh := make(chan error, 1)
	closeCh := make(chan struct{})

	return &PriceChannels{
		Price: priceCh,
		Error: errorCh,
		close: closeCh,
	}, nil
}

func (ch *PriceChannels) Close() {
	close(ch.close)
}
