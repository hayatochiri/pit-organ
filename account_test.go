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

func Test_AccountSummary(t *testing.T) {
	connection := newConnection(t, OandaPractice)
	accountID := Getenv("ACCOUNT_ID")
	data, err := connection.Accounts().AccountID(accountID).Summary().Get()

	if err != nil {
		t.Fatalf("Get account summary failed.\n%+v", err)
	}

	t.Logf("Response:\n%s", spew.Sdump(data))
}

func Test_AccountInstruments(t *testing.T) {
	connection := newConnection(t, OandaPractice)
	accountID := Getenv("ACCOUNT_ID")
	params := &GetAccountInstrumentsParams{
		Instruments: []string{"USD_JPY", "EUR_JPY", "EUR_USD"},
	}
	data, err := connection.Accounts().AccountID(accountID).Instruments().Get(params)

	if err != nil {
		t.Fatalf("Get account instruments failed.\n%+v", err)
	}

	t.Logf("Response:\n%s", spew.Sdump(data))
}

func Test_AccountConfiguration(t *testing.T) {
	t.Run("AliasSettingSuccess", func(t *testing.T) {
		expect := "Test Account #1"

		connection := newConnection(t, OandaPractice)
		accountID := Getenv("ACCOUNT_ID")
		params := &PatchAccountConfigurationParams{
			&PatchAccountConfigurationBodyParams{
				Alias: expect,
			},
		}
		data, err := connection.Accounts().AccountID(accountID).Configuration().Patch(params)

		if err != nil {
			t.Fatalf("Patch account configuration failed.\n%+v", err)
		}

		if actual := data.ClientConfigureTransaction.Alias; actual != expect {
			t.Fatalf("Failed to set alias.\nExpect: %#v\nActual: %#v", expect, actual)
		}

		t.Logf("Response:\n%s", spew.Sdump(data))
	})

	t.Run("MarginRateSettingSuccess", func(t *testing.T) {
		expect := "0.04"

		connection := newConnection(t, OandaPractice)
		accountID := Getenv("ACCOUNT_ID")
		params := &PatchAccountConfigurationParams{
			&PatchAccountConfigurationBodyParams{
				MarginRate: expect,
			},
		}
		data, err := connection.Accounts().AccountID(accountID).Configuration().Patch(params)

		if err != nil {
			t.Fatalf("Patch account configuration failed.\n%+v", err)
		}

		if actual := string(data.ClientConfigureTransaction.MarginRate); actual != expect {
			t.Fatalf("Failed to set margin rate.\nExpect: %#v\nActual: %#v", expect, actual)
		}

		t.Logf("Response:\n%s", spew.Sdump(data))
	})
}
