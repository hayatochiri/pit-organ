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

type ReceiverOpenTrades struct {
	AccountID  string
	Connection *Connection
}

func (r *ReceiverAccountID) OpenTrades() *ReceiverOpenTrades {
	return &ReceiverOpenTrades{
		AccountID:  r.AccountID,
		Connection: r.Connection,
	}
}

type ReceiverTradeSpecifier struct {
	AccountID      string
	Connection     *Connection
	TradeSpecifier string
}

func (r *ReceiverTrades) TradeSpecifier(tradeSpecifier string) *ReceiverTradeSpecifier {
	return &ReceiverTradeSpecifier{
		AccountID:      r.AccountID,
		Connection:     r.Connection,
		TradeSpecifier: tradeSpecifier,
	}
}

type ReceiverTradeSpecifierClose struct {
	AccountID      string
	Connection     *Connection
	TradeSpecifier string
}

func (r *ReceiverTradeSpecifier) Close() *ReceiverTradeSpecifierClose {
	return &ReceiverTradeSpecifierClose{
		AccountID:      r.AccountID,
		Connection:     r.Connection,
		TradeSpecifier: r.TradeSpecifier,
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

type PutTradeSpecifierCloseBodyParams struct {
	Units string `json:"units,omitempty"`
}

type PutTradeSpecifierCloseParams struct {
	Body *PutTradeSpecifierCloseBodyParams
}

/* Schemas */

type GetTradesSchema struct {
	Trades            []*TradeDefinition      `json:"trades,omitempty"`
	LastTransactionID TransactionIDDefinition `json:"lastTransactionID,omitempty"`
}

type GetOpenTradesSchema struct {
	Trades            []*TradeDefinition      `json:"trades,omitempty"`
	LastTransactionID TransactionIDDefinition `json:"lastTransactionID,omitempty"`
}

type GetTradeSpecifierSchema struct {
	Trade             *TradeDefinition        `json:"trade,omitempty"`
	LastTransactionID TransactionIDDefinition `json:"lastTransactionID,omitempty"`
}

type PutTradeSpecifierCloseSchema struct {
	OrderCreateTransaction *TransactionDefinition    `json:"orderCreateTransaction,omitempty"`
	OrderFillTransaction   *TransactionDefinition    `json:"orderFillTransaction,omitempty"`
	OrderCancelTransaction *TransactionDefinition    `json:"orderCancelTransaction,omitempty"`
	RelatedTransactionIDs  []TransactionIDDefinition `json:"relatedTransactionIDs,omitempty"`
	LastTransactionID      TransactionIDDefinition   `json:"lastTransactionID,omitempty"`
}

/* Errors */

type PutTradeSpecifierCloseBadRequestError struct {
	OrderRejectTransaction *TransactionDefinition `json:"orderRejectTransaction,omitempty"`
	ErrorCode              string                 `json:"errorCode,omitempty"`
	ErrorMessage           string                 `json:"errorMessage,omitempty"`
}

func (r *PutTradeSpecifierCloseBadRequestError) Error() string {
	// TODO: エラーを整える
	return r.ErrorMessage
}

type PutTradeSpecifierCloseNotFoundError struct {
	OrderRejectTransaction *TransactionDefinition    `json:"orderRejectTransaction,omitempty"`
	LastTransactionID      TransactionIDDefinition   `json:"lastTransactionID,omitempty"`
	RelatedTransactionIDs  []TransactionIDDefinition `json:"relatedTransactionIDs,omitempty"`
	ErrorCode              string                    `json:"errorCode,omitempty"`
	ErrorMessage           string                    `json:"errorMessage,omitempty"`
}

func (r *PutTradeSpecifierCloseNotFoundError) Error() string {
	// TODO: エラーを整える
	return r.ErrorMessage
}

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

// GET /v3/accounts/{accountID}/openTrades
func (r *ReceiverOpenTrades) Get() (*GetOpenTradesSchema, error) {
	resp, err := r.Connection.request(
		&requestParams{
			method:   "Get",
			endPoint: "/v3/accounts/" + r.AccountID + "/openTrades",
			headers: []header{
				{key: "Accept-Datetime-Format", value: "RFC3339"},
			},
		},
	)
	if err != nil {
		return nil, xerrors.Errorf("Get open trades canceled: %w", err)
	}
	defer resp.Body.Close()

	var data interface{}
	switch resp.StatusCode {
	case 200:
		data = new(GetOpenTradesSchema)
	}

	data, err = parseResponse(resp, data, r.Connection.strict)
	if err != nil {
		return nil, xerrors.Errorf("Get open trades failed: %w", err)
	}
	return data.(*GetOpenTradesSchema), nil
}

// GET /v3/accounts/{accountID}/trades/{tradeSpecifier}
func (r *ReceiverTradeSpecifier) Get() (*GetTradeSpecifierSchema, error) {
	resp, err := r.Connection.request(
		&requestParams{
			method:   "Get",
			endPoint: "/v3/accounts/" + r.AccountID + "/trades/" + r.TradeSpecifier,
			headers: []header{
				{key: "Accept-Datetime-Format", value: "RFC3339"},
			},
		},
	)
	if err != nil {
		return nil, xerrors.Errorf("Get trade specifier canceled: %w", err)
	}
	defer resp.Body.Close()

	var data interface{}
	switch resp.StatusCode {
	case 200:
		data = new(GetTradeSpecifierSchema)
	}

	data, err = parseResponse(resp, data, r.Connection.strict)
	if err != nil {
		return nil, xerrors.Errorf("Get trade specifier failed: %w", err)
	}
	return data.(*GetTradeSpecifierSchema), nil
}

// PUT /v3/accounts/{accountID}/trades/{tradeSpecifier}/close
func (r *ReceiverTradeSpecifierClose) Put(params *PutTradeSpecifierCloseParams) (*PutTradeSpecifierCloseSchema, error) {
	resp, err := r.Connection.request(
		&requestParams{
			method:   "PUT",
			endPoint: "/v3/accounts/" + r.AccountID + "/trades/" + r.TradeSpecifier + "/close",
			headers: []header{
				{key: "Accept-Datetime-Format", value: "RFC3339"},
			},
			body: params.Body,
		},
	)
	if err != nil {
		return nil, xerrors.Errorf("Put trade specifier close canceled: %w", err)
	}
	defer resp.Body.Close()

	var data interface{}
	switch resp.StatusCode {
	case 200:
		data = new(PutTradeSpecifierCloseSchema)
	case 400:
		data = new(PutTradeSpecifierCloseBadRequestError)
	case 404:
		data = new(PutTradeSpecifierCloseNotFoundError)
	}

	data, err = parseResponse(resp, data, r.Connection.strict)
	if err != nil {
		return nil, xerrors.Errorf("Put trade specifier close failed: %w", err)
	}
	return data.(*PutTradeSpecifierCloseSchema), nil
}

// TODO: PUT /v3/accounts/{accountID}/trades/{tradeSpecifier}/clientExtensions

// TODO: PUT /v3/accounts/{accountID}/trades/{tradeSpecifier}/orders
