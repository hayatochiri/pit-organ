package pitOrgan

import (
	"encoding/json"
)

//
// Account Definitions
//

type AccountIDDefinition string

type AccountDefinition struct {
	ID                          AccountIDDefinition                   `json:"id"`
	Alias                       string                                `json:"alias"`
	Currency                    CurrencyDefinition                    `json:"currency"`
	Balance                     AccountUnitsDefinition                `json:"balance"`
	CreatedByUserID             int                                   `json:"createdByUserID"`
	CreatedTime                 DateTimeDefinition                    `json:"createdTime"`
	GuaranteedStopLossOrderMode GuaranteedStopLossOrderModeDefinition `json:"guaranteedStopLossOrderMode"`
	PL                          AccountUnitsDefinition                `json:"pl"`
	ResettablePL                AccountUnitsDefinition                `json:"resettablePL"`
	ResettablePLTime            DateTimeDefinition                    `json:"resettablePLTime"`
	Financing                   AccountUnitsDefinition                `json:"financing"`
	Commission                  AccountUnitsDefinition                `json:"commission"`
	GuaranteedExecutionFees     AccountUnitsDefinition                `json:"guaranteedExecutionFees"`
	MarginRate                  DecimalNumberDefinition               `json:"marginRate"`
	MarginCallEnterTime         DateTimeDefinition                    `json:"marginCallEnterTime"`
	MarginCallExtensionCount    int                                   `json:"marginCallExtensionCount"`
	LastMarginCallExtensionTime DateTimeDefinition                    `json:"lastMarginCallExtensionTime"`
	OpenTradeCount              int                                   `json:"openTradeCount"`
	OpenPositionCount           int                                   `json:"openPositionCount"`
	PendingOrderCount           int                                   `json:"pendingOrderCount"`
	HedgingEnabled              bool                                  `json:"hedgingEnabled"`
	UnrealizedPL                AccountUnitsDefinition                `json:"unrealizedPL"`
	NAV                         AccountUnitsDefinition                `json:"NAV"`
	MarginUsed                  AccountUnitsDefinition                `json:"marginUsed"`
	MarginAvailable             AccountUnitsDefinition                `json:"marginAvailable"`
	PositionValue               AccountUnitsDefinition                `json:"positionValue"`
	MarginCloseoutUnrealizedPL  AccountUnitsDefinition                `json:"marginCloseoutUnrealizedPL"`
	MarginCloseoutNAV           AccountUnitsDefinition                `json:"marginCloseoutNAV"`
	MarginCloseoutMarginUsed    AccountUnitsDefinition                `json:"marginCloseoutMarginUsed"`
	MarginCloseoutPercent       DecimalNumberDefinition               `json:"marginCloseoutPercent"`
	MarginCloseoutPositionValue DecimalNumberDefinition               `json:"marginCloseoutPositionValue"`
	WithdrawalLimit             AccountUnitsDefinition                `json:"withdrawalLimit"`
	MarginCallMarginUsed        AccountUnitsDefinition                `json:"marginCallMarginUsed"`
	MarginCallPercent           DecimalNumberDefinition               `json:"marginCallPercent"`
	LastTransactionID           TransactionIDDefinition               `json:"lastTransactionID"`
	Trades                      []TradeSummaryDefinition              `json:"trades"`
	Positions                   []PositionDefinition                  `json:"positions"`
	Orders                      []OrderDefinition                     `json:"orders"`
}

type AccountChangesStateDefinition struct {
	UnrealizedPL                AccountUnitsDefinition              `json:"unrealizedPL"`
	NAV                         AccountUnitsDefinition              `json:"NAV"`
	MarginUsed                  AccountUnitsDefinition              `json:"marginUsed"`
	MarginAvailable             AccountUnitsDefinition              `json:"marginAvailable"`
	PositionValue               AccountUnitsDefinition              `json:"positionValue"`
	MarginCloseoutUnrealizedPL  AccountUnitsDefinition              `json:"marginCloseoutUnrealizedPL"`
	MarginCloseoutNAV           AccountUnitsDefinition              `json:"marginCloseoutNAV"`
	MarginCloseoutMarginUsed    AccountUnitsDefinition              `json:"marginCloseoutMarginUsed"`
	MarginCloseoutPercent       DecimalNumberDefinition             `json:"marginCloseoutPercent"`
	MarginCloseoutPositionValue DecimalNumberDefinition             `json:"marginCloseoutPositionValue"`
	WithdrawalLimit             AccountUnitsDefinition              `json:"withdrawalLimit"`
	MarginCallMarginUsed        AccountUnitsDefinition              `json:"marginCallMarginUsed"`
	MarginCallPercent           DecimalNumberDefinition             `json:"marginCallPercent"`
	Orders                      []DynamicOrderStateDefinition       `json:"orders"`
	Trades                      []CalculatedTradeStateDefinition    `json:"trades"`
	Positions                   []CalculatedPositionStateDefinition `json:"positions"`
}

type AccountPropertiesDefinition struct {
	ID           AccountIDDefinition `json:"id"`
	MT4AccountID int                 `json:"mt4AccountID"`
	Tags         []string            `json:"tags"`
}

type GuaranteedStopLossOrderModeDefinition string

type AccountSummaryDefinition struct {
	Id                          AccountIDDefinition                   `json:"id"`
	Alias                       string                                `json:"alias"`
	Currency                    CurrencyDefinition                    `json:"currency"`
	Balance                     AccountUnitsDefinition                `json:"balance"`
	CreatedByUserID             int                                   `json:"createdByUserID"`
	CreatedTime                 DateTimeDefinition                    `json:"createdTime"`
	GuaranteedStopLossOrderMode GuaranteedStopLossOrderModeDefinition `json:"guaranteedStopLossOrderMode"`
	PL                          AccountUnitsDefinition                `json:"pl"`
	ResettablePL                AccountUnitsDefinition                `json:"resettablePL"`
	ResettablePLTime            DateTimeDefinition                    `json:"resettablePLTime"`
	Financing                   AccountUnitsDefinition                `json:"financing"`
	Commission                  AccountUnitsDefinition                `json:"commission"`
	GuaranteedExecutionFees     AccountUnitsDefinition                `json:"guaranteedExecutionFees"`
	MarginRate                  DecimalNumberDefinition               `json:"marginRate"`
	MarginCallEnterTime         DateTimeDefinition                    `json:"marginCallEnterTime"`
	MarginCallExtensionCount    int                                   `json:"marginCallExtensionCount"`
	LastMarginCallExtensionTime DateTimeDefinition                    `json:"lastMarginCallExtensionTime"`
	OpenTradeCount              int                                   `json:"openTradeCount"`
	OpenPositionCount           int                                   `json:"openPositionCount"`
	PendingOrderCount           int                                   `json:"pendingOrderCount"`
	HedgingEnabled              bool                                  `json:"hedgingEnabled"`
	UnrealizedPL                AccountUnitsDefinition                `json:"unrealizedPL"`
	NAV                         AccountUnitsDefinition                `json:"NAV"`
	MarginUsed                  AccountUnitsDefinition                `json:"marginUsed"`
	MarginAvailable             AccountUnitsDefinition                `json:"marginAvailable"`
	PositionValue               AccountUnitsDefinition                `json:"positionValue"`
	MarginCloseoutUnrealizedPL  AccountUnitsDefinition                `json:"marginCloseoutUnrealizedPL"`
	MarginCloseoutNAV           AccountUnitsDefinition                `json:"marginCloseoutNAV"`
	MarginCloseoutMarginUsed    AccountUnitsDefinition                `json:"marginCloseoutMarginUsed"`
	MarginCloseoutPercent       DecimalNumberDefinition               `json:"marginCloseoutPercent"`
	MarginCloseoutPositionValue DecimalNumberDefinition               `json:"marginCloseoutPositionValue"`
	WithdrawalLimit             AccountUnitsDefinition                `json:"withdrawalLimit"`
	MarginCallMarginUsed        AccountUnitsDefinition                `json:"marginCallMarginUsed"`
	MarginCallPercent           DecimalNumberDefinition               `json:"marginCallPercent"`
	LastTransactionID           TransactionIDDefinition               `json:"lastTransactionID"`
}

