package pitOrgan

import (
	"time"
)

type OandaEnvironment int

type Connection struct {
	Token       string
	Environemnt OandaEnvironment
	Timeout     time.Duration
}

