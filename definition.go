package pitOrgan

import (
	"encoding/json"
)

//
// Account Definitions
//

type AccountIDDefinition string

type AccountDefinition struct {
	ID                          AccountIDDefinition                   `json:"id,omitempty"`
	Alias                       string                                `json:"alias,omitempty"`
	Currency                    CurrencyDefinition                    `json:"currency,omitempty"`
	Balance                     AccountUnitsDefinition                `json:"balance,omitempty"`
	CreatedByUserID             int                                   `json:"createdByUserID,omitempty"`
	CreatedTime                 DateTimeDefinition                    `json:"createdTime,omitempty"`
	GuaranteedStopLossOrderMode GuaranteedStopLossOrderModeDefinition `json:"guaranteedStopLossOrderMode,omitempty"`
	PL                          AccountUnitsDefinition                `json:"pl,omitempty"`
	ResettablePL                AccountUnitsDefinition                `json:"resettablePL,omitempty"`
	ResettablePLTime            DateTimeDefinition                    `json:"resettablePLTime,omitempty"`
	Financing                   AccountUnitsDefinition                `json:"financing,omitempty"`
	Commission                  AccountUnitsDefinition                `json:"commission,omitempty"`
	GuaranteedExecutionFees     AccountUnitsDefinition                `json:"guaranteedExecutionFees,omitempty"`
	MarginRate                  DecimalNumberDefinition               `json:"marginRate,omitempty"`
	MarginCallEnterTime         DateTimeDefinition                    `json:"marginCallEnterTime,omitempty"`
	MarginCallExtensionCount    int                                   `json:"marginCallExtensionCount,omitempty"`
	LastMarginCallExtensionTime DateTimeDefinition                    `json:"lastMarginCallExtensionTime,omitempty"`
	OpenTradeCount              int                                   `json:"openTradeCount,omitempty"`
	OpenPositionCount           int                                   `json:"openPositionCount,omitempty"`
	PendingOrderCount           int                                   `json:"pendingOrderCount,omitempty"`
	HedgingEnabled              bool                                  `json:"hedgingEnabled,omitempty"`
	UnrealizedPL                AccountUnitsDefinition                `json:"unrealizedPL,omitempty"`
	NAV                         AccountUnitsDefinition                `json:"NAV,omitempty"`
	MarginUsed                  AccountUnitsDefinition                `json:"marginUsed,omitempty"`
	MarginAvailable             AccountUnitsDefinition                `json:"marginAvailable,omitempty"`
	PositionValue               AccountUnitsDefinition                `json:"positionValue,omitempty"`
	MarginCloseoutUnrealizedPL  AccountUnitsDefinition                `json:"marginCloseoutUnrealizedPL,omitempty"`
	MarginCloseoutNAV           AccountUnitsDefinition                `json:"marginCloseoutNAV,omitempty"`
	MarginCloseoutMarginUsed    AccountUnitsDefinition                `json:"marginCloseoutMarginUsed,omitempty"`
	MarginCloseoutPercent       DecimalNumberDefinition               `json:"marginCloseoutPercent,omitempty"`
	MarginCloseoutPositionValue DecimalNumberDefinition               `json:"marginCloseoutPositionValue,omitempty"`
	WithdrawalLimit             AccountUnitsDefinition                `json:"withdrawalLimit,omitempty"`
	MarginCallMarginUsed        AccountUnitsDefinition                `json:"marginCallMarginUsed,omitempty"`
	MarginCallPercent           DecimalNumberDefinition               `json:"marginCallPercent,omitempty"`
	LastTransactionID           TransactionIDDefinition               `json:"lastTransactionID,omitempty"`
	Trades                      []*TradeSummaryDefinition             `json:"trades,omitempty"`
	Positions                   []*PositionDefinition                 `json:"positions,omitempty"`
	Orders                      []*OrderDefinition                    `json:"orders,omitempty"`
}

type AccountChangesStateDefinition struct {
	UnrealizedPL                AccountUnitsDefinition               `json:"unrealizedPL,omitempty"`
	NAV                         AccountUnitsDefinition               `json:"NAV,omitempty"`
	MarginUsed                  AccountUnitsDefinition               `json:"marginUsed,omitempty"`
	MarginAvailable             AccountUnitsDefinition               `json:"marginAvailable,omitempty"`
	PositionValue               AccountUnitsDefinition               `json:"positionValue,omitempty"`
	MarginCloseoutUnrealizedPL  AccountUnitsDefinition               `json:"marginCloseoutUnrealizedPL,omitempty"`
	MarginCloseoutNAV           AccountUnitsDefinition               `json:"marginCloseoutNAV,omitempty"`
	MarginCloseoutMarginUsed    AccountUnitsDefinition               `json:"marginCloseoutMarginUsed,omitempty"`
	MarginCloseoutPercent       DecimalNumberDefinition              `json:"marginCloseoutPercent,omitempty"`
	MarginCloseoutPositionValue DecimalNumberDefinition              `json:"marginCloseoutPositionValue,omitempty"`
	WithdrawalLimit             AccountUnitsDefinition               `json:"withdrawalLimit,omitempty"`
	MarginCallMarginUsed        AccountUnitsDefinition               `json:"marginCallMarginUsed,omitempty"`
	MarginCallPercent           DecimalNumberDefinition              `json:"marginCallPercent,omitempty"`
	Orders                      []*DynamicOrderStateDefinition       `json:"orders,omitempty"`
	Trades                      []*CalculatedTradeStateDefinition    `json:"trades,omitempty"`
	Positions                   []*CalculatedPositionStateDefinition `json:"positions,omitempty"`
}

type AccountPropertiesDefinition struct {
	ID           AccountIDDefinition `json:"id,omitempty"`
	MT4AccountID int                 `json:"mt4AccountID,omitempty"`
	Tags         []string            `json:"tags,omitempty"`
}

type GuaranteedStopLossOrderModeDefinition string

type AccountSummaryDefinition struct {
	ID                          AccountIDDefinition                   `json:"id,omitempty"`
	Alias                       string                                `json:"alias,omitempty"`
	Currency                    CurrencyDefinition                    `json:"currency,omitempty"`
	Balance                     AccountUnitsDefinition                `json:"balance,omitempty"`
	CreatedByUserID             int                                   `json:"createdByUserID,omitempty"`
	CreatedTime                 DateTimeDefinition                    `json:"createdTime,omitempty"`
	GuaranteedStopLossOrderMode GuaranteedStopLossOrderModeDefinition `json:"guaranteedStopLossOrderMode,omitempty"`
	PL                          AccountUnitsDefinition                `json:"pl,omitempty"`
	ResettablePL                AccountUnitsDefinition                `json:"resettablePL,omitempty"`
	ResettablePLTime            DateTimeDefinition                    `json:"resettablePLTime,omitempty"`
	Financing                   AccountUnitsDefinition                `json:"financing,omitempty"`
	Commission                  AccountUnitsDefinition                `json:"commission,omitempty"`
	GuaranteedExecutionFees     AccountUnitsDefinition                `json:"guaranteedExecutionFees,omitempty"`
	MarginRate                  DecimalNumberDefinition               `json:"marginRate,omitempty"`
	MarginCallEnterTime         DateTimeDefinition                    `json:"marginCallEnterTime,omitempty"`
	MarginCallExtensionCount    int                                   `json:"marginCallExtensionCount,omitempty"`
	LastMarginCallExtensionTime DateTimeDefinition                    `json:"lastMarginCallExtensionTime,omitempty"`
	OpenTradeCount              int                                   `json:"openTradeCount,omitempty"`
	OpenPositionCount           int                                   `json:"openPositionCount,omitempty"`
	PendingOrderCount           int                                   `json:"pendingOrderCount,omitempty"`
	HedgingEnabled              bool                                  `json:"hedgingEnabled,omitempty"`
	UnrealizedPL                AccountUnitsDefinition                `json:"unrealizedPL,omitempty"`
	NAV                         AccountUnitsDefinition                `json:"NAV,omitempty"`
	MarginUsed                  AccountUnitsDefinition                `json:"marginUsed,omitempty"`
	MarginAvailable             AccountUnitsDefinition                `json:"marginAvailable,omitempty"`
	PositionValue               AccountUnitsDefinition                `json:"positionValue,omitempty"`
	MarginCloseoutUnrealizedPL  AccountUnitsDefinition                `json:"marginCloseoutUnrealizedPL,omitempty"`
	MarginCloseoutNAV           AccountUnitsDefinition                `json:"marginCloseoutNAV,omitempty"`
	MarginCloseoutMarginUsed    AccountUnitsDefinition                `json:"marginCloseoutMarginUsed,omitempty"`
	MarginCloseoutPercent       DecimalNumberDefinition               `json:"marginCloseoutPercent,omitempty"`
	MarginCloseoutPositionValue DecimalNumberDefinition               `json:"marginCloseoutPositionValue,omitempty"`
	WithdrawalLimit             AccountUnitsDefinition                `json:"withdrawalLimit,omitempty"`
	MarginCallMarginUsed        AccountUnitsDefinition                `json:"marginCallMarginUsed,omitempty"`
	MarginCallPercent           DecimalNumberDefinition               `json:"marginCallPercent,omitempty"`
	LastTransactionID           TransactionIDDefinition               `json:"lastTransactionID,omitempty"`
}

type CalculatedAccountStateDefinition struct {
	UnrealizedPL                AccountUnitsDefinition  `json:"unrealizedPL,omitempty"`
	NAV                         AccountUnitsDefinition  `json:"NAV,omitempty"`
	MarginUsed                  AccountUnitsDefinition  `json:"marginUsed,omitempty"`
	MarginAvailable             AccountUnitsDefinition  `json:"marginAvailable,omitempty"`
	PositionValue               AccountUnitsDefinition  `json:"positionValue,omitempty"`
	MarginCloseoutUnrealizedPL  AccountUnitsDefinition  `json:"marginCloseoutUnrealizedPL,omitempty"`
	MarginCloseoutNAV           AccountUnitsDefinition  `json:"marginCloseoutNAV,omitempty"`
	MarginCloseoutMarginUsed    AccountUnitsDefinition  `json:"marginCloseoutMarginUsed,omitempty"`
	MarginCloseoutPercent       DecimalNumberDefinition `json:"marginCloseoutPercent,omitempty"`
	MarginCloseoutPositionValue DecimalNumberDefinition `json:"marginCloseoutPositionValue,omitempty"`
	WithdrawalLimit             AccountUnitsDefinition  `json:"withdrawalLimit,omitempty"`
	MarginCallMarginUsed        AccountUnitsDefinition  `json:"marginCallMarginUsed,omitempty"`
	MarginCallPercent           DecimalNumberDefinition `json:"marginCallPercent,omitempty"`
}

type AccountChangesDefinition struct {
	OrdersCreated   []*OrderDefinition        `json:"ordersCreated,omitempty"`
	OrdersCancelled []*OrderDefinition        `json:"ordersCancelled,omitempty"`
	OrdersFilled    []*OrderDefinition        `json:"ordersFilled,omitempty"`
	OrdersTriggered []*OrderDefinition        `json:"ordersTriggered,omitempty"`
	TradesOpened    []*TradeSummaryDefinition `json:"tradesOpened,omitempty"`
	TradesReduced   []*TradeSummaryDefinition `json:"tradesReduced,omitempty"`
	TradesClosed    []*TradeSummaryDefinition `json:"tradesClosed,omitempty"`
	Positions       []*PositionDefinition     `json:"positions,omitempty"`
	Transactions    []*TransactionDefinition  `json:"transactions,omitempty"`
}

type AccountFinancingModeDefinition string

type PositionAggregationModeDefinition string

//
// Instrument Definitions
//

type CandlestickGranularityDefinition string

const (
	S5  CandlestickGranularityDefinition = "S5"  // 5 second candlesticks, minute alignment
	S10 CandlestickGranularityDefinition = "S10" // 10 second candlesticks, minute alignment
	S15 CandlestickGranularityDefinition = "S15" // 15 second candlesticks, minute alignment
	S30 CandlestickGranularityDefinition = "S30" // 30 second candlesticks, minute alignment
	M1  CandlestickGranularityDefinition = "M1"  // 1 minute candlesticks, minute alignment
	M2  CandlestickGranularityDefinition = "M2"  // 2 minute candlesticks, hour alignment
	M4  CandlestickGranularityDefinition = "M4"  // 4 minute candlesticks, hour alignment
	M5  CandlestickGranularityDefinition = "M5"  // 5 minute candlesticks, hour alignment
	M10 CandlestickGranularityDefinition = "M10" // 10 minute candlesticks, hour alignment
	M15 CandlestickGranularityDefinition = "M15" // 15 minute candlesticks, hour alignment
	M30 CandlestickGranularityDefinition = "M30" // 30 minute candlesticks, hour alignment
	H1  CandlestickGranularityDefinition = "H1"  // 1 hour candlesticks, hour alignment
	H2  CandlestickGranularityDefinition = "H2"  // 2 hour candlesticks, day alignment
	H3  CandlestickGranularityDefinition = "H3"  // 3 hour candlesticks, day alignment
	H4  CandlestickGranularityDefinition = "H4"  // 4 hour candlesticks, day alignment
	H6  CandlestickGranularityDefinition = "H6"  // 6 hour candlesticks, day alignment
	H8  CandlestickGranularityDefinition = "H8"  // 8 hour candlesticks, day alignment
	H12 CandlestickGranularityDefinition = "H12" // 12 hour candlesticks, day alignment
	D   CandlestickGranularityDefinition = "D"   // 1 day candlesticks, day alignment
	W   CandlestickGranularityDefinition = "W"   // 1 week candlesticks, aligned to start of week
	M   CandlestickGranularityDefinition = "M"   // 1 month candlesticks, aligned to first day of the month
)

type WeeklyAlignmentDefinition string

const (
	Monday    WeeklyAlignmentDefinition = "Monday"
	Tuesday   WeeklyAlignmentDefinition = "Tuesday"
	Wednesday WeeklyAlignmentDefinition = "Wednesday"
	Thursday  WeeklyAlignmentDefinition = "Thursday"
	Friday    WeeklyAlignmentDefinition = "Friday"
	Saturday  WeeklyAlignmentDefinition = "Saturday"
	Sunday    WeeklyAlignmentDefinition = "Sunday"
)

