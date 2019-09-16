package pitOrgan

import (
	"golang.org/x/xerrors"
	"strconv"
	"strings"
)

/* Receivers */

type ReceiverOrders struct {
	AccountID  string
	Connection *Connection
}

func (r *ReceiverAccountID) Orders() *ReceiverOrders {
	return &ReceiverOrders{
		AccountID:  r.AccountID,
		Connection: r.Connection,
	}
}

type ReceiverPendingOrders struct {
	AccountID  string
	Connection *Connection
}

func (r *ReceiverAccountID) PendingOrders() *ReceiverPendingOrders {
	return &ReceiverPendingOrders{
		AccountID:  r.AccountID,
		Connection: r.Connection,
	}
}

type ReceiverOrderSpecifier struct {
	AccountID      string
	Connection     *Connection
	OrderSpecifier string
}

func (r *ReceiverOrders) OrderSpecifier(orderSpecifier string) *ReceiverOrderSpecifier {
	return &ReceiverOrderSpecifier{
		AccountID:      r.AccountID,
		Connection:     r.Connection,
		OrderSpecifier: orderSpecifier,
	}
}

// Params

type PostOrdersBodyParams struct {
	Order OrderRequestDefinition `json:"order"`
}

type PostOrdersParams struct {
	Body PostOrdersBodyParams
}

type GetOrdersParams struct {
	IDs        []string
	State      interface{}
	Instrument interface{}
	Count      int
	BeforeID   interface{}
}

type PutOrderSpecifierBodyParams struct {
	Order OrderRequestDefinition `json:"order"`
}

type PutOrderSpecifierParams struct {
	Body PutOrderSpecifierBodyParams
}

/* Schemas */

type PostOrdersSchema struct {
	OrderCreateTransaction        *TransactionDefinition    `json:"orderCreateTransaction,omitempty"`
	OrderFillTransaction          *TransactionDefinition    `json:"orderFillTransaction,omitempty"`
	OrderCancelTransaction        *TransactionDefinition    `json:"orderCancelTransaction,omitempty"`
	OrderReissueTransaction       *TransactionDefinition    `json:"orderReissueTransaction,omitempty"`
	OrderReissueRejectTransaction *TransactionDefinition    `json:"orderReissueRejectTransaction,omitempty"`
	RelatedTransactionIDs         []TransactionIDDefinition `json:"relatedTransactionIDs,omitempty"`
	LastTransactionID             TransactionIDDefinition   `json:"lastTransactionID,omitempty"`
}

type GetOrdersSchema struct {
	Orders            []*OrderDefinition      `json:"orders,omitempty"`
	LastTransactionID TransactionIDDefinition `json:"lastTransactionID,omitempty"`
}

type GetPendingOrdersSchema struct {
	Orders            []*OrderDefinition      `json:"orders,omitempty"`
	LastTransactionID TransactionIDDefinition `json:"lastTransactionID,omitempty"`
}

type GetOrderSpecifierSchema struct {
	Order             *OrderDefinition        `json:"order,omitempty"`
	LastTransactionID TransactionIDDefinition `json:"lastTransactionID,omitempty"`
}

type PutOrderSpecifierSchema struct {
	OrderCancelTransaction          *TransactionDefinition    `json:"orderCancelTransaction,omitempty"`
	OrderCreateTransaction          *TransactionDefinition    `json:"orderCreateTransaction,omitempty"`
	OrderFillTransaction            *TransactionDefinition    `json:"orderFillTransaction,omitempty"`
	OrderReissueTransaction         *TransactionDefinition    `json:"orderReissueTransaction,omitempty"`
	OrderReissueRejectTransaction   *TransactionDefinition    `json:"orderReissueRejectTransaction,omitempty"`
	ReplacingOrderCancelTransaction *TransactionDefinition    `json:"replacingOrderCancelTransaction,omitempty"`
	RelatedTransactionIDs           []TransactionIDDefinition `json:"relatedTransactionIDs,omitempty"`
	LastTransactionID               TransactionIDDefinition   `json:"lastTransactionID,omitempty"`
}

/* Errors */

