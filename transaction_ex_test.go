package pitOrgan_test

import (
	"context"
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
	data, err := connection.Accounts().AccountID("AccountID").Transactions().Get(context.Background(), params)
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
	data, err := connection.Accounts().AccountID("AccountID").Transactions().Idrange().Get(context.Background(), params)
	if err != nil {
		log.Fatalf("%+v", err)
	}

	spew.Dump(data)
}

func ExampleReceiverTransactionsIdrange_Get_idrangeParams() {
	connection := &pitOrgan.Connection{
		Token:       "OANDA_API_TOKEN",
		Environemnt: pitOrgan.OandaPractice,
		Timeout:     time.Second * 10,
	}
	api := connection.Accounts().AccountID("accountID").Transactions()

	data, err := api.Get(
		context.Background(),
		&pitOrgan.GetTransactionsParams{Type: []pitOrgan.TransactionFilterDefinition{
			pitOrgan.OrderTransaction,
			pitOrgan.CloseTransaction,
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
		data, err := api.Idrange().Get(context.Background(), param)
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

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	params := &pitOrgan.GetTransactionsStreamParams{
		BufferSize: 100,
	}
	chs, err := connection.Accounts().AccountID("AccountID").Transactions().Stream().Get(ctx, params)
	if err != nil {
		log.Fatalf("%+v", err)
	}
	defer chs.Close()

	for {
		select {
		case transaction := <-chs.TransactionCh:
			spew.Dump(transaction)
		}
	}

	if err := chs.Err(); err != nil {
		log.Fatalf("%+v", err)
	}
}
