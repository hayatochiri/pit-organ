package pitOrgan

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

// TODO: GET /v3/instruments/{instrument}/candles

// TODO: GET /v3/instruments/{instrument}/orderBook

// TODO: GET /v3/instruments/{instrument}/positionBook
