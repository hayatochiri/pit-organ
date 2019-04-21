package pitOrgan

// Receivers

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

// TODO: POST /v3/accounts/{accountID}/orders

// TODO: GET /v3/accounts/{accountID}/orders

// TODO: GET /v3/accounts/{accountID}/pendingOrders

// TODO: GET /v3/accounts/{accountID}/orders/{orderSpecifier}

// TODO: PUT /v3/accounts/{accountID}/orders/{orderSpecifier}

// TODO: PUT /v3/accounts/{accountID}/orders/{orderSpecifier}/cancel

// TODO: PUT /v3/accounts/{accountID}/orders/{orderSpecifier}/clientExtensions
