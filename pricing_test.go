package pitOrgan

import (
	"context"
	"github.com/davecgh/go-spew/spew"
	"testing"
	"time"
)

func Test_Pricing(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		connection := newConnection(t, OandaPractice)
		accountID := Getenv("ACCOUNT_ID")

		params := &GetPricingParams{Instruments: []string{"EUR_USD"}}
		data, err := connection.Accounts().AccountID(accountID).Pricing().Get(context.Background(), params)
		if err != nil {
			t.Fatalf("Error occurred.\n%+v", err)
		}

		t.Logf("Response:\n%s", spew.Sdump(data))
	})
}

func Test_PricingStream(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		connection := newConnection(t, OandaPractice)
		accountID := Getenv("ACCOUNT_ID")

		params := &GetPricingStreamParams{
			BufferSize:  300,
			Instruments: []string{"EUR_ZAR", "EUR_PLN", "AUD_JPY", "USD_CAD", "USD_NOK", "CAD_SGD", "HKD_JPY", "NZD_JPY", "USD_HUF", "CHF_ZAR", "EUR_CZK", "AUD_HKD", "GBP_NZD", "NZD_HKD", "NZD_CHF", "USD_SAR", "GBP_CAD", "CAD_JPY", "ZAR_JPY", "NZD_SGD", "GBP_ZAR", "NZD_CAD", "USD_INR", "CAD_HKD", "SGD_CHF", "CAD_CHF", "AUD_SGD", "EUR_NOK", "EUR_CHF", "GBP_USD", "USD_MXN", "USD_CHF", "AUD_CHF", "EUR_DKK", "AUD_USD", "CHF_HKD", "USD_THB", "GBP_CHF", "TRY_JPY", "AUD_CAD", "SGD_JPY", "EUR_NZD", "USD_HKD", "EUR_AUD", "USD_DKK", "CHF_JPY", "EUR_SGD", "USD_SGD", "EUR_SEK", "USD_JPY", "EUR_TRY", "USD_CZK", "GBP_AUD", "USD_PLN", "EUR_USD", "AUD_NZD", "SGD_HKD", "EUR_HUF", "NZD_USD", "USD_CNH", "EUR_HKD", "EUR_JPY", "GBP_PLN", "GBP_JPY", "USD_TRY", "EUR_CAD", "USD_SEK", "GBP_SGD", "EUR_GBP", "GBP_HKD", "USD_ZAR"},
		}
		ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)
		chs, err := connection.Accounts().AccountID(accountID).Pricing().Stream().Get(ctx, params)
		if err != nil {
			t.Fatalf("Error occurred.\n%+v", err)
		}

		for range chs.PriceCh {
		}

		chs.Close()

		if err := chs.Err(); err != nil {
			t.Fatalf("Error occurred.\n%+v", err)
		}
	})

	t.Run("Unauthorized", func(t *testing.T) {
		connection := newConnection(t, OandaPractice)
		connection.Token = "foo"
		accountID := Getenv("ACCOUNT_ID")

		params := &GetPricingStreamParams{
			BufferSize:  100,
			Instruments: []string{"USD_JPY", "EUR_JPY", "EUR_USD"},
		}

		chs, err := connection.Accounts().AccountID(accountID).Pricing().Stream().Get(context.Background(), params)

		if err == nil {
			chs.Close()
			t.Fatal("Got 200 OK but unauthorized.")
		}

		if _, ok := nakedError(err).(*UnauthorizedError); !ok {
			t.Fatalf("Got unexpected error.\n%+v", err)
		}

		t.Logf("Error occurred as expected.\n%+v", err)
	})
}