type CalculatedAccountStateDefinition struct {
	UnrealizedPL                AccountUnitsDefinition  `json:"unrealizedPL"`
	NAV                         AccountUnitsDefinition  `json:"NAV"`
	MarginUsed                  AccountUnitsDefinition  `json:"marginUsed"`
	MarginAvailable             AccountUnitsDefinition  `json:"marginAvailable"`
	PositionValue               AccountUnitsDefinition  `json:"positionValue"`
	MarginCloseoutUnrealizedPL  AccountUnitsDefinition  `json:"marginCloseoutUnrealizedPL"`
	MarginCloseoutNAV           AccountUnitsDefinition  `json:"marginCloseoutNAV"`
	MarginCloseoutMarginUsed    AccountUnitsDefinition  `json:"marginCloseoutMarginUsed"`
	MarginCloseoutPercent       DecimalNumberDefinition `json:"marginCloseoutPercent"`
	MarginCloseoutPositionValue DecimalNumberDefinition `json:"marginCloseoutPositionValue"`
	WithdrawalLimit             AccountUnitsDefinition  `json:"withdrawalLimit"`
	MarginCallMarginUsed        AccountUnitsDefinition  `json:"marginCallMarginUsed"`
	MarginCallPercent           DecimalNumberDefinition `json:"marginCallPercent"`
}

type AccountChangesDefinition struct {
	OrdersCreated   []OrderDefinition        `json:"ordersCreated"`
	OrdersCancelled []OrderDefinition        `json:"ordersCancelled"`
	OrdersFilled    []OrderDefinition        `json:"ordersFilled"`
	OrdersTriggered []OrderDefinition        `json:"ordersTriggered"`
	TradesOpened    []TradeSummaryDefinition `json:"tradesOpened"`
	TradesReduced   []TradeSummaryDefinition `json:"tradesReduced"`
	TradesClosed    []TradeSummaryDefinition `json:"tradesClosed"`
	Positions       []PositionDefinition     `json:"positions"`
	Transactions    []TransactionDefinition  `json:"transactions"`
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
	Instrument  InstrumentNameDefinition    `json:"instrument"`
	Time        DateTimeDefinition          `json:"time"`
	Price       PriceValueDefinition        `json:"price"`
	BucketWidth PriceValueDefinition        `json:"bucketWidth"`
	Buckets     []OrderBookBucketDefinition `json:"buckets"`
}

type OrderBookBucketDefinition struct {
	Price             PriceValueDefinition    `json:"price"`
	LongCountPercent  DecimalNumberDefinition `json:"longCountPercent"`
	ShortCountPercent DecimalNumberDefinition `json:"shortCountPercent"`
}

type PositionBookDefinition struct {
	Instrument  InstrumentNameDefinition       `json:"instrument"`
	Time        DateTimeDefinition             `json:"time"`
	Price       PriceValueDefinition           `json:"price"`
	BucketWidth PriceValueDefinition           `json:"bucketWidth"`
	Buckets     []PositionBookBucketDefinition `json:"buckets"`
}

type PositionBookBucketDefinition struct {
	Price             PriceValueDefinition    `json:"price"`
	LongCountPercent  DecimalNumberDefinition `json:"longCountPercent"`
	ShortCountPercent DecimalNumberDefinition `json:"shortCountPercent"`
}

//
// Order Definitions
//

// Orders

type OrderDefinition struct {
	Id               OrderIDDefinition          `json:"id"`
	CreateTime       DateTimeDefinition         `json:"createTime"`
	State            OrderStateDefinition       `json:"state"`
	ClientExtensions ClientExtensionsDefinition `json:"clientExtensions"`
}

type MarketOrderDefinition struct {
	Id                      OrderIDDefinition                      `json:"id"`
	CreateTime              DateTimeDefinition                     `json:"createTime"`
	State                   OrderStateDefinition                   `json:"state"`
	ClientExtensions        ClientExtensionsDefinition             `json:"clientExtensions"`
	Type                    OrderTypeDefinition                    `json:"type"`
	Instrument              InstrumentNameDefinition               `json:"instrument"`
	Units                   DecimalNumberDefinition                `json:"units"`
	TimeInForce             TimeInForceDefinition                  `json:"timeInForce"`
	PriceBound              PriceValueDefinition                   `json:"priceBound"`
	PositionFill            OrderPositionFillDefinition            `json:"positionFill"`
	TradeClose              MarketOrderTradeCloseDefinition        `json:"tradeClose"`
	LongPositionCloseout    MarketOrderPositionCloseoutDefinition  `json:"longPositionCloseout"`
	ShortPositionCloseout   MarketOrderPositionCloseoutDefinition  `json:"shortPositionCloseout"`
	MarginCloseout          MarketOrderMarginCloseoutDefinition    `json:"marginCloseout"`
	DelayedTradeClose       MarketOrderDelayedTradeCloseDefinition `json:"delayedTradeClose"`
	TakeProfitOnFill        TakeProfitDetailsDefinition            `json:"takeProfitOnFill"`
	StopLossOnFill          StopLossDetailsDefinition              `json:"stopLossOnFill"`
	TrailingStopLossOnFill  TrailingStopLossDetailsDefinition      `json:"trailingStopLossOnFill"`
	TradeClientExtensions   ClientExtensionsDefinition             `json:"tradeClientExtensions"`
	FillingTransactionID    TransactionIDDefinition                `json:"fillingTransactionID"`
	FilledTime              DateTimeDefinition                     `json:"filledTime"`
	TradeOpenedID           TradeIDDefinition                      `json:"tradeOpenedID"`
	TradeReducedID          TradeIDDefinition                      `json:"tradeReducedID"`
	TradeClosedIDs          []TradeIDDefinition                    `json:"tradeClosedIDs"`
	CancellingTransactionID TransactionIDDefinition                `json:"cancellingTransactionID"`
	CancelledTime           DateTimeDefinition                     `json:"cancelledTime"`
}

