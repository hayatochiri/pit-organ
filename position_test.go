package pitOrgan

import (
	"github.com/davecgh/go-spew/spew"
	"testing"
)

func Test_Positions(t *testing.T) {
	connection := newConnection(t, OandaPractice)
	accountID := Getenv("ACCOUNT_ID")

	data, err := connection.Accounts().AccountID(accountID).Positions().Get()
	if err != nil {
		t.Fatalf("Error occurred.\n%+v", err)
	}

	t.Logf("Response:\n%s", spew.Sdump(data))
}
