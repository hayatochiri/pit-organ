package pitOrgan

import (
	"github.com/davecgh/go-spew/spew"
	"testing"
	"time"
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

func Test_TransactionsStream(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		connection := newConnection(t, OandaPractice)
		accountID := Getenv("ACCOUNT_ID")
		api := connection.Accounts().AccountID(accountID)

		orderParams := &PostOrdersParams{
			Body: PostOrdersBodyParams{
				Order: MarketOrderRequestDefinition{
					Type:        "MARKET",
					Instrument:  "EUR_USD",
					Units:       "100",
					TimeInForce: "FOK",
				},
			},
		}

		params := &GetTransactionsStreamParams{
			BufferSize: 100,
		}
		chs, err := api.Transactions().Stream().Get(params)
		if err != nil {
			t.Fatalf("Error occurred.\n%+v", err)
		}
		defer chs.Close()

		for i := 1; i <= 5; {
			select {
			case transaction := <-chs.Transaction:
				t.Logf("Stream:\n%s", spew.Sdump(transaction))
				i++
			case err := <-chs.Error:
				t.Fatalf("Error occurred.\n%+v", err)
			default:
				if _, err := api.Orders().Post(orderParams); err != nil {
					t.Fatalf("Error occurred.\n%+v", err)
				}
				time.Sleep(time.Second)
			}
		}
	})
}
