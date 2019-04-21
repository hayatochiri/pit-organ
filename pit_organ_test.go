package pitOrgan

import (
	"github.com/joho/godotenv"
	"golang.org/x/xerrors"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(xerrors.Errorf("Error loading .env file: %w", err))
	}
}