type FixedPriceOrderDefinition struct {
	Id                      OrderIDDefinition                 `json:"id"`
	CreateTime              DateTimeDefinition                `json:"createTime"`
	State                   OrderStateDefinition              `json:"state"`
	ClientExtensions        ClientExtensionsDefinition        `json:"clientExtensions"`
	Type                    OrderTypeDefinition               `json:"type"`
	Instrument              InstrumentNameDefinition          `json:"instrument"`
	Units                   DecimalNumberDefinition           `json:"units"`
	Price                   PriceValueDefinition              `json:"price"`
	PositionFill            OrderPositionFillDefinition       `json:"positionFill"`
	TradeState              string                            `json:"tradeState"`
	TakeProfitOnFill        TakeProfitDetailsDefinition       `json:"takeProfitOnFill"`
	StopLossOnFill          StopLossDetailsDefinition         `json:"stopLossOnFill"`
	TrailingStopLossOnFill  TrailingStopLossDetailsDefinition `json:"trailingStopLossOnFill"`
	TradeClientExtensions   ClientExtensionsDefinition        `json:"tradeClientExtensions"`
	FillingTransactionID    TransactionIDDefinition           `json:"fillingTransactionID"`
	FilledTime              DateTimeDefinition                `json:"filledTime"`
	TradeOpenedID           TradeIDDefinition                 `json:"tradeOpenedID"`
	TradeReducedID          TradeIDDefinition                 `json:"tradeReducedID"`
	TradeClosedIDs          []TradeIDDefinition               `json:"tradeClosedIDs"`
	CancellingTransactionID TransactionIDDefinition           `json:"cancellingTransactionID"`
	CancelledTime           DateTimeDefinition                `json:"cancelledTime"`
}

type LimitOrderDefinition struct {
	Id                      OrderIDDefinition                 `json:"id"`
	CreateTime              DateTimeDefinition                `json:"createTime"`
	State                   OrderStateDefinition              `json:"state"`
	ClientExtensions        ClientExtensionsDefinition        `json:"clientExtensions"`
	Type                    OrderTypeDefinition               `json:"type"`
	Instrument              InstrumentNameDefinition          `json:"instrument"`
	Units                   DecimalNumberDefinition           `json:"units"`
	Price                   PriceValueDefinition              `json:"price"`
	TimeInForce             TimeInForceDefinition             `json:"timeInForce"`
	GtdTime                 DateTimeDefinition                `json:"gtdTime"`
	PositionFill            OrderPositionFillDefinition       `json:"positionFill"`
	TriggerCondition        OrderTriggerConditionDefinition   `json:"triggerCondition"`
	TakeProfitOnFill        TakeProfitDetailsDefinition       `json:"takeProfitOnFill"`
	StopLossOnFill          StopLossDetailsDefinition         `json:"stopLossOnFill"`
	TrailingStopLossOnFill  TrailingStopLossDetailsDefinition `json:"trailingStopLossOnFill"`
	TradeClientExtensions   ClientExtensionsDefinition        `json:"tradeClientExtensions"`
	FillingTransactionID    TransactionIDDefinition           `json:"fillingTransactionID"`
	FilledTime              DateTimeDefinition                `json:"filledTime"`
	TradeOpenedID           TradeIDDefinition                 `json:"tradeOpenedID"`
	TradeReducedID          TradeIDDefinition                 `json:"tradeReducedID"`
	TradeClosedIDs          []TradeIDDefinition               `json:"tradeClosedIDs"`
	CancellingTransactionID TransactionIDDefinition           `json:"cancellingTransactionID"`
	CancelledTime           DateTimeDefinition                `json:"cancelledTime"`
	ReplacesOrderID         OrderIDDefinition                 `json:"replacesOrderID"`
	ReplacedByOrderID       OrderIDDefinition                 `json:"replacedByOrderID"`
}

type StopOrderDefinition struct {
	Id                      OrderIDDefinition                 `json:"id"`
	CreateTime              DateTimeDefinition                `json:"createTime"`
	State                   OrderStateDefinition              `json:"state"`
	ClientExtensions        ClientExtensionsDefinition        `json:"clientExtensions"`
	Type                    OrderTypeDefinition               `json:"type"`
	Instrument              InstrumentNameDefinition          `json:"instrument"`
	Units                   DecimalNumberDefinition           `json:"units"`
	Price                   PriceValueDefinition              `json:"price"`
	PriceBound              PriceValueDefinition              `json:"priceBound"`
	TimeInForce             TimeInForceDefinition             `json:"timeInForce"`
	GtdTime                 DateTimeDefinition                `json:"gtdTime"`
	PositionFill            OrderPositionFillDefinition       `json:"positionFill"`
	TriggerCondition        OrderTriggerConditionDefinition   `json:"triggerCondition"`
	TakeProfitOnFill        TakeProfitDetailsDefinition       `json:"takeProfitOnFill"`
	StopLossOnFill          StopLossDetailsDefinition         `json:"stopLossOnFill"`
	TrailingStopLossOnFill  TrailingStopLossDetailsDefinition `json:"trailingStopLossOnFill"`
	TradeClientExtensions   ClientExtensionsDefinition        `json:"tradeClientExtensions"`
	FillingTransactionID    TransactionIDDefinition           `json:"fillingTransactionID"`
	FilledTime              DateTimeDefinition                `json:"filledTime"`
	TradeOpenedID           TradeIDDefinition                 `json:"tradeOpenedID"`
	TradeReducedID          TradeIDDefinition                 `json:"tradeReducedID"`
	TradeClosedIDs          []TradeIDDefinition               `json:"tradeClosedIDs"`
	CancellingTransactionID TransactionIDDefinition           `json:"cancellingTransactionID"`
	CancelledTime           DateTimeDefinition                `json:"cancelledTime"`
	ReplacesOrderID         OrderIDDefinition                 `json:"replacesOrderID"`
	ReplacedByOrderID       OrderIDDefinition                 `json:"replacedByOrderID"`
}

