package pitOrgan

import (
	"golang.org/x/xerrors"
)

// Receivers

type ReceiverAccounts struct {
	Connection *Connection
}

// Schemas

type GetAccountsSchema struct {
	Accounts []AccountPropertiesDefinition `json:"accounts"`
}

func (c *Connection) Accounts() *ReceiverAccounts {
	return &ReceiverAccounts{
		Connection: c,
	}
}

// GET /v3/accounts
func (r *ReceiverAccounts) Get() (*GetAccountsSchema, error) {
	resp, err := r.Connection.request(
		&requestParams{
			method:   "GET",
			endPoint: "/v3/accounts",
		},
	)
	if err != nil {
		return nil, xerrors.Errorf("Get accounts canceled: %w", err)
	}
	defer resp.Body.Close()

	var data interface{}
	switch resp.StatusCode {
	case 200:
		data = new(GetAccountsSchema)
	}

	data, err = parseResponse(resp, data)
	if err != nil {
		return nil, xerrors.Errorf("Get accounts failed: %w", err)
	}
	return data.(*GetAccountsSchema), nil
}

// TODO: GET /v3/accounts/{accountID}

// TODO: GET /v3/accounts/{accountID}/summary

// TODO: GET /v3/accounts/{accountID}/instruments

// TODO: PATCH /v3/accounts/{accountID}/configuration

// TODO: GET /v3/accounts/{accountID}/changes
