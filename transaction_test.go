package pitOrgan

import (
	"github.com/davecgh/go-spew/spew"
	"testing"
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

