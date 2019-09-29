package pitOrgan

import (
	"context"
	"github.com/davecgh/go-spew/spew"
	"testing"
)

func Test_Positions(t *testing.T) {
	connection := newConnection(t, OandaPractice)
	accountID := Getenv("ACCOUNT_ID")

	data, err := connection.Accounts().AccountID(accountID).Positions().Get(context.Background())
	if err != nil {
		t.Fatalf("Error occurred.\n%+v", err)
	}

	t.Logf("Response:\n%s", spew.Sdump(data))
}

func Test_OpenPositions(t *testing.T) {
	connection := newConnection(t, OandaPractice)
	accountID := Getenv("ACCOUNT_ID")
	accountIDReceiver := connection.Accounts().AccountID(accountID)

	paramsArray := []MarketOrderRequestDefinition{
		{Type: "MARKET", Instrument: "EUR_USD", Units: "100"},
		{Type: "MARKET", Instrument: "USD_JPY", Units: "150"},
		{Type: "MARKET", Instrument: "EUR_JPY", Units: "200"},
	}
	for _, orderRequest := range paramsArray {
		params := &PostOrdersParams{
			Body: PostOrdersBodyParams{
				Order: orderRequest,
			},
		}

		_, err := accountIDReceiver.Orders().Post(context.Background(), params)
		if err != nil {
			t.Fatalf("Error occurred.\n%+v", err)
		}
	}

	data, err := accountIDReceiver.OpenPositions().Get(context.Background())
	if err != nil {
		t.Fatalf("Error occurred.\n%+v", err)
	}

	t.Logf("Response:\n%s", spew.Sdump(data))
}

func Test_PositionsInstrument(t *testing.T) {
	connection := newConnection(t, OandaPractice)
	accountID := Getenv("ACCOUNT_ID")
	accountIDReceiver := connection.Accounts().AccountID(accountID)

	{
		params := &PostOrdersParams{
			Body: PostOrdersBodyParams{
				Order: MarketOrderRequestDefinition{Type: "MARKET", Instrument: "EUR_JPY", Units: "300"},
			},
		}

		_, err := accountIDReceiver.Orders().Post(context.Background(), params)
		if err != nil {
			t.Fatalf("Error occurred.\n%+v", err)
		}
	}

	data, err := accountIDReceiver.Positions().Instrument("EUR_JPY").Get(context.Background())
	if err != nil {
		t.Fatalf("Error occurred.\n%+v", err)
	}

	t.Logf("Response:\n%s", spew.Sdump(data))
}

func Test_PositionsInstrumentClose(t *testing.T) {
	connection := newConnection(t, OandaPractice)
	accountID := Getenv("ACCOUNT_ID")
	accountIDReceiver := connection.Accounts().AccountID(accountID)

	{
		params := &PostOrdersParams{
			Body: PostOrdersBodyParams{
				Order: MarketOrderRequestDefinition{Type: "MARKET", Instrument: "USD_JPY", Units: "1"},
			},
		}

		_, err := accountIDReceiver.Orders().Post(context.Background(), params)
		if err != nil {
			t.Fatalf("Error occurred.\n%+v", err)
		}
	}

	params := &PutPositionsInstrumentCloseParams{
		Body: &PutPositionsInstrumentCloseBodyParams{
			LongUnits: "ALL",
		},
	}

	data, err := accountIDReceiver.Positions().Instrument("USD_JPY").Close().Put(context.Background(), params)
	if err != nil {
		t.Fatalf("Error occurred.\n%+v", err)
	}

	t.Logf("Response:\n%s", spew.Sdump(data))
}