type CandlestickDefinition struct {
	Time     DateTimeDefinition         `json:"time,omitempty"`
	Bid      *CandlestickDataDefinition `json:"bid,omitempty"`
	Ask      *CandlestickDataDefinition `json:"ask,omitempty"`
	Mid      *CandlestickDataDefinition `json:"mid,omitempty"`
	Volume   int                        `json:"volume,omitempty"`
	Complete bool                       `json:"complete,omitempty"`
}

type CandlestickDataDefinition struct {
	O PriceValueDefinition `json:"o,omitempty"`
	H PriceValueDefinition `json:"h,omitempty"`
	L PriceValueDefinition `json:"l,omitempty"`
	C PriceValueDefinition `json:"c,omitempty"`
}

type OrderBookDefinition struct {
	Instrument  InstrumentNameDefinition     `json:"instrument,omitempty"`
	Time        DateTimeDefinition           `json:"time,omitempty"`
	Price       PriceValueDefinition         `json:"price,omitempty"`
	BucketWidth PriceValueDefinition         `json:"bucketWidth,omitempty"`
	Buckets     []*OrderBookBucketDefinition `json:"buckets,omitempty"`
}

type OrderBookBucketDefinition struct {
	Price             PriceValueDefinition    `json:"price,omitempty"`
	LongCountPercent  DecimalNumberDefinition `json:"longCountPercent,omitempty"`
	ShortCountPercent DecimalNumberDefinition `json:"shortCountPercent,omitempty"`
}

type PositionBookDefinition struct {
	Instrument  InstrumentNameDefinition        `json:"instrument,omitempty"`
	Time        DateTimeDefinition              `json:"time,omitempty"`
	Price       PriceValueDefinition            `json:"price,omitempty"`
	BucketWidth PriceValueDefinition            `json:"bucketWidth,omitempty"`
	Buckets     []*PositionBookBucketDefinition `json:"buckets,omitempty"`
}

type PositionBookBucketDefinition struct {
	Price             PriceValueDefinition    `json:"price,omitempty"`
	LongCountPercent  DecimalNumberDefinition `json:"longCountPercent,omitempty"`
	ShortCountPercent DecimalNumberDefinition `json:"shortCountPercent,omitempty"`
}

//
// Order Definitions
//

// Orders

type OrderDefinition struct {
	ID               OrderIDDefinition           `json:"id,omitempty"`
	CreateTime       DateTimeDefinition          `json:"createTime,omitempty"`
	State            OrderStateDefinition        `json:"state,omitempty"`
	ClientExtensions *ClientExtensionsDefinition `json:"clientExtensions,omitempty"`
}

type MarketOrderDefinition struct {
	ID                      OrderIDDefinition                       `json:"id,omitempty"`
	CreateTime              DateTimeDefinition                      `json:"createTime,omitempty"`
	State                   OrderStateDefinition                    `json:"state,omitempty"`
	ClientExtensions        *ClientExtensionsDefinition             `json:"clientExtensions,omitempty"`
	Type                    OrderTypeDefinition                     `json:"type,omitempty"`
	Instrument              InstrumentNameDefinition                `json:"instrument,omitempty"`
	Units                   DecimalNumberDefinition                 `json:"units,omitempty"`
	TimeInForce             TimeInForceDefinition                   `json:"timeInForce,omitempty"`
	PriceBound              PriceValueDefinition                    `json:"priceBound,omitempty"`
	PositionFill            OrderPositionFillDefinition             `json:"positionFill,omitempty"`
	TradeClose              *MarketOrderTradeCloseDefinition        `json:"tradeClose,omitempty"`
	LongPositionCloseout    *MarketOrderPositionCloseoutDefinition  `json:"longPositionCloseout,omitempty"`
	ShortPositionCloseout   *MarketOrderPositionCloseoutDefinition  `json:"shortPositionCloseout,omitempty"`
	MarginCloseout          *MarketOrderMarginCloseoutDefinition    `json:"marginCloseout,omitempty"`
	DelayedTradeClose       *MarketOrderDelayedTradeCloseDefinition `json:"delayedTradeClose,omitempty"`
	TakeProfitOnFill        *TakeProfitDetailsDefinition            `json:"takeProfitOnFill,omitempty"`
	StopLossOnFill          *StopLossDetailsDefinition              `json:"stopLossOnFill,omitempty"`
	TrailingStopLossOnFill  *TrailingStopLossDetailsDefinition      `json:"trailingStopLossOnFill,omitempty"`
	TradeClientExtensions   *ClientExtensionsDefinition             `json:"tradeClientExtensions,omitempty"`
	FillingTransactionID    TransactionIDDefinition                 `json:"fillingTransactionID,omitempty"`
	FilledTime              DateTimeDefinition                      `json:"filledTime,omitempty"`
	TradeOpenedID           TradeIDDefinition                       `json:"tradeOpenedID,omitempty"`
	TradeReducedID          TradeIDDefinition                       `json:"tradeReducedID,omitempty"`
	TradeClosedIDs          []TradeIDDefinition                     `json:"tradeClosedIDs,omitempty"`
	CancellingTransactionID TransactionIDDefinition                 `json:"cancellingTransactionID,omitempty"`
	CancelledTime           DateTimeDefinition                      `json:"cancelledTime,omitempty"`
}

type FixedPriceOrderDefinition struct {
	ID                      OrderIDDefinition                  `json:"id,omitempty"`
	CreateTime              DateTimeDefinition                 `json:"createTime,omitempty"`
	State                   OrderStateDefinition               `json:"state,omitempty"`
	ClientExtensions        *ClientExtensionsDefinition        `json:"clientExtensions,omitempty"`
	Type                    OrderTypeDefinition                `json:"type,omitempty"`
	Instrument              InstrumentNameDefinition           `json:"instrument,omitempty"`
	Units                   DecimalNumberDefinition            `json:"units,omitempty"`
	Price                   PriceValueDefinition               `json:"price,omitempty"`
	PositionFill            OrderPositionFillDefinition        `json:"positionFill,omitempty"`
	TradeState              string                             `json:"tradeState,omitempty"`
	TakeProfitOnFill        *TakeProfitDetailsDefinition       `json:"takeProfitOnFill,omitempty"`
	StopLossOnFill          *StopLossDetailsDefinition         `json:"stopLossOnFill,omitempty"`
	TrailingStopLossOnFill  *TrailingStopLossDetailsDefinition `json:"trailingStopLossOnFill,omitempty"`
	TradeClientExtensions   *ClientExtensionsDefinition        `json:"tradeClientExtensions,omitempty"`
	FillingTransactionID    TransactionIDDefinition            `json:"fillingTransactionID,omitempty"`
	FilledTime              DateTimeDefinition                 `json:"filledTime,omitempty"`
	TradeOpenedID           TradeIDDefinition                  `json:"tradeOpenedID,omitempty"`
	TradeReducedID          TradeIDDefinition                  `json:"tradeReducedID,omitempty"`
	TradeClosedIDs          []TradeIDDefinition                `json:"tradeClosedIDs,omitempty"`
	CancellingTransactionID TransactionIDDefinition            `json:"cancellingTransactionID,omitempty"`
	CancelledTime           DateTimeDefinition                 `json:"cancelledTime,omitempty"`
}

type LimitOrderDefinition struct {
	ID                      OrderIDDefinition                  `json:"id,omitempty"`
	CreateTime              DateTimeDefinition                 `json:"createTime,omitempty"`
	State                   OrderStateDefinition               `json:"state,omitempty"`
	ClientExtensions        *ClientExtensionsDefinition        `json:"clientExtensions,omitempty"`
	Type                    OrderTypeDefinition                `json:"type,omitempty"`
	Instrument              InstrumentNameDefinition           `json:"instrument,omitempty"`
	Units                   DecimalNumberDefinition            `json:"units,omitempty"`
	Price                   PriceValueDefinition               `json:"price,omitempty"`
	TimeInForce             TimeInForceDefinition              `json:"timeInForce,omitempty"`
	GtdTime                 DateTimeDefinition                 `json:"gtdTime,omitempty"`
	PositionFill            OrderPositionFillDefinition        `json:"positionFill,omitempty"`
	TriggerCondition        OrderTriggerConditionDefinition    `json:"triggerCondition,omitempty"`
	TakeProfitOnFill        *TakeProfitDetailsDefinition       `json:"takeProfitOnFill,omitempty"`
	StopLossOnFill          *StopLossDetailsDefinition         `json:"stopLossOnFill,omitempty"`
	TrailingStopLossOnFill  *TrailingStopLossDetailsDefinition `json:"trailingStopLossOnFill,omitempty"`
	TradeClientExtensions   *ClientExtensionsDefinition        `json:"tradeClientExtensions,omitempty"`
	FillingTransactionID    TransactionIDDefinition            `json:"fillingTransactionID,omitempty"`
	FilledTime              DateTimeDefinition                 `json:"filledTime,omitempty"`
	TradeOpenedID           TradeIDDefinition                  `json:"tradeOpenedID,omitempty"`
	TradeReducedID          TradeIDDefinition                  `json:"tradeReducedID,omitempty"`
	TradeClosedIDs          []TradeIDDefinition                `json:"tradeClosedIDs,omitempty"`
	CancellingTransactionID TransactionIDDefinition            `json:"cancellingTransactionID,omitempty"`
	CancelledTime           DateTimeDefinition                 `json:"cancelledTime,omitempty"`
	ReplacesOrderID         OrderIDDefinition                  `json:"replacesOrderID,omitempty"`
	ReplacedByOrderID       OrderIDDefinition                  `json:"replacedByOrderID,omitempty"`
}

type StopOrderDefinition struct {
	ID                      OrderIDDefinition                  `json:"id,omitempty"`
	CreateTime              DateTimeDefinition                 `json:"createTime,omitempty"`
	State                   OrderStateDefinition               `json:"state,omitempty"`
	ClientExtensions        *ClientExtensionsDefinition        `json:"clientExtensions,omitempty"`
	Type                    OrderTypeDefinition                `json:"type,omitempty"`
	Instrument              InstrumentNameDefinition           `json:"instrument,omitempty"`
	Units                   DecimalNumberDefinition            `json:"units,omitempty"`
	Price                   PriceValueDefinition               `json:"price,omitempty"`
	PriceBound              PriceValueDefinition               `json:"priceBound,omitempty"`
	TimeInForce             TimeInForceDefinition              `json:"timeInForce,omitempty"`
	GtdTime                 DateTimeDefinition                 `json:"gtdTime,omitempty"`
	PositionFill            OrderPositionFillDefinition        `json:"positionFill,omitempty"`
	TriggerCondition        OrderTriggerConditionDefinition    `json:"triggerCondition,omitempty"`
	TakeProfitOnFill        *TakeProfitDetailsDefinition       `json:"takeProfitOnFill,omitempty"`
	StopLossOnFill          *StopLossDetailsDefinition         `json:"stopLossOnFill,omitempty"`
	TrailingStopLossOnFill  *TrailingStopLossDetailsDefinition `json:"trailingStopLossOnFill,omitempty"`
	TradeClientExtensions   *ClientExtensionsDefinition        `json:"tradeClientExtensions,omitempty"`
	FillingTransactionID    TransactionIDDefinition            `json:"fillingTransactionID,omitempty"`
	FilledTime              DateTimeDefinition                 `json:"filledTime,omitempty"`
	TradeOpenedID           TradeIDDefinition                  `json:"tradeOpenedID,omitempty"`
	TradeReducedID          TradeIDDefinition                  `json:"tradeReducedID,omitempty"`
	TradeClosedIDs          []TradeIDDefinition                `json:"tradeClosedIDs,omitempty"`
	CancellingTransactionID TransactionIDDefinition            `json:"cancellingTransactionID,omitempty"`
	CancelledTime           DateTimeDefinition                 `json:"cancelledTime,omitempty"`
	ReplacesOrderID         OrderIDDefinition                  `json:"replacesOrderID,omitempty"`
	ReplacedByOrderID       OrderIDDefinition                  `json:"replacedByOrderID,omitempty"`
}

type MarketIfTouchedOrderDefinition struct {
	ID                      OrderIDDefinition                  `json:"id,omitempty"`
	CreateTime              DateTimeDefinition                 `json:"createTime,omitempty"`
	State                   OrderStateDefinition               `json:"state,omitempty"`
	ClientExtensions        *ClientExtensionsDefinition        `json:"clientExtensions,omitempty"`
	Type                    OrderTypeDefinition                `json:"type,omitempty"`
	Instrument              InstrumentNameDefinition           `json:"instrument,omitempty"`
	Units                   DecimalNumberDefinition            `json:"units,omitempty"`
	Price                   PriceValueDefinition               `json:"price,omitempty"`
	PriceBound              PriceValueDefinition               `json:"priceBound,omitempty"`
	TimeInForce             TimeInForceDefinition              `json:"timeInForce,omitempty"`
	GtdTime                 DateTimeDefinition                 `json:"gtdTime,omitempty"`
	PositionFill            OrderPositionFillDefinition        `json:"positionFill,omitempty"`
	TriggerCondition        OrderTriggerConditionDefinition    `json:"triggerCondition,omitempty"`
	InitialMarketPrice      PriceValueDefinition               `json:"initialMarketPrice,omitempty"`
	TakeProfitOnFill        *TakeProfitDetailsDefinition       `json:"takeProfitOnFill,omitempty"`
	StopLossOnFill          *StopLossDetailsDefinition         `json:"stopLossOnFill,omitempty"`
	TrailingStopLossOnFill  *TrailingStopLossDetailsDefinition `json:"trailingStopLossOnFill,omitempty"`
	TradeClientExtensions   *ClientExtensionsDefinition        `json:"tradeClientExtensions,omitempty"`
	FillingTransactionID    TransactionIDDefinition            `json:"fillingTransactionID,omitempty"`
	FilledTime              DateTimeDefinition                 `json:"filledTime,omitempty"`
	TradeOpenedID           TradeIDDefinition                  `json:"tradeOpenedID,omitempty"`
	TradeReducedID          TradeIDDefinition                  `json:"tradeReducedID,omitempty"`
	TradeClosedIDs          []TradeIDDefinition                `json:"tradeClosedIDs,omitempty"`
	CancellingTransactionID TransactionIDDefinition            `json:"cancellingTransactionID,omitempty"`
	CancelledTime           DateTimeDefinition                 `json:"cancelledTime,omitempty"`
	ReplacesOrderID         OrderIDDefinition                  `json:"replacesOrderID,omitempty"`
	ReplacedByOrderID       OrderIDDefinition                  `json:"replacedByOrderID,omitempty"`
}

