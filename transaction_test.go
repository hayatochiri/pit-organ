package pitOrgan

import (
	"github.com/davecgh/go-spew/spew"
	"testing"
)

func Test_Transactions(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		connection := newConnection(t, OandaPractice)
		accountID := Getenv("ACCOUNT_ID")
		params := &GetTransactionsParams{
			PageSize: 1000,
		}
		data, err := connection.Accounts().AccountID(accountID).Transactions().Get(params)
		if err != nil {
			t.Fatalf("Get transactions failed.\n%+v", err)
		}

		t.Logf("Response:\n%s", spew.Sdump(data))
	})
}

func Test_TransactionsIdrange(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		connection := newConnection(t, OandaPractice)
		accountID := Getenv("ACCOUNT_ID")
		params := &GetTransactionsIdrangeParams{
			From: 1,
			To:   165,
		}
		data, err := connection.Accounts().AccountID(accountID).Transactions().Idrange().Get(params)
		if err != nil {
			t.Fatalf("Get transactions idrange failed.\n%+v", err)
		}

		for _, transaction := range data.Transactions {
			switch tr := transaction.(type) {
			case *CreateTransactionDefinition:
			case *CloseTransactionDefinition:
			case *ReopenTransactionDefinition:
			case *ClientConfigureTransactionDefinition:
			case *ClientConfigureRejectTransactionDefinition:
			case *TransferFundsTransactionDefinition:
			case *TransferFundsRejectTransactionDefinition:
			case *MarketOrderTransactionDefinition:
			case *MarketOrderRejectTransactionDefinition:
			case *FixedPriceOrderTransactionDefinition:
			case *LimitOrderTransactionDefinition:
			case *LimitOrderRejectTransactionDefinition:
			case *StopOrderTransactionDefinition:
			case *StopOrderRejectTransactionDefinition:
			case *MarketIfTouchedOrderTransactionDefinition:
			case *MarketIfTouchedOrderRejectTransactionDefinition:
			case *TakeProfitOrderTransactionDefinition:
			case *TakeProfitOrderRejectTransactionDefinition:
			case *StopLossOrderTransactionDefinition:
			case *StopLossOrderRejectTransactionDefinition:
			case *TrailingStopLossOrderTransactionDefinition:
			case *TrailingStopLossOrderRejectTransactionDefinition:
			case *OrderFillTransactionDefinition:
			case *OrderCancelTransactionDefinition:
			case *OrderCancelRejectTransactionDefinition:
			case *OrderClientExtensionsModifyTransactionDefinition:
			case *OrderClientExtensionsModifyRejectTransactionDefinition:
			case *TradeClientExtensionsModifyTransactionDefinition:
			case *TradeClientExtensionsModifyRejectTransactionDefinition:
			case *MarginCallEnterTransactionDefinition:
			case *MarginCallExtendTransactionDefinition:
			case *MarginCallExitTransactionDefinition:
			case *DelayedTradeClosureTransactionDefinition:
			case *DailyFinancingTransactionDefinition:
			case *ResetResettablePLTransactionDefinition:
			default:
				t.Errorf("Unexpected transaction type detected.\n%s", spew.Sdump(tr))
			}
		}

		t.Logf("Response:\n%s", spew.Sdump(data))
	})
}

func Test_IdrangeParams(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		connection := newConnection(t, OandaPractice)
		accountID := Getenv("ACCOUNT_ID")
		transactionAPI := connection.Accounts().AccountID(accountID).Transactions()
		data, err := transactionAPI.Get(
			&GetTransactionsParams{Type: []TransactionFilterDefinition{
				OrderTransaction,
				CloseTransaction,
			}},
		)
		if err != nil {
			t.Fatalf("Get transactions failed.\n%+v", err)
		}

		params, err := data.IdrangeParams()
		if err != nil {
			t.Fatalf("Get transactions idrange params failed.\n%+v", err)
		}

		for _, param := range params {
			_, err := transactionAPI.Idrange().Get(param)
			if err != nil {
				t.Errorf("Get transactions idrange failed.\nParam:\n%s\nError:\n%+v", spew.Sdump(param), err)
			}
		}
	})
}

}
