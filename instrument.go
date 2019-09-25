package pitOrgan

import (
	"github.com/peterhellberg/link"
	"golang.org/x/xerrors"
	"net/url"
	"strconv"
	"time"
)

/* Receivers */

type ReceiverInstruments struct {
	Connection *Connection
}

func (c *Connection) Instruments() *ReceiverInstruments {
	return &ReceiverInstruments{
		Connection: c,
	}
}

type ReceiverInstrument struct {
	Instrument string
	Connection *Connection
}

func (r *ReceiverInstruments) Instrument(i string) *ReceiverInstrument {
	return &ReceiverInstrument{
		Instrument: i,
		Connection: r.Connection,
	}
}

type ReceiverInstrumentCandles struct {
	Instrument string
	Connection *Connection
}

func (r *ReceiverInstrument) Candles() *ReceiverInstrumentCandles {
	return &ReceiverInstrumentCandles{
		Instrument: r.Instrument,
		Connection: r.Connection,
	}
}

type ReceiverInstrumentOrderBook struct {
	Instrument string
	Connection *Connection
}

func (r *ReceiverInstrument) OrderBook() *ReceiverInstrumentOrderBook {
	return &ReceiverInstrumentOrderBook{
		Instrument: r.Instrument,
		Connection: r.Connection,
	}
}

type ReceiverInstrumentPositionBook struct {
	Instrument string
	Connection *Connection
}

func (r *ReceiverInstrument) PositionBook() *ReceiverInstrumentPositionBook {
	return &ReceiverInstrumentPositionBook{
		Instrument: r.Instrument,
		Connection: r.Connection,
	}
}

/* Params */

type GetInstrumentCandlesParams struct {
	PriceMid          bool
	PriceBid          bool
	PriceAsk          bool
	Granularity       CandlestickGranularityDefinition
	Count             int
	From              time.Time
	To                time.Time
	Smooth            *bool
	IncludeFirst      *bool
	DailyAlignment    *int
	AlignmentTimezone string
	WeeklyAlignment   WeeklyAlignmentDefinition
}

type GetInstrumentOrderBookParams struct {
	Time time.Time // The time of the snapshot to fetch. If not specified, then the most recent snapshot is fetched.
}

type GetInstrumentPositionBookParams struct {
	Time time.Time // The time of the snapshot to fetch. If not specified, then the most recent snapshot is fetched.
}

/* Schemas */

type GetInstrumentCandlesSchema struct {
	Instrument  InstrumentNameDefinition         `json:"instrument,omitempty"`
	Granularity CandlestickGranularityDefinition `json:"granularity,omitempty"`
	Candles     []*CandlestickDefinition         `json:"candles,omitempty"`
}

type GetInstrumentOrderBookSchema struct {
	// The instrument’s order book
	OrderBook *OrderBookDefinition `json:"orderBook,omitempty"`
	Headers   struct {
		RequestID string
		Link      struct {
			Prev time.Time
			Next time.Time
		}
	}
}

type GetInstrumentPositionBookSchema struct {
	// The instrument’s position book
	PositionBook *PositionBookDefinition `json:"positionBook,omitempty"`
	Headers      struct {
		RequestID string
		Link      struct {
			Prev time.Time
			Next time.Time
		}
	}
}

/* API */

// GET /v3/instruments/{instrument}/candles
//
// Fetch candlestick data for an instrument.
func (r *ReceiverInstrumentCandles) Get(params *GetInstrumentCandlesParams) (*GetInstrumentCandlesSchema, error) {
	resp, err := r.Connection.request(
		&requestParams{
			method:   "GET",
			endPoint: "/v3/instruments/" + r.Instrument + "/candles",
			headers: []header{
				{key: "Accept-Datetime-Format", value: "RFC3339"},
			},
			queries: func() []query {
				q := make([]query, 0, 10)
				if params == nil {
					return q
				}

				// price
				price := make([]byte, 0, 3)
				if params.PriceMid {
					price = append(price, 'M')
				}
				if params.PriceBid {
					price = append(price, 'B')
				}
				if params.PriceAsk {
					price = append(price, 'A')
				}
				if len(price) > 0 {
					q = append(q, query{key: "price", value: string(price)})
				}

				// granularity
				if params.Granularity != "" {
					q = append(q, query{key: "granularity", value: string(params.Granularity)})
				}

				// count
				if params.Count > 0 {
					q = append(q, query{key: "count", value: strconv.Itoa(params.Count)})
				}

				// from
				if !params.From.IsZero() {
					q = append(q, query{key: "from", value: params.From.Format(time.RFC3339Nano)})
				}

				// to
				if !params.To.IsZero() {
					q = append(q, query{key: "to", value: params.To.Format(time.RFC3339Nano)})
				}

				// smooth
				if params.Smooth != nil {
					q = append(q, query{key: "smooth", value: strconv.FormatBool(*params.Smooth)})
				}

				// includeFirst
				if params.IncludeFirst != nil {
					q = append(q, query{key: "includeFirst", value: strconv.FormatBool(*params.IncludeFirst)})
				}

				// dailyAlignment
				if params.DailyAlignment != nil {
					q = append(q, query{key: "dailyAlignment", value: strconv.Itoa(*params.DailyAlignment)})
				}

				// alignmentTimezone
				if params.AlignmentTimezone != "" {
					q = append(q, query{key: "alignmentTimezone", value: params.AlignmentTimezone})
				}

				// weeklyAlignment
				if params.WeeklyAlignment != "" {
					q = append(q, query{key: "weeklyAlignment", value: string(params.WeeklyAlignment)})
				}

				return q
			}(),
		},
	)
	if err != nil {
		return nil, xerrors.Errorf("Get instrument candles canceled: %w", err)
	}
	defer resp.Body.Close()

	var data interface{}
	switch resp.StatusCode {
	case 200:
		data = new(GetInstrumentCandlesSchema)
	}

	data, err = parseResponse(resp, data, r.Connection.strict)
	if err != nil {
		return nil, xerrors.Errorf("Get instrument candles failed: %w", err)
	}
	return data.(*GetInstrumentCandlesSchema), nil
}