type TakeProfitOrderDefinition struct {
	ID                      OrderIDDefinition               `json:"id,omitempty"`
	CreateTime              DateTimeDefinition              `json:"createTime,omitempty"`
	State                   OrderStateDefinition            `json:"state,omitempty"`
	ClientExtensions        *ClientExtensionsDefinition     `json:"clientExtensions,omitempty"`
	Type                    OrderTypeDefinition             `json:"type,omitempty"`
	TradeID                 TradeIDDefinition               `json:"tradeID,omitempty"`
	ClientTradeID           ClientIDDefinition              `json:"clientTradeID,omitempty"`
	Price                   PriceValueDefinition            `json:"price,omitempty"`
	TimeInForce             TimeInForceDefinition           `json:"timeInForce,omitempty"`
	GtdTime                 DateTimeDefinition              `json:"gtdTime,omitempty"`
	TriggerCondition        OrderTriggerConditionDefinition `json:"triggerCondition,omitempty"`
	FillingTransactionID    TransactionIDDefinition         `json:"fillingTransactionID,omitempty"`
	FilledTime              DateTimeDefinition              `json:"filledTime,omitempty"`
	TradeOpenedID           TradeIDDefinition               `json:"tradeOpenedID,omitempty"`
	TradeReducedID          TradeIDDefinition               `json:"tradeReducedID,omitempty"`
	TradeClosedIDs          []TradeIDDefinition             `json:"tradeClosedIDs,omitempty"`
	CancellingTransactionID TransactionIDDefinition         `json:"cancellingTransactionID,omitempty"`
	CancelledTime           DateTimeDefinition              `json:"cancelledTime,omitempty"`
	ReplacesOrderID         OrderIDDefinition               `json:"replacesOrderID,omitempty"`
	ReplacedByOrderID       OrderIDDefinition               `json:"replacedByOrderID,omitempty"`
}

type StopLossOrderDefinition struct {
	ID                         OrderIDDefinition               `json:"id,omitempty"`
	CreateTime                 DateTimeDefinition              `json:"createTime,omitempty"`
	State                      OrderStateDefinition            `json:"state,omitempty"`
	ClientExtensions           *ClientExtensionsDefinition     `json:"clientExtensions,omitempty"`
	Type                       OrderTypeDefinition             `json:"type,omitempty"`
	GuaranteedExecutionPremium DecimalNumberDefinition         `json:"guaranteedExecutionPremium,omitempty"`
	TradeID                    TradeIDDefinition               `json:"tradeID,omitempty"`
	ClientTradeID              ClientIDDefinition              `json:"clientTradeID,omitempty"`
	Price                      PriceValueDefinition            `json:"price,omitempty"`
	Distance                   DecimalNumberDefinition         `json:"distance,omitempty"`
	TimeInForce                TimeInForceDefinition           `json:"timeInForce,omitempty"`
	GtdTime                    DateTimeDefinition              `json:"gtdTime,omitempty"`
	TriggerCondition           OrderTriggerConditionDefinition `json:"triggerCondition,omitempty"`
	Guaranteed                 bool                            `json:"guaranteed,omitempty"`
	FillingTransactionID       TransactionIDDefinition         `json:"fillingTransactionID,omitempty"`
	FilledTime                 DateTimeDefinition              `json:"filledTime,omitempty"`
	TradeOpenedID              TradeIDDefinition               `json:"tradeOpenedID,omitempty"`
	TradeReducedID             TradeIDDefinition               `json:"tradeReducedID,omitempty"`
	TradeClosedIDs             []TradeIDDefinition             `json:"tradeClosedIDs,omitempty"`
	CancellingTransactionID    TransactionIDDefinition         `json:"cancellingTransactionID,omitempty"`
	CancelledTime              DateTimeDefinition              `json:"cancelledTime,omitempty"`
	ReplacesOrderID            OrderIDDefinition               `json:"replacesOrderID,omitempty"`
	ReplacedByOrderID          OrderIDDefinition               `json:"replacedByOrderID,omitempty"`
}

type TrailingStopLossOrderDefinition struct {
	ID                      OrderIDDefinition               `json:"id,omitempty"`
	CreateTime              DateTimeDefinition              `json:"createTime,omitempty"`
	State                   OrderStateDefinition            `json:"state,omitempty"`
	ClientExtensions        *ClientExtensionsDefinition     `json:"clientExtensions,omitempty"`
	Type                    OrderTypeDefinition             `json:"type,omitempty"`
	TradeID                 TradeIDDefinition               `json:"tradeID,omitempty"`
	ClientTradeID           ClientIDDefinition              `json:"clientTradeID,omitempty"`
	Distance                DecimalNumberDefinition         `json:"distance,omitempty"`
	TimeInForce             TimeInForceDefinition           `json:"timeInForce,omitempty"`
	GtdTime                 DateTimeDefinition              `json:"gtdTime,omitempty"`
	TriggerCondition        OrderTriggerConditionDefinition `json:"triggerCondition,omitempty"`
	TrailingStopValue       PriceValueDefinition            `json:"trailingStopValue,omitempty"`
	FillingTransactionID    TransactionIDDefinition         `json:"fillingTransactionID,omitempty"`
	FilledTime              DateTimeDefinition              `json:"filledTime,omitempty"`
	TradeOpenedID           TradeIDDefinition               `json:"tradeOpenedID,omitempty"`
	TradeReducedID          TradeIDDefinition               `json:"tradeReducedID,omitempty"`
	TradeClosedIDs          []TradeIDDefinition             `json:"tradeClosedIDs,omitempty"`
	CancellingTransactionID TransactionIDDefinition         `json:"cancellingTransactionID,omitempty"`
	CancelledTime           DateTimeDefinition              `json:"cancelledTime,omitempty"`
	ReplacesOrderID         OrderIDDefinition               `json:"replacesOrderID,omitempty"`
	ReplacedByOrderID       OrderIDDefinition               `json:"replacedByOrderID,omitempty"`
}

// Order Requests

// type OrderRequestDefinition
// TODO: Implemented by: MarketOrderRequest, LimitOrderRequest, StopOrderRequest, MarketIfTouchedOrderRequest, TakeProfitOrderRequest, StopLossOrderRequest, TrailingStopLossOrderRequest

type MarketOrderRequestDefinition struct {
	Type                   OrderTypeDefinition                `json:"type,omitempty"`
	Instrument             InstrumentNameDefinition           `json:"instrument,omitempty"`
	Units                  DecimalNumberDefinition            `json:"units,omitempty"`
	TimeInForce            TimeInForceDefinition              `json:"timeInForce,omitempty"`
	PriceBound             PriceValueDefinition               `json:"priceBound,omitempty"`
	PositionFill           OrderPositionFillDefinition        `json:"positionFill,omitempty"`
	ClientExtensions       *ClientExtensionsDefinition        `json:"clientExtensions,omitempty"`
	TakeProfitOnFill       *TakeProfitDetailsDefinition       `json:"takeProfitOnFill,omitempty"`
	StopLossOnFill         *StopLossDetailsDefinition         `json:"stopLossOnFill,omitempty"`
	TrailingStopLossOnFill *TrailingStopLossDetailsDefinition `json:"trailingStopLossOnFill,omitempty"`
	TradeClientExtensions  *ClientExtensionsDefinition        `json:"tradeClientExtensions,omitempty"`
}

type LimitOrderRequestDefinition struct {
	Type                   OrderTypeDefinition                `json:"type,omitempty"`
	Instrument             InstrumentNameDefinition           `json:"instrument,omitempty"`
	Units                  DecimalNumberDefinition            `json:"units,omitempty"`
	Price                  PriceValueDefinition               `json:"price,omitempty"`
	TimeInForce            TimeInForceDefinition              `json:"timeInForce,omitempty"`
	GtdTime                DateTimeDefinition                 `json:"gtdTime,omitempty"`
	PositionFill           OrderPositionFillDefinition        `json:"positionFill,omitempty"`
	TriggerCondition       OrderTriggerConditionDefinition    `json:"triggerCondition,omitempty"`
	ClientExtensions       *ClientExtensionsDefinition        `json:"clientExtensions,omitempty"`
	TakeProfitOnFill       *TakeProfitDetailsDefinition       `json:"takeProfitOnFill,omitempty"`
	StopLossOnFill         *StopLossDetailsDefinition         `json:"stopLossOnFill,omitempty"`
	TrailingStopLossOnFill *TrailingStopLossDetailsDefinition `json:"trailingStopLossOnFill,omitempty"`
	TradeClientExtensions  *ClientExtensionsDefinition        `json:"tradeClientExtensions,omitempty"`
}

type StopOrderRequestDefinition struct {
	Type                   OrderTypeDefinition                `json:"type,omitempty"`
	Instrument             InstrumentNameDefinition           `json:"instrument,omitempty"`
	Units                  DecimalNumberDefinition            `json:"units,omitempty"`
	Price                  PriceValueDefinition               `json:"price,omitempty"`
	PriceBound             PriceValueDefinition               `json:"priceBound,omitempty"`
	TimeInForce            TimeInForceDefinition              `json:"timeInForce,omitempty"`
	GtdTime                DateTimeDefinition                 `json:"gtdTime,omitempty"`
	PositionFill           OrderPositionFillDefinition        `json:"positionFill,omitempty"`
	TriggerCondition       OrderTriggerConditionDefinition    `json:"triggerCondition,omitempty"`
	ClientExtensions       *ClientExtensionsDefinition        `json:"clientExtensions,omitempty"`
	TakeProfitOnFill       *TakeProfitDetailsDefinition       `json:"takeProfitOnFill,omitempty"`
	StopLossOnFill         *StopLossDetailsDefinition         `json:"stopLossOnFill,omitempty"`
	TrailingStopLossOnFill *TrailingStopLossDetailsDefinition `json:"trailingStopLossOnFill,omitempty"`
	TradeClientExtensions  *ClientExtensionsDefinition        `json:"tradeClientExtensions,omitempty"`
}

type MarketIfTouchedOrderRequestDefinition struct {
	Type                   OrderTypeDefinition                `json:"type,omitempty"`
	Instrument             InstrumentNameDefinition           `json:"instrument,omitempty"`
	Units                  DecimalNumberDefinition            `json:"units,omitempty"`
	Price                  PriceValueDefinition               `json:"price,omitempty"`
	PriceBound             PriceValueDefinition               `json:"priceBound,omitempty"`
	TimeInForce            TimeInForceDefinition              `json:"timeInForce,omitempty"`
	GtdTime                DateTimeDefinition                 `json:"gtdTime,omitempty"`
	PositionFill           OrderPositionFillDefinition        `json:"positionFill,omitempty"`
	TriggerCondition       OrderTriggerConditionDefinition    `json:"triggerCondition,omitempty"`
	ClientExtensions       *ClientExtensionsDefinition        `json:"clientExtensions,omitempty"`
	TakeProfitOnFill       *TakeProfitDetailsDefinition       `json:"takeProfitOnFill,omitempty"`
	StopLossOnFill         *StopLossDetailsDefinition         `json:"stopLossOnFill,omitempty"`
	TrailingStopLossOnFill *TrailingStopLossDetailsDefinition `json:"trailingStopLossOnFill,omitempty"`
	TradeClientExtensions  *ClientExtensionsDefinition        `json:"tradeClientExtensions,omitempty"`
}

type TakeProfitOrderRequestDefinition struct {
	Type             OrderTypeDefinition             `json:"type,omitempty"`
	TradeID          TradeIDDefinition               `json:"tradeID,omitempty"`
	ClientTradeID    ClientIDDefinition              `json:"clientTradeID,omitempty"`
	Price            PriceValueDefinition            `json:"price,omitempty"`
	TimeInForce      TimeInForceDefinition           `json:"timeInForce,omitempty"`
	GtdTime          DateTimeDefinition              `json:"gtdTime,omitempty"`
	TriggerCondition OrderTriggerConditionDefinition `json:"triggerCondition,omitempty"`
	ClientExtensions *ClientExtensionsDefinition     `json:"clientExtensions,omitempty"`
}

type StopLossOrderRequestDefinition struct {
	Type             OrderTypeDefinition             `json:"type,omitempty"`
	TradeID          TradeIDDefinition               `json:"tradeID,omitempty"`
	ClientTradeID    ClientIDDefinition              `json:"clientTradeID,omitempty"`
	Price            PriceValueDefinition            `json:"price,omitempty"`
	Distance         DecimalNumberDefinition         `json:"distance,omitempty"`
	TimeInForce      TimeInForceDefinition           `json:"timeInForce,omitempty"`
	GtdTime          DateTimeDefinition              `json:"gtdTime,omitempty"`
	TriggerCondition OrderTriggerConditionDefinition `json:"triggerCondition,omitempty"`
	Guaranteed       bool                            `json:"guaranteed,omitempty"`
	ClientExtensions *ClientExtensionsDefinition     `json:"clientExtensions,omitempty"`
}

type TrailingStopLossOrderRequestDefinition struct {
	Type             OrderTypeDefinition             `json:"type,omitempty"`
	TradeID          TradeIDDefinition               `json:"tradeID,omitempty"`
	ClientTradeID    ClientIDDefinition              `json:"clientTradeID,omitempty"`
	Distance         DecimalNumberDefinition         `json:"distance,omitempty"`
	TimeInForce      TimeInForceDefinition           `json:"timeInForce,omitempty"`
	GtdTime          DateTimeDefinition              `json:"gtdTime,omitempty"`
	TriggerCondition OrderTriggerConditionDefinition `json:"triggerCondition,omitempty"`
	ClientExtensions *ClientExtensionsDefinition     `json:"clientExtensions,omitempty"`
}

// Order-related Definitions

type OrderIDDefinition string

type OrderTypeDefinition string

type CancellableOrderTypeDefinition string

type OrderStateDefinition string

type OrderStateFilterDefinition string

type OrderIdentifierDefinition struct {
	OrderID       OrderIDDefinition  `json:"orderID,omitempty"`
	ClientOrderID ClientIDDefinition `json:"clientOrderID,omitempty"`
}

type OrderSpecifierDefinition string

type TimeInForceDefinition string

type OrderPositionFillDefinition string

