package pitOrgan

import (
	"bufio"
	"context"
	"encoding/json"
	"golang.org/x/xerrors"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"
)

/* Receivers */

type ReceiverTransactions struct {
	AccountID  string
	Connection *Connection
}

func (r *ReceiverAccountID) Transactions() *ReceiverTransactions {
	return &ReceiverTransactions{
		AccountID:  r.AccountID,
		Connection: r.Connection,
	}
}

type ReceiverTransactionID struct {
	AccountID     string
	Connection    *Connection
	TransactionID string
}

func (r *ReceiverTransactions) TransactionID(transactionID string) *ReceiverTransactionID {
	return &ReceiverTransactionID{
		AccountID:     r.AccountID,
		Connection:    r.Connection,
		TransactionID: transactionID,
	}
}

type ReceiverTransactionsIdrange struct {
	AccountID  string
	Connection *Connection
}

func (r *ReceiverTransactions) Idrange() *ReceiverTransactionsIdrange {
	return &ReceiverTransactionsIdrange{
		AccountID:  r.AccountID,
		Connection: r.Connection,
	}
}

type ReceiverTransactionsSinceID struct {
	AccountID  string
	Connection *Connection
}

func (r *ReceiverTransactions) SinceID() *ReceiverTransactionsSinceID {
	return &ReceiverTransactionsSinceID{
		AccountID:  r.AccountID,
		Connection: r.Connection,
	}
}

type ReceiverTransactionsStream struct {
	AccountID  string
	Connection *Connection
}

func (r *ReceiverTransactions) Stream() *ReceiverTransactionsStream {
	return &ReceiverTransactionsStream{
		AccountID:  r.AccountID,
		Connection: r.Connection,
	}
}

/* Params */

type GetTransactionsParams struct {
	From     time.Time
	To       time.Time
	PageSize int
	Type     []TransactionFilterDefinition
}

type GetTransactionsIdrangeParams struct {
	From int
	To   int
	Type []TransactionFilterDefinition
}

type GetTransactionsSinceIDParams struct {
	ID string
}

type GetTransactionsStreamParams struct {
	BufferSize int
}

/* Schemas */

type GetTransactionsSchema struct {
	From              DateTimeDefinition            `json:"from,omitempty"`
	To                DateTimeDefinition            `json:"to,omitempty"`
	PageSize          int                           `json:"pageSize,omitempty"`
	Type              []TransactionFilterDefinition `json:"type,omitempty"`
	Count             int                           `json:"count,omitempty"`
	Pages             []string                      `json:"pages,omitempty"`
	LastTransactionID TransactionIDDefinition       `json:"lastTransactionID,omitempty"`
}

type GetTransactionsIdrangeSchema struct {
	Transactions      []*TransactionDefinition `json:"transactions,omitempty"`
	LastTransactionID TransactionIDDefinition  `json:"lastTransactionID,omitempty"`
}

type GetTransactionIDSchema struct {
	Transaction       *TransactionDefinition  `json:"transaction,omitempty"`
	LastTransactionID TransactionIDDefinition `json:"lastTransactionID,omitempty"`
}

type GetTransactionsSinceIDSchema struct {
	Transactions      []*TransactionDefinition `json:"transactions,omitempty"`
	LastTransactionID TransactionIDDefinition  `json:"lastTransactionID,omitempty"`
}

/* Streams */

type TransactionsChannels struct {
	TransactionCh <-chan *TransactionDefinition
	lastError     error
	errorCh       <-chan error
	close         context.CancelFunc
	closeWait     *sync.WaitGroup
}

/* API */

// GET /v3/accounts/{accountID}/transactions
//
// Get a list of Transactions pages that satisfy a time-based Transaction query.
func (r *ReceiverTransactions) Get(ctx context.Context, params *GetTransactionsParams) (*GetTransactionsSchema, error) {
	resp, err := r.Connection.request(
		ctx,
		&requestParams{
			method:   "GET",
			endPoint: "/v3/accounts/" + r.AccountID + "/transactions",
			headers: []header{
				{key: "Accept-Datetime-Format", value: "RFC3339"},
			},
			queries: func() []query {
				q := make([]query, 0, 4)

				// from
				if !params.From.IsZero() {
					q = append(q, query{key: "from", value: params.From.Format(time.RFC3339Nano)})
				}

				// to
				if !params.To.IsZero() {
					q = append(q, query{key: "to", value: params.To.Format(time.RFC3339Nano)})
				}

				// pageSize
				if params.PageSize != 0 {
					q = append(q, query{key: "pageSize", value: strconv.Itoa(params.PageSize)})
				}

				// type
				if params.Type != nil {
					types := make([]string, len(params.Type))
					for n, t := range params.Type {
						types[n] = string(t)
					}
					q = append(q, query{key: "type", value: strings.Join(types, ",")})
				}

				return q
			}(),
		},
	)
	if err != nil {
		return nil, xerrors.Errorf("Get transactions canceled: %w", err)
	}
	defer resp.Body.Close()

	var data interface{}
	switch resp.StatusCode {
	case 200:
		data = new(GetTransactionsSchema)
	}

	data, err = parseResponse(resp, data, r.Connection.strict)
	if err != nil {
		return nil, xerrors.Errorf("Get transactions failed: %w", err)
	}
	return data.(*GetTransactionsSchema), nil
}

