package pitOrgan

import (
	"github.com/peterhellberg/link"
	"golang.org/x/xerrors"
	"net/url"
	"strconv"
	"time"
)

// Receivers

type ReceiverInstruments struct {
	Connection *Connection
}

type ReceiverInstrument struct {
	Instrument string
	Connection *Connection
}

type ReceiverInstrumentCandles struct {
	Instrument string
	Connection *Connection
}

type ReceiverInstrumentOrderBook struct {
	Instrument string
	Connection *Connection
}

type ReceiverInstrumentPositionBook struct {
	Instrument string
	Connection *Connection
}

// Params

type GetInstrumentCandlesParams struct {
	PriceMid          interface{} // bool or nil
	PriceBid          interface{} // bool or nil
	PriceAsk          interface{} // bool or nil
	Granularity       CandlestickGranularityDefinition
	Count             int
	From              time.Time
	To                time.Time
	Smooth            interface{} // bool or nil(default=false)
	IncludeFirst      interface{} // bool or nil(default=true)
	DailyAlignment    interface{} // int or nil(default=17)
	AlignmentTimezone string
	WeeklyAlignment   WeeklyAlignmentDefinition
}

type GetInstrumentOrderBookParams struct {
	Time time.Time // The time of the snapshot to fetch. If not specified, then the most recent snapshot is fetched.
}

type GetInstrumentPositionBookParams struct {
	Time time.Time // The time of the snapshot to fetch. If not specified, then the most recent snapshot is fetched.
}

// Schemas

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

func (c *Connection) Instruments() *ReceiverInstruments {
	return &ReceiverInstruments{
		Connection: c,
	}
}

func (r *ReceiverInstruments) Instruments(i string) *ReceiverInstrument {
	return &ReceiverInstrument{
		Instrument: i,
		Connection: r.Connection,
	}
}

func (r *ReceiverInstrument) Candles() *ReceiverInstrumentCandles {
	return &ReceiverInstrumentCandles{
		Instrument: r.Instrument,
		Connection: r.Connection,
	}
}

// GET /v3/instruments/{instrument}/candles
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
				if b, ok := params.PriceMid.(bool); ok && b {
					price = append(price, 'M')
				}
				if b, ok := params.PriceBid.(bool); ok && b {
					price = append(price, 'B')
				}
				if b, ok := params.PriceAsk.(bool); ok && b {
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
				if b, ok := params.Smooth.(bool); ok {
					q = append(q, query{key: "smooth", value: strconv.FormatBool(b)})
				}

				// includeFirst
				if b, ok := params.IncludeFirst.(bool); ok {
					q = append(q, query{key: "includeFirst", value: strconv.FormatBool(b)})
				}

				// dailyAlignment
				if da, ok := params.DailyAlignment.(int); ok {
					q = append(q, query{key: "dailyAlignment", value: strconv.Itoa(da)})
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

	data, err = parseResponse(resp, data)
	if err != nil {
		return nil, xerrors.Errorf("Get instrument candles failed: %w", err)
	}
	return data.(*GetInstrumentCandlesSchema), nil
}

func (r *ReceiverInstrument) OrderBook() *ReceiverInstrumentOrderBook {
	return &ReceiverInstrumentOrderBook{
		Instrument: r.Instrument,
		Connection: r.Connection,
	}
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

	data, err = parseResponse(resp, data)
	if err != nil {
		return nil, xerrors.Errorf("Get instrument order book failed: %w", err)
	}
	return data.(*GetInstrumentOrderBookSchema), nil
}

func (r *ReceiverInstrument) PositionBook() *ReceiverInstrumentPositionBook {
	return &ReceiverInstrumentPositionBook{
		Instrument: r.Instrument,
		Connection: r.Connection,
	}
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

	data, err = parseResponse(resp, data)
	if err != nil {
		return nil, xerrors.Errorf("Get instrument position book failed: %w", err)
	}
	return data.(*GetInstrumentPositionBookSchema), nil
}

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
