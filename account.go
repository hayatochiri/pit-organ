package pitOrgan

import (
	"golang.org/x/xerrors"
	"strings"
)

// Receivers

type ReceiverAccounts struct {
	Connection *Connection
}

type ReceiverAccountID struct {
	AccountID  string
	Connection *Connection
}

type ReceiverAccountSummary struct {
	AccountID  string
	Connection *Connection
}

type ReceiverAccountInstruments struct {
	AccountID  string
	Connection *Connection
}

// Params

type GetAccountInstrumentsParams struct {
	Instruments []string
}

// Schemas

type GetAccountsSchema struct {
	Accounts []AccountPropertiesDefinition `json:"accounts"`
}

type GetAccountIDSchema struct {
	Account           AccountDefinition       `json:"account"`
	LastTransactionID TransactionIDDefinition `json:"lastTransactionID"`
}

type GetAccountSummarySchema struct {
	Account           AccountSummaryDefinition `json:"account"`
	LastTransactionID TransactionIDDefinition  `json:"lastTransactionID"`
}

type GetAccountInstrumentsSchema struct {
	Instruments       []InstrumentDefinition  `json:"instruments"`
	LastTransactionID TransactionIDDefinition `json:"lastTransactionID"`
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

func (r *ReceiverAccounts) AccountID(id string) *ReceiverAccountID {
	return &ReceiverAccountID{
		AccountID:  id,
		Connection: r.Connection,
	}
}

// GET /v3/accounts/{accountID}
func (r *ReceiverAccountID) Get() (*GetAccountIDSchema, error) {
	resp, err := r.Connection.request(
		&requestParams{
			method:   "GET",
			endPoint: "/v3/accounts/" + r.AccountID,
			headers: []header{
				{key: "Accept-Datetime-Format", value: "RFC3339"},
			},
		},
	)
	if err != nil {
		return nil, xerrors.Errorf("Get account ID canceled: %w", err)
	}
	defer resp.Body.Close()

	var data interface{}
	switch resp.StatusCode {
	case 200:
		data = new(GetAccountIDSchema)
	}

	data, err = parseResponse(resp, data)
	if err != nil {
		return nil, xerrors.Errorf("Get account ID failed: %w", err)
	}
	return data.(*GetAccountIDSchema), nil
}

func (r *ReceiverAccountID) Summary() *ReceiverAccountSummary {
	return &ReceiverAccountSummary{
		AccountID:  r.AccountID,
		Connection: r.Connection,
	}
}

// GET /v3/accounts/{accountID}/summary
func (r *ReceiverAccountSummary) Get() (*GetAccountSummarySchema, error) {
	resp, err := r.Connection.request(
		&requestParams{
			method:   "GET",
			endPoint: "/v3/accounts/" + r.AccountID + "/summary",
			headers: []header{
				{key: "Accept-Datetime-Format", value: "RFC3339"},
			},
		},
	)
	if err != nil {
		return nil, xerrors.Errorf("Get account summary canceled: %w", err)
	}
	defer resp.Body.Close()

	var data interface{}
	switch resp.StatusCode {
	case 200:
		data = new(GetAccountSummarySchema)
	}

	data, err = parseResponse(resp, data)
	if err != nil {
		return nil, xerrors.Errorf("Get account summary failed: %w", err)
	}
	return data.(*GetAccountSummarySchema), nil
}

func (r *ReceiverAccountID) Instruments() *ReceiverAccountInstruments {
	return &ReceiverAccountInstruments{
		AccountID:  r.AccountID,
		Connection: r.Connection,
	}
}

// GET /v3/accounts/{accountID}/instruments
func (r *ReceiverAccountInstruments) Get(params *GetAccountInstrumentsParams) (*GetAccountInstrumentsSchema, error) {
	resp, err := r.Connection.request(
		&requestParams{
			method:   "GET",
			endPoint: "/v3/accounts/" + r.AccountID + "/instruments",
			headers: []header{
				{key: "Accept-Datetime-Format", value: "RFC3339"},
			},
			queries: []query{
				{key: "instruments", value: strings.Join(params.Instruments, ",")},
			},
		},
	)
	if err != nil {
		return nil, xerrors.Errorf("Get account instruments canceled: %w", err)
	}
	defer resp.Body.Close()

	var data interface{}
	switch resp.StatusCode {
	case 200:
		data = new(GetAccountInstrumentsSchema)
	}

	data, err = parseResponse(resp, data)
	if err != nil {
		return nil, xerrors.Errorf("Get account instruments failed: %w", err)
	}
	return data.(*GetAccountInstrumentsSchema), nil
}

// TODO: PATCH /v3/accounts/{accountID}/configuration

// TODO: GET /v3/accounts/{accountID}/changes
