package pitOrgan

import (
	"golang.org/x/xerrors"
)

// Receivers

type ReceiverOrders struct {
	AccountID  string
	Connection *Connection
}

// Params

type PostOrdersBodyParams struct {
	Order interface{} `json:"order"`
}

type PostOrdersParams struct {
	Body PostOrdersBodyParams
}

// Schemas

type PostOrdersSchema struct {
	OrderCreateTransaction        TransactionDefinition            `json:"orderCreateTransaction"`
	OrderFillTransaction          OrderFillTransactionDefinition   `json:"orderFillTransaction"`
	OrderCancelTransaction        OrderCancelTransactionDefinition `json:"orderCancelTransaction"`
	OrderReissueTransaction       TransactionDefinition            `json:"orderReissueTransaction"`
	OrderReissueRejectTransaction TransactionDefinition            `json:"orderReissueRejectTransaction"`
	RelatedTransactionIDs         []TransactionIDDefinition        `json:"relatedTransactionIDs"`
	LastTransactionID             TransactionIDDefinition          `json:"lastTransactionID"`
}

// Errors

type PostOrdersBadRequestError struct {
	OrderRejectTransaction TransactionDefinition     `json:"orderRejectTransaction"`
	RelatedTransactionIDs  []TransactionIDDefinition `json:"relatedTransactionIDs"`
	LastTransactionID      TransactionIDDefinition   `json:"lastTransactionID"`
	ErrorCode              string                    `json:"errorCode"`
	ErrorMessage           string                    `json:"errorMessage"`
}

func (r *PostOrdersBadRequestError) Error() string {
	// TODO: エラーを整える
	return r.ErrorMessage
}

type PostOrdersNotFoundError struct {
	OrderRejectTransaction TransactionDefinition     `json:"orderRejectTransaction"`
	RelatedTransactionIDs  []TransactionIDDefinition `json:"relatedTransactionIDs"`
	LastTransactionID      TransactionIDDefinition   `json:"lastTransactionID"`
	ErrorCode              string                    `json:"errorCode"`
	ErrorMessage           string                    `json:"errorMessage"`
}

func (r *PostOrdersNotFoundError) Error() string {
	// TODO: エラーを整える
	return r.ErrorMessage
}

func (r *ReceiverAccountID) Orders() *ReceiverOrders {
	return &ReceiverOrders{
		AccountID:  r.AccountID,
		Connection: r.Connection,
	}
}

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

	data, err = parseResponse(resp, data)
	if err != nil {
		return nil, xerrors.Errorf("Post orders failed: %w", err)
	}
	return data.(*PostOrdersSchema), nil
}

// TODO: GET /v3/accounts/{accountID}/orders

// TODO: GET /v3/accounts/{accountID}/pendingOrders

// TODO: GET /v3/accounts/{accountID}/orders/{orderSpecifier}

// TODO: PUT /v3/accounts/{accountID}/orders/{orderSpecifier}

// TODO: PUT /v3/accounts/{accountID}/orders/{orderSpecifier}/cancel

// TODO: PUT /v3/accounts/{accountID}/orders/{orderSpecifier}/clientExtensions
