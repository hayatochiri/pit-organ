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