type MarketIfTouchedOrderDefinition struct {
	Id                      OrderIDDefinition                 `json:"id"`
	CreateTime              DateTimeDefinition                `json:"createTime"`
	State                   OrderStateDefinition              `json:"state"`
	ClientExtensions        ClientExtensionsDefinition        `json:"clientExtensions"`
	Type                    OrderTypeDefinition               `json:"type"`
	Instrument              InstrumentNameDefinition          `json:"instrument"`
	Units                   DecimalNumberDefinition           `json:"units"`
	Price                   PriceValueDefinition              `json:"price"`
	PriceBound              PriceValueDefinition              `json:"priceBound"`
	TimeInForce             TimeInForceDefinition             `json:"timeInForce"`
	GtdTime                 DateTimeDefinition                `json:"gtdTime"`
	PositionFill            OrderPositionFillDefinition       `json:"positionFill"`
	TriggerCondition        OrderTriggerConditionDefinition   `json:"triggerCondition"`
	InitialMarketPrice      PriceValueDefinition              `json:"initialMarketPrice"`
	TakeProfitOnFill        TakeProfitDetailsDefinition       `json:"takeProfitOnFill"`
	StopLossOnFill          StopLossDetailsDefinition         `json:"stopLossOnFill"`
	TrailingStopLossOnFill  TrailingStopLossDetailsDefinition `json:"trailingStopLossOnFill"`
	TradeClientExtensions   ClientExtensionsDefinition        `json:"tradeClientExtensions"`
	FillingTransactionID    TransactionIDDefinition           `json:"fillingTransactionID"`
	FilledTime              DateTimeDefinition                `json:"filledTime"`
	TradeOpenedID           TradeIDDefinition                 `json:"tradeOpenedID"`
	TradeReducedID          TradeIDDefinition                 `json:"tradeReducedID"`
	TradeClosedIDs          []TradeIDDefinition               `json:"tradeClosedIDs"`
	CancellingTransactionID TransactionIDDefinition           `json:"cancellingTransactionID"`
	CancelledTime           DateTimeDefinition                `json:"cancelledTime"`
	ReplacesOrderID         OrderIDDefinition                 `json:"replacesOrderID"`
	ReplacedByOrderID       OrderIDDefinition                 `json:"replacedByOrderID"`
}

type TakeProfitOrderDefinition struct {
	Id                      OrderIDDefinition               `json:"id"`
	CreateTime              DateTimeDefinition              `json:"createTime"`
	State                   OrderStateDefinition            `json:"state"`
	ClientExtensions        ClientExtensionsDefinition      `json:"clientExtensions"`
	Type                    OrderTypeDefinition             `json:"type"`
	TradeID                 TradeIDDefinition               `json:"tradeID"`
	ClientTradeID           ClientIDDefinition              `json:"clientTradeID"`
	Price                   PriceValueDefinition            `json:"price"`
	TimeInForce             TimeInForceDefinition           `json:"timeInForce"`
	GtdTime                 DateTimeDefinition              `json:"gtdTime"`
	TriggerCondition        OrderTriggerConditionDefinition `json:"triggerCondition"`
	FillingTransactionID    TransactionIDDefinition         `json:"fillingTransactionID"`
	FilledTime              DateTimeDefinition              `json:"filledTime"`
	TradeOpenedID           TradeIDDefinition               `json:"tradeOpenedID"`
	TradeReducedID          TradeIDDefinition               `json:"tradeReducedID"`
	TradeClosedIDs          []TradeIDDefinition             `json:"tradeClosedIDs"`
	CancellingTransactionID TransactionIDDefinition         `json:"cancellingTransactionID"`
	CancelledTime           DateTimeDefinition              `json:"cancelledTime"`
	ReplacesOrderID         OrderIDDefinition               `json:"replacesOrderID"`
	ReplacedByOrderID       OrderIDDefinition               `json:"replacedByOrderID"`
}

type StopLossOrderDefinition struct {
	Id                         OrderIDDefinition               `json:"id"`
	CreateTime                 DateTimeDefinition              `json:"createTime"`
	State                      OrderStateDefinition            `json:"state"`
	ClientExtensions           ClientExtensionsDefinition      `json:"clientExtensions"`
	Type                       OrderTypeDefinition             `json:"type"`
	GuaranteedExecutionPremium DecimalNumberDefinition         `json:"guaranteedExecutionPremium"`
	TradeID                    TradeIDDefinition               `json:"tradeID"`
	ClientTradeID              ClientIDDefinition              `json:"clientTradeID"`
	Price                      PriceValueDefinition            `json:"price"`
	Distance                   DecimalNumberDefinition         `json:"distance"`
	TimeInForce                TimeInForceDefinition           `json:"timeInForce"`
	GtdTime                    DateTimeDefinition              `json:"gtdTime"`
	TriggerCondition           OrderTriggerConditionDefinition `json:"triggerCondition"`
	Guaranteed                 bool                            `json:"guaranteed"`
	FillingTransactionID       TransactionIDDefinition         `json:"fillingTransactionID"`
	FilledTime                 DateTimeDefinition              `json:"filledTime"`
	TradeOpenedID              TradeIDDefinition               `json:"tradeOpenedID"`
	TradeReducedID             TradeIDDefinition               `json:"tradeReducedID"`
	TradeClosedIDs             []TradeIDDefinition             `json:"tradeClosedIDs"`
	CancellingTransactionID    TransactionIDDefinition         `json:"cancellingTransactionID"`
	CancelledTime              DateTimeDefinition              `json:"cancelledTime"`
	ReplacesOrderID            OrderIDDefinition               `json:"replacesOrderID"`
	ReplacedByOrderID          OrderIDDefinition               `json:"replacedByOrderID"`
}