// GET /v3/accounts/{accountID}/transactions/{transactionID}
func (r *ReceiverTransactionID) Get(ctx context.Context) (*GetTransactionIDSchema, error) {
	resp, err := r.Connection.request(
		ctx,
		&requestParams{
			method:   "GET",
			endPoint: "/v3/accounts/" + r.AccountID + "/transactions/" + r.TransactionID,
			headers: []header{
				{key: "Accept-Datetime-Format", value: "RFC3339"},
			},
		},
	)
	if err != nil {
		return nil, xerrors.Errorf("Get transactions id canceled: %w", err)
	}
	defer resp.Body.Close()

	var data interface{}
	switch resp.StatusCode {
	case 200:
		data = new(GetTransactionIDSchema)
	}

	data, err = parseResponse(resp, data, r.Connection.strict)
	if err != nil {
		return nil, xerrors.Errorf("Get transactions id failed: %w", err)
	}
	return data.(*GetTransactionIDSchema), nil
}

// GET /v3/accounts/{accountID}/transactions/idrange
//
// Get a range of Transactions for an Account based on the Transaction IDs.
func (r *ReceiverTransactionsIdrange) Get(ctx context.Context, params *GetTransactionsIdrangeParams) (*GetTransactionsIdrangeSchema, error) {
	resp, err := r.Connection.request(
		ctx,
		&requestParams{
			method:   "GET",
			endPoint: "/v3/accounts/" + r.AccountID + "/transactions/idrange",
			headers: []header{
				{key: "Accept-Datetime-Format", value: "RFC3339"},
			},
			queries: func() []query {
				q := make([]query, 0, 3)

				// from
				q = append(q, query{key: "from", value: strconv.Itoa(params.From)})

				// to
				q = append(q, query{key: "to", value: strconv.Itoa(params.To)})

				// type
				if params.Type != nil {
					types := make([]string, len(params.Type))
					for n, t := range params.Type {
						types[n] = string(t)
					}
					q = append(q, query{key: "type", value: strings.Join(types, ",")})
				}

				return q
			}(),
		},
	)
	if err != nil {
		return nil, xerrors.Errorf("Get transactions idrange canceled: %w", err)
	}
	defer resp.Body.Close()

	var data interface{}
	switch resp.StatusCode {
	case 200:
		data = new(GetTransactionsIdrangeSchema)
	}
	data, err = parseResponse(resp, data, r.Connection.strict)
	if err != nil {
		return nil, xerrors.Errorf("Get transactions idrange failed: %w", err)
	}

	return data.(*GetTransactionsIdrangeSchema), nil
}

// GET /v3/accounts/{accountID}/transactions/sinceid
func (r *ReceiverTransactionsSinceID) Get(ctx context.Context, params *GetTransactionsSinceIDParams) (*GetTransactionsSinceIDSchema, error) {
	resp, err := r.Connection.request(
		ctx,
		&requestParams{
			method:   "GET",
			endPoint: "/v3/accounts/" + r.AccountID + "/transactions/sinceid",
			headers: []header{
				{key: "Accept-Datetime-Format", value: "RFC3339"},
			},
			queries: func() []query {
				q := make([]query, 0, 1)

				// id
				if params.ID != "" {
					q = append(q, query{key: "id", value: params.ID})
				}

				return q
			}(),
		},
	)
	if err != nil {
		return nil, xerrors.Errorf("Get transactions sinceid canceled: %w", err)
	}
	defer resp.Body.Close()

	var data interface{}
	switch resp.StatusCode {
	case 200:
		data = new(GetTransactionsSinceIDSchema)
	}

	data, err = parseResponse(resp, data, r.Connection.strict)
	if err != nil {
		return nil, xerrors.Errorf("Get transactions sinceid failed: %w", err)
	}
	return data.(*GetTransactionsSinceIDSchema), nil
}

