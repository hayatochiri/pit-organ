package pitOrgan

import (
	"golang.org/x/xerrors"
	"net/http"
	"strings"
)

/* Receivers */

type ReceiverAccounts struct {
	Connection *Connection
}

func (c *Connection) Accounts() *ReceiverAccounts {
	return &ReceiverAccounts{
		Connection: c,
	}
}

type ReceiverAccountID struct {
	AccountID  string
	Connection *Connection
}

func (r *ReceiverAccounts) AccountID(id string) *ReceiverAccountID {
	return &ReceiverAccountID{
		AccountID:  id,
		Connection: r.Connection,
	}
}

type ReceiverAccountSummary struct {
	AccountID  string
	Connection *Connection
}

func (r *ReceiverAccountID) Summary() *ReceiverAccountSummary {
	return &ReceiverAccountSummary{
		AccountID:  r.AccountID,
		Connection: r.Connection,
	}
}

type ReceiverAccountInstruments struct {
	AccountID  string
	Connection *Connection
}

func (r *ReceiverAccountID) Instruments() *ReceiverAccountInstruments {
	return &ReceiverAccountInstruments{
		AccountID:  r.AccountID,
		Connection: r.Connection,
	}
}

type ReceiverAccountConfiguration struct {
	AccountID  string
	Connection *Connection
}

func (r *ReceiverAccountID) Configuration() *ReceiverAccountConfiguration {
	return &ReceiverAccountConfiguration{
		AccountID:  r.AccountID,
		Connection: r.Connection,
	}
}

type ReceiverAccountChanges struct {
	AccountID  string
	Connection *Connection
}

func (r *ReceiverAccountID) Changes() *ReceiverAccountChanges {
	return &ReceiverAccountChanges{
		AccountID:  r.AccountID,
		Connection: r.Connection,
	}
}

/* Params */

type GetAccountInstrumentsParams struct {
	Instruments []string
}

type PatchAccountConfigurationBodyParams struct {
	Alias      string `json:"alias,omitempty"`
	MarginRate string `json:"marginRate,omitempty"`
}

type PatchAccountConfigurationParams struct {
	Body *PatchAccountConfigurationBodyParams
}

type GetAccountChangesParams struct {
	SinceTransactionID TransactionIDDefinition
}

/* Headers */

type GetAccountsHeaders struct {
	RequestID string
}

func (s *GetAccountsSchema) setHeaders(resp *http.Response) error {
	s.Headers = new(GetAccountsHeaders)
	if h, err := copyHeader(resp, "Requestid"); err == nil {
		s.Headers.RequestID = h[0]
	} else {
		return xerrors.Errorf("Parse headers failed: %w", err)
	}
	return nil
}

type GetAccountIDHeaders struct {
	RequestID string
}

func (s *GetAccountIDSchema) setHeaders(resp *http.Response) error {
	s.Headers = new(GetAccountIDHeaders)
	if h, err := copyHeader(resp, "Requestid"); err == nil {
		s.Headers.RequestID = h[0]
	} else {
		return xerrors.Errorf("Parse headers failed: %w", err)
	}
	return nil
}

type GetAccountSummaryHeaders struct {
	RequestID string
}

func (s *GetAccountSummarySchema) setHeaders(resp *http.Response) error {
	s.Headers = new(GetAccountSummaryHeaders)
	if h, err := copyHeader(resp, "Requestid"); err == nil {
		s.Headers.RequestID = h[0]
	} else {
		return xerrors.Errorf("Parse headers failed: %w", err)
	}
	return nil
}

type GetAccountInstrumentsHeaders struct {
	RequestID string
}

func (s *GetAccountInstrumentsSchema) setHeaders(resp *http.Response) error {
	s.Headers = new(GetAccountInstrumentsHeaders)
	if h, err := copyHeader(resp, "Requestid"); err == nil {
		s.Headers.RequestID = h[0]
	} else {
		return xerrors.Errorf("Parse headers failed: %w", err)
	}
	return nil
}

type PatchAccountConfigurationHeaders struct {
	RequestID string
}

func (s *PatchAccountConfigurationSchema) setHeaders(resp *http.Response) error {
	s.Headers = new(PatchAccountConfigurationHeaders)
	if h, err := copyHeader(resp, "Requestid"); err == nil {
		s.Headers.RequestID = h[0]
	} else {
		return xerrors.Errorf("Parse headers failed: %w", err)
	}
	return nil
}

type GetAccountChangesHeaders struct {
	RequestID string
}

func (s *GetAccountChangesSchema) setHeaders(resp *http.Response) error {
	s.Headers = new(GetAccountChangesHeaders)
	if h, err := copyHeader(resp, "Requestid"); err == nil {
		s.Headers.RequestID = h[0]
	} else {
		return xerrors.Errorf("Parse headers failed: %w", err)
	}
	return nil
}

/* Schemas */

type GetAccountsSchema struct {
	Headers  *GetAccountsHeaders
	Accounts []*AccountPropertiesDefinition `json:"accounts,omitempty"`
}

type GetAccountIDSchema struct {
	Headers           *GetAccountIDHeaders
	Account           *AccountDefinition      `json:"account,omitempty"`
	LastTransactionID TransactionIDDefinition `json:"lastTransactionID,omitempty"`
}

type GetAccountSummarySchema struct {
	Headers           *GetAccountSummaryHeaders
	Account           *AccountSummaryDefinition `json:"account,omitempty"`
	LastTransactionID TransactionIDDefinition   `json:"lastTransactionID,omitempty"`
}