type TrailingStopLossOrderDefinition struct {
	Id                      OrderIDDefinition               `json:"id"`
	CreateTime              DateTimeDefinition              `json:"createTime"`
	State                   OrderStateDefinition            `json:"state"`
	ClientExtensions        ClientExtensionsDefinition      `json:"clientExtensions"`
	Type                    OrderTypeDefinition             `json:"type"`
	TradeID                 TradeIDDefinition               `json:"tradeID"`
	ClientTradeID           ClientIDDefinition              `json:"clientTradeID"`
	Distance                DecimalNumberDefinition         `json:"distance"`
	TimeInForce             TimeInForceDefinition           `json:"timeInForce"`
	GtdTime                 DateTimeDefinition              `json:"gtdTime"`
	TriggerCondition        OrderTriggerConditionDefinition `json:"triggerCondition"`
	TrailingStopValue       PriceValueDefinition            `json:"trailingStopValue"`
	FillingTransactionID    TransactionIDDefinition         `json:"fillingTransactionID"`
	FilledTime              DateTimeDefinition              `json:"filledTime"`
	TradeOpenedID           TradeIDDefinition               `json:"tradeOpenedID"`
	TradeReducedID          TradeIDDefinition               `json:"tradeReducedID"`
	TradeClosedIDs          []TradeIDDefinition             `json:"tradeClosedIDs"`
	CancellingTransactionID TransactionIDDefinition         `json:"cancellingTransactionID"`
	CancelledTime           DateTimeDefinition              `json:"cancelledTime"`
	ReplacesOrderID         OrderIDDefinition               `json:"replacesOrderID"`
	ReplacedByOrderID       OrderIDDefinition               `json:"replacedByOrderID"`
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
	Type                   OrderTypeDefinition               `json:"type"`
	Instrument             InstrumentNameDefinition          `json:"instrument"`
	Units                  DecimalNumberDefinition           `json:"units"`
	Price                  PriceValueDefinition              `json:"price"`
	TimeInForce            TimeInForceDefinition             `json:"timeInForce"`
	GtdTime                DateTimeDefinition                `json:"gtdTime"`
	PositionFill           OrderPositionFillDefinition       `json:"positionFill"`
	TriggerCondition       OrderTriggerConditionDefinition   `json:"triggerCondition"`
	ClientExtensions       ClientExtensionsDefinition        `json:"clientExtensions"`
	TakeProfitOnFill       TakeProfitDetailsDefinition       `json:"takeProfitOnFill"`
	StopLossOnFill         StopLossDetailsDefinition         `json:"stopLossOnFill"`
	TrailingStopLossOnFill TrailingStopLossDetailsDefinition `json:"trailingStopLossOnFill"`
	TradeClientExtensions  ClientExtensionsDefinition        `json:"tradeClientExtensions"`
}

type StopOrderRequestDefinition struct {
	Type                   OrderTypeDefinition               `json:"type"`
	Instrument             InstrumentNameDefinition          `json:"instrument"`
	Units                  DecimalNumberDefinition           `json:"units"`
	Price                  PriceValueDefinition              `json:"price"`
	PriceBound             PriceValueDefinition              `json:"priceBound"`
	TimeInForce            TimeInForceDefinition             `json:"timeInForce"`
	GtdTime                DateTimeDefinition                `json:"gtdTime"`
	PositionFill           OrderPositionFillDefinition       `json:"positionFill"`
	TriggerCondition       OrderTriggerConditionDefinition   `json:"triggerCondition"`
	ClientExtensions       ClientExtensionsDefinition        `json:"clientExtensions"`
	TakeProfitOnFill       TakeProfitDetailsDefinition       `json:"takeProfitOnFill"`
	StopLossOnFill         StopLossDetailsDefinition         `json:"stopLossOnFill"`
	TrailingStopLossOnFill TrailingStopLossDetailsDefinition `json:"trailingStopLossOnFill"`
	TradeClientExtensions  ClientExtensionsDefinition        `json:"tradeClientExtensions"`
}

type MarketIfTouchedOrderRequestDefinition struct {
	Type                   OrderTypeDefinition               `json:"type"`
	Instrument             InstrumentNameDefinition          `json:"instrument"`
	Units                  DecimalNumberDefinition           `json:"units"`
	Price                  PriceValueDefinition              `json:"price"`
	PriceBound             PriceValueDefinition              `json:"priceBound"`
	TimeInForce            TimeInForceDefinition             `json:"timeInForce"`
	GtdTime                DateTimeDefinition                `json:"gtdTime"`
	PositionFill           OrderPositionFillDefinition       `json:"positionFill"`
	TriggerCondition       OrderTriggerConditionDefinition   `json:"triggerCondition"`
	ClientExtensions       ClientExtensionsDefinition        `json:"clientExtensions"`
	TakeProfitOnFill       TakeProfitDetailsDefinition       `json:"takeProfitOnFill"`
	StopLossOnFill         StopLossDetailsDefinition         `json:"stopLossOnFill"`
	TrailingStopLossOnFill TrailingStopLossDetailsDefinition `json:"trailingStopLossOnFill"`
	TradeClientExtensions  ClientExtensionsDefinition        `json:"tradeClientExtensions"`
}

type TakeProfitOrderRequestDefinition struct {
	Type             OrderTypeDefinition             `json:"type"`
	TradeID          TradeIDDefinition               `json:"tradeID"`
	ClientTradeID    ClientIDDefinition              `json:"clientTradeID"`
	Price            PriceValueDefinition            `json:"price"`
	TimeInForce      TimeInForceDefinition           `json:"timeInForce"`
	GtdTime          DateTimeDefinition              `json:"gtdTime"`
	TriggerCondition OrderTriggerConditionDefinition `json:"triggerCondition"`
	ClientExtensions ClientExtensionsDefinition      `json:"clientExtensions"`
}

type StopLossOrderRequestDefinition struct {
	Type             OrderTypeDefinition             `json:"type"`
	TradeID          TradeIDDefinition               `json:"tradeID"`
	ClientTradeID    ClientIDDefinition              `json:"clientTradeID"`
	Price            PriceValueDefinition            `json:"price"`
	Distance         DecimalNumberDefinition         `json:"distance"`
	TimeInForce      TimeInForceDefinition           `json:"timeInForce"`
	GtdTime          DateTimeDefinition              `json:"gtdTime"`
	TriggerCondition OrderTriggerConditionDefinition `json:"triggerCondition"`
	Guaranteed       bool                            `json:"guaranteed"`
	ClientExtensions ClientExtensionsDefinition      `json:"clientExtensions"`
}

type TrailingStopLossOrderRequestDefinition struct {
	Type             OrderTypeDefinition             `json:"type"`
	TradeID          TradeIDDefinition               `json:"tradeID"`
	ClientTradeID    ClientIDDefinition              `json:"clientTradeID"`
	Distance         DecimalNumberDefinition         `json:"distance"`
	TimeInForce      TimeInForceDefinition           `json:"timeInForce"`
	GtdTime          DateTimeDefinition              `json:"gtdTime"`
	TriggerCondition OrderTriggerConditionDefinition `json:"triggerCondition"`
	ClientExtensions ClientExtensionsDefinition      `json:"clientExtensions"`
}

// Order-related Definitions

type OrderIDDefinition string

type OrderTypeDefinition string

type CancellableOrderTypeDefinition string

type OrderStateDefinition string

type OrderStateFilterDefinition string

type OrderIdentifierDefinition struct {
	OrderID       OrderIDDefinition  `json:"orderID"`
	ClientOrderID ClientIDDefinition `json:"clientOrderID"`
}

type OrderSpecifierDefinition string

type TimeInForceDefinition string

type OrderPositionFillDefinition string

type OrderTriggerConditionDefinition string

