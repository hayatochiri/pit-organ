package pitOrgan

import (
	"encoding/json"
	"golang.org/x/xerrors"
	"io/ioutil"
	"net/url"
	"strconv"
	"strings"
	"time"
)

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

// Params

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

// Schemas

type GetTransactionsSchema struct {
	From              DateTimeDefinition            `json:"from"`
	To                DateTimeDefinition            `json:"to"`
	PageSize          int                           `json:"pageSize"`
	Type              []TransactionFilterDefinition `json:"type"`
	Count             int                           `json:"count"`
	Pages             []string                      `json:"pages"`
	LastTransactionID TransactionIDDefinition       `json:"lastTransactionID"`
}

type GetTransactionsIdrangeSchema struct {
	Transactions      []interface{}           `json:"transactions"`
	LastTransactionID TransactionIDDefinition `json:"lastTransactionID"`
}

type getTransactionsIdrangeParser struct {
	Transactions []TransactionDefinition `json:"transactions"`
}
type getTransactionsIdrangeRawMessage struct {
	Message []json.RawMessage `json:"transactions"`
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

// GET /v3/accounts/{accountID}/transactions
//
// Get a list of Transactions pages that satisfy a time-based Transaction query.
func (r *ReceiverTransactions) Get(params *GetTransactionsParams) (*GetTransactionsSchema, error) {
	resp, err := r.Connection.request(
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

	data, err = parseResponse(resp, data)
	if err != nil {
		return nil, xerrors.Errorf("Get transactions failed: %w", err)
	}
	return data.(*GetTransactionsSchema), nil
}

// TODO: GET /v3/accounts/{accountID}/transactions/{transactionID}

// GET /v3/accounts/{accountID}/transactions/idrange
//
// Get a range of Transactions for an Account based on the Transaction IDs.
func (r *ReceiverTransactionsIdrange) Get(params *GetTransactionsIdrangeParams) (*GetTransactionsIdrangeSchema, error) {
	resp, err := r.Connection.request(
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

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, xerrors.Errorf("Read response body failed: %w", err)
	}

	var data interface{}
	switch resp.StatusCode {
	case 200:
		data = new(GetTransactionsIdrangeSchema)
	}
	data, err = parseBody(body, resp.StatusCode, data)
	// HTTPステータスコードが200以外ならエラーとしてreturnされる
	if err != nil {
		return nil, xerrors.Errorf("Get transactions idrange failed: %w", err)
	}

	var parser = new(getTransactionsIdrangeParser)
	if err := json.Unmarshal(body, parser); err != nil {
		return nil, xerrors.Errorf("Unmarshal response body failed: %w", err)
	}

	var rawMessage = new(getTransactionsIdrangeRawMessage)
	if err := json.Unmarshal(body, rawMessage); err != nil {
		return nil, xerrors.Errorf("Unmarshal response body failed: %w", err)
	}

	data.(*GetTransactionsIdrangeSchema).Transactions = make([]interface{}, len(parser.Transactions), len(parser.Transactions))
	for n, transaction := range parser.Transactions {
		var tData interface{}
		if tData, err = unmarshalTransaction(rawMessage.Message[n], transaction.Type); err != nil {
			return nil, xerrors.Errorf("Unmarshal response body failed(type=%s): %w", transaction.Type, err)
		}
		data.(*GetTransactionsIdrangeSchema).Transactions[n] = tData
	}

	return data.(*GetTransactionsIdrangeSchema), nil
}

// TODO: GET /v3/accounts/{accountID}/transactions/sinceid

// TODO: GET /v3/accounts/{accountID}/transactions/stream

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

func unmarshalTransaction(data []byte, tType TransactionTypeDefinition) (interface{}, error) {
	var v interface{}

	switch tType {
	case "CREATE":
		v = new(CreateTransactionDefinition)
	case "CLOSE":
		v = new(CloseTransactionDefinition)
	case "REOPEN":
		v = new(ReopenTransactionDefinition)
	case "CLIENT_CONFIGURE":
		v = new(ClientConfigureTransactionDefinition)
	case "CLIENT_CONFIGURE_REJECT":
		v = new(ClientConfigureRejectTransactionDefinition)
	case "TRANSFER_FUNDS":
		v = new(TransferFundsTransactionDefinition)
	case "TRANSFER_FUNDS_REJECT":
		v = new(TransferFundsRejectTransactionDefinition)
	case "MARKET_ORDER":
		v = new(MarketOrderTransactionDefinition)
	case "MARKET_ORDER_REJECT":
		v = new(MarketOrderRejectTransactionDefinition)
	case "FIXED_PRICE_ORDER":
		v = new(FixedPriceOrderTransactionDefinition)
	case "LIMIT_ORDER":
		v = new(LimitOrderTransactionDefinition)
	case "LIMIT_ORDER_REJECT":
		v = new(LimitOrderRejectTransactionDefinition)
	case "STOP_ORDER":
		v = new(StopOrderTransactionDefinition)
	case "STOP_ORDER_REJECT":
		v = new(StopOrderRejectTransactionDefinition)
	case "MARKET_IF_TOUCHED_ORDER":
		v = new(MarketIfTouchedOrderTransactionDefinition)
	case "MARKET_IF_TOUCHED_ORDER_REJECT":
		v = new(MarketIfTouchedOrderRejectTransactionDefinition)
	case "TAKE_PROFIT_ORDER":
		v = new(TakeProfitOrderTransactionDefinition)
	case "TAKE_PROFIT_ORDER_REJECT":
		v = new(TakeProfitOrderRejectTransactionDefinition)
	case "STOP_LOSS_ORDER":
		v = new(StopLossOrderTransactionDefinition)
	case "STOP_LOSS_ORDER_REJECT":
		v = new(StopLossOrderRejectTransactionDefinition)
	case "TRAILING_STOP_LOSS_ORDER":
		v = new(TrailingStopLossOrderTransactionDefinition)
	case "TRAILING_STOP_LOSS_ORDER_REJECT":
		v = new(TrailingStopLossOrderRejectTransactionDefinition)
	case "ORDER_FILL":
		v = new(OrderFillTransactionDefinition)
	case "ORDER_CANCEL":
		v = new(OrderCancelTransactionDefinition)
	case "ORDER_CANCEL_REJECT":
		v = new(OrderCancelRejectTransactionDefinition)
	case "ORDER_CLIENT_EXTENSIONS_MODIFY":
		v = new(OrderClientExtensionsModifyTransactionDefinition)
	case "ORDER_CLIENT_EXTENSIONS_MODIFY_REJECT":
		v = new(OrderClientExtensionsModifyRejectTransactionDefinition)
	case "TRADE_CLIENT_EXTENSIONS_MODIFY":
		v = new(TradeClientExtensionsModifyTransactionDefinition)
	case "TRADE_CLIENT_EXTENSIONS_MODIFY_REJECT":
		v = new(TradeClientExtensionsModifyRejectTransactionDefinition)
	case "MARGIN_CALL_ENTER":
		v = new(MarginCallEnterTransactionDefinition)
	case "MARGIN_CALL_EXTEND":
		v = new(MarginCallExtendTransactionDefinition)
	case "MARGIN_CALL_EXIT":
		v = new(MarginCallExitTransactionDefinition)
	case "DELAYED_TRADE_CLOSURE":
		v = new(DelayedTradeClosureTransactionDefinition)
	case "DAILY_FINANCING":
		v = new(DailyFinancingTransactionDefinition)
	case "RESET_RESETTABLE_PL":
		v = new(ResetResettablePLTransactionDefinition)
	}

	if err := json.Unmarshal(data, v); err != nil {
		return nil, xerrors.Errorf("Unmarshal transaction failed(type=%s): %w", tType, err)
	}

	return v, nil
}
