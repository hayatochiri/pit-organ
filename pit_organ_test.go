package pitOrgan

import (
	. "github.com/MakeNowJust/heredoc/dot"
	"github.com/joho/godotenv"
	"golang.org/x/xerrors"
	"os"
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

func newConnection(t *testing.T, env OandaEnvironment) *Connection {

	connection := &Connection{
		Token:       os.Getenv("TOKEN"),
		Environemnt: env,
		Timeout:     time.Second * 10,
	}

	return connection
}

func Getenv(k string) string {
	return os.Getenv(k)
}
