package pitOrgan

import (
	"golang.org/x/xerrors"
	"strconv"
	"strings"
)

/* Receivers */

type ReceiverTrades struct {
	AccountID  string
	Connection *Connection
}

func (r *ReceiverAccountID) Trades() *ReceiverTrades {
	return &ReceiverTrades{
		AccountID:  r.AccountID,
		Connection: r.Connection,
	}
}

/* Params */

type GetTradesParams struct {
	IDs        []string
	State      string
	Instrument string
	Count      int
	BeforeID   string
}

/* Schemas */

type GetTradesSchema struct {
	Trades            []*TradeDefinition      `json:"trades,omitempty"`
	LastTransactionID TransactionIDDefinition `json:"lastTransactionID,omitempty"`
}

/* Errors */

/* API */

// GET /v3/accounts/{accountID}/trades
func (r *ReceiverTrades) Get(params *GetTradesParams) (*GetTradesSchema, error) {
	resp, err := r.Connection.request(
		&requestParams{
			method:   "Get",
			endPoint: "/v3/accounts/" + r.AccountID + "/trades",
			headers: []header{
				{key: "Accept-Datetime-Format", value: "RFC3339"},
			},
			queries: func() []query {
				q := make([]query, 0, 5)

				if len(params.IDs) > 0 {
					q = append(q, query{key: "ids", value: strings.Join(params.IDs, ",")})
				}

				if len(params.State) > 0 {
					q = append(q, query{key: "state", value: params.State})
				}

				if len(params.Instrument) > 0 {
					q = append(q, query{key: "instrument", value: params.Instrument})
				}

				if params.Count != 0 {
					q = append(q, query{key: "count", value: strconv.Itoa(params.Count)})
				}

				if len(params.BeforeID) > 0 {
					q = append(q, query{key: "beforeID", value: params.BeforeID})
				}

				return q
			}(),
		},
	)
	if err != nil {
		return nil, xerrors.Errorf("Get trades canceled: %w", err)
	}
	defer resp.Body.Close()

	var data interface{}
	switch resp.StatusCode {
	case 200:
		data = new(GetTradesSchema)
	}

	data, err = parseResponse(resp, data, r.Connection.strict)
	if err != nil {
		return nil, xerrors.Errorf("Get trades failed: %w", err)
	}
	return data.(*GetTradesSchema), nil
}

// TODO: GET /v3/accounts/{accountID}/openTrades

// TODO: GET /v3/accounts/{accountID}/trades/{tradeSpecifier}

// TODO: PUT /v3/accounts/{accountID}/trades/{tradeSpecifier}/close

// TODO: PUT /v3/accounts/{accountID}/trades/{tradeSpecifier}/clientExtensions

// TODO: PUT /v3/accounts/{accountID}/trades/{tradeSpecifier}/orders
