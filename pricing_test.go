package pitOrgan

import (
	"github.com/davecgh/go-spew/spew"
	"testing"
)

func Test_Pricing(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		connection := newConnection(t, OandaPractice)
		accountID := Getenv("ACCOUNT_ID")

		params := &GetPricingParams{Instruments: []string{"EUR_USD"}}
		data, err := connection.Accounts().AccountID(accountID).Pricing().Get(params)
		if err != nil {
			t.Fatalf("Error occurred.\n%+v", err)
		}

		t.Logf("Response:\n%s", spew.Sdump(data))
	})
}