type DynamicOrderStateDefinition struct {
	Id                     OrderIDDefinition    `json:"id"`
	TrailingStopValue      PriceValueDefinition `json:"trailingStopValue"`
	TriggerDistance        PriceValueDefinition `json:"triggerDistance"`
	IsTriggerDistanceExact bool                 `json:"isTriggerDistanceExact"`
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
	MinimumDistance  DecimalNumberDefinition                           `json:"minimumDistance"`
	Premium          DecimalNumberDefinition                           `json:"premium"`
	LevelRestriction GuaranteedStopLossOrderLevelRestrictionDefinition `json:"levelRestriction"`
}

//
// Trade Definitions
//

type TradeIDDefinition string

type TradeStateDefinition string

type TradeStateFilterDefinition string

type TradeSpecifierDefinition string

type TradeDefinition struct {
	Id                    TradeIDDefinition               `json:"id"`
	Instrument            InstrumentNameDefinition        `json:"instrument"`
	Price                 PriceValueDefinition            `json:"price"`
	OpenTime              DateTimeDefinition              `json:"openTime"`
	State                 TradeStateDefinition            `json:"state"`
	InitialUnits          DecimalNumberDefinition         `json:"initialUnits"`
	InitialMarginRequired AccountUnitsDefinition          `json:"initialMarginRequired"`
	CurrentUnits          DecimalNumberDefinition         `json:"currentUnits"`
	RealizedPL            AccountUnitsDefinition          `json:"realizedPL"`
	UnrealizedPL          AccountUnitsDefinition          `json:"unrealizedPL"`
	MarginUsed            AccountUnitsDefinition          `json:"marginUsed"`
	AverageClosePrice     PriceValueDefinition            `json:"averageClosePrice"`
	ClosingTransactionIDs []TransactionIDDefinition       `json:"closingTransactionIDs"`
	Financing             AccountUnitsDefinition          `json:"financing"`
	CloseTime             DateTimeDefinition              `json:"closeTime"`
	ClientExtensions      ClientExtensionsDefinition      `json:"clientExtensions"`
	TakeProfitOrder       TakeProfitOrderDefinition       `json:"takeProfitOrder"`
	StopLossOrder         StopLossOrderDefinition         `json:"stopLossOrder"`
	TrailingStopLossOrder TrailingStopLossOrderDefinition `json:"trailingStopLossOrder"`
}

type TradeSummaryDefinition struct {
	Id                      TradeIDDefinition          `json:"id"`
	Instrument              InstrumentNameDefinition   `json:"instrument"`
	Price                   PriceValueDefinition       `json:"price"`
	OpenTime                DateTimeDefinition         `json:"openTime"`
	State                   TradeStateDefinition       `json:"state"`
	InitialUnits            DecimalNumberDefinition    `json:"initialUnits"`
	InitialMarginRequired   AccountUnitsDefinition     `json:"initialMarginRequired"`
	CurrentUnits            DecimalNumberDefinition    `json:"currentUnits"`
	RealizedPL              AccountUnitsDefinition     `json:"realizedPL"`
	UnrealizedPL            AccountUnitsDefinition     `json:"unrealizedPL"`
	MarginUsed              AccountUnitsDefinition     `json:"marginUsed"`
	AverageClosePrice       PriceValueDefinition       `json:"averageClosePrice"`
	ClosingTransactionIDs   []TransactionIDDefinition  `json:"closingTransactionIDs"`
	Financing               AccountUnitsDefinition     `json:"financing"`
	CloseTime               DateTimeDefinition         `json:"closeTime"`
	ClientExtensions        ClientExtensionsDefinition `json:"clientExtensions"`
	TakeProfitOrderID       OrderIDDefinition          `json:"takeProfitOrderID"`
	StopLossOrderID         OrderIDDefinition          `json:"stopLossOrderID"`
	TrailingStopLossOrderID OrderIDDefinition          `json:"trailingStopLossOrderID"`
}

type CalculatedTradeStateDefinition struct {
	Id           TradeIDDefinition      `json:"id"`
	UnrealizedPL AccountUnitsDefinition `json:"unrealizedPL"`
	MarginUsed   AccountUnitsDefinition `json:"marginUsed"`
}

type TradePLDefinition string

//
// Position Definitions
//

type PositionDefinition struct {
	Instrument              InstrumentNameDefinition `json:"instrument"`
	PL                      AccountUnitsDefinition   `json:"pl"`
	UnrealizedPL            AccountUnitsDefinition   `json:"unrealizedPL"`
	MarginUsed              AccountUnitsDefinition   `json:"marginUsed"`
	ResettablePL            AccountUnitsDefinition   `json:"resettablePL"`
	Financing               AccountUnitsDefinition   `json:"financing"`
	Commission              AccountUnitsDefinition   `json:"commission"`
	GuaranteedExecutionFees AccountUnitsDefinition   `json:"guaranteedExecutionFees"`
	Long                    PositionSideDefinition   `json:"long"`
	Short                   PositionSideDefinition   `json:"short"`
}

type PositionSideDefinition struct {
	Units                   DecimalNumberDefinition `json:"units"`
	AveragePrice            PriceValueDefinition    `json:"averagePrice"`
	TradeIDs                []TradeIDDefinition     `json:"tradeIDs"`
	PL                      AccountUnitsDefinition  `json:"pl"`
	UnrealizedPL            AccountUnitsDefinition  `json:"unrealizedPL"`
	ResettablePL            AccountUnitsDefinition  `json:"resettablePL"`
	Financing               AccountUnitsDefinition  `json:"financing"`
	GuaranteedExecutionFees AccountUnitsDefinition  `json:"guaranteedExecutionFees"`
}

type CalculatedPositionStateDefinition struct {
	Instrument        InstrumentNameDefinition `json:"instrument"`
	NetUnrealizedPL   AccountUnitsDefinition   `json:"netUnrealizedPL"`
	LongUnrealizedPL  AccountUnitsDefinition   `json:"longUnrealizedPL"`
	ShortUnrealizedPL AccountUnitsDefinition   `json:"shortUnrealizedPL"`
	MarginUsed        AccountUnitsDefinition   `json:"marginUsed"`
}

//
// Transaction Definitions
//

// Transactions

