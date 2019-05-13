package pitOrgan

import (
	"golang.org/x/xerrors"
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

// Schemas

type GetInstrumentCandlesSchema struct {
	Instrument  InstrumentNameDefinition         `json:"instrument,omitempty"`
	Granularity CandlestickGranularityDefinition `json:"granularity,omitempty"`
	Candles     []*CandlestickDefinition         `json:"candles,omitempty"`
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

// TODO: GET /v3/instruments/{instrument}/orderBook

func (r *ReceiverInstrument) PositionBook() *ReceiverInstrumentPositionBook {
	return &ReceiverInstrumentPositionBook{
		Instrument: r.Instrument,
		Connection: r.Connection,
	}
}

// TODO: GET /v3/instruments/{instrument}/positionBook
