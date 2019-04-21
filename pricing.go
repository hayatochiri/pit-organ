package pitOrgan

import (
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

// Schemas

type GetPricingSchema struct {
	Prices          []PriceDefinition           `json:"prices"`
	HomeConversions []HomeConversionsDefinition `json:"homeConversions"`
	Time            DateTimeDefinition          `json:"time"`
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

// TODO: GET /v3/accounts/{accountID}/pricing/stream

