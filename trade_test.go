package pitOrgan

import (
	"github.com/davecgh/go-spew/spew"
	"testing"
)

func Test_Trades(t *testing.T) {
	connection := newConnection(t, OandaPractice)
	accountID := Getenv("ACCOUNT_ID")
	accountIDReceiver := connection.Accounts().AccountID(accountID)

	paramsArray := []MarketOrderRequestDefinition{
		{Type: "MARKET", Instrument: "EUR_USD", Units: "1"},
		{Type: "MARKET", Instrument: "USD_JPY", Units: "1"},
		{Type: "MARKET", Instrument: "EUR_JPY", Units: "1"},
	}
	for _, orderRequest := range paramsArray {
		params := &PostOrdersParams{
			Body: PostOrdersBodyParams{
				Order: orderRequest,
			},
		}

		_, err := accountIDReceiver.Orders().Post(params)
		if err != nil {
			t.Fatalf("Error occurred.\n%+v", err)
		}
	}

	params := &GetTradesParams{}

	data, err := accountIDReceiver.Trades().Get(params)
	if err != nil {
		t.Fatalf("Error occurred.\n%+v", err)
	}

	t.Logf("Response:\n%s", spew.Sdump(data))
}