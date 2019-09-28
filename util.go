package pitOrgan

import (
	"bytes"
	"encoding/json"
	"github.com/davecgh/go-spew/spew"
	"golang.org/x/xerrors"
	"io/ioutil"
	"net/http"
	"path"
	"reflect"
	"strconv"
	"strings"
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

func Int(v int) *int          { return &v }
func String(v string) *string { return &v }
func Bool(v bool) *bool       { return &v }

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

func parseResponse(resp *http.Response, data interface{}, strict bool) (interface{}, error) {
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, xerrors.Errorf("Read response body failed: %w", err)
	}

	var errMessage string
	switch resp.StatusCode {
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
		return nil, xerrors.Errorf("Unexpected status code(%d)", resp.StatusCode)
	}

	if err := json.Unmarshal(body, data); err != nil {
		return nil, xerrors.Errorf("Unmarshal response body failed: %w", err)
	}

	if strict {
		if err := compareJson(data, body); err != nil {
			return nil, xerrors.Errorf("Response body JSON is different from unmarshalled it: %w", err)
		}
	}

	if resp.StatusCode/100 != 2 {
		return nil, xerrors.Errorf("%s: %w", errMessage, data)
	}

	return data, nil
}

func compareJson(jsonObj interface{}, jsonStr []byte) error {
	bytes, err := json.Marshal(jsonObj)
	if err != nil {
		return xerrors.Errorf("Marshal JSON object failed: %w", err)
	}

	actual := new(interface{})
	if err = json.Unmarshal(bytes, actual); err != nil {
		return xerrors.Errorf("Reunmarshal JSON string failed: %w", err)
	}

	expect := new(interface{})
	if err = json.Unmarshal(jsonStr, expect); err != nil {
		return xerrors.Errorf("Unmarshal JSON string failed: %w", err)
	}

	if err := deepEqual(expect, actual, []string{reflect.TypeOf(jsonObj).String()}); err != nil {
		return xerrors.Errorf("Unmarshalled JSON is lacking:\nExpect: %sActual: %s\n: %w", string(jsonStr), string(bytes), err)
	}

	return nil
}

func deepEqual(expect, actual interface{}, breadcrumbs []string) error {
	if reflect.TypeOf(expect).String() == "*interface {}" {
		expect = reflect.Indirect(reflect.ValueOf(expect)).Interface()
	}

	if actual == nil {
		return xerrors.Errorf("Actual value is not equal to expect\nBreadcrumbs: %s\nExpect: %sActual: %s", strings.Join(breadcrumbs, " > "), spew.Sdump(expect), spew.Sdump(actual))
	}
	if reflect.TypeOf(actual).String() == "*interface {}" {
		actual = reflect.Indirect(reflect.ValueOf(actual)).Interface()
	}

	switch expectValue := expect.(type) {
	case map[string]interface{}:
		actualValue, ok := actual.(map[string]interface{})
		if !ok {
			return xerrors.Errorf("Actual value is not equal to expect\nBreadcrumbs: %s\nExpect: %+v\nActual: %+v", strings.Join(breadcrumbs, " > "), expect, actual)
		}

		for k := range expectValue {
			if v, ok := expectValue[k].([]interface{}); ok && len(v) == 0 {
				if _, ok := actualValue[k]; !ok {
					return nil
				}
			}

			if err := deepEqual(expectValue[k], actualValue[k], append(breadcrumbs, k)); err != nil {
				return err
			}
		}
	case []interface{}:
		actualValue, ok := actual.([]interface{})
		if !ok {
			return xerrors.Errorf("Actual value is not equal to expect\nBreadcrumbs: %s\nExpect: %+v\nActual: %+v", strings.Join(breadcrumbs, " > "), expect, actual)
		}

		for n := range expectValue {
			if err := deepEqual(expectValue[n], actualValue[n], append(breadcrumbs, "["+strconv.Itoa(n)+"]")); err != nil {
				return err
			}
		}
	case string:
		switch actualValue := actual.(type) {
		case string:
			if expectValue != actualValue {
				return xerrors.Errorf("Actual value is not equal to expect\nBreadcrumbs: %s\nExpect: %sActual: %s", strings.Join(breadcrumbs, " > "), spew.Sdump(expect), spew.Sdump(actual))
			}
		case float64:
			ev, _ := strconv.ParseFloat(expectValue, 64)
			if ev != actualValue {
				return xerrors.Errorf("Actual value is not equal to expect\nBreadcrumbs: %s\nExpect: %sActual: %s", strings.Join(breadcrumbs, " > "), spew.Sdump(expect), spew.Sdump(actual))
			}
		default:
			return xerrors.Errorf("Actual value is not equal to expect\nBreadcrumbs: %s\nExpect: %sActual: %s", strings.Join(breadcrumbs, " > "), spew.Sdump(expect), spew.Sdump(actual))
		}
	case float64:
		switch actualValue := actual.(type) {
		case float64:
			if expectValue != actualValue {
				return xerrors.Errorf("Actual value is not equal to expect\nBreadcrumbs: %s\nExpect: %sActual: %s", strings.Join(breadcrumbs, " > "), spew.Sdump(expect), spew.Sdump(actual))
			}
		default:
			return xerrors.Errorf("Actual value is not equal to expect\nBreadcrumbs: %s\nExpect: %sActual: %s", strings.Join(breadcrumbs, " > "), spew.Sdump(expect), spew.Sdump(actual))
		}
	case bool:
		switch actualValue := actual.(type) {
		case bool:
			if expectValue != actualValue {
				return xerrors.Errorf("Actual value is not equal to expect\nBreadcrumbs: %s\nExpect: %sActual: %s", strings.Join(breadcrumbs, " > "), spew.Sdump(expect), spew.Sdump(actual))
			}
		default:
			return xerrors.Errorf("Actual value is not equal to expect\nBreadcrumbs: %s\nExpect: %sActual: %s", strings.Join(breadcrumbs, " > "), spew.Sdump(expect), spew.Sdump(actual))
		}
	default:
		return xerrors.Errorf("Unexpected type was given\nBreadcrumbs: %s\nType: %s\nExpect: %sActual: %s", strings.Join(breadcrumbs, " > "), reflect.TypeOf(expect).String(), spew.Sdump(expect), spew.Sdump(actual))
	}

	return nil
}
