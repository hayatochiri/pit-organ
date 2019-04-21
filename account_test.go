package pitOrgan

import (
	"github.com/davecgh/go-spew/spew"
	"testing"
)

func Test_Accounts(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		connection := newConnection(t, OandaPractice)
		data, err := connection.Accounts().Get()

		if err != nil {
			t.Fatalf("Get accounts failed.\n%+v", err)
		}

		t.Logf("Response:\n%s", spew.Sdump(data))
	})
}

func Test_AccountID(t *testing.T) {
	connection := newConnection(t, OandaPractice)
	accountID := Getenv("ACCOUNT_ID")
	data, err := connection.Accounts().AccountID(accountID).Get()

	if err != nil {
		t.Fatalf("Get account ID failed.\n%+v", err)
	}

	t.Logf("Response:\n%s", spew.Sdump(data))
}
