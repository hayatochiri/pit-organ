package pitOrgan

import (
	"net/url"
	"time"
)

type OandaEnvironment int

const (
	OandaPractice OandaEnvironment = iota + 1
	OandaLive
	oandaDummy
)

type Connection struct {
	Token       string
	Environemnt OandaEnvironment
	Timeout     time.Duration
}

type baseURLs struct {
	rest   *url.URL
	stream *url.URL
}

func oandaBaseURL(env OandaEnvironment) *baseURLs {
	var urls *baseURLs

	switch env {
	case OandaPractice:
		urls = &baseURLs{
			rest:   parseURL("https://api-fxpractice.oanda.com"),
			stream: parseURL("https://stream-fxpractice.oanda.com"),
		}
	case OandaLive:
		urls = &baseURLs{
			rest:   parseURL("https://api-fxtrade.oanda.com"),
			stream: parseURL("https://stream-fxtrade.oanda.com"),
		}
	case oandaDummy:
		urls = &baseURLs{
			rest:   parseURL("https://192.0.2.1"),
			stream: parseURL("https://192.0.2.2"),
		}
	}

	return urls
}

func parseURL(urlString string) *url.URL {
	urlStructure, _ := url.Parse(urlString)

	return urlStructure
}
