package pitOrgan

import (
	"github.com/davecgh/go-spew/spew"
	"testing"
	"time"
)

func Test_InstrumentCandles(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		connection := newConnection(t, OandaPractice)
		params := &GetInstrumentCandlesParams{}

		data, err := connection.Instruments().Instruments("EUR_USD").Candles().Get(params)
		if err != nil {
			t.Fatalf("Error occurred.\n%+v", err)
		}

		t.Logf("Response:\n%s", spew.Sdump(data))
	})

	successPatterns := []struct {
		name   string                                               // サブテスト SuccessPatterns の接尾辞
		params *GetInstrumentCandlesParams                          // APIリクエストする際の値
		verify func(*testing.T, *GetInstrumentCandlesSchema, error) // レスポンスを検証するコールバック
	}{
		{
			name:   "OneRecord",
			params: &GetInstrumentCandlesParams{Count: 1},
			verify: func(t *testing.T, s *GetInstrumentCandlesSchema, err error) {
				if expect := 1; len(s.Candles) != expect {
					t.Errorf("There are %d records, not %d records.\n", len(s.Candles), expect)
				}
			},
		},
		{
			name:   "BidAsk",
			params: &GetInstrumentCandlesParams{Count: 1, PriceBid: true, PriceAsk: true},
			verify: func(t *testing.T, s *GetInstrumentCandlesSchema, err error) {
				if s.Candles[0].Mid != nil {
					t.Error("The candlestick mid value is not nil.")
				}
				if s.Candles[0].Bid == nil {
					t.Error("The candlestick bid value is nil.")
				}
				if s.Candles[0].Ask == nil {
					t.Error("The candlestick ask value is nil.")
				}
			},
		},
	}

	for _, pattern := range successPatterns {
		t.Run("SuccessPatterns"+pattern.name, func(t *testing.T) {
			connection := newConnection(t, OandaPractice)

			data, err := connection.Instruments().Instruments("USD_JPY").Candles().Get(pattern.params)
			if err != nil {
				t.Fatalf("Error occurred.\n%+v", err)
			}
			pattern.verify(t, data, err)
			t.Logf("Response:\n%s", spew.Sdump(data))
		})
	}

}

func Test_InstrumentOrderBook(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		connection := newConnection(t, OandaPractice)
		params := &GetInstrumentOrderBookParams{}

		data, err := connection.Instruments().Instrument("USD_JPY").OrderBook().Get(params)
		if err != nil {
			t.Fatalf("Error occurred.\n%+v", err)
		}

		// TODO: 時刻(Prev/Next)の比較

		t.Logf("Response:\n%s", spew.Sdump(data))
	})

	t.Run("LinkParams", func(t *testing.T) {
		connection := newConnection(t, OandaPractice)
		refTime := time.Date(2018, time.Month(8), 14, 9, 0, 0, 0, time.UTC)
		params := &GetInstrumentOrderBookParams{Time: refTime}

		data, err := connection.Instruments().Instrument("USD_JPY").OrderBook().Get(params)
		if err != nil {
			t.Fatalf("Error occurred.\n%+v", err)
		}

		t.Logf("Response:\n%s", spew.Sdump(data.PrevParams()))
		t.Logf("Response:\n%s", spew.Sdump(data.NextParams()))
	})
}

func Test_InstrumentPositionBook(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		connection := newConnection(t, OandaPractice)
		params := &GetInstrumentPositionBookParams{}

		data, err := connection.Instruments().Instrument("USD_JPY").PositionBook().Get(params)
		if err != nil {
			t.Fatalf("Error occurred.\n%+v", err)
		}

		// TODO: 時刻(Prev/Next)の比較

		t.Logf("Response:\n%s", spew.Sdump(data))
	})

	t.Run("LinkParams", func(t *testing.T) {
		connection := newConnection(t, OandaPractice)
		refTime := time.Date(2018, time.Month(8), 14, 9, 0, 0, 0, time.UTC)
		params := &GetInstrumentPositionBookParams{Time: refTime}

		data, err := connection.Instruments().Instrument("USD_JPY").PositionBook().Get(params)
		if err != nil {
			t.Fatalf("Error occurred.\n%+v", err)
		}

		t.Logf("Response:\n%s", spew.Sdump(data.PrevParams()))
		t.Logf("Response:\n%s", spew.Sdump(data.NextParams()))
	})
}
