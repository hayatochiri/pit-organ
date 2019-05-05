package pitOrgan_test

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/hayatochiri/pit-organ"
	"log"
	"time"
)

func ExampleReceiverTransactions_Get() {
	connection := &pitOrgan.Connection{
		Token:       "OANDA_API_TOKEN",
		Environemnt: pitOrgan.OandaPractice,
		Timeout:     time.Second * 10,
	}

	params := &pitOrgan.GetTransactionsParams{
		PageSize: 1000,
	}
	data, err := connection.Accounts().AccountID("AccountID").Transactions().Get(params)
	if err != nil {
		log.Fatalf("%+v", err)
	}

	spew.Dump(data)
}

func ExampleReceiverTransactionsIdrange_Get() {
	connection := &pitOrgan.Connection{
		Token:       "OANDA_API_TOKEN",
		Environemnt: pitOrgan.OandaPractice,
		Timeout:     time.Second * 10,
	}

	params := &pitOrgan.GetTransactionsIdrangeParams{
		From: 1,
		To:   165,
	}
	data, err := connection.Accounts().AccountID("AccountID").Transactions().Idrange().Get(params)
	if err != nil {
		log.Fatalf("%+v", err)
	}

	for _, transaction := range data.Transactions {
		switch t := transaction.(type) {
		case *pitOrgan.CreateTransactionDefinition:
			fmt.Printf("Transaction type = %s\n", t.Type)
		case *pitOrgan.CloseTransactionDefinition:
			fmt.Printf("Transaction type = %s\n", t.Type)
		case *pitOrgan.ReopenTransactionDefinition:
			fmt.Printf("Transaction type = %s\n", t.Type)
		case *pitOrgan.ClientConfigureTransactionDefinition:
			fmt.Printf("Transaction type = %s\n", t.Type)
		case *pitOrgan.ClientConfigureRejectTransactionDefinition:
			fmt.Printf("Transaction type = %s\n", t.Type)
		case *pitOrgan.TransferFundsTransactionDefinition:
			fmt.Printf("Transaction type = %s\n", t.Type)
		case *pitOrgan.TransferFundsRejectTransactionDefinition:
			fmt.Printf("Transaction type = %s\n", t.Type)
		case *pitOrgan.MarketOrderTransactionDefinition:
			fmt.Printf("Transaction type = %s\n", t.Type)
		case *pitOrgan.MarketOrderRejectTransactionDefinition:
			fmt.Printf("Transaction type = %s\n", t.Type)
		case *pitOrgan.FixedPriceOrderTransactionDefinition:
			fmt.Printf("Transaction type = %s\n", t.Type)
		case *pitOrgan.LimitOrderTransactionDefinition:
			fmt.Printf("Transaction type = %s\n", t.Type)
		case *pitOrgan.LimitOrderRejectTransactionDefinition:
			fmt.Printf("Transaction type = %s\n", t.Type)
		case *pitOrgan.StopOrderTransactionDefinition:
			fmt.Printf("Transaction type = %s\n", t.Type)
		case *pitOrgan.StopOrderRejectTransactionDefinition:
			fmt.Printf("Transaction type = %s\n", t.Type)
		case *pitOrgan.MarketIfTouchedOrderTransactionDefinition:
			fmt.Printf("Transaction type = %s\n", t.Type)
		case *pitOrgan.MarketIfTouchedOrderRejectTransactionDefinition:
			fmt.Printf("Transaction type = %s\n", t.Type)
		case *pitOrgan.TakeProfitOrderTransactionDefinition:
			fmt.Printf("Transaction type = %s\n", t.Type)
		case *pitOrgan.TakeProfitOrderRejectTransactionDefinition:
			fmt.Printf("Transaction type = %s\n", t.Type)
		case *pitOrgan.StopLossOrderTransactionDefinition:
			fmt.Printf("Transaction type = %s\n", t.Type)
		case *pitOrgan.StopLossOrderRejectTransactionDefinition:
			fmt.Printf("Transaction type = %s\n", t.Type)
		case *pitOrgan.TrailingStopLossOrderTransactionDefinition:
			fmt.Printf("Transaction type = %s\n", t.Type)
		case *pitOrgan.TrailingStopLossOrderRejectTransactionDefinition:
			fmt.Printf("Transaction type = %s\n", t.Type)
		case *pitOrgan.OrderFillTransactionDefinition:
			fmt.Printf("Transaction type = %s\n", t.Type)
		case *pitOrgan.OrderCancelTransactionDefinition:
			fmt.Printf("Transaction type = %s\n", t.Type)
		case *pitOrgan.OrderCancelRejectTransactionDefinition:
			fmt.Printf("Transaction type = %s\n", t.Type)
		case *pitOrgan.OrderClientExtensionsModifyTransactionDefinition:
			fmt.Printf("Transaction type = %s\n", t.Type)
		case *pitOrgan.OrderClientExtensionsModifyRejectTransactionDefinition:
			fmt.Printf("Transaction type = %s\n", t.Type)
		case *pitOrgan.TradeClientExtensionsModifyTransactionDefinition:
			fmt.Printf("Transaction type = %s\n", t.Type)
		case *pitOrgan.TradeClientExtensionsModifyRejectTransactionDefinition:
			fmt.Printf("Transaction type = %s\n", t.Type)
		case *pitOrgan.MarginCallEnterTransactionDefinition:
			fmt.Printf("Transaction type = %s\n", t.Type)
		case *pitOrgan.MarginCallExtendTransactionDefinition:
			fmt.Printf("Transaction type = %s\n", t.Type)
		case *pitOrgan.MarginCallExitTransactionDefinition:
			fmt.Printf("Transaction type = %s\n", t.Type)
		case *pitOrgan.DelayedTradeClosureTransactionDefinition:
			fmt.Printf("Transaction type = %s\n", t.Type)
		case *pitOrgan.DailyFinancingTransactionDefinition:
			fmt.Printf("Transaction type = %s\n", t.Type)
		case *pitOrgan.ResetResettablePLTransactionDefinition:
			fmt.Printf("Transaction type = %s\n", t.Type)
		}
	}
}