// TODO: Implemented by: OrderFillTransaction, OrderCancelTransaction, OrderCancelRejectTransaction, OrderClientExtensionsModifyTransaction, OrderClientExtensionsModifyRejectTransaction, CreateTransaction, CloseTransaction, ReopenTransaction, ClientConfigureTransaction, ClientConfigureRejectTransaction, TransferFundsTransaction, TransferFundsRejectTransaction, MarketOrderTransaction, MarketOrderRejectTransaction, FixedPriceOrderTransaction, LimitOrderTransaction, LimitOrderRejectTransaction, StopOrderTransaction, StopOrderRejectTransaction, MarketIfTouchedOrderTransaction, MarketIfTouchedOrderRejectTransaction, TakeProfitOrderTransaction, TakeProfitOrderRejectTransaction, StopLossOrderTransaction, StopLossOrderRejectTransaction, TrailingStopLossOrderTransaction, TrailingStopLossOrderRejectTransaction, TradeClientExtensionsModifyTransaction, TradeClientExtensionsModifyRejectTransaction, MarginCallEnterTransaction, MarginCallExtendTransaction, MarginCallExitTransaction, DelayedTradeClosureTransaction, DailyFinancingTransaction, ResetResettablePLTransaction
type TransactionDefinition struct {
	Id        TransactionIDDefinition   `json:"id,omitempty"`
	Time      DateTimeDefinition        `json:"time,omitempty"`
	UserID    int                       `json:"userID,omitempty"`
	AccountID AccountIDDefinition       `json:"accountID,omitempty"`
	BatchID   TransactionIDDefinition   `json:"batchID,omitempty"`
	RequestID RequestIDDefinition       `json:"requestID,omitempty"`
	Type      TransactionTypeDefinition `json:"type,omitempty"`
}

type CreateTransactionDefinition struct {
	Id            TransactionIDDefinition   `json:"id,omitempty"`
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
	Id        TransactionIDDefinition   `json:"id,omitempty"`
	Time      DateTimeDefinition        `json:"time,omitempty"`
	UserID    int                       `json:"userID,omitempty"`
	AccountID AccountIDDefinition       `json:"accountID,omitempty"`
	BatchID   TransactionIDDefinition   `json:"batchID,omitempty"`
	RequestID RequestIDDefinition       `json:"requestID,omitempty"`
	Type      TransactionTypeDefinition `json:"type,omitempty"`
}

type ReopenTransactionDefinition struct {
	Id        TransactionIDDefinition   `json:"id,omitempty"`
	Time      DateTimeDefinition        `json:"time,omitempty"`
	UserID    int                       `json:"userID,omitempty"`
	AccountID AccountIDDefinition       `json:"accountID,omitempty"`
	BatchID   TransactionIDDefinition   `json:"batchID,omitempty"`
	RequestID RequestIDDefinition       `json:"requestID,omitempty"`
	Type      TransactionTypeDefinition `json:"type,omitempty"`
}

type ClientConfigureTransactionDefinition struct {
	Id         TransactionIDDefinition   `json:"id,omitempty"`
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
	Id           TransactionIDDefinition           `json:"id,omitempty"`
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
	Id             TransactionIDDefinition   `json:"id,omitempty"`
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
	Id            TransactionIDDefinition           `json:"id,omitempty"`
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
	Id                     TransactionIDDefinition                 `json:"id,omitempty"`
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
	Id                     TransactionIDDefinition                 `json:"id,omitempty"`
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
	Id                     TransactionIDDefinition            `json:"id,omitempty"`
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
	Id                      TransactionIDDefinition            `json:"id,omitempty"`
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
	Id                      TransactionIDDefinition            `json:"id,omitempty"`
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
	Id                      TransactionIDDefinition            `json:"id,omitempty"`
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
	Id                      TransactionIDDefinition            `json:"id,omitempty"`
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
	Id                      TransactionIDDefinition              `json:"id,omitempty"`
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
	Id                      TransactionIDDefinition              `json:"id,omitempty"`
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
	Id                      TransactionIDDefinition         `json:"id,omitempty"`
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
	Id                      TransactionIDDefinition           `json:"id,omitempty"`
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
	Id                         TransactionIDDefinition         `json:"id,omitempty"`
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
	Id                      TransactionIDDefinition           `json:"id,omitempty"`
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
	Id                      TransactionIDDefinition               `json:"id,omitempty"`
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
	Id                      TransactionIDDefinition               `json:"id,omitempty"`
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
	Id                            TransactionIDDefinition   `json:"id,omitempty"`
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
	Id                TransactionIDDefinition     `json:"id,omitempty"`
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
	Id            TransactionIDDefinition           `json:"id,omitempty"`
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
	Id                          TransactionIDDefinition     `json:"id,omitempty"`
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
	Id                          TransactionIDDefinition           `json:"id,omitempty"`
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
	Id                          TransactionIDDefinition     `json:"id,omitempty"`
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
	Id                          TransactionIDDefinition           `json:"id,omitempty"`
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
	Id        TransactionIDDefinition   `json:"id,omitempty"`
	Time      DateTimeDefinition        `json:"time,omitempty"`
	UserID    int                       `json:"userID,omitempty"`
	AccountID AccountIDDefinition       `json:"accountID,omitempty"`
	BatchID   TransactionIDDefinition   `json:"batchID,omitempty"`
	RequestID RequestIDDefinition       `json:"requestID,omitempty"`
	Type      TransactionTypeDefinition `json:"type,omitempty"`
}

type MarginCallExtendTransactionDefinition struct {
	Id              TransactionIDDefinition   `json:"id,omitempty"`
	Time            DateTimeDefinition        `json:"time,omitempty"`
	UserID          int                       `json:"userID,omitempty"`
	AccountID       AccountIDDefinition       `json:"accountID,omitempty"`
	BatchID         TransactionIDDefinition   `json:"batchID,omitempty"`
	RequestID       RequestIDDefinition       `json:"requestID,omitempty"`
	Type            TransactionTypeDefinition `json:"type,omitempty"`
	ExtensionNumber int                       `json:"extensionNumber,omitempty"`
}

type MarginCallExitTransactionDefinition struct {
	Id        TransactionIDDefinition   `json:"id,omitempty"`
	Time      DateTimeDefinition        `json:"time,omitempty"`
	UserID    int                       `json:"userID,omitempty"`
	AccountID AccountIDDefinition       `json:"accountID,omitempty"`
	BatchID   TransactionIDDefinition   `json:"batchID,omitempty"`
	RequestID RequestIDDefinition       `json:"requestID,omitempty"`
	Type      TransactionTypeDefinition `json:"type,omitempty"`
}

type DelayedTradeClosureTransactionDefinition struct {
	Id        TransactionIDDefinition     `json:"id,omitempty"`
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
	Id                   TransactionIDDefinition        `json:"id,omitempty"`
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
	Id        TransactionIDDefinition   `json:"id,omitempty"`
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
	Id      ClientIDDefinition      `json:"id,omitempty"`
	Tag     ClientTagDefinition     `json:"tag,omitempty"`
	Comment ClientCommentDefinition `json:"comment,omitempty"`
}

type TakeProfitDetailsDefinition struct {
	Price            PriceValueDefinition       `json:"price,omitempty"`
	TimeInForce      TimeInForceDefinition      `json:"timeInForce,omitempty"`
	GtdTime          DateTimeDefinition         `json:"gtdTime,omitempty"`
	ClientExtensions ClientExtensionsDefinition `json:"clientExtensions,omitempty"`
}