type OrderTriggerConditionDefinition string

type DynamicOrderStateDefinition struct {
	ID                     OrderIDDefinition    `json:"id,omitempty"`
	TrailingStopValue      PriceValueDefinition `json:"trailingStopValue,omitempty"`
	TriggerDistance        PriceValueDefinition `json:"triggerDistance,omitempty"`
	IsTriggerDistanceExact bool                 `json:"isTriggerDistanceExact,omitempty"`
}

type UnitsAvailableDetailsDefinition struct {
	Long  DecimalNumberDefinition `json:"long,omitempty"`
	Short DecimalNumberDefinition `json:"short,omitempty"`
}

type UnitsAvailableDefinition struct {
	Default     *UnitsAvailableDetailsDefinition `json:"default,omitempty"`
	ReduceFirst *UnitsAvailableDetailsDefinition `json:"reduceFirst,omitempty"`
	ReduceOnly  *UnitsAvailableDetailsDefinition `json:"reduceOnly,omitempty"`
	OpenOnly    *UnitsAvailableDetailsDefinition `json:"openOnly,omitempty"`
}

type GuaranteedStopLossOrderEntryDataDefinition struct {
	MinimumDistance  DecimalNumberDefinition                            `json:"minimumDistance,omitempty"`
	Premium          DecimalNumberDefinition                            `json:"premium,omitempty"`
	LevelRestriction *GuaranteedStopLossOrderLevelRestrictionDefinition `json:"levelRestriction,omitempty"`
}

//
// Trade Definitions
//

type TradeIDDefinition string

type TradeStateDefinition string

type TradeStateFilterDefinition string

type TradeSpecifierDefinition string

type TradeDefinition struct {
	ID                    TradeIDDefinition                `json:"id,omitempty"`
	Instrument            InstrumentNameDefinition         `json:"instrument,omitempty"`
	Price                 PriceValueDefinition             `json:"price,omitempty"`
	OpenTime              DateTimeDefinition               `json:"openTime,omitempty"`
	State                 TradeStateDefinition             `json:"state,omitempty"`
	InitialUnits          DecimalNumberDefinition          `json:"initialUnits,omitempty"`
	InitialMarginRequired AccountUnitsDefinition           `json:"initialMarginRequired,omitempty"`
	CurrentUnits          DecimalNumberDefinition          `json:"currentUnits,omitempty"`
	RealizedPL            AccountUnitsDefinition           `json:"realizedPL,omitempty"`
	UnrealizedPL          AccountUnitsDefinition           `json:"unrealizedPL,omitempty"`
	MarginUsed            AccountUnitsDefinition           `json:"marginUsed,omitempty"`
	AverageClosePrice     PriceValueDefinition             `json:"averageClosePrice,omitempty"`
	ClosingTransactionIDs []TransactionIDDefinition        `json:"closingTransactionIDs,omitempty"`
	Financing             AccountUnitsDefinition           `json:"financing,omitempty"`
	CloseTime             DateTimeDefinition               `json:"closeTime,omitempty"`
	ClientExtensions      *ClientExtensionsDefinition      `json:"clientExtensions,omitempty"`
	TakeProfitOrder       *TakeProfitOrderDefinition       `json:"takeProfitOrder,omitempty"`
	StopLossOrder         *StopLossOrderDefinition         `json:"stopLossOrder,omitempty"`
	TrailingStopLossOrder *TrailingStopLossOrderDefinition `json:"trailingStopLossOrder,omitempty"`
}

type TradeSummaryDefinition struct {
	ID                      TradeIDDefinition           `json:"id,omitempty"`
	Instrument              InstrumentNameDefinition    `json:"instrument,omitempty"`
	Price                   PriceValueDefinition        `json:"price,omitempty"`
	OpenTime                DateTimeDefinition          `json:"openTime,omitempty"`
	State                   TradeStateDefinition        `json:"state,omitempty"`
	InitialUnits            DecimalNumberDefinition     `json:"initialUnits,omitempty"`
	InitialMarginRequired   AccountUnitsDefinition      `json:"initialMarginRequired,omitempty"`
	CurrentUnits            DecimalNumberDefinition     `json:"currentUnits,omitempty"`
	RealizedPL              AccountUnitsDefinition      `json:"realizedPL,omitempty"`
	UnrealizedPL            AccountUnitsDefinition      `json:"unrealizedPL,omitempty"`
	MarginUsed              AccountUnitsDefinition      `json:"marginUsed,omitempty"`
	AverageClosePrice       PriceValueDefinition        `json:"averageClosePrice,omitempty"`
	ClosingTransactionIDs   []TransactionIDDefinition   `json:"closingTransactionIDs,omitempty"`
	Financing               AccountUnitsDefinition      `json:"financing,omitempty"`
	CloseTime               DateTimeDefinition          `json:"closeTime,omitempty"`
	ClientExtensions        *ClientExtensionsDefinition `json:"clientExtensions,omitempty"`
	TakeProfitOrderID       OrderIDDefinition           `json:"takeProfitOrderID,omitempty"`
	StopLossOrderID         OrderIDDefinition           `json:"stopLossOrderID,omitempty"`
	TrailingStopLossOrderID OrderIDDefinition           `json:"trailingStopLossOrderID,omitempty"`
}

type CalculatedTradeStateDefinition struct {
	ID           TradeIDDefinition      `json:"id,omitempty"`
	UnrealizedPL AccountUnitsDefinition `json:"unrealizedPL,omitempty"`
	MarginUsed   AccountUnitsDefinition `json:"marginUsed,omitempty"`
}

type TradePLDefinition string

//
// Position Definitions
//

type PositionDefinition struct {
	Instrument              InstrumentNameDefinition `json:"instrument,omitempty"`
	PL                      AccountUnitsDefinition   `json:"pl,omitempty"`
	UnrealizedPL            AccountUnitsDefinition   `json:"unrealizedPL,omitempty"`
	MarginUsed              AccountUnitsDefinition   `json:"marginUsed,omitempty"`
	ResettablePL            AccountUnitsDefinition   `json:"resettablePL,omitempty"`
	Financing               AccountUnitsDefinition   `json:"financing,omitempty"`
	Commission              AccountUnitsDefinition   `json:"commission,omitempty"`
	GuaranteedExecutionFees AccountUnitsDefinition   `json:"guaranteedExecutionFees,omitempty"`
	Long                    *PositionSideDefinition  `json:"long,omitempty"`
	Short                   *PositionSideDefinition  `json:"short,omitempty"`
}

type PositionSideDefinition struct {
	Units                   DecimalNumberDefinition `json:"units,omitempty"`
	AveragePrice            PriceValueDefinition    `json:"averagePrice,omitempty"`
	TradeIDs                []TradeIDDefinition     `json:"tradeIDs,omitempty"`
	PL                      AccountUnitsDefinition  `json:"pl,omitempty"`
	UnrealizedPL            AccountUnitsDefinition  `json:"unrealizedPL,omitempty"`
	ResettablePL            AccountUnitsDefinition  `json:"resettablePL,omitempty"`
	Financing               AccountUnitsDefinition  `json:"financing,omitempty"`
	GuaranteedExecutionFees AccountUnitsDefinition  `json:"guaranteedExecutionFees,omitempty"`
}

type CalculatedPositionStateDefinition struct {
	Instrument        InstrumentNameDefinition `json:"instrument,omitempty"`
	NetUnrealizedPL   AccountUnitsDefinition   `json:"netUnrealizedPL,omitempty"`
	LongUnrealizedPL  AccountUnitsDefinition   `json:"longUnrealizedPL,omitempty"`
	ShortUnrealizedPL AccountUnitsDefinition   `json:"shortUnrealizedPL,omitempty"`
	MarginUsed        AccountUnitsDefinition   `json:"marginUsed,omitempty"`
}

//
// Transaction Definitions
//

// Transactions

// TODO: Implemented by: OrderFillTransaction, OrderCancelTransaction, OrderCancelRejectTransaction, OrderClientExtensionsModifyTransaction, OrderClientExtensionsModifyRejectTransaction, CreateTransaction, CloseTransaction, ReopenTransaction, ClientConfigureTransaction, ClientConfigureRejectTransaction, TransferFundsTransaction, TransferFundsRejectTransaction, MarketOrderTransaction, MarketOrderRejectTransaction, FixedPriceOrderTransaction, LimitOrderTransaction, LimitOrderRejectTransaction, StopOrderTransaction, StopOrderRejectTransaction, MarketIfTouchedOrderTransaction, MarketIfTouchedOrderRejectTransaction, TakeProfitOrderTransaction, TakeProfitOrderRejectTransaction, StopLossOrderTransaction, StopLossOrderRejectTransaction, TrailingStopLossOrderTransaction, TrailingStopLossOrderRejectTransaction, TradeClientExtensionsModifyTransaction, TradeClientExtensionsModifyRejectTransaction, MarginCallEnterTransaction, MarginCallExtendTransaction, MarginCallExitTransaction, DelayedTradeClosureTransaction, DailyFinancingTransaction, ResetResettablePLTransaction
type TransactionDefinition struct {
	ID        TransactionIDDefinition   `json:"id,omitempty"`
	Time      DateTimeDefinition        `json:"time,omitempty"`
	UserID    int                       `json:"userID,omitempty"`
	AccountID AccountIDDefinition       `json:"accountID,omitempty"`
	BatchID   TransactionIDDefinition   `json:"batchID,omitempty"`
	RequestID RequestIDDefinition       `json:"requestID,omitempty"`
	Type      TransactionTypeDefinition `json:"type,omitempty"`
}

type CreateTransactionDefinition struct {
	ID            TransactionIDDefinition   `json:"id,omitempty"`
	Time          DateTimeDefinition        `json:"time,omitempty"`
	UserID        int                       `json:"userID,omitempty"`
	AccountID     AccountIDDefinition       `json:"accountID,omitempty"`
	BatchID       TransactionIDDefinition   `json:"batchID,omitempty"`
	RequestID     RequestIDDefinition       `json:"requestID,omitempty"`
	Type          TransactionTypeDefinition `json:"type,omitempty"`
	DivisionID    int                       `json:"divisionID,omitempty"`
	SiteID        int                       `json:"siteID,omitempty"`
	AccountUserID int                       `json:"accountUserID,omitempty"`
	AccountNumber int                       `json:"accountNumber,omitempty"`
	HomeCurrency  CurrencyDefinition        `json:"homeCurrency,omitempty"`
}

type CloseTransactionDefinition struct {
	ID        TransactionIDDefinition   `json:"id,omitempty"`
	Time      DateTimeDefinition        `json:"time,omitempty"`
	UserID    int                       `json:"userID,omitempty"`
	AccountID AccountIDDefinition       `json:"accountID,omitempty"`
	BatchID   TransactionIDDefinition   `json:"batchID,omitempty"`
	RequestID RequestIDDefinition       `json:"requestID,omitempty"`
	Type      TransactionTypeDefinition `json:"type,omitempty"`
}

type ReopenTransactionDefinition struct {
	ID        TransactionIDDefinition   `json:"id,omitempty"`
	Time      DateTimeDefinition        `json:"time,omitempty"`
	UserID    int                       `json:"userID,omitempty"`
	AccountID AccountIDDefinition       `json:"accountID,omitempty"`
	BatchID   TransactionIDDefinition   `json:"batchID,omitempty"`
	RequestID RequestIDDefinition       `json:"requestID,omitempty"`
	Type      TransactionTypeDefinition `json:"type,omitempty"`
}

type ClientConfigureTransactionDefinition struct {
	ID         TransactionIDDefinition   `json:"id,omitempty"`
	Time       DateTimeDefinition        `json:"time,omitempty"`
	UserID     int                       `json:"userID,omitempty"`
	AccountID  AccountIDDefinition       `json:"accountID,omitempty"`
	BatchID    TransactionIDDefinition   `json:"batchID,omitempty"`
	RequestID  RequestIDDefinition       `json:"requestID,omitempty"`
	Type       TransactionTypeDefinition `json:"type,omitempty"`
	Alias      string                    `json:"alias,omitempty"`
	MarginRate DecimalNumberDefinition   `json:"marginRate,omitempty"`
}

type ClientConfigureRejectTransactionDefinition struct {
	ID           TransactionIDDefinition           `json:"id,omitempty"`
	Time         DateTimeDefinition                `json:"time,omitempty"`
	UserID       int                               `json:"userID,omitempty"`
	AccountID    AccountIDDefinition               `json:"accountID,omitempty"`
	BatchID      TransactionIDDefinition           `json:"batchID,omitempty"`
	RequestID    RequestIDDefinition               `json:"requestID,omitempty"`
	Type         TransactionTypeDefinition         `json:"type,omitempty"`
	Alias        string                            `json:"alias,omitempty"`
	MarginRate   DecimalNumberDefinition           `json:"marginRate,omitempty"`
	RejectReason TransactionRejectReasonDefinition `json:"rejectReason,omitempty"`
}

type TransferFundsTransactionDefinition struct {
	ID             TransactionIDDefinition   `json:"id,omitempty"`
	Time           DateTimeDefinition        `json:"time,omitempty"`
	UserID         int                       `json:"userID,omitempty"`
	AccountID      AccountIDDefinition       `json:"accountID,omitempty"`
	BatchID        TransactionIDDefinition   `json:"batchID,omitempty"`
	RequestID      RequestIDDefinition       `json:"requestID,omitempty"`
	Type           TransactionTypeDefinition `json:"type,omitempty"`
	Amount         AccountUnitsDefinition    `json:"amount,omitempty"`
	FundingReason  FundingReasonDefinition   `json:"fundingReason,omitempty"`
	Comment        string                    `json:"comment,omitempty"`
	AccountBalance AccountUnitsDefinition    `json:"accountBalance,omitempty"`
}

type TransferFundsRejectTransactionDefinition struct {
	ID            TransactionIDDefinition           `json:"id,omitempty"`
	Time          DateTimeDefinition                `json:"time,omitempty"`
	UserID        int                               `json:"userID,omitempty"`
	AccountID     AccountIDDefinition               `json:"accountID,omitempty"`
	BatchID       TransactionIDDefinition           `json:"batchID,omitempty"`
	RequestID     RequestIDDefinition               `json:"requestID,omitempty"`
	Type          TransactionTypeDefinition         `json:"type,omitempty"`
	Amount        AccountUnitsDefinition            `json:"amount,omitempty"`
	FundingReason FundingReasonDefinition           `json:"fundingReason,omitempty"`
	Comment       string                            `json:"comment,omitempty"`
	RejectReason  TransactionRejectReasonDefinition `json:"rejectReason,omitempty"`
}

