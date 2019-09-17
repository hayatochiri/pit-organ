package pitOrgan

import (
	"golang.org/x/xerrors"
)

/* Receivers */

type ReceiverPositions struct {
	AccountID  string
	Connection *Connection
}

func (r *ReceiverAccountID) Positions() *ReceiverPositions {
	return &ReceiverPositions{
		AccountID:  r.AccountID,
		Connection: r.Connection,
	}
}

type ReceiverOpenPositions struct {
	AccountID  string
	Connection *Connection
}

func (r *ReceiverAccountID) OpenPositions() *ReceiverOpenPositions {
	return &ReceiverOpenPositions{
		AccountID:  r.AccountID,
		Connection: r.Connection,
	}
}

/* Params */

/* Schemas */

type GetPositionsSchema struct {
	Positions         []PositionDefinition    `json:"positions,omitempty"`
	LastTransactionID TransactionIDDefinition `json:"lastTransactionID,omitempty"`
}

type GetOpenPositionsSchema struct {
	Positions         []PositionDefinition    `json:"positions,omitempty"`
	LastTransactionID TransactionIDDefinition `json:"lastTransactionID,omitempty"`
}

/* Errors */

/* API */

// GET /v3/accounts/{accountID}/positions
func (r *ReceiverPositions) Get() (*GetPositionsSchema, error) {
	resp, err := r.Connection.request(
		&requestParams{
			method:   "Get",
			endPoint: "/v3/accounts/" + r.AccountID + "/positions",
			headers: []header{
				{key: "Accept-Datetime-Format", value: "RFC3339"},
			},
		},
	)
	if err != nil {
		return nil, xerrors.Errorf("Get positions canceled: %w", err)
	}
	defer resp.Body.Close()

	var data interface{}
	switch resp.StatusCode {
	case 200:
		data = new(GetPositionsSchema)
	}

	data, err = parseResponse(resp, data, r.Connection.strict)
	if err != nil {
		return nil, xerrors.Errorf("Get positions failed: %w", err)
	}
	return data.(*GetPositionsSchema), nil
}

// GET /v3/accounts/{accountID}/openPositions
func (r *ReceiverOpenPositions) Get() (*GetOpenPositionsSchema, error) {
	resp, err := r.Connection.request(
		&requestParams{
			method:   "Get",
			endPoint: "/v3/accounts/" + r.AccountID + "/openPositions",
			headers: []header{
				{key: "Accept-Datetime-Format", value: "RFC3339"},
			},
		},
	)
	if err != nil {
		return nil, xerrors.Errorf("Get open positions canceled: %w", err)
	}
	defer resp.Body.Close()

	var data interface{}
	switch resp.StatusCode {
	case 200:
		data = new(GetOpenPositionsSchema)
	}

	data, err = parseResponse(resp, data, r.Connection.strict)
	if err != nil {
		return nil, xerrors.Errorf("Get open positions failed: %w", err)
	}
	return data.(*GetOpenPositionsSchema), nil
}

// TODO: GET /v3/accounts/{accountID}/positions/{instrument}

// TODO: PUT /v3/accounts/{accountID}/positions/{instrument}/close
