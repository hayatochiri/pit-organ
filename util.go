package pitOrgan

import (
	"bytes"
	"encoding/json"
	"golang.org/x/xerrors"
	"net/http"
	"path"
)

type header struct {
	key   string
	value string
}

type query struct {
	key   string
	value string
}

type requestParams struct {
	method   string
	endPoint string
	headers  []header
	queries  []query
	body     interface{}
}

func (c *Connection) request(params *requestParams) (*http.Response, error) {

	destURL := oandaBaseURL(c.Environemnt).rest
	destURL.Path = path.Join(destURL.Path, params.endPoint)

	body, _ := json.Marshal(params.body)
	req, err := http.NewRequest(params.method, destURL.String(), bytes.NewBuffer(body))
	if err != nil {
		return nil, xerrors.Errorf("Prepare new request failed: %w", err)
	}

	// req.Header.Set("User-Agent", "Go 1.1 package http")
	req.Header.Set("Authorization", "Bearer "+c.Token)
	req.Header.Set("Content-Type", "application/json")
	for _, h := range params.headers {
		req.Header.Set(h.key, h.value)
	}

	reqQ := req.URL.Query()
	for _, q := range params.queries {
		reqQ.Add(q.key, q.value)
	}
	req.URL.RawQuery = reqQ.Encode()

	client := http.Client{
		Timeout: c.Timeout,
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, xerrors.Errorf("Request canceled: %w", err)
	}

	return resp, nil
}

func (c *Connection) stream(params *requestParams) (*http.Response, error) {
	destURL := oandaBaseURL(c.Environemnt).stream

	destURL.Path = path.Join(destURL.Path, params.endPoint)

	req, err := http.NewRequest(params.method, destURL.String(), nil)
	if err != nil {
		return nil, xerrors.Errorf("error in stream method: %w", err)
	}

	// req.Header.Set("User-Agent", "Go 1.1 package http")
	req.Header.Set("Authorization", "Bearer "+c.Token)
	req.Header.Set("Content-Type", "application/json")
	for _, h := range params.headers {
		req.Header.Set(h.key, h.value)
	}

	reqQ := req.URL.Query()
	for _, q := range params.queries {
		reqQ.Add(q.key, q.value)
	}
	req.URL.RawQuery = reqQ.Encode()

	client := http.Client{
		Timeout: 0,
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, xerrors.Errorf("error in stream method: %w", err)
	}

	return resp, nil
}
