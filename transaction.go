package pitOrgan

import (
	"encoding/json"
	"golang.org/x/xerrors"
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
