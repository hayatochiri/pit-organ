package pitOrgan

// Receivers

type ReceiverAccounts struct {
	Connection *Connection
}

func (c *Connection) Accounts() *ReceiverAccounts {
	return &ReceiverAccounts{
		Connection: c,
	}
}

// TODO: GET /v3/accounts

// TODO: GET /v3/accounts/{accountID}

// TODO: GET /v3/accounts/{accountID}/summary

// TODO: GET /v3/accounts/{accountID}/instruments

// TODO: PATCH /v3/accounts/{accountID}/configuration

// TODO: GET /v3/accounts/{accountID}/changes