type MarketOrderTransactionDefinition struct {
	ID                     TransactionIDDefinition                 `json:"id,omitempty"`
	Time                   DateTimeDefinition                      `json:"time,omitempty"`
	UserID                 int                                     `json:"userID,omitempty"`
	AccountID              AccountIDDefinition                     `json:"accountID,omitempty"`
	BatchID                TransactionIDDefinition                 `json:"batchID,omitempty"`
	RequestID              RequestIDDefinition                     `json:"requestID,omitempty"`
	Type                   TransactionTypeDefinition               `json:"type,omitempty"`
	Instrument             InstrumentNameDefinition                `json:"instrument,omitempty"`
	Units                  DecimalNumberDefinition                 `json:"units,omitempty"`
	TimeInForce            TimeInForceDefinition                   `json:"timeInForce,omitempty"`
	PriceBound             PriceValueDefinition                    `json:"priceBound,omitempty"`
	PositionFill           OrderPositionFillDefinition             `json:"positionFill,omitempty"`
	TradeClose             *MarketOrderTradeCloseDefinition        `json:"tradeClose,omitempty"`
	LongPositionCloseout   *MarketOrderPositionCloseoutDefinition  `json:"longPositionCloseout,omitempty"`
	ShortPositionCloseout  *MarketOrderPositionCloseoutDefinition  `json:"shortPositionCloseout,omitempty"`
	MarginCloseout         *MarketOrderMarginCloseoutDefinition    `json:"marginCloseout,omitempty"`
	DelayedTradeClose      *MarketOrderDelayedTradeCloseDefinition `json:"delayedTradeClose,omitempty"`
	Reason                 MarketOrderReasonDefinition             `json:"reason,omitempty"`
	ClientExtensions       *ClientExtensionsDefinition             `json:"clientExtensions,omitempty"`
	TakeProfitOnFill       *TakeProfitDetailsDefinition            `json:"takeProfitOnFill,omitempty"`
	StopLossOnFill         *StopLossDetailsDefinition              `json:"stopLossOnFill,omitempty"`
	TrailingStopLossOnFill *TrailingStopLossDetailsDefinition      `json:"trailingStopLossOnFill,omitempty"`
	TradeClientExtensions  *ClientExtensionsDefinition             `json:"tradeClientExtensions,omitempty"`
}

type MarketOrderRejectTransactionDefinition struct {
	ID                     TransactionIDDefinition                 `json:"id,omitempty"`
	Time                   DateTimeDefinition                      `json:"time,omitempty"`
	UserID                 int                                     `json:"userID,omitempty"`
	AccountID              AccountIDDefinition                     `json:"accountID,omitempty"`
	BatchID                TransactionIDDefinition                 `json:"batchID,omitempty"`
	RequestID              RequestIDDefinition                     `json:"requestID,omitempty"`
	Type                   TransactionTypeDefinition               `json:"type,omitempty"`
	Instrument             InstrumentNameDefinition                `json:"instrument,omitempty"`
	Units                  DecimalNumberDefinition                 `json:"units,omitempty"`
	TimeInForce            TimeInForceDefinition                   `json:"timeInForce,omitempty"`
	PriceBound             PriceValueDefinition                    `json:"priceBound,omitempty"`
	PositionFill           OrderPositionFillDefinition             `json:"positionFill,omitempty"`
	TradeClose             *MarketOrderTradeCloseDefinition        `json:"tradeClose,omitempty"`
	LongPositionCloseout   *MarketOrderPositionCloseoutDefinition  `json:"longPositionCloseout,omitempty"`
	ShortPositionCloseout  *MarketOrderPositionCloseoutDefinition  `json:"shortPositionCloseout,omitempty"`
	MarginCloseout         *MarketOrderMarginCloseoutDefinition    `json:"marginCloseout,omitempty"`
	DelayedTradeClose      *MarketOrderDelayedTradeCloseDefinition `json:"delayedTradeClose,omitempty"`
	Reason                 MarketOrderReasonDefinition             `json:"reason,omitempty"`
	ClientExtensions       *ClientExtensionsDefinition             `json:"clientExtensions,omitempty"`
	TakeProfitOnFill       *TakeProfitDetailsDefinition            `json:"takeProfitOnFill,omitempty"`
	StopLossOnFill         *StopLossDetailsDefinition              `json:"stopLossOnFill,omitempty"`
	TrailingStopLossOnFill *TrailingStopLossDetailsDefinition      `json:"trailingStopLossOnFill,omitempty"`
	TradeClientExtensions  *ClientExtensionsDefinition             `json:"tradeClientExtensions,omitempty"`
	RejectReason           TransactionRejectReasonDefinition       `json:"rejectReason,omitempty"`
}

type FixedPriceOrderTransactionDefinition struct {
	ID                     TransactionIDDefinition            `json:"id,omitempty"`
	Time                   DateTimeDefinition                 `json:"time,omitempty"`
	UserID                 int                                `json:"userID,omitempty"`
	AccountID              AccountIDDefinition                `json:"accountID,omitempty"`
	BatchID                TransactionIDDefinition            `json:"batchID,omitempty"`
	RequestID              RequestIDDefinition                `json:"requestID,omitempty"`
	Type                   TransactionTypeDefinition          `json:"type,omitempty"`
	Instrument             InstrumentNameDefinition           `json:"instrument,omitempty"`
	Units                  DecimalNumberDefinition            `json:"units,omitempty"`
	Price                  PriceValueDefinition               `json:"price,omitempty"`
	PositionFill           OrderPositionFillDefinition        `json:"positionFill,omitempty"`
	TradeState             string                             `json:"tradeState,omitempty"`
	Reason                 FixedPriceOrderReasonDefinition    `json:"reason,omitempty"`
	ClientExtensions       *ClientExtensionsDefinition        `json:"clientExtensions,omitempty"`
	TakeProfitOnFill       *TakeProfitDetailsDefinition       `json:"takeProfitOnFill,omitempty"`
	StopLossOnFill         *StopLossDetailsDefinition         `json:"stopLossOnFill,omitempty"`
	TrailingStopLossOnFill *TrailingStopLossDetailsDefinition `json:"trailingStopLossOnFill,omitempty"`
	TradeClientExtensions  *ClientExtensionsDefinition        `json:"tradeClientExtensions,omitempty"`
}

type LimitOrderTransactionDefinition struct {
	ID                      TransactionIDDefinition            `json:"id,omitempty"`
	Time                    DateTimeDefinition                 `json:"time,omitempty"`
	UserID                  int                                `json:"userID,omitempty"`
	AccountID               AccountIDDefinition                `json:"accountID,omitempty"`
	BatchID                 TransactionIDDefinition            `json:"batchID,omitempty"`
	RequestID               RequestIDDefinition                `json:"requestID,omitempty"`
	Type                    TransactionTypeDefinition          `json:"type,omitempty"`
	Instrument              InstrumentNameDefinition           `json:"instrument,omitempty"`
	Units                   DecimalNumberDefinition            `json:"units,omitempty"`
	Price                   PriceValueDefinition               `json:"price,omitempty"`
	TimeInForce             TimeInForceDefinition              `json:"timeInForce,omitempty"`
	GtdTime                 DateTimeDefinition                 `json:"gtdTime,omitempty"`
	PositionFill            OrderPositionFillDefinition        `json:"positionFill,omitempty"`
	TriggerCondition        OrderTriggerConditionDefinition    `json:"triggerCondition,omitempty"`
	Reason                  LimitOrderReasonDefinition         `json:"reason,omitempty"`
	ClientExtensions        *ClientExtensionsDefinition        `json:"clientExtensions,omitempty"`
	TakeProfitOnFill        *TakeProfitDetailsDefinition       `json:"takeProfitOnFill,omitempty"`
	StopLossOnFill          *StopLossDetailsDefinition         `json:"stopLossOnFill,omitempty"`
	TrailingStopLossOnFill  *TrailingStopLossDetailsDefinition `json:"trailingStopLossOnFill,omitempty"`
	TradeClientExtensions   *ClientExtensionsDefinition        `json:"tradeClientExtensions,omitempty"`
	ReplacesOrderID         OrderIDDefinition                  `json:"replacesOrderID,omitempty"`
	CancellingTransactionID TransactionIDDefinition            `json:"cancellingTransactionID,omitempty"`
}

type LimitOrderRejectTransactionDefinition struct {
	ID                      TransactionIDDefinition            `json:"id,omitempty"`
	Time                    DateTimeDefinition                 `json:"time,omitempty"`
	UserID                  int                                `json:"userID,omitempty"`
	AccountID               AccountIDDefinition                `json:"accountID,omitempty"`
	BatchID                 TransactionIDDefinition            `json:"batchID,omitempty"`
	RequestID               RequestIDDefinition                `json:"requestID,omitempty"`
	Type                    TransactionTypeDefinition          `json:"type,omitempty"`
	Instrument              InstrumentNameDefinition           `json:"instrument,omitempty"`
	Units                   DecimalNumberDefinition            `json:"units,omitempty"`
	Price                   PriceValueDefinition               `json:"price,omitempty"`
	TimeInForce             TimeInForceDefinition              `json:"timeInForce,omitempty"`
	GtdTime                 DateTimeDefinition                 `json:"gtdTime,omitempty"`
	PositionFill            OrderPositionFillDefinition        `json:"positionFill,omitempty"`
	TriggerCondition        OrderTriggerConditionDefinition    `json:"triggerCondition,omitempty"`
	Reason                  LimitOrderReasonDefinition         `json:"reason,omitempty"`
	ClientExtensions        *ClientExtensionsDefinition        `json:"clientExtensions,omitempty"`
	TakeProfitOnFill        *TakeProfitDetailsDefinition       `json:"takeProfitOnFill,omitempty"`
	StopLossOnFill          *StopLossDetailsDefinition         `json:"stopLossOnFill,omitempty"`
	TrailingStopLossOnFill  *TrailingStopLossDetailsDefinition `json:"trailingStopLossOnFill,omitempty"`
	TradeClientExtensions   *ClientExtensionsDefinition        `json:"tradeClientExtensions,omitempty"`
	IntendedReplacesOrderID OrderIDDefinition                  `json:"intendedReplacesOrderID,omitempty"`
	RejectReason            TransactionRejectReasonDefinition  `json:"rejectReason,omitempty"`
}

type StopOrderTransactionDefinition struct {
	ID                      TransactionIDDefinition            `json:"id,omitempty"`
	Time                    DateTimeDefinition                 `json:"time,omitempty"`
	UserID                  int                                `json:"userID,omitempty"`
	AccountID               AccountIDDefinition                `json:"accountID,omitempty"`
	BatchID                 TransactionIDDefinition            `json:"batchID,omitempty"`
	RequestID               RequestIDDefinition                `json:"requestID,omitempty"`
	Type                    TransactionTypeDefinition          `json:"type,omitempty"`
	Instrument              InstrumentNameDefinition           `json:"instrument,omitempty"`
	Units                   DecimalNumberDefinition            `json:"units,omitempty"`
	Price                   PriceValueDefinition               `json:"price,omitempty"`
	PriceBound              PriceValueDefinition               `json:"priceBound,omitempty"`
	TimeInForce             TimeInForceDefinition              `json:"timeInForce,omitempty"`
	GtdTime                 DateTimeDefinition                 `json:"gtdTime,omitempty"`
	PositionFill            OrderPositionFillDefinition        `json:"positionFill,omitempty"`
	TriggerCondition        OrderTriggerConditionDefinition    `json:"triggerCondition,omitempty"`
	Reason                  StopOrderReasonDefinition          `json:"reason,omitempty"`
	ClientExtensions        *ClientExtensionsDefinition        `json:"clientExtensions,omitempty"`
	TakeProfitOnFill        *TakeProfitDetailsDefinition       `json:"takeProfitOnFill,omitempty"`
	StopLossOnFill          *StopLossDetailsDefinition         `json:"stopLossOnFill,omitempty"`
	TrailingStopLossOnFill  *TrailingStopLossDetailsDefinition `json:"trailingStopLossOnFill,omitempty"`
	TradeClientExtensions   *ClientExtensionsDefinition        `json:"tradeClientExtensions,omitempty"`
	ReplacesOrderID         OrderIDDefinition                  `json:"replacesOrderID,omitempty"`
	CancellingTransactionID TransactionIDDefinition            `json:"cancellingTransactionID,omitempty"`
}

type StopOrderRejectTransactionDefinition struct {
	ID                      TransactionIDDefinition            `json:"id,omitempty"`
	Time                    DateTimeDefinition                 `json:"time,omitempty"`
	UserID                  int                                `json:"userID,omitempty"`
	AccountID               AccountIDDefinition                `json:"accountID,omitempty"`
	BatchID                 TransactionIDDefinition            `json:"batchID,omitempty"`
	RequestID               RequestIDDefinition                `json:"requestID,omitempty"`
	Type                    TransactionTypeDefinition          `json:"type,omitempty"`
	Instrument              InstrumentNameDefinition           `json:"instrument,omitempty"`
	Units                   DecimalNumberDefinition            `json:"units,omitempty"`
	Price                   PriceValueDefinition               `json:"price,omitempty"`
	PriceBound              PriceValueDefinition               `json:"priceBound,omitempty"`
	TimeInForce             TimeInForceDefinition              `json:"timeInForce,omitempty"`
	GtdTime                 DateTimeDefinition                 `json:"gtdTime,omitempty"`
	PositionFill            OrderPositionFillDefinition        `json:"positionFill,omitempty"`
	TriggerCondition        OrderTriggerConditionDefinition    `json:"triggerCondition,omitempty"`
	Reason                  StopOrderReasonDefinition          `json:"reason,omitempty"`
	ClientExtensions        *ClientExtensionsDefinition        `json:"clientExtensions,omitempty"`
	TakeProfitOnFill        *TakeProfitDetailsDefinition       `json:"takeProfitOnFill,omitempty"`
	StopLossOnFill          *StopLossDetailsDefinition         `json:"stopLossOnFill,omitempty"`
	TrailingStopLossOnFill  *TrailingStopLossDetailsDefinition `json:"trailingStopLossOnFill,omitempty"`
	TradeClientExtensions   *ClientExtensionsDefinition        `json:"tradeClientExtensions,omitempty"`
	IntendedReplacesOrderID OrderIDDefinition                  `json:"intendedReplacesOrderID,omitempty"`
	RejectReason            TransactionRejectReasonDefinition  `json:"rejectReason,omitempty"`
}

