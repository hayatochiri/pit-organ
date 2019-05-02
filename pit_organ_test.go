package pitOrgan

import (
	. "github.com/MakeNowJust/heredoc/dot"
	"github.com/joho/godotenv"
	"golang.org/x/xerrors"
	"net/url"
	"os"
	"reflect"
	"testing"
	"time"
)

func init() {
	if _, err := os.Stat("./.env"); err != nil {
		file, fErr := os.OpenFile("./.env", os.O_WRONLY|os.O_CREATE, 0644)
		if fErr != nil {
			panic(xerrors.Errorf("Create .env failed: %w", err))
		}
		_, fErr = file.WriteString(D(`
			TOKEN=
			ACCOUNT_ID=
		`))
		if fErr != nil {
			panic(xerrors.Errorf("Write .env failed: %w", err))
		}

		panic(xerrors.Errorf(".env file not exist: %w", err))
	}

	err := godotenv.Load()
	if err != nil {
		panic(xerrors.Errorf("Error loading .env file: %w", err))
	}

	if os.Getenv("TOKEN") == "" {
		panic(xerrors.Errorf("Env 'TOKEN' is empty"))
	}

	if os.Getenv("ACCOUNT_ID") == "" {
		panic(xerrors.Errorf("Env 'ACCOUNT_ID' is empty"))
	}
}

func Test_oandaBaseURL(t *testing.T) {
	var expect *baseURLs

	expect = new(baseURLs)
	expect.rest, _ = url.Parse("https://api-fxpractice.oanda.com")
	expect.stream, _ = url.Parse("https://stream-fxpractice.oanda.com")
	if actual := oandaBaseURL(OandaPractice); !reflect.DeepEqual(actual, expect) {
		if !reflect.DeepEqual(actual.rest, expect.rest) {
			t.Errorf("\ngot:  %#v\nwant: %#v", actual.rest, expect.rest)
		}
		if !reflect.DeepEqual(actual.stream, expect.stream) {
			t.Errorf("\ngot:  %#v\nwant: %#v", actual.stream, expect.stream)
		}
	}

	expect = new(baseURLs)
	expect.rest, _ = url.Parse("https://api-fxtrade.oanda.com")
	expect.stream, _ = url.Parse("https://stream-fxtrade.oanda.com")
	if actual := oandaBaseURL(OandaLive); !reflect.DeepEqual(actual, expect) {
		if !reflect.DeepEqual(actual.rest, expect.rest) {
			t.Errorf("\ngot:  %#v\nwant: %#v", actual.rest, expect.rest)
		}
		if !reflect.DeepEqual(actual.stream, expect.stream) {
			t.Errorf("\ngot:  %#v\nwant: %#v", actual.stream, expect.stream)
		}
	}
}

func newConnection(t *testing.T, env OandaEnvironment) *Connection {

	connection := &Connection{
		Token:       Getenv("TOKEN"),
		Environemnt: env,
		Timeout:     time.Second * 10,
	}

	return connection
}

func Getenv(k string) string {
	return os.Getenv(k)
}