type GetAccountInstrumentsSchema struct {
	Headers           *GetAccountInstrumentsHeaders
	Instruments       []*InstrumentDefinition `json:"instruments,omitempty"`
	LastTransactionID TransactionIDDefinition `json:"lastTransactionID,omitempty"`
}

type PatchAccountConfigurationSchema struct {
	Headers                    *PatchAccountConfigurationHeaders
	ClientConfigureTransaction *TransactionDefinition  `json:"clientConfigureTransaction,omitempty"`
	LastTransactionID          TransactionIDDefinition `json:"lastTransactionID,omitempty"`
}

type GetAccountChangesSchema struct {
	Headers           *GetAccountChangesHeaders
	Changes           *AccountChangesDefinition      `json:"changes,omitempty"`
	State             *AccountChangesStateDefinition `json:"state,omitempty"`
	LastTransactionID TransactionIDDefinition        `json:"lastTransactionID,omitempty"`
}

/* Errors */

type PatchAccountConfigurationBadRequestError struct {
	ClientConfigureRejectTransaction *TransactionDefinition  `json:"clientConfigureRejectTransaction,omitempty"`
	LastTransactionID                TransactionIDDefinition `json:"lastTransactionID,omitempty"`
	ErrorCode                        string                  `json:"errorCode,omitempty"`
	ErrorMessage                     string                  `json:"errorMessage,omitempty"`
}

func (r *PatchAccountConfigurationBadRequestError) Error() string {
	// TODO: エラーを整える
	return r.ErrorMessage
}

type PatchAccountConfigurationForbiddenError struct {
	ClientConfigureRejectTransaction *TransactionDefinition  `json:"clientConfigureRejectTransaction,omitempty"`
	LastTransactionID                TransactionIDDefinition `json:"lastTransactionID,omitempty"`
	ErrorCode                        string                  `json:"errorCode,omitempty"`
	ErrorMessage                     string                  `json:"errorMessage,omitempty"`
}

func (r *PatchAccountConfigurationForbiddenError) Error() string {
	// TODO: エラーを整える
	return r.ErrorMessage
}

/* API */

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

	data, err = parseResponse(resp, data, r.Connection.strict)
	if err != nil {
		return nil, xerrors.Errorf("Get accounts failed: %w", err)
	}
	return data.(*GetAccountsSchema), nil
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

	data, err = parseResponse(resp, data, r.Connection.strict)
	if err != nil {
		return nil, xerrors.Errorf("Get account ID failed: %w", err)
	}
	return data.(*GetAccountIDSchema), nil
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

	data, err = parseResponse(resp, data, r.Connection.strict)
	if err != nil {
		return nil, xerrors.Errorf("Get account summary failed: %w", err)
	}
	return data.(*GetAccountSummarySchema), nil
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

	data, err = parseResponse(resp, data, r.Connection.strict)
	if err != nil {
		return nil, xerrors.Errorf("Get account instruments failed: %w", err)
	}
	return data.(*GetAccountInstrumentsSchema), nil
}

// PATCH /v3/accounts/{accountID}/configuration
func (r *ReceiverAccountConfiguration) Patch(params *PatchAccountConfigurationParams) (*PatchAccountConfigurationSchema, error) {
	resp, err := r.Connection.request(
		&requestParams{
			method:   "PATCH",
			endPoint: "/v3/accounts/" + r.AccountID + "/configuration",
			headers: []header{
				{key: "Accept-Datetime-Format", value: "RFC3339"},
			},
			body: params.Body,
		},
	)
	if err != nil {
		return nil, xerrors.Errorf("Patch account configuration canceled: %w", err)
	}
	defer resp.Body.Close()

	var data interface{}
	switch resp.StatusCode {
	case 200:
		data = new(PatchAccountConfigurationSchema)
	case 400:
		data = new(PatchAccountConfigurationBadRequestError)
	case 403:
		data = new(PatchAccountConfigurationForbiddenError)
	}

	data, err = parseResponse(resp, data, r.Connection.strict)
	if err != nil {
		return nil, xerrors.Errorf("Patch account configuration failed: %w", err)
	}
	return data.(*PatchAccountConfigurationSchema), nil
}

// GET /v3/accounts/{accountID}/changes
func (r *ReceiverAccountChanges) Get(params *GetAccountChangesParams) (*GetAccountChangesSchema, error) {
	resp, err := r.Connection.request(
		&requestParams{
			method:   "GET",
			endPoint: "/v3/accounts/" + r.AccountID + "/changes",
			headers: []header{
				{key: "Accept-Datetime-Format", value: "RFC3339"},
			},
			queries: func() []query {
				q := make([]query, 0, 1)

				// sinceTransactionID
				q = append(q, query{key: "sinceTransactionID", value: string(params.SinceTransactionID)})

				return q
			}(),
		},
	)
	if err != nil {
		return nil, xerrors.Errorf("Get account changes canceled: %w", err)
	}
	defer resp.Body.Close()

	var data interface{}
	switch resp.StatusCode {
	case 200:
		data = new(GetAccountChangesSchema)
	}

	data, err = parseResponse(resp, data, r.Connection.strict)
	if err != nil {
		return nil, xerrors.Errorf("Get account changes failed: %w", err)
	}
	return data.(*GetAccountChangesSchema), nil
}