type MarketIfTouchedOrderTransactionDefinition struct {
	ID                      TransactionIDDefinition              `json:"id,omitempty"`
	Time                    DateTimeDefinition                   `json:"time,omitempty"`
	UserID                  int                                  `json:"userID,omitempty"`
	AccountID               AccountIDDefinition                  `json:"accountID,omitempty"`
	BatchID                 TransactionIDDefinition              `json:"batchID,omitempty"`
	RequestID               RequestIDDefinition                  `json:"requestID,omitempty"`
	Type                    TransactionTypeDefinition            `json:"type,omitempty"`
	Instrument              InstrumentNameDefinition             `json:"instrument,omitempty"`
	Units                   DecimalNumberDefinition              `json:"units,omitempty"`
	Price                   PriceValueDefinition                 `json:"price,omitempty"`
	PriceBound              PriceValueDefinition                 `json:"priceBound,omitempty"`
	TimeInForce             TimeInForceDefinition                `json:"timeInForce,omitempty"`
	GtdTime                 DateTimeDefinition                   `json:"gtdTime,omitempty"`
	PositionFill            OrderPositionFillDefinition          `json:"positionFill,omitempty"`
	TriggerCondition        OrderTriggerConditionDefinition      `json:"triggerCondition,omitempty"`
	Reason                  MarketIfTouchedOrderReasonDefinition `json:"reason,omitempty"`
	ClientExtensions        *ClientExtensionsDefinition          `json:"clientExtensions,omitempty"`
	TakeProfitOnFill        *TakeProfitDetailsDefinition         `json:"takeProfitOnFill,omitempty"`
	StopLossOnFill          *StopLossDetailsDefinition           `json:"stopLossOnFill,omitempty"`
	TrailingStopLossOnFill  *TrailingStopLossDetailsDefinition   `json:"trailingStopLossOnFill,omitempty"`
	TradeClientExtensions   *ClientExtensionsDefinition          `json:"tradeClientExtensions,omitempty"`
	ReplacesOrderID         OrderIDDefinition                    `json:"replacesOrderID,omitempty"`
	CancellingTransactionID TransactionIDDefinition              `json:"cancellingTransactionID,omitempty"`
}

type MarketIfTouchedOrderRejectTransactionDefinition struct {
	ID                      TransactionIDDefinition              `json:"id,omitempty"`
	Time                    DateTimeDefinition                   `json:"time,omitempty"`
	UserID                  int                                  `json:"userID,omitempty"`
	AccountID               AccountIDDefinition                  `json:"accountID,omitempty"`
	BatchID                 TransactionIDDefinition              `json:"batchID,omitempty"`
	RequestID               RequestIDDefinition                  `json:"requestID,omitempty"`
	Type                    TransactionTypeDefinition            `json:"type,omitempty"`
	Instrument              InstrumentNameDefinition             `json:"instrument,omitempty"`
	Units                   DecimalNumberDefinition              `json:"units,omitempty"`
	Price                   PriceValueDefinition                 `json:"price,omitempty"`
	PriceBound              PriceValueDefinition                 `json:"priceBound,omitempty"`
	TimeInForce             TimeInForceDefinition                `json:"timeInForce,omitempty"`
	GtdTime                 DateTimeDefinition                   `json:"gtdTime,omitempty"`
	PositionFill            OrderPositionFillDefinition          `json:"positionFill,omitempty"`
	TriggerCondition        OrderTriggerConditionDefinition      `json:"triggerCondition,omitempty"`
	Reason                  MarketIfTouchedOrderReasonDefinition `json:"reason,omitempty"`
	ClientExtensions        *ClientExtensionsDefinition          `json:"clientExtensions,omitempty"`
	TakeProfitOnFill        *TakeProfitDetailsDefinition         `json:"takeProfitOnFill,omitempty"`
	StopLossOnFill          *StopLossDetailsDefinition           `json:"stopLossOnFill,omitempty"`
	TrailingStopLossOnFill  *TrailingStopLossDetailsDefinition   `json:"trailingStopLossOnFill,omitempty"`
	TradeClientExtensions   *ClientExtensionsDefinition          `json:"tradeClientExtensions,omitempty"`
	IntendedReplacesOrderID OrderIDDefinition                    `json:"intendedReplacesOrderID,omitempty"`
	RejectReason            TransactionRejectReasonDefinition    `json:"rejectReason,omitempty"`
}

type TakeProfitOrderTransactionDefinition struct {
	ID                      TransactionIDDefinition         `json:"id,omitempty"`
	Time                    DateTimeDefinition              `json:"time,omitempty"`
	UserID                  int                             `json:"userID,omitempty"`
	AccountID               AccountIDDefinition             `json:"accountID,omitempty"`
	BatchID                 TransactionIDDefinition         `json:"batchID,omitempty"`
	RequestID               RequestIDDefinition             `json:"requestID,omitempty"`
	Type                    TransactionTypeDefinition       `json:"type,omitempty"`
	TradeID                 TradeIDDefinition               `json:"tradeID,omitempty"`
	ClientTradeID           ClientIDDefinition              `json:"clientTradeID,omitempty"`
	Price                   PriceValueDefinition            `json:"price,omitempty"`
	TimeInForce             TimeInForceDefinition           `json:"timeInForce,omitempty"`
	GtdTime                 DateTimeDefinition              `json:"gtdTime,omitempty"`
	TriggerCondition        OrderTriggerConditionDefinition `json:"triggerCondition,omitempty"`
	Reason                  TakeProfitOrderReasonDefinition `json:"reason,omitempty"`
	ClientExtensions        *ClientExtensionsDefinition     `json:"clientExtensions,omitempty"`
	OrderFillTransactionID  TransactionIDDefinition         `json:"orderFillTransactionID,omitempty"`
	ReplacesOrderID         OrderIDDefinition               `json:"replacesOrderID,omitempty"`
	CancellingTransactionID TransactionIDDefinition         `json:"cancellingTransactionID,omitempty"`
}

type TakeProfitOrderRejectTransactionDefinition struct {
	ID                      TransactionIDDefinition           `json:"id,omitempty"`
	Time                    DateTimeDefinition                `json:"time,omitempty"`
	UserID                  int                               `json:"userID,omitempty"`
	AccountID               AccountIDDefinition               `json:"accountID,omitempty"`
	BatchID                 TransactionIDDefinition           `json:"batchID,omitempty"`
	RequestID               RequestIDDefinition               `json:"requestID,omitempty"`
	Type                    TransactionTypeDefinition         `json:"type,omitempty"`
	TradeID                 TradeIDDefinition                 `json:"tradeID,omitempty"`
	ClientTradeID           ClientIDDefinition                `json:"clientTradeID,omitempty"`
	Price                   PriceValueDefinition              `json:"price,omitempty"`
	TimeInForce             TimeInForceDefinition             `json:"timeInForce,omitempty"`
	GtdTime                 DateTimeDefinition                `json:"gtdTime,omitempty"`
	TriggerCondition        OrderTriggerConditionDefinition   `json:"triggerCondition,omitempty"`
	Reason                  TakeProfitOrderReasonDefinition   `json:"reason,omitempty"`
	ClientExtensions        *ClientExtensionsDefinition       `json:"clientExtensions,omitempty"`
	OrderFillTransactionID  TransactionIDDefinition           `json:"orderFillTransactionID,omitempty"`
	IntendedReplacesOrderID OrderIDDefinition                 `json:"intendedReplacesOrderID,omitempty"`
	RejectReason            TransactionRejectReasonDefinition `json:"rejectReason,omitempty"`
}

type StopLossOrderTransactionDefinition struct {
	ID                         TransactionIDDefinition         `json:"id,omitempty"`
	Time                       DateTimeDefinition              `json:"time,omitempty"`
	UserID                     int                             `json:"userID,omitempty"`
	AccountID                  AccountIDDefinition             `json:"accountID,omitempty"`
	BatchID                    TransactionIDDefinition         `json:"batchID,omitempty"`
	RequestID                  RequestIDDefinition             `json:"requestID,omitempty"`
	Type                       TransactionTypeDefinition       `json:"type,omitempty"`
	TradeID                    TradeIDDefinition               `json:"tradeID,omitempty"`
	ClientTradeID              ClientIDDefinition              `json:"clientTradeID,omitempty"`
	Price                      PriceValueDefinition            `json:"price,omitempty"`
	Distance                   DecimalNumberDefinition         `json:"distance,omitempty"`
	TimeInForce                TimeInForceDefinition           `json:"timeInForce,omitempty"`
	GtdTime                    DateTimeDefinition              `json:"gtdTime,omitempty"`
	TriggerCondition           OrderTriggerConditionDefinition `json:"triggerCondition,omitempty"`
	Guaranteed                 bool                            `json:"guaranteed,omitempty"`
	GuaranteedExecutionPremium DecimalNumberDefinition         `json:"guaranteedExecutionPremium,omitempty"`
	Reason                     StopLossOrderReasonDefinition   `json:"reason,omitempty"`
	ClientExtensions           *ClientExtensionsDefinition     `json:"clientExtensions,omitempty"`
	OrderFillTransactionID     TransactionIDDefinition         `json:"orderFillTransactionID,omitempty"`
	ReplacesOrderID            OrderIDDefinition               `json:"replacesOrderID,omitempty"`
	CancellingTransactionID    TransactionIDDefinition         `json:"cancellingTransactionID,omitempty"`
}

type StopLossOrderRejectTransactionDefinition struct {
	ID                      TransactionIDDefinition           `json:"id,omitempty"`
	Time                    DateTimeDefinition                `json:"time,omitempty"`
	UserID                  int                               `json:"userID,omitempty"`
	AccountID               AccountIDDefinition               `json:"accountID,omitempty"`
	BatchID                 TransactionIDDefinition           `json:"batchID,omitempty"`
	RequestID               RequestIDDefinition               `json:"requestID,omitempty"`
	Type                    TransactionTypeDefinition         `json:"type,omitempty"`
	TradeID                 TradeIDDefinition                 `json:"tradeID,omitempty"`
	ClientTradeID           ClientIDDefinition                `json:"clientTradeID,omitempty"`
	Price                   PriceValueDefinition              `json:"price,omitempty"`
	Distance                DecimalNumberDefinition           `json:"distance,omitempty"`
	TimeInForce             TimeInForceDefinition             `json:"timeInForce,omitempty"`
	GtdTime                 DateTimeDefinition                `json:"gtdTime,omitempty"`
	TriggerCondition        OrderTriggerConditionDefinition   `json:"triggerCondition,omitempty"`
	Guaranteed              bool                              `json:"guaranteed,omitempty"`
	Reason                  StopLossOrderReasonDefinition     `json:"reason,omitempty"`
	ClientExtensions        *ClientExtensionsDefinition       `json:"clientExtensions,omitempty"`
	OrderFillTransactionID  TransactionIDDefinition           `json:"orderFillTransactionID,omitempty"`
	IntendedReplacesOrderID OrderIDDefinition                 `json:"intendedReplacesOrderID,omitempty"`
	RejectReason            TransactionRejectReasonDefinition `json:"rejectReason,omitempty"`
}

type TrailingStopLossOrderTransactionDefinition struct {
	ID                      TransactionIDDefinition               `json:"id,omitempty"`
	Time                    DateTimeDefinition                    `json:"time,omitempty"`
	UserID                  int                                   `json:"userID,omitempty"`
	AccountID               AccountIDDefinition                   `json:"accountID,omitempty"`
	BatchID                 TransactionIDDefinition               `json:"batchID,omitempty"`
	RequestID               RequestIDDefinition                   `json:"requestID,omitempty"`
	Type                    TransactionTypeDefinition             `json:"type,omitempty"`
	TradeID                 TradeIDDefinition                     `json:"tradeID,omitempty"`
	ClientTradeID           ClientIDDefinition                    `json:"clientTradeID,omitempty"`
	Distance                DecimalNumberDefinition               `json:"distance,omitempty"`
	TimeInForce             TimeInForceDefinition                 `json:"timeInForce,omitempty"`
	GtdTime                 DateTimeDefinition                    `json:"gtdTime,omitempty"`
	TriggerCondition        OrderTriggerConditionDefinition       `json:"triggerCondition,omitempty"`
	Reason                  TrailingStopLossOrderReasonDefinition `json:"reason,omitempty"`
	ClientExtensions        *ClientExtensionsDefinition           `json:"clientExtensions,omitempty"`
	OrderFillTransactionID  TransactionIDDefinition               `json:"orderFillTransactionID,omitempty"`
	ReplacesOrderID         OrderIDDefinition                     `json:"replacesOrderID,omitempty"`
	CancellingTransactionID TransactionIDDefinition               `json:"cancellingTransactionID,omitempty"`
}

type TrailingStopLossOrderRejectTransactionDefinition struct {
	ID                      TransactionIDDefinition               `json:"id,omitempty"`
	Time                    DateTimeDefinition                    `json:"time,omitempty"`
	UserID                  int                                   `json:"userID,omitempty"`
	AccountID               AccountIDDefinition                   `json:"accountID,omitempty"`
	BatchID                 TransactionIDDefinition               `json:"batchID,omitempty"`
	RequestID               RequestIDDefinition                   `json:"requestID,omitempty"`
	Type                    TransactionTypeDefinition             `json:"type,omitempty"`
	TradeID                 TradeIDDefinition                     `json:"tradeID,omitempty"`
	ClientTradeID           ClientIDDefinition                    `json:"clientTradeID,omitempty"`
	Distance                DecimalNumberDefinition               `json:"distance,omitempty"`
	TimeInForce             TimeInForceDefinition                 `json:"timeInForce,omitempty"`
	GtdTime                 DateTimeDefinition                    `json:"gtdTime,omitempty"`
	TriggerCondition        OrderTriggerConditionDefinition       `json:"triggerCondition,omitempty"`
	Reason                  TrailingStopLossOrderReasonDefinition `json:"reason,omitempty"`
	ClientExtensions        *ClientExtensionsDefinition           `json:"clientExtensions,omitempty"`
	OrderFillTransactionID  TransactionIDDefinition               `json:"orderFillTransactionID,omitempty"`
	IntendedReplacesOrderID OrderIDDefinition                     `json:"intendedReplacesOrderID,omitempty"`
	RejectReason            TransactionRejectReasonDefinition     `json:"rejectReason,omitempty"`
}