func ExampleReceiverTransactionsIdrange_Get_idrangeParams() {
	connection := &pitOrgan.Connection{
		Token:       "OANDA_API_TOKEN",
		Environemnt: pitOrgan.OandaPractice,
		Timeout:     time.Second * 10,
	}
	api := connection.Accounts().AccountID("accountID").Transactions()

	params := &pitOrgan.GetTransactionsIdrangeParams{
		From: 1,
		To:   165,
	}
	data, err := api.Get(
		&GetTransactionsParams{Type: []TransactionFilterDefinition{
			OrderTransaction,
			CloseTransaction,
		}},
	)
	if err != nil {
		log.Fatalf("%+v", err)
	}

	params, err := data.IdrangeParams()
	if err != nil {
		log.Fatalf("%+v", err)
	}

	for _, param := range params {
		data, err := transactionAPI.Idrange().Get(param)
		if err != nil {
			log.Fatalf("%+v", err)
		}
		spew.Dump(data)
	}
}

func ExampleReceiverTransactionsStream_Get() {
	connection := &pitOrgan.Connection{
		Token:       "OANDA_API_TOKEN",
		Environemnt: pitOrgan.OandaPractice,
		Timeout:     time.Second * 10,
	}

	params := &GetTransactionsStreamParams{
		BufferSize: 100,
	}
	chs, err := connection.Accounts().AccountID("AccountID").Transactions().Get(params)
	if err != nil {
		log.Fatalf("%+v", err)
	}
	defer chs.Close()

	for {
		select {
		case transaction := <-chs.Transaction:
			spew.Dump(transaction)
		case err := <-chs.Error:
			log.Fatalf("%+v", err)
		}
	}
}
