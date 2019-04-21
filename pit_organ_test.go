package pitOrgan

import (
	"github.com/joho/godotenv"
	"golang.org/x/xerrors"
	"os"
	"testing"
	"time"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(xerrors.Errorf("Error loading .env file: %w", err))
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