type PostOrdersBadRequestError struct {
	OrderRejectTransaction *TransactionDefinition    `json:"orderRejectTransaction,omitempty"`
	RelatedTransactionIDs  []TransactionIDDefinition `json:"relatedTransactionIDs,omitempty"`
	LastTransactionID      TransactionIDDefinition   `json:"lastTransactionID,omitempty"`
	ErrorCode              string                    `json:"errorCode,omitempty"`
	ErrorMessage           string                    `json:"errorMessage,omitempty"`
}

func (r *PostOrdersBadRequestError) Error() string {
	// TODO: エラーを整える
	return r.ErrorMessage
}

type PostOrdersNotFoundError struct {
	OrderRejectTransaction *TransactionDefinition    `json:"orderRejectTransaction,omitempty"`
	RelatedTransactionIDs  []TransactionIDDefinition `json:"relatedTransactionIDs,omitempty"`
	LastTransactionID      TransactionIDDefinition   `json:"lastTransactionID,omitempty"`
	ErrorCode              string                    `json:"errorCode,omitempty"`
	ErrorMessage           string                    `json:"errorMessage,omitempty"`
}

func (r *PostOrdersNotFoundError) Error() string {
	// TODO: エラーを整える
	return r.ErrorMessage
}

type PutOrderSpecifierBadRequestError struct {
	OrderRejectTransaction *TransactionDefinition    `json:"orderRejectTransaction"`
	RelatedTransactionIDs  []TransactionIDDefinition `json:"relatedTransactionIDs"`
	LastTransactionID      TransactionIDDefinition   `json:"lastTransactionID"`
	ErrorCode              string                    `json:"errorCode"`
	ErrorMessage           string                    `json:"errorMessage"`
}

func (r *PutOrderSpecifierBadRequestError) Error() string {
	// TODO: エラーを整える
	return r.ErrorMessage
}

type PutOrderSpecifierNotFoundError struct {
	OrderCancelRejectTransaction *TransactionDefinition    `json:"orderCancelRejectTransaction"`
	RelatedTransactionIDs        []TransactionIDDefinition `json:"relatedTransactionIDs"`
	LastTransactionID            TransactionIDDefinition   `json:"lastTransactionID"`
	ErrorCode                    string                    `json:"errorCode"`
	ErrorMessage                 string                    `json:"errorMessage"`
}

func (r *PutOrderSpecifierNotFoundError) Error() string {
	// TODO: エラーを整える
	return r.ErrorMessage
}

/* API */

// POST /v3/accounts/{accountID}/orders
func (r *ReceiverOrders) Post(params *PostOrdersParams) (*PostOrdersSchema, error) {
	resp, err := r.Connection.request(
		&requestParams{
			method:   "POST",
			endPoint: "/v3/accounts/" + r.AccountID + "/orders",
			headers: []header{
				{key: "Accept-Datetime-Format", value: "RFC3339"},
			},
			body: params.Body,
		},
	)
	if err != nil {
		return nil, xerrors.Errorf("Post orders canceled: %w", err)
	}
	defer resp.Body.Close()

	var data interface{}
	switch resp.StatusCode {
	case 201:
		data = new(PostOrdersSchema)
	case 400:
		data = new(PostOrdersBadRequestError)
	case 404:
		data = new(PostOrdersNotFoundError)
	}

	data, err = parseResponse(resp, data, r.Connection.strict)
	if err != nil {
		return nil, xerrors.Errorf("Post orders failed: %w", err)
	}
	return data.(*PostOrdersSchema), nil
}