// GET /v3/accounts/{accountID}/transactions/stream
//
// Get a stream of Transactions for an Account starting from when the request is made.
func (r *ReceiverTransactionsStream) Get(ctx context.Context, params *GetTransactionsStreamParams) (*TransactionsChannels, error) {
	childCtx, cancel := context.WithCancel(ctx)

	resp, err := r.Connection.stream(
		childCtx,
		&requestParams{
			method:   "GET",
			endPoint: "/v3/accounts/" + r.AccountID + "/transactions/stream",
			headers: []header{
				{key: "Accept-Datetime-Format", value: "RFC3339"},
			},
		},
	)
	if err != nil {
		return nil, xerrors.Errorf("Get transactions stream canceled: %w", err)
	}

	if resp.StatusCode != 200 {
		defer func() {
			resp.Body.Close()
			cancel()
		}()
		var err interface{}
		_, err = parseResponse(resp, err, r.Connection.strict)
		return nil, xerrors.Errorf("Get transactions stream failed: %w", err)
	}

	closeWait := new(sync.WaitGroup)

	transactionCh := make(chan *TransactionDefinition, params.BufferSize)
	errorCh := make(chan error, 3)

	// closeChがcloseされたらstreamを終了するgoroutine
	// 下記の readre.ReadBytes はデータを受信しないと処理が進まないため
	// streamの途中で途切れると永遠に待ち続けてしまうので終了用goroutineを用意。
	go func() {
		defer func() {
			resp.Body.Close()
			cancel()
			closeWait.Done()
		}()
		closeWait.Add(1)
		<-childCtx.Done()
	}()

	// 受信したデータ(JSON)を構造体にしてreaderCh channelに送信するgoroutine
	readerCh := make(chan []byte, params.BufferSize)
	go func() {
		defer func() {
			close(readerCh)
			cancel()
			closeWait.Done()
		}()
		closeWait.Add(1)

		reader := bufio.NewReader(resp.Body)
		for {
			line, err := reader.ReadBytes('\n')
			if err != nil {
				select {
				case <-childCtx.Done():
				default:
					errorCh <- xerrors.Errorf("Read response stream failed: %w", err)
				}
				return
			}
			readerCh <- line
		}
	}()

	// readeerCh channleから受信した構造体をユーザーに渡すgoroutine
	// heartbeat(5秒間隔)が途切れた場合を検知するためにタイムアウトも管理する。
	go func() {
		defer func() {
			close(transactionCh)
			cancel()
			closeWait.Done()
		}()
		closeWait.Add(1)

		timeout := time.NewTimer(0)
		received := true

		for {
			select {
			case <-childCtx.Done():
				return
			case line := <-readerCh:
				received = true

				data := new(TransactionDefinition)
				if err := json.Unmarshal(line, data); err != nil {
					errorCh <- xerrors.Errorf("Unmarshal response stream failed: %w", err)
					return
				}

				if data.Type == "HEARTBEAT" {
					continue
				}

				transactionCh <- data
			case <-timeout.C:
				timeout.Reset(r.Connection.Timeout)
				if !received {
					var err error = &StreamHeartbeatBroken{ErrorMessage: "Heartbeat was broken"}
					errorCh <- xerrors.Errorf("Get pricing stream heartbeat was broken: %w", err)
					return
				}
				received = false
			}
		}
	}()

	return &TransactionsChannels{
		TransactionCh: transactionCh,
		lastError:     nil,
		errorCh:       errorCh,
		close:         cancel,
		closeWait:     closeWait,
	}, nil
}

/* Utils */

func (ch *TransactionsChannels) Close() {
	ch.close()
	ch.closeWait.Wait()
}

func (ch *TransactionsChannels) Err() error {
	if ch.lastError == nil {
		select {
		case ch.lastError = <-ch.errorCh:
		default:
		}
	}
	return ch.lastError
}

func (s *GetTransactionsSchema) IdrangeParams() ([]*GetTransactionsIdrangeParams, error) {
	params := make([]*GetTransactionsIdrangeParams, 0, len(s.Pages))

	for _, page := range s.Pages {
		u, err := url.Parse(page)
		if err != nil {
			return nil, xerrors.Errorf("Parse transactions idrange params failed: %w", err)
		}
		query := u.Query()

		param := new(GetTransactionsIdrangeParams)
		if v, ok := query["from"]; ok {
			param.From, err = strconv.Atoi(v[0])
			if err != nil {
				return nil, xerrors.Errorf("Parse transactions idrange from query failed: %w", err)
			}
		}
		if v, ok := query["to"]; ok {
			param.To, err = strconv.Atoi(v[0])
			if err != nil {
				return nil, xerrors.Errorf("Parse transactions idrange to query failed: %w", err)
			}
		}
		if v, ok := query["type"]; ok {
			filters := strings.Split(v[0], ",")
			param.Type = make([]TransactionFilterDefinition, 0, len(filters))
			for _, f := range filters {
				param.Type = append(param.Type, TransactionFilterDefinition(f))
			}
		}

		params = append(params, param)
	}

	return params, nil
}
