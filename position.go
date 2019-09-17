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

type ReceiverPositionsInstrument struct {
	AccountID  string
	Connection *Connection
	Instrument string
}

func (r *ReceiverPositions) Instrument(instrument string) *ReceiverPositionsInstrument {
	return &ReceiverPositionsInstrument{
		AccountID:  r.AccountID,
		Connection: r.Connection,
		Instrument: instrument,
	}
}

type ReceiverPositionsInstrumentClose struct {
	AccountID  string
	Connection *Connection
	Instrument string
}

func (r *ReceiverPositionsInstrument) Close() *ReceiverPositionsInstrumentClose {
	return &ReceiverPositionsInstrumentClose{
		AccountID:  r.AccountID,
		Connection: r.Connection,
		Instrument: r.Instrument,
	}
}

/* Params */

type PutPositionsInstrumentCloseBodyParams struct {
	LongUnits             string                      `json:"longUnits,omitempty"`
	LongClientExtensions  *ClientExtensionsDefinition `json:"longClientExtensions,omitempty"`
	ShortUnits            string                      `json:"shortUnits,omitempty"`
	ShortClientExtensions *ClientExtensionsDefinition `json:"shortClientExtensions,omitempty"`
}

type PutPositionsInstrumentCloseParams struct {
	Body *PutPositionsInstrumentCloseBodyParams
}

/* Schemas */

type GetPositionsSchema struct {
	Positions         []*PositionDefinition   `json:"positions,omitempty"`
	LastTransactionID TransactionIDDefinition `json:"lastTransactionID,omitempty"`
}

type GetOpenPositionsSchema struct {
	Positions         []*PositionDefinition   `json:"positions,omitempty"`
	LastTransactionID TransactionIDDefinition `json:"lastTransactionID,omitempty"`
}

type GetPositionsInstrumentSchema struct {
	Position          *PositionDefinition     `json:"position,omitempty"`
	LastTransactionID TransactionIDDefinition `json:"lastTransactionID,omitempty"`
}

type PutPositionsInstrumentCloseSchema struct {
	LongOrderCreateTransaction  *TransactionDefinition    `json:"longOrderCreateTransaction,omitempty"`
	LongOrderFillTransaction    *TransactionDefinition    `json:"longOrderFillTransaction,omitempty"`
	LongOrderCancelTransaction  *TransactionDefinition    `json:"longOrderCancelTransaction,omitempty"`
	ShortOrderCreateTransaction *TransactionDefinition    `json:"shortOrderCreateTransaction,omitempty"`
	ShortOrderFillTransaction   *TransactionDefinition    `json:"shortOrderFillTransaction,omitempty"`
	ShortOrderCancelTransaction *TransactionDefinition    `json:"shortOrderCancelTransaction,omitempty"`
	RelatedTransactionIDs       []TransactionIDDefinition `json:"relatedTransactionIDs,omitempty"`
	LastTransactionID           TransactionIDDefinition   `json:"lastTransactionID,omitempty"`
}

/* Errors */

type PutPositionsInstrumentCloseBadRequestError struct {
	LongOrderRejectTransaction  *TransactionDefinition    `json:"longOrderRejectTransaction,omitempty"`
	ShortOrderRejectTransaction *TransactionDefinition    `json:"shortOrderRejectTransaction,omitempty"`
	RelatedTransactionIDs       []TransactionIDDefinition `json:"relatedTransactionIDs,omitempty"`
	LastTransactionID           TransactionIDDefinition   `json:"lastTransactionID,omitempty"`
	ErrorCode                   string                    `json:"errorCode,omitempty"`
	ErrorMessage                string                    `json:"errorMessage,omitempty"`
}

func (r *PutPositionsInstrumentCloseBadRequestError) Error() string {
	// TODO: エラーを整える
	return r.ErrorMessage
}

type PutPositionsInstrumentCloseNotFoundError struct {
	LongOrderRejectTransaction  *TransactionDefinition    `json:"longOrderRejectTransaction,omitempty"`
	ShortOrderRejectTransaction *TransactionDefinition    `json:"shortOrderRejectTransaction,omitempty"`
	RelatedTransactionIDs       []TransactionIDDefinition `json:"relatedTransactionIDs,omitempty"`
	LastTransactionID           TransactionIDDefinition   `json:"lastTransactionID,omitempty"`
	ErrorCode                   string                    `json:"errorCode,omitempty"`
	ErrorMessage                string                    `json:"errorMessage,omitempty"`
}

func (r *PutPositionsInstrumentCloseNotFoundError) Error() string {
	// TODO: エラーを整える
	return r.ErrorMessage
}

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

// GET /v3/accounts/{accountID}/positions/{instrument}
func (r *ReceiverPositionsInstrument) Get() (*GetPositionsInstrumentSchema, error) {
	resp, err := r.Connection.request(
		&requestParams{
			method:   "Get",
			endPoint: "/v3/accounts/" + r.AccountID + "/positions/" + r.Instrument,
			headers: []header{
				{key: "Accept-Datetime-Format", value: "RFC3339"},
			},
		},
	)
	if err != nil {
		return nil, xerrors.Errorf("Get positions instrument canceled: %w", err)
	}
	defer resp.Body.Close()

	var data interface{}
	switch resp.StatusCode {
	case 200:
		data = new(GetPositionsInstrumentSchema)
	}

	data, err = parseResponse(resp, data, r.Connection.strict)
	if err != nil {
		return nil, xerrors.Errorf("Get positions instrument failed: %w", err)
	}
	return data.(*GetPositionsInstrumentSchema), nil
}

// PUT /v3/accounts/{accountID}/positions/{instrument}/close
func (r *ReceiverPositionsInstrumentClose) Put(params *PutPositionsInstrumentCloseParams) (*PutPositionsInstrumentCloseSchema, error) {
	resp, err := r.Connection.request(
		&requestParams{
			method:   "PUT",
			endPoint: "/v3/accounts/" + r.AccountID + "/positions/" + r.Instrument + "/close",
			headers: []header{
				{key: "Accept-Datetime-Format", value: "RFC3339"},
			},
			body: params.Body,
		},
	)
	if err != nil {
		return nil, xerrors.Errorf("Put positions instrument close canceled: %w", err)
	}
	defer resp.Body.Close()

	var data interface{}
	switch resp.StatusCode {
	case 200:
		data = new(PutPositionsInstrumentCloseSchema)
	case 400:
		data = new(PutPositionsInstrumentCloseBadRequestError)
	case 404:
		data = new(PutPositionsInstrumentCloseNotFoundError)
	}

	data, err = parseResponse(resp, data, r.Connection.strict)
	if err != nil {
		return nil, xerrors.Errorf("Put positions instrument close failed: %w", err)
	}
	return data.(*PutPositionsInstrumentCloseSchema), nil
}
