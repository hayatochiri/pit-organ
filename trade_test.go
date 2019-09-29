package pitOrgan

import (
	"context"
	"github.com/davecgh/go-spew/spew"
	"strconv"
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

		_, err := accountIDReceiver.Orders().Post(context.Background(), params)
		if err != nil {
			t.Fatalf("Error occurred.\n%+v", err)
		}
	}

	params := &GetTradesParams{}

	data, err := accountIDReceiver.Trades().Get(context.Background(), params)
	if err != nil {
		t.Fatalf("Error occurred.\n%+v", err)
	}

	t.Logf("Response:\n%s", spew.Sdump(data))
}

func Test_OpenTrades(t *testing.T) {
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

		_, err := accountIDReceiver.Orders().Post(context.Background(), params)
		if err != nil {
			t.Fatalf("Error occurred.\n%+v", err)
		}
	}

	data, err := accountIDReceiver.OpenTrades().Get(context.Background())
	if err != nil {
		t.Fatalf("Error occurred.\n%+v", err)
	}

	t.Logf("Response:\n%s", spew.Sdump(data))
}

func Test_TradeSpecifier(t *testing.T) {
	connection := newConnection(t, OandaPractice)
	accountID := Getenv("ACCOUNT_ID")
	accountIDReceiver := connection.Accounts().AccountID(accountID)

	order := func() *PostOrdersSchema {
		params := &PostOrdersParams{
			Body: PostOrdersBodyParams{
				Order: MarketOrderRequestDefinition{
					Type:       "MARKET",
					Instrument: "EUR_USD",
					Units:      "1",
				},
			},
		}
		order, err := accountIDReceiver.Orders().Post(context.Background(), params)
		if err != nil {
			t.Fatalf("Error occurred.\n%+v", err)
		}

		return order
	}()

	data, err := accountIDReceiver.Trades().TradeSpecifier(order.LastTransactionID).Get(context.Background())
	if err != nil {
		t.Fatalf("Error occurred.\n%+v", err)
	}

	t.Logf("Response:\n%s", spew.Sdump(data))
}

func Test_TradeSpecifierClose(t *testing.T) {
	connection := newConnection(t, OandaPractice)
	accountID := Getenv("ACCOUNT_ID")
	accountIDReceiver := connection.Accounts().AccountID(accountID)

	order := func() *PostOrdersSchema {
		params := &PostOrdersParams{
			Body: PostOrdersBodyParams{
				Order: MarketOrderRequestDefinition{Type: "MARKET", Instrument: "USD_JPY", Units: "100"},
			},
		}

		order, err := accountIDReceiver.Orders().Post(context.Background(), params)
		if err != nil {
			t.Fatalf("Error occurred.\n%+v", err)
		}

		return order
	}()

	params := &PutTradeSpecifierCloseParams{
		Body: &PutTradeSpecifierCloseBodyParams{},
	}

	data, err := accountIDReceiver.Trades().TradeSpecifier(order.LastTransactionID).Close().Put(context.Background(), params)
	if err != nil {
		t.Fatalf("Error occurred.\n%+v", err)
	}

	t.Logf("Response:\n%s", spew.Sdump(data))
}

func Test_TradeSpecifierClientExtensions(t *testing.T) {
	connection := newConnection(t, OandaPractice)
	accountID := Getenv("ACCOUNT_ID")
	accountIDReceiver := connection.Accounts().AccountID(accountID)

	order := func() *PostOrdersSchema {
		params := &PostOrdersParams{
			Body: PostOrdersBodyParams{
				Order: MarketOrderRequestDefinition{Type: "MARKET", Instrument: "USD_JPY", Units: "1"},
			},
		}

		order, err := accountIDReceiver.Orders().Post(context.Background(), params)
		if err != nil {
			t.Fatalf("Error occurred.\n%+v", err)
		}

		return order
	}()

	params := &PutTradeSpecifierClientExtensionsParams{
		Body: &PutTradeSpecifierClientExtensionsBodyParams{
			ClientExtensions: &ClientExtensionsDefinition{
				ID:      "Hoge" + order.LastTransactionID,
				Tag:     "Huga",
				Comment: "Piyo",
			},
		},
	}

	data, err := accountIDReceiver.Trades().TradeSpecifier(order.LastTransactionID).ClientExtensions().Put(context.Background(), params)
	if err != nil {
		t.Fatalf("Error occurred.\n%+v", err)
	}

	t.Logf("Response:\n%s", spew.Sdump(data))
}

func Test_TradeSpecifierOrders(t *testing.T) {
	connection := newConnection(t, OandaPractice)
	accountID := Getenv("ACCOUNT_ID")
	accountIDReceiver := connection.Accounts().AccountID(accountID)

	order := func() *PostOrdersSchema {
		params := &PostOrdersParams{
			Body: PostOrdersBodyParams{
				Order: MarketOrderRequestDefinition{Type: "MARKET", Instrument: "USD_JPY", Units: "10000"},
			},
		}

		order, err := accountIDReceiver.Orders().Post(context.Background(), params)
		if err != nil {
			t.Fatalf("Error occurred.\n%+v", err)
		}

		return order
	}()

	price, err := strconv.ParseFloat(order.OrderFillTransaction.Price, 32)
	if err != nil {
		t.Fatalf("Error occurred.\n%+v", err)
	}

	params := &PutTradeSpecifierOrdersParams{
		Body: &PutTradeSpecifierOrdersBodyParams{
			TakeProfit: &TakeProfitDetailsDefinition{
				Price: strconv.FormatFloat(price+1.0, 'f', 4, 32),
			},
			StopLoss: &StopLossDetailsDefinition{
				Price: strconv.FormatFloat(price-1.0, 'f', 4, 32),
			},
			TrailingStopLoss: &TrailingStopLossDetailsDefinition{
				Distance: "0.05",
			},
		},
	}

	data, err := accountIDReceiver.Trades().TradeSpecifier(order.LastTransactionID).Orders().Put(context.Background(), params)
	if err != nil {
		t.Fatalf("Error occurred.\n%+v", err)
	}

	t.Logf("Response:\n%s", spew.Sdump(data))
}