type OrderFillTransactionDefinition struct {
	ID                            TransactionIDDefinition   `json:"id,omitempty"`
	Time                          DateTimeDefinition        `json:"time,omitempty"`
	UserID                        int                       `json:"userID,omitempty"`
	AccountID                     AccountIDDefinition       `json:"accountID,omitempty"`
	BatchID                       TransactionIDDefinition   `json:"batchID,omitempty"`
	RequestID                     RequestIDDefinition       `json:"requestID,omitempty"`
	Type                          TransactionTypeDefinition `json:"type,omitempty"`
	OrderID                       OrderIDDefinition         `json:"orderID,omitempty"`
	ClientOrderID                 ClientIDDefinition        `json:"clientOrderID,omitempty"`
	Instrument                    InstrumentNameDefinition  `json:"instrument,omitempty"`
	Units                         DecimalNumberDefinition   `json:"units,omitempty"`
	GainQuoteHomeConversionFactor DecimalNumberDefinition   `json:"gainQuoteHomeConversionFactor,omitempty"`
	LossQuoteHomeConversionFactor DecimalNumberDefinition   `json:"lossQuoteHomeConversionFactor,omitempty"`
	Price                         PriceValueDefinition      `json:"price,omitempty"`
	FullPrice                     *ClientPriceDefinition    `json:"fullPrice,omitempty"`
	Reason                        OrderFillReasonDefinition `json:"reason,omitempty"`
	PL                            AccountUnitsDefinition    `json:"pl,omitempty"`
	Financing                     AccountUnitsDefinition    `json:"financing,omitempty"`
	Commission                    AccountUnitsDefinition    `json:"commission,omitempty"`
	GuaranteedExecutionFee        AccountUnitsDefinition    `json:"guaranteedExecutionFee,omitempty"`
	AccountBalance                AccountUnitsDefinition    `json:"accountBalance,omitempty"`
	TradeOpened                   *TradeOpenDefinition      `json:"tradeOpened,omitempty"`
	TradesClosed                  []*TradeReduceDefinition  `json:"tradesClosed,omitempty"`
	TradeReduced                  *TradeReduceDefinition    `json:"tradeReduced,omitempty"`
	HalfSpreadCost                AccountUnitsDefinition    `json:"halfSpreadCost,omitempty"`
}

type OrderCancelTransactionDefinition struct {
	ID                TransactionIDDefinition     `json:"id,omitempty"`
	Time              DateTimeDefinition          `json:"time,omitempty"`
	UserID            int                         `json:"userID,omitempty"`
	AccountID         AccountIDDefinition         `json:"accountID,omitempty"`
	BatchID           TransactionIDDefinition     `json:"batchID,omitempty"`
	RequestID         RequestIDDefinition         `json:"requestID,omitempty"`
	Type              TransactionTypeDefinition   `json:"type,omitempty"`
	OrderID           OrderIDDefinition           `json:"orderID,omitempty"`
	ClientOrderID     OrderIDDefinition           `json:"clientOrderID,omitempty"`
	Reason            OrderCancelReasonDefinition `json:"reason,omitempty"`
	ReplacedByOrderID OrderIDDefinition           `json:"replacedByOrderID,omitempty"`
}

type OrderCancelRejectTransactionDefinition struct {
	ID            TransactionIDDefinition           `json:"id,omitempty"`
	Time          DateTimeDefinition                `json:"time,omitempty"`
	UserID        int                               `json:"userID,omitempty"`
	AccountID     AccountIDDefinition               `json:"accountID,omitempty"`
	BatchID       TransactionIDDefinition           `json:"batchID,omitempty"`
	RequestID     RequestIDDefinition               `json:"requestID,omitempty"`
	Type          TransactionTypeDefinition         `json:"type,omitempty"`
	OrderID       OrderIDDefinition                 `json:"orderID,omitempty"`
	ClientOrderID OrderIDDefinition                 `json:"clientOrderID,omitempty"`
	RejectReason  TransactionRejectReasonDefinition `json:"rejectReason,omitempty"`
}

type OrderClientExtensionsModifyTransactionDefinition struct {
	ID                          TransactionIDDefinition     `json:"id,omitempty"`
	Time                        DateTimeDefinition          `json:"time,omitempty"`
	UserID                      int                         `json:"userID,omitempty"`
	AccountID                   AccountIDDefinition         `json:"accountID,omitempty"`
	BatchID                     TransactionIDDefinition     `json:"batchID,omitempty"`
	RequestID                   RequestIDDefinition         `json:"requestID,omitempty"`
	Type                        TransactionTypeDefinition   `json:"type,omitempty"`
	OrderID                     OrderIDDefinition           `json:"orderID,omitempty"`
	ClientOrderID               ClientIDDefinition          `json:"clientOrderID,omitempty"`
	ClientExtensionsModify      *ClientExtensionsDefinition `json:"clientExtensionsModify,omitempty"`
	TradeClientExtensionsModify *ClientExtensionsDefinition `json:"tradeClientExtensionsModify,omitempty"`
}

type OrderClientExtensionsModifyRejectTransactionDefinition struct {
	ID                          TransactionIDDefinition           `json:"id,omitempty"`
	Time                        DateTimeDefinition                `json:"time,omitempty"`
	UserID                      int                               `json:"userID,omitempty"`
	AccountID                   AccountIDDefinition               `json:"accountID,omitempty"`
	BatchID                     TransactionIDDefinition           `json:"batchID,omitempty"`
	RequestID                   RequestIDDefinition               `json:"requestID,omitempty"`
	Type                        TransactionTypeDefinition         `json:"type,omitempty"`
	OrderID                     OrderIDDefinition                 `json:"orderID,omitempty"`
	ClientOrderID               ClientIDDefinition                `json:"clientOrderID,omitempty"`
	ClientExtensionsModify      *ClientExtensionsDefinition       `json:"clientExtensionsModify,omitempty"`
	TradeClientExtensionsModify *ClientExtensionsDefinition       `json:"tradeClientExtensionsModify,omitempty"`
	RejectReason                TransactionRejectReasonDefinition `json:"rejectReason,omitempty"`
}

type TradeClientExtensionsModifyTransactionDefinition struct {
	ID                          TransactionIDDefinition     `json:"id,omitempty"`
	Time                        DateTimeDefinition          `json:"time,omitempty"`
	UserID                      int                         `json:"userID,omitempty"`
	AccountID                   AccountIDDefinition         `json:"accountID,omitempty"`
	BatchID                     TransactionIDDefinition     `json:"batchID,omitempty"`
	RequestID                   RequestIDDefinition         `json:"requestID,omitempty"`
	Type                        TransactionTypeDefinition   `json:"type,omitempty"`
	TradeID                     TradeIDDefinition           `json:"tradeID,omitempty"`
	ClientTradeID               ClientIDDefinition          `json:"clientTradeID,omitempty"`
	TradeClientExtensionsModify *ClientExtensionsDefinition `json:"tradeClientExtensionsModify,omitempty"`
}

type TradeClientExtensionsModifyRejectTransactionDefinition struct {
	ID                          TransactionIDDefinition           `json:"id,omitempty"`
	Time                        DateTimeDefinition                `json:"time,omitempty"`
	UserID                      int                               `json:"userID,omitempty"`
	AccountID                   AccountIDDefinition               `json:"accountID,omitempty"`
	BatchID                     TransactionIDDefinition           `json:"batchID,omitempty"`
	RequestID                   RequestIDDefinition               `json:"requestID,omitempty"`
	Type                        TransactionTypeDefinition         `json:"type,omitempty"`
	TradeID                     TradeIDDefinition                 `json:"tradeID,omitempty"`
	ClientTradeID               ClientIDDefinition                `json:"clientTradeID,omitempty"`
	TradeClientExtensionsModify *ClientExtensionsDefinition       `json:"tradeClientExtensionsModify,omitempty"`
	RejectReason                TransactionRejectReasonDefinition `json:"rejectReason,omitempty"`
}

type MarginCallEnterTransactionDefinition struct {
	ID        TransactionIDDefinition   `json:"id,omitempty"`
	Time      DateTimeDefinition        `json:"time,omitempty"`
	UserID    int                       `json:"userID,omitempty"`
	AccountID AccountIDDefinition       `json:"accountID,omitempty"`
	BatchID   TransactionIDDefinition   `json:"batchID,omitempty"`
	RequestID RequestIDDefinition       `json:"requestID,omitempty"`
	Type      TransactionTypeDefinition `json:"type,omitempty"`
}

type MarginCallExtendTransactionDefinition struct {
	ID              TransactionIDDefinition   `json:"id,omitempty"`
	Time            DateTimeDefinition        `json:"time,omitempty"`
	UserID          int                       `json:"userID,omitempty"`
	AccountID       AccountIDDefinition       `json:"accountID,omitempty"`
	BatchID         TransactionIDDefinition   `json:"batchID,omitempty"`
	RequestID       RequestIDDefinition       `json:"requestID,omitempty"`
	Type            TransactionTypeDefinition `json:"type,omitempty"`
	ExtensionNumber int                       `json:"extensionNumber,omitempty"`
}

type MarginCallExitTransactionDefinition struct {
	ID        TransactionIDDefinition   `json:"id,omitempty"`
	Time      DateTimeDefinition        `json:"time,omitempty"`
	UserID    int                       `json:"userID,omitempty"`
	AccountID AccountIDDefinition       `json:"accountID,omitempty"`
	BatchID   TransactionIDDefinition   `json:"batchID,omitempty"`
	RequestID RequestIDDefinition       `json:"requestID,omitempty"`
	Type      TransactionTypeDefinition `json:"type,omitempty"`
}

type DelayedTradeClosureTransactionDefinition struct {
	ID        TransactionIDDefinition     `json:"id,omitempty"`
	Time      DateTimeDefinition          `json:"time,omitempty"`
	UserID    int                         `json:"userID,omitempty"`
	AccountID AccountIDDefinition         `json:"accountID,omitempty"`
	BatchID   TransactionIDDefinition     `json:"batchID,omitempty"`
	RequestID RequestIDDefinition         `json:"requestID,omitempty"`
	Type      TransactionTypeDefinition   `json:"type,omitempty"`
	Reason    MarketOrderReasonDefinition `json:"reason,omitempty"`
	TradeIDs  TradeIDDefinition           `json:"tradeIDs,omitempty"`
}

type DailyFinancingTransactionDefinition struct {
	ID                   TransactionIDDefinition        `json:"id,omitempty"`
	Time                 DateTimeDefinition             `json:"time,omitempty"`
	UserID               int                            `json:"userID,omitempty"`
	AccountID            AccountIDDefinition            `json:"accountID,omitempty"`
	BatchID              TransactionIDDefinition        `json:"batchID,omitempty"`
	RequestID            RequestIDDefinition            `json:"requestID,omitempty"`
	Type                 TransactionTypeDefinition      `json:"type,omitempty"`
	Financing            AccountUnitsDefinition         `json:"financing,omitempty"`
	AccountBalance       AccountUnitsDefinition         `json:"accountBalance,omitempty"`
	AccountFinancingMode AccountFinancingModeDefinition `json:"accountFinancingMode,omitempty"`
	PositionFinancings   []*PositionFinancingDefinition `json:"positionFinancings,omitempty"`
}

type ResetResettablePLTransactionDefinition struct {
	ID        TransactionIDDefinition   `json:"id,omitempty"`
	Time      DateTimeDefinition        `json:"time,omitempty"`
	UserID    int                       `json:"userID,omitempty"`
	AccountID AccountIDDefinition       `json:"accountID,omitempty"`
	BatchID   TransactionIDDefinition   `json:"batchID,omitempty"`
	RequestID RequestIDDefinition       `json:"requestID,omitempty"`
	Type      TransactionTypeDefinition `json:"type,omitempty"`
}

// Transaction-related Definitions

type TransactionIDDefinition string

type TransactionTypeDefinition string

type FundingReasonDefinition string

type MarketOrderReasonDefinition string

type FixedPriceOrderReasonDefinition string

type LimitOrderReasonDefinition string

type StopOrderReasonDefinition string

type MarketIfTouchedOrderReasonDefinition string

type TakeProfitOrderReasonDefinition string

type StopLossOrderReasonDefinition string

type TrailingStopLossOrderReasonDefinition string

type OrderFillReasonDefinition string

type OrderCancelReasonDefinition string

type ClientIDDefinition string

type ClientTagDefinition string

type ClientCommentDefinition string

type ClientExtensionsDefinition struct {
	ID      ClientIDDefinition      `json:"id,omitempty"`
	Tag     ClientTagDefinition     `json:"tag,omitempty"`
	Comment ClientCommentDefinition `json:"comment,omitempty"`
}

type TakeProfitDetailsDefinition struct {
	Price            PriceValueDefinition        `json:"price,omitempty"`
	TimeInForce      TimeInForceDefinition       `json:"timeInForce,omitempty"`
	GtdTime          DateTimeDefinition          `json:"gtdTime,omitempty"`
	ClientExtensions *ClientExtensionsDefinition `json:"clientExtensions,omitempty"`
}

type StopLossDetailsDefinition struct {
	Price            PriceValueDefinition        `json:"price,omitempty"`
	Distance         DecimalNumberDefinition     `json:"distance,omitempty"`
	TimeInForce      TimeInForceDefinition       `json:"timeInForce,omitempty"`
	GtdTime          DateTimeDefinition          `json:"gtdTime,omitempty"`
	ClientExtensions *ClientExtensionsDefinition `json:"clientExtensions,omitempty"`
	Guaranteed       bool                        `json:"guaranteed,omitempty"`
}

type TrailingStopLossDetailsDefinition struct {
	Distance         DecimalNumberDefinition     `json:"distance,omitempty"`
	TimeInForce      TimeInForceDefinition       `json:"timeInForce,omitempty"`
	GtdTime          DateTimeDefinition          `json:"gtdTime,omitempty"`
	ClientExtensions *ClientExtensionsDefinition `json:"clientExtensions,omitempty"`
}