// GET /v3/instruments/{instrument}/orderBook
//
// Fetch an order book for an instrument.
func (r *ReceiverInstrumentOrderBook) Get(params *GetInstrumentOrderBookParams) (*GetInstrumentOrderBookSchema, error) {
	resp, err := r.Connection.request(
		&requestParams{
			method:   "GET",
			endPoint: "/v3/instruments/" + r.Instrument + "/orderBook",
			headers: []header{
				{key: "Accept-Datetime-Format", value: "RFC3339"},
			},
			queries: func() []query {
				q := make([]query, 0, 1)
				if params == nil {
					return q
				}

				// time
				if !params.Time.IsZero() {
					q = append(q, query{key: "time", value: params.Time.Format(time.RFC3339Nano)})
				}

				return q
			}(),
		},
	)
	if err != nil {
		return nil, xerrors.Errorf("Get instrument order book canceled: %w", err)
	}
	defer resp.Body.Close()

	var data interface{}
	switch resp.StatusCode {
	case 200:
		schema := new(GetInstrumentOrderBookSchema)
		if reqID, ok := resp.Header["Requestid"]; ok {
			schema.Headers.RequestID = reqID[0]
		}
		if links, ok := resp.Header["Link"]; ok {
			for _, l := range link.Parse(links[0]) {
				u, err := url.Parse(l.URI)
				if err != nil {
					return nil, xerrors.Errorf("Parse orderbook header failed: %w", err)
				}
				if refTimes, ok := u.Query()["time"]; ok {
					refTime, err := time.Parse(time.RFC3339, refTimes[0])
					if err != nil {
						return nil, xerrors.Errorf("Parse orderbook %#v time: %w", l.Rel, err)
					}
					switch l.Rel {
					case "prev":
						schema.Headers.Link.Prev = refTime
					case "next":
						schema.Headers.Link.Next = refTime
					default:
						return nil, xerrors.Errorf("Unexpected orderbook link %#v <%#v>", l.Rel, l.URI)
					}
				}
			}
		}
		data = schema
	}

	data, err = parseResponse(resp, data, r.Connection.strict)
	if err != nil {
		return nil, xerrors.Errorf("Get instrument order book failed: %w", err)
	}
	return data.(*GetInstrumentOrderBookSchema), nil
}

// GET /v3/instruments/{instrument}/positionBook
//
// Fetch a position book for an instrument.
func (r *ReceiverInstrumentPositionBook) Get(params *GetInstrumentPositionBookParams) (*GetInstrumentPositionBookSchema, error) {
	resp, err := r.Connection.request(
		&requestParams{
			method:   "GET",
			endPoint: "/v3/instruments/" + r.Instrument + "/positionBook",
			headers: []header{
				{key: "Accept-Datetime-Format", value: "RFC3339"},
			},
			queries: func() []query {
				q := make([]query, 0, 1)
				if params == nil {
					return q
				}

				// time
				if !params.Time.IsZero() {
					q = append(q, query{key: "time", value: params.Time.Format(time.RFC3339Nano)})
				}

				return q
			}(),
		},
	)
	if err != nil {
		return nil, xerrors.Errorf("Get instrument position book canceled: %w", err)
	}
	defer resp.Body.Close()

	var data interface{}
	switch resp.StatusCode {
	case 200:
		schema := new(GetInstrumentPositionBookSchema)
		if reqID, ok := resp.Header["Requestid"]; ok {
			schema.Headers.RequestID = reqID[0]
		}
		if links, ok := resp.Header["Link"]; ok {
			for _, l := range link.Parse(links[0]) {
				u, err := url.Parse(l.URI)
				if err != nil {
					return nil, xerrors.Errorf("Parse positionbook header failed: %w", err)
				}
				if refTimes, ok := u.Query()["time"]; ok {
					refTime, err := time.Parse(time.RFC3339, refTimes[0])
					if err != nil {
						return nil, xerrors.Errorf("Parse positionbook %#v time: %w", l.Rel, err)
					}
					switch l.Rel {
					case "prev":
						schema.Headers.Link.Prev = refTime
					case "next":
						schema.Headers.Link.Next = refTime
					default:
						return nil, xerrors.Errorf("Unexpected positionbook link %#v <%#v>", l.Rel, l.URI)
					}
				}
			}
		}
		data = schema
	}

	data, err = parseResponse(resp, data, r.Connection.strict)
	if err != nil {
		return nil, xerrors.Errorf("Get instrument position book failed: %w", err)
	}
	return data.(*GetInstrumentPositionBookSchema), nil
}

/* Utils */

func (s *GetInstrumentOrderBookSchema) PrevParams() *GetInstrumentOrderBookParams {
	return &GetInstrumentOrderBookParams{Time: s.Headers.Link.Prev}
}

func (s *GetInstrumentOrderBookSchema) NextParams() *GetInstrumentOrderBookParams {
	return &GetInstrumentOrderBookParams{Time: s.Headers.Link.Next}
}

func (s *GetInstrumentPositionBookSchema) PrevParams() *GetInstrumentPositionBookParams {
	return &GetInstrumentPositionBookParams{Time: s.Headers.Link.Prev}
}

func (s *GetInstrumentPositionBookSchema) NextParams() *GetInstrumentPositionBookParams {
	return &GetInstrumentPositionBookParams{Time: s.Headers.Link.Next}
}