type StopLossDetailsDefinition struct {
	Price            PriceValueDefinition       `json:"price,omitempty"`
	Distance         DecimalNumberDefinition    `json:"distance,omitempty"`
	TimeInForce      TimeInForceDefinition      `json:"timeInForce,omitempty"`
	GtdTime          DateTimeDefinition         `json:"gtdTime,omitempty"`
	ClientExtensions ClientExtensionsDefinition `json:"clientExtensions,omitempty"`
	Guaranteed       bool                       `json:"guaranteed,omitempty"`
}

type TrailingStopLossDetailsDefinition struct {
	Distance         DecimalNumberDefinition    `json:"distance,omitempty"`
	TimeInForce      TimeInForceDefinition      `json:"timeInForce,omitempty"`
	GtdTime          DateTimeDefinition         `json:"gtdTime,omitempty"`
	ClientExtensions ClientExtensionsDefinition `json:"clientExtensions,omitempty"`
}

type TradeOpenDefinition struct {
	TradeID                TradeIDDefinition          `json:"tradeID"`
	Units                  DecimalNumberDefinition    `json:"units"`
	Price                  PriceValueDefinition       `json:"price"`
	GuaranteedExecutionFee AccountUnitsDefinition     `json:"guaranteedExecutionFee"`
	ClientExtensions       ClientExtensionsDefinition `json:"clientExtensions"`
	HalfSpreadCost         AccountUnitsDefinition     `json:"halfSpreadCost"`
	InitialMarginRequired  AccountUnitsDefinition     `json:"initialMarginRequired"`
}

type TradeReduceDefinition struct {
	TradeID                TradeIDDefinition       `json:"tradeID"`
	Units                  DecimalNumberDefinition `json:"units"`
	Price                  PriceValueDefinition    `json:"price"`
	RealizedPL             AccountUnitsDefinition  `json:"realizedPL"`
	Financing              AccountUnitsDefinition  `json:"financing"`
	GuaranteedExecutionFee AccountUnitsDefinition  `json:"guaranteedExecutionFee"`
	HalfSpreadCost         AccountUnitsDefinition  `json:"halfSpreadCost"`
}

type MarketOrderTradeCloseDefinition struct {
	TradeID       TradeIDDefinition `json:"tradeID"`
	ClientTradeID string            `json:"clientTradeID"`
	Units         string            `json:"units"`
}

type MarketOrderMarginCloseoutDefinition struct {
	Reason MarketOrderMarginCloseoutReasonDefinition `json:"reason"`
}

type MarketOrderMarginCloseoutReasonDefinition string

type MarketOrderDelayedTradeCloseDefinition struct {
	TradeID             TradeIDDefinition       `json:"tradeID"`
	ClientTradeID       TradeIDDefinition       `json:"clientTradeID"`
	SourceTransactionID TransactionIDDefinition `json:"sourceTransactionID"`
}

type MarketOrderPositionCloseoutDefinition struct {
	Instrument InstrumentNameDefinition `json:"instrument"`
	Units      string                   `json:"units"`
}

type LiquidityRegenerationScheduleDefinition struct {
	Steps []LiquidityRegenerationScheduleStepDefinition `json:"steps"`
}

type LiquidityRegenerationScheduleStepDefinition struct {
	Timestamp        DateTimeDefinition      `json:"timestamp"`
	BidLiquidityUsed DecimalNumberDefinition `json:"bidLiquidityUsed"`
	AskLiquidityUsed DecimalNumberDefinition `json:"askLiquidityUsed"`
}

type OpenTradeFinancingDefinition struct {
	TradeID   TradeIDDefinition      `json:"tradeID"`
	Financing AccountUnitsDefinition `json:"financing"`
}

type PositionFinancingDefinition struct {
	Instrument          InstrumentNameDefinition       `json:"instrument"`
	Financing           AccountUnitsDefinition         `json:"financing"`
	OpenTradeFinancings []OpenTradeFinancingDefinition `json:"openTradeFinancings"`
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
	Type              string                  `json:"type"`
	LastTransactionID TransactionIDDefinition `json:"lastTransactionID"`
	Time              DateTimeDefinition      `json:"time"`
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
	Currency      CurrencyDefinition      `json:"currency"`
	AccountGain   DecimalNumberDefinition `json:"accountGain"`
	AccountLoss   DecimalNumberDefinition `json:"accountLoss"`
	PositionValue DecimalNumberDefinition `json:"positionValue"`
}

type ClientPriceDefinition struct {
	Bids        []PriceBucketDefinition `json:"bids"`
	Asks        []PriceBucketDefinition `json:"asks"`
	CloseoutBid PriceValueDefinition    `json:"closeoutBid"`
	CloseoutAsk PriceValueDefinition    `json:"closeoutAsk"`
	Timestamp   DateTimeDefinition      `json:"timestamp"`
}

type PricingHeartbeatDefinition struct {
	Type string             `json:"type"`
	Time DateTimeDefinition `json:"time"`
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
	Name                        InstrumentNameDefinition       `json:"name"`
	Type                        InstrumentTypeDefinition       `json:"type"`
	DisplayName                 string                         `json:"displayName"`
	PipLocation                 int                            `json:"pipLocation"`
	DisplayPrecision            int                            `json:"displayPrecision"`
	TradeUnitsPrecision         int                            `json:"tradeUnitsPrecision"`
	MinimumTradeSize            DecimalNumberDefinition        `json:"minimumTradeSize"`
	MaximumTrailingStopDistance DecimalNumberDefinition        `json:"maximumTrailingStopDistance"`
	MinimumTrailingStopDistance DecimalNumberDefinition        `json:"minimumTrailingStopDistance"`
	MaximumPositionSize         DecimalNumberDefinition        `json:"maximumPositionSize"`
	MaximumOrderUnits           DecimalNumberDefinition        `json:"maximumOrderUnits"`
	MarginRate                  DecimalNumberDefinition        `json:"marginRate"`
	Commission                  InstrumentCommissionDefinition `json:"commission"`
}

type DateTimeDefinition string

type AcceptDatetimeFormatDefinition string

type InstrumentCommissionDefinition struct {
	Commission        DecimalNumberDefinition `json:"commission"`
	UnitsTraded       DecimalNumberDefinition `json:"unitsTraded"`
	MinimumCommission DecimalNumberDefinition `json:"minimumCommission"`
}

type GuaranteedStopLossOrderLevelRestrictionDefinition struct {
	Volume     DecimalNumberDefinition `json:"volume"`
	PriceRange DecimalNumberDefinition `json:"priceRange"`
}

type DirectionDefinition string
