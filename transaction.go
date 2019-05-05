package pitOrgan

// Receivers

type ReceiverTransactions struct {
	AccountID  string
	Connection *Connection
}

type ReceiverTransactionsIdrange struct {
	AccountID  string
	Connection *Connection
}

type ReceiverTransactionsStream struct {
	AccountID  string
	Connection *Connection
}


func (r *ReceiverAccountID) Transactions() *ReceiverTransactions {
	return &ReceiverTransactions{
		AccountID:  r.AccountID,
		Connection: r.Connection,
	}
}

func (r *ReceiverTransactions) Idrange() *ReceiverTransactionsIdrange {
	return &ReceiverTransactionsIdrange{
		AccountID:  r.AccountID,
		Connection: r.Connection,
	}
}

func (r *ReceiverTransactions) Stream() *ReceiverTransactionsStream {
	return &ReceiverTransactionsStream{
		AccountID:  r.AccountID,
		Connection: r.Connection,
	}
}

// TODO: GET /v3/accounts/{accountID}/transactions

// TODO: GET /v3/accounts/{accountID}/transactions/{transactionID}

// TODO: GET /v3/accounts/{accountID}/transactions/idrange

// TODO: GET /v3/accounts/{accountID}/transactions/sinceid

// TODO: GET /v3/accounts/{accountID}/transactions/stream
