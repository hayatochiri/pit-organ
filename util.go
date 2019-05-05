package pitOrgan

import (
	"bytes"
	"encoding/json"
	"golang.org/x/xerrors"
	"io/ioutil"
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

func parseBody(body []byte, statusCode int, data interface{}) (interface{}, error) {
	var errMessage string
	switch statusCode {
	case 200, 201:
		errMessage = ""
		if data == nil {
			return nil, xerrors.New("Variable that receives the response is nil")
		}
	case 400:
		errMessage = "400 bad request"
		if data == nil {
			data = new(BadRequestError)
		}
	case 401:
		errMessage = "401 unauthorized"
		if data == nil {
			data = new(UnauthorizedError)
		}
	case 403:
		errMessage = "403 forbidden"
		if data == nil {
			data = new(ForbiddenError)
		}
	case 404:
		errMessage = "404 not found"
		if data == nil {
			data = new(NotFoundError)
		}
	// TODO: 405
	// TODO: 416
	default:
		return nil, xerrors.Errorf("Unexpected status code(%d)", statusCode)
	}

	if err := json.Unmarshal(body, data); err != nil {
		return nil, xerrors.Errorf("Unmarshal response body failed: %w", err)
	}

	if statusCode/100 != 2 {
		return nil, xerrors.Errorf("%s: %w", errMessage, data)
	}

	return data, nil
}

func parseResponse(resp *http.Response, data interface{}) (interface{}, error) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, xerrors.Errorf("Read response body failed: %w", err)
	}

	return parseBody(body, resp.StatusCode, data)
}
