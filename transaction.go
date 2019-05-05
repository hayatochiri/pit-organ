package pitOrgan

import (
	"encoding/json"
	"golang.org/x/xerrors"
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

// TODO: GET /v3/accounts/{accountID}/transactions/idrange

// TODO: GET /v3/accounts/{accountID}/transactions/sinceid

// TODO: GET /v3/accounts/{accountID}/transactions/stream

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