// GET /v3/accounts/{accountID}/orders
func (r *ReceiverOrders) Get(params *GetOrdersParams) (*GetOrdersSchema, error) {
	resp, err := r.Connection.request(
		&requestParams{
			method:   "GET",
			endPoint: "/v3/accounts/" + r.AccountID + "/orders",
			headers: []header{
				{key: "Accept-Datetime-Format", value: "RFC3339"},
			},
			queries: func() []query {
				q := make([]query, 0, 5)

				if len(params.IDs) > 0 {
					q = append(q, query{key: "ids", value: strings.Join(params.IDs, ",")})
				}

				if str, ok := params.State.(string); ok {
					q = append(q, query{key: "state", value: str})
				}

				if str, ok := params.Instrument.(string); ok {
					q = append(q, query{key: "instrument", value: str})
				}

				if params.Count != 0 {
					q = append(q, query{key: "count", value: strconv.Itoa(params.Count)})
				}

				if str, ok := params.BeforeID.(string); ok {
					q = append(q, query{key: "beforeID", value: str})
				}

				return q
			}(),
		},
	)
	if err != nil {
		return nil, xerrors.Errorf("Get orders canceled: %w", err)
	}
	defer resp.Body.Close()

	var data interface{}
	switch resp.StatusCode {
	case 200:
		data = new(GetOrdersSchema)
	}

	data, err = parseResponse(resp, data, r.Connection.strict)
	if err != nil {
		return nil, xerrors.Errorf("Get orders failed: %w", err)
	}
	return data.(*GetOrdersSchema), nil
}

// GET /v3/accounts/{accountID}/pendingOrders
func (r *ReceiverPendingOrders) Get() (*GetPendingOrdersSchema, error) {
	resp, err := r.Connection.request(
		&requestParams{
			method:   "GET",
			endPoint: "/v3/accounts/" + r.AccountID + "/pendingOrders",
			headers: []header{
				{key: "Accept-Datetime-Format", value: "RFC3339"},
			},
		},
	)
	if err != nil {
		return nil, xerrors.Errorf("Get pending orders canceled: %w", err)
	}
	defer resp.Body.Close()

	var data interface{}
	switch resp.StatusCode {
	case 200:
		data = new(GetPendingOrdersSchema)
	}

	data, err = parseResponse(resp, data, r.Connection.strict)
	if err != nil {
		return nil, xerrors.Errorf("Get pending orders failed: %w", err)
	}
	return data.(*GetPendingOrdersSchema), nil
}

// GET /v3/accounts/{accountID}/orders/{orderSpecifier}
func (r *ReceiverOrderSpecifier) Get() (*GetOrderSpecifierSchema, error) {
	resp, err := r.Connection.request(
		&requestParams{
			method:   "GET",
			endPoint: "/v3/accounts/" + r.AccountID + "/orders/" + r.OrderSpecifier,
			headers: []header{
				{key: "Accept-Datetime-Format", value: "RFC3339"},
			},
		},
	)
	if err != nil {
		return nil, xerrors.Errorf("Get order specifier canceled: %w", err)
	}
	defer resp.Body.Close()

	var data interface{}
	switch resp.StatusCode {
	case 200:
		data = new(GetOrderSpecifierSchema)
	}

	data, err = parseResponse(resp, data, r.Connection.strict)
	if err != nil {
		return nil, xerrors.Errorf("Get order specifier failed: %w", err)
	}
	return data.(*GetOrderSpecifierSchema), nil
}

// PUT /v3/accounts/{accountID}/orders/{orderSpecifier}
func (r *ReceiverOrderSpecifier) Put(params *PutOrderSpecifierParams) (*PutOrderSpecifierSchema, error) {
	resp, err := r.Connection.request(
		&requestParams{
			method:   "PUT",
			endPoint: "/v3/accounts/" + r.AccountID + "/orders/" + r.OrderSpecifier,
			headers: []header{
				{key: "Accept-Datetime-Format", value: "RFC3339"},
			},
			body: params.Body,
		},
	)
	if err != nil {
		return nil, xerrors.Errorf("Put order specifier canceled: %w", err)
	}
	defer resp.Body.Close()

	var data interface{}
	switch resp.StatusCode {
	case 201:
		data = new(PutOrderSpecifierSchema)
	case 400:
		data = new(PutOrderSpecifierBadRequestError)
	case 404:
		data = new(PutOrderSpecifierNotFoundError)
	}

	data, err = parseResponse(resp, data, r.Connection.strict)
	if err != nil {
		return nil, xerrors.Errorf("Put order specifier failed: %w", err)
	}
	return data.(*PutOrderSpecifierSchema), nil
}

// TODO: PUT /v3/accounts/{accountID}/orders/{orderSpecifier}/cancel

// TODO: PUT /v3/accounts/{accountID}/orders/{orderSpecifier}/clientExtensions
