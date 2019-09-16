package pitOrgan

import (
	"github.com/davecgh/go-spew/spew"
	"testing"
)

func Test_Orders(t *testing.T) {
	t.Run("PostMarketOrder", func(t *testing.T) {
		connection := newConnection(t, OandaPractice)
		accountID := Getenv("ACCOUNT_ID")

		params := &PostOrdersParams{
			Body: PostOrdersBodyParams{
				Order: MarketOrderRequestDefinition{
					Type:        "MARKET",
					Instrument:  "EUR_USD",
					Units:       "100",
					TimeInForce: "FOK",
				},
			},
		}

		data, err := connection.Accounts().AccountID(accountID).Orders().Post(params)
		if err != nil {
			t.Fatalf("Error occurred.\n%+v", err)
		}

		t.Logf("Response:\n%s", spew.Sdump(data))
	})

	t.Run("GetMarketOrder", func(t *testing.T) {
		paramsPatterns := []*GetOrdersParams{
			{},
			{IDs: []string{"1", "2", "3"}},
			{State: "ALL"},
			{Instrument: "EUR_USD"},
			{Count: 100},
			{BeforeID: "1"},
		}

		connection := newConnection(t, OandaPractice)
		accountID := Getenv("ACCOUNT_ID")

		for _, params := range paramsPatterns {
			data, err := connection.Accounts().AccountID(accountID).Orders().Get(params)
			if err != nil {
				t.Fatalf("Error occurred.\n%+v", err)
			}
			t.Logf("Response:\n%s", spew.Sdump(data))
		}
	})
}

func Test_PendingOrders(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		connection := newConnection(t, OandaPractice)
		accountID := Getenv("ACCOUNT_ID")

		data, err := connection.Accounts().AccountID(accountID).PendingOrders().Get()
		if err != nil {
			t.Fatalf("Error occurred.\n%+v", err)
		}

		t.Logf("Response:\n%s", spew.Sdump(data))
	})
}

func Test_OrderSpecifier(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		connection := newConnection(t, OandaPractice)
		accountID := Getenv("ACCOUNT_ID")
		accountIDPath := connection.Accounts().AccountID(accountID)

		orderID := func() string {
			data, err := accountIDPath.Orders().Get(&GetOrdersParams{State: "ALL"})
			if err != nil {
				t.Fatalf("Error occurred.\n%+v", err)
			}
			return string(data.Orders[0].ID)
		}()

		data, err := accountIDPath.Orders().OrderSpecifier(orderID).Get()
		if err != nil {
			t.Fatalf("Error occurred.\n%+v", err)
		}

		t.Logf("Response:\n%s", spew.Sdump(data))
	})

	t.Run("PutSuccess", func(t *testing.T) {
		connection := newConnection(t, OandaPractice)
		accountID := Getenv("ACCOUNT_ID")
		accountIDPath := connection.Accounts().AccountID(accountID)

		order := func() *PostOrdersSchema {
			params := &PostOrdersParams{
				Body: PostOrdersBodyParams{
					Order: MarketIfTouchedOrderRequestDefinition{
						Type:        "MARKET_IF_TOUCHED",
						Instrument:  "USD_JPY",
						Units:       "100",
						Price:       "150.00",
						TimeInForce: "GTC",
					},
				},
			}

			data, err := accountIDPath.Orders().Post(params)
			if err != nil {
				t.Fatalf("Error occurred.\n%+v", err)
			}

			return data
		}()

		params := &PutOrderSpecifierParams{
			Body: PutOrderSpecifierBodyParams{
				Order: MarketIfTouchedOrderRequestDefinition{
					Type:        "MARKET_IF_TOUCHED",
					Instrument:  "USD_JPY",
					Units:       "200",
					Price:       "110.0",
					TimeInForce: "GTC",
				},
			},
		}

		data, err := accountIDPath.Orders().OrderSpecifier(order.LastTransactionID).Put(params)
		if err != nil {
			t.Fatalf("Error occurred.\n%+v", err)
		}

		t.Logf("Response:\n%s", spew.Sdump(data))
	})
}