type TradeOpenDefinition struct {
	TradeID                TradeIDDefinition           `json:"tradeID,omitempty"`
	Units                  DecimalNumberDefinition     `json:"units,omitempty"`
	Price                  PriceValueDefinition        `json:"price,omitempty"`
	GuaranteedExecutionFee AccountUnitsDefinition      `json:"guaranteedExecutionFee,omitempty"`
	ClientExtensions       *ClientExtensionsDefinition `json:"clientExtensions,omitempty"`
	HalfSpreadCost         AccountUnitsDefinition      `json:"halfSpreadCost,omitempty"`
	InitialMarginRequired  AccountUnitsDefinition      `json:"initialMarginRequired,omitempty"`
}

type TradeReduceDefinition struct {
	TradeID                TradeIDDefinition       `json:"tradeID,omitempty"`
	Units                  DecimalNumberDefinition `json:"units,omitempty"`
	Price                  PriceValueDefinition    `json:"price,omitempty"`
	RealizedPL             AccountUnitsDefinition  `json:"realizedPL,omitempty"`
	Financing              AccountUnitsDefinition  `json:"financing,omitempty"`
	GuaranteedExecutionFee AccountUnitsDefinition  `json:"guaranteedExecutionFee,omitempty"`
	HalfSpreadCost         AccountUnitsDefinition  `json:"halfSpreadCost,omitempty"`
}

type MarketOrderTradeCloseDefinition struct {
	TradeID       TradeIDDefinition `json:"tradeID,omitempty"`
	ClientTradeID string            `json:"clientTradeID,omitempty"`
	Units         string            `json:"units,omitempty"`
}

type MarketOrderMarginCloseoutDefinition struct {
	Reason MarketOrderMarginCloseoutReasonDefinition `json:"reason,omitempty"`
}

type MarketOrderMarginCloseoutReasonDefinition string

type MarketOrderDelayedTradeCloseDefinition struct {
	TradeID             TradeIDDefinition       `json:"tradeID,omitempty"`
	ClientTradeID       TradeIDDefinition       `json:"clientTradeID,omitempty"`
	SourceTransactionID TransactionIDDefinition `json:"sourceTransactionID,omitempty"`
}

type MarketOrderPositionCloseoutDefinition struct {
	Instrument InstrumentNameDefinition `json:"instrument,omitempty"`
	Units      string                   `json:"units,omitempty"`
}

type LiquidityRegenerationScheduleDefinition struct {
	Steps []*LiquidityRegenerationScheduleStepDefinition `json:"steps,omitempty"`
}

type LiquidityRegenerationScheduleStepDefinition struct {
	Timestamp        DateTimeDefinition      `json:"timestamp,omitempty"`
	BidLiquidityUsed DecimalNumberDefinition `json:"bidLiquidityUsed,omitempty"`
	AskLiquidityUsed DecimalNumberDefinition `json:"askLiquidityUsed,omitempty"`
}

type OpenTradeFinancingDefinition struct {
	TradeID   TradeIDDefinition      `json:"tradeID,omitempty"`
	Financing AccountUnitsDefinition `json:"financing,omitempty"`
}

type PositionFinancingDefinition struct {
	Instrument          InstrumentNameDefinition        `json:"instrument,omitempty"`
	Financing           AccountUnitsDefinition          `json:"financing,omitempty"`
	OpenTradeFinancings []*OpenTradeFinancingDefinition `json:"openTradeFinancings,omitempty"`
}

type RequestIDDefinition string

type TransactionRejectReasonDefinition string

type TransactionFilterDefinition string

const (
	OrderTransaction                             TransactionFilterDefinition = "ORDER"                                 //	Order-related Transactions. These are the Transactions that create, cancel, fill or trigger Orders
	FundingTransaction                           TransactionFilterDefinition = "FUNDING"                               //	Funding-related Transactions
	AdminTransaction                             TransactionFilterDefinition = "ADMIN"                                 //	Administrative Transactions
	CreateTransaction                            TransactionFilterDefinition = "CREATE"                                //	Account Create Transaction
	CloseTransaction                             TransactionFilterDefinition = "CLOSE"                                 //	Account Close Transaction
	ReopenTransaction                            TransactionFilterDefinition = "REOPEN"                                //	Account Reopen Transaction
	ClientConfigureTransaction                   TransactionFilterDefinition = "CLIENT_CONFIGURE"                      //	Client Configuration Transaction
	ClientConfigureRejectTransaction             TransactionFilterDefinition = "CLIENT_CONFIGURE_REJECT"               //	Client Configuration Reject Transaction
	TransferFundsTransaction                     TransactionFilterDefinition = "TRANSFER_FUNDS"                        //	Transfer Funds Transaction
	TransferFundsRejectTransaction               TransactionFilterDefinition = "TRANSFER_FUNDS_REJECT"                 //	Transfer Funds Reject Transaction
	MarketOrderTransaction                       TransactionFilterDefinition = "MARKET_ORDER"                          //	Market Order Transaction
	MarketOrderRejectTransaction                 TransactionFilterDefinition = "MARKET_ORDER_REJECT"                   //	Market Order Reject Transaction
	LimitOrderTransaction                        TransactionFilterDefinition = "LIMIT_ORDER"                           //	Limit Order Transaction
	LimitOrderRejectTransaction                  TransactionFilterDefinition = "LIMIT_ORDER_REJECT"                    //	Limit Order Reject Transaction
	StopOrderTransaction                         TransactionFilterDefinition = "STOP_ORDER"                            //	Stop Order Transaction
	StopOrderRejectTransaction                   TransactionFilterDefinition = "STOP_ORDER_REJECT"                     //	Stop Order Reject Transaction
	MarketIfTouchedOrderTransaction              TransactionFilterDefinition = "MARKET_IF_TOUCHED_ORDER"               //	Market if Touched Order Transaction
	MarketIfTouchedOrderRejectTransaction        TransactionFilterDefinition = "MARKET_IF_TOUCHED_ORDER_REJECT"        //	Market if Touched Order Reject Transaction
	TakeProfitOrderTransaction                   TransactionFilterDefinition = "TAKE_PROFIT_ORDER"                     //	Take Profit Order Transaction
	TakeProfitOrderRejectTransaction             TransactionFilterDefinition = "TAKE_PROFIT_ORDER_REJECT"              //	Take Profit Order Reject Transaction
	StopLossOrderTransaction                     TransactionFilterDefinition = "STOP_LOSS_ORDER"                       //	Stop Loss Order Transaction
	StopLossOrderRejectTransaction               TransactionFilterDefinition = "STOP_LOSS_ORDER_REJECT"                //	Stop Loss Order Reject Transaction
	TrailingStopLossOrderTransaction             TransactionFilterDefinition = "TRAILING_STOP_LOSS_ORDER"              //	Trailing Stop Loss Order Transaction
	TrailingStopLossOrderRejectTransaction       TransactionFilterDefinition = "TRAILING_STOP_LOSS_ORDER_REJECT"       //	Trailing Stop Loss Order Reject Transaction
	OneCancelsAllOrderTransaction                TransactionFilterDefinition = "ONE_CANCELS_ALL_ORDER"                 //	One Cancels All Order Transaction
	OneCancelsAllOrderRejectTransaction          TransactionFilterDefinition = "ONE_CANCELS_ALL_ORDER_REJECT"          //	One Cancels All Order Reject Transaction
	OneCancelsAllOrderTriggeredTransaction       TransactionFilterDefinition = "ONE_CANCELS_ALL_ORDER_TRIGGERED"       //	One Cancels All Order Trigger Transaction
	OrderFillTransaction                         TransactionFilterDefinition = "ORDER_FILL"                            //	Order Fill Transaction
	OrderCancelTransaction                       TransactionFilterDefinition = "ORDER_CANCEL"                          //	Order Cancel Transaction
	OrderCancelRejectTransaction                 TransactionFilterDefinition = "ORDER_CANCEL_REJECT"                   //	Order Cancel Reject Transaction
	OrderClientExtensionsModifyTransaction       TransactionFilterDefinition = "ORDER_CLIENT_EXTENSIONS_MODIFY"        //	Order Client Extensions Modify Transaction
	OrderClientExtensionsModifyRejectTransaction TransactionFilterDefinition = "ORDER_CLIENT_EXTENSIONS_MODIFY_REJECT" //	Order Client Extensions Modify Reject Transaction
	TradeClientExtensionsModifyTransaction       TransactionFilterDefinition = "TRADE_CLIENT_EXTENSIONS_MODIFY"        //	Trade Client Extensions Modify Transaction
	TradeClientExtensionsModifyRejectTransaction TransactionFilterDefinition = "TRADE_CLIENT_EXTENSIONS_MODIFY_REJECT" //	Trade Client Extensions Modify Reject Transaction
	MarginCallEnterTransaction                   TransactionFilterDefinition = "MARGIN_CALL_ENTER"                     //	Margin Call Enter Transaction
	MarginCallExtendTransaction                  TransactionFilterDefinition = "MARGIN_CALL_EXTEND"                    //	Margin Call Extend Transaction
	MarginCallExitTransaction                    TransactionFilterDefinition = "MARGIN_CALL_EXIT"                      //	Margin Call Exit Transaction
	DelayedTradeClosureTransaction               TransactionFilterDefinition = "DELAYED_TRADE_CLOSURE"                 //	Delayed Trade Closure Transaction
	DailyFinancingTransaction                    TransactionFilterDefinition = "DAILY_FINANCING"                       //	Daily Financing Transaction
	ResetResettablePLTransaction                 TransactionFilterDefinition = "RESET_RESETTABLE_PL"                   //	Reset Resettable PL Transaction
)

type TransactionHeartbeatDefinition struct {
	Type              string                  `json:"type,omitempty"`
	LastTransactionID TransactionIDDefinition `json:"lastTransactionID,omitempty"`
	Time              DateTimeDefinition      `json:"time,omitempty"`
}

//
// Pricing Definitions
//

type PriceDefinition struct {
	Type                       string                                `json:"type,omitempty"`
	Instrument                 InstrumentNameDefinition              `json:"instrument,omitempty"`
	Time                       DateTimeDefinition                    `json:"time,omitempty"`
	Status                     PriceStatusDefinition                 `json:"status,omitempty"`
	Tradeable                  bool                                  `json:"tradeable,omitempty"`
	Bids                       []*PriceBucketDefinition              `json:"bids,omitempty"`
	Asks                       []*PriceBucketDefinition              `json:"asks,omitempty"`
	CloseoutBid                PriceValueDefinition                  `json:"closeoutBid,omitempty"`
	CloseoutAsk                PriceValueDefinition                  `json:"closeoutAsk,omitempty"`
	QuoteHomeConversionFactors *QuoteHomeConversionFactorsDefinition `json:"quoteHomeConversionFactors,omitempty"`
	UnitsAvailable             *UnitsAvailableDefinition             `json:"unitsAvailable,omitempty"`
}

type PriceValueDefinition string

type PriceBucketDefinition struct {
	Price     PriceValueDefinition `json:"price,omitempty"`
	Liquidity json.Number          `json:"liquidity,omitempty"`
}

type PriceStatusDefinition string

type QuoteHomeConversionFactorsDefinition struct {
	PositiveUnits DecimalNumberDefinition `json:"positiveUnits,omitempty"`
	NegativeUnits DecimalNumberDefinition `json:"negativeUnits,omitempty"`
}

type HomeConversionsDefinition struct {
	Currency      CurrencyDefinition      `json:"currency,omitempty"`
	AccountGain   DecimalNumberDefinition `json:"accountGain,omitempty"`
	AccountLoss   DecimalNumberDefinition `json:"accountLoss,omitempty"`
	PositionValue DecimalNumberDefinition `json:"positionValue,omitempty"`
}

type ClientPriceDefinition struct {
	Bids        []*PriceBucketDefinition `json:"bids,omitempty"`
	Asks        []*PriceBucketDefinition `json:"asks,omitempty"`
	CloseoutBid PriceValueDefinition     `json:"closeoutBid,omitempty"`
	CloseoutAsk PriceValueDefinition     `json:"closeoutAsk,omitempty"`
	Timestamp   DateTimeDefinition       `json:"timestamp,omitempty"`
}

type PricingHeartbeatDefinition struct {
	Type string             `json:"type,omitempty"`
	Time DateTimeDefinition `json:"time,omitempty"`
}

//
// Primitives Definitions
//

type DecimalNumberDefinition string

type AccountUnitsDefinition string

type CurrencyDefinition string

type InstrumentNameDefinition string

type InstrumentTypeDefinition string

type InstrumentDefinition struct {
	Name                        InstrumentNameDefinition        `json:"name,omitempty"`
	Type                        InstrumentTypeDefinition        `json:"type,omitempty"`
	DisplayName                 string                          `json:"displayName,omitempty"`
	PipLocation                 int                             `json:"pipLocation,omitempty"`
	DisplayPrecision            int                             `json:"displayPrecision,omitempty"`
	TradeUnitsPrecision         int                             `json:"tradeUnitsPrecision,omitempty"`
	MinimumTradeSize            DecimalNumberDefinition         `json:"minimumTradeSize,omitempty"`
	MaximumTrailingStopDistance DecimalNumberDefinition         `json:"maximumTrailingStopDistance,omitempty"`
	MinimumTrailingStopDistance DecimalNumberDefinition         `json:"minimumTrailingStopDistance,omitempty"`
	MaximumPositionSize         DecimalNumberDefinition         `json:"maximumPositionSize,omitempty"`
	MaximumOrderUnits           DecimalNumberDefinition         `json:"maximumOrderUnits,omitempty"`
	MarginRate                  DecimalNumberDefinition         `json:"marginRate,omitempty"`
	Commission                  *InstrumentCommissionDefinition `json:"commission,omitempty"`
}

type DateTimeDefinition string

type AcceptDatetimeFormatDefinition string

type InstrumentCommissionDefinition struct {
	Commission        DecimalNumberDefinition `json:"commission,omitempty"`
	UnitsTraded       DecimalNumberDefinition `json:"unitsTraded,omitempty"`
	MinimumCommission DecimalNumberDefinition `json:"minimumCommission,omitempty"`
}

type GuaranteedStopLossOrderLevelRestrictionDefinition struct {
	Volume     DecimalNumberDefinition `json:"volume,omitempty"`
	PriceRange DecimalNumberDefinition `json:"priceRange,omitempty"`
}

type DirectionDefinition string
