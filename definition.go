package pitOrgan

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
	Pl                          AccountUnitsDefinition                `json:"pl"`
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
	Pl                          AccountUnitsDefinition                `json:"pl"`
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

type WeeklyAlignmentDefinition string

type CandlestickDefinition struct {
	Time     DateTimeDefinition        `json:"time"`
	Bid      CandlestickDataDefinition `json:"bid"`
	Ask      CandlestickDataDefinition `json:"ask"`
	Mid      CandlestickDataDefinition `json:"mid"`
	Volume   int                       `json:"volume"`
	Complete bool                      `json:"complete"`
}

type CandlestickDataDefinition struct {
	O PriceValueDefinition `json:"o"`
	H PriceValueDefinition `json:"h"`
	L PriceValueDefinition `json:"l"`
	C PriceValueDefinition `json:"c"`
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
	Long  DecimalNumberDefinition `json:"long"`
	Short DecimalNumberDefinition `json:"short"`
}

type UnitsAvailableDefinition struct {
	Default     UnitsAvailableDetailsDefinition `json:"default"`
	ReduceFirst UnitsAvailableDetailsDefinition `json:"reduceFirst"`
	ReduceOnly  UnitsAvailableDetailsDefinition `json:"reduceOnly"`
	OpenOnly    UnitsAvailableDetailsDefinition `json:"openOnly"`
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
	Pl                      AccountUnitsDefinition   `json:"pl"`
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
	Pl                      AccountUnitsDefinition  `json:"pl"`
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
	Id        TransactionIDDefinition `json:"id"`
	Time      DateTimeDefinition      `json:"time"`
	UserID    int                     `json:"userID"`
	AccountID AccountIDDefinition     `json:"accountID"`
	BatchID   TransactionIDDefinition `json:"batchID"`
	RequestID RequestIDDefinition     `json:"requestID"`
}

type CreateTransactionDefinition struct {
	Id            TransactionIDDefinition   `json:"id"`
	Time          DateTimeDefinition        `json:"time"`
	UserID        int                       `json:"userID"`
	AccountID     AccountIDDefinition       `json:"accountID"`
	BatchID       TransactionIDDefinition   `json:"batchID"`
	RequestID     RequestIDDefinition       `json:"requestID"`
	Type          TransactionTypeDefinition `json:"type"`
	DivisionID    int                       `json:"divisionID"`
	SiteID        int                       `json:"siteID"`
	AccountUserID int                       `json:"accountUserID"`
	AccountNumber int                       `json:"accountNumber"`
	HomeCurrency  CurrencyDefinition        `json:"homeCurrency"`
}

type CloseTransactionDefinition struct {
	Id        TransactionIDDefinition   `json:"id"`
	Time      DateTimeDefinition        `json:"time"`
	UserID    int                       `json:"userID"`
	AccountID AccountIDDefinition       `json:"accountID"`
	BatchID   TransactionIDDefinition   `json:"batchID"`
	RequestID RequestIDDefinition       `json:"requestID"`
	Type      TransactionTypeDefinition `json:"type"`
}

type ReopenTransactionDefinition struct {
	Id        TransactionIDDefinition   `json:"id"`
	Time      DateTimeDefinition        `json:"time"`
	UserID    int                       `json:"userID"`
	AccountID AccountIDDefinition       `json:"accountID"`
	BatchID   TransactionIDDefinition   `json:"batchID"`
	RequestID RequestIDDefinition       `json:"requestID"`
	Type      TransactionTypeDefinition `json:"type"`
}

type ClientConfigureTransactionDefinition struct {
	Id         TransactionIDDefinition   `json:"id"`
	Time       DateTimeDefinition        `json:"time"`
	UserID     int                       `json:"userID"`
	AccountID  AccountIDDefinition       `json:"accountID"`
	BatchID    TransactionIDDefinition   `json:"batchID"`
	RequestID  RequestIDDefinition       `json:"requestID"`
	Type       TransactionTypeDefinition `json:"type"`
	Alias      string                    `json:"alias"`
	MarginRate DecimalNumberDefinition   `json:"marginRate"`
}

type ClientConfigureRejectTransactionDefinition struct {
	Id           TransactionIDDefinition           `json:"id"`
	Time         DateTimeDefinition                `json:"time"`
	UserID       int                               `json:"userID"`
	AccountID    AccountIDDefinition               `json:"accountID"`
	BatchID      TransactionIDDefinition           `json:"batchID"`
	RequestID    RequestIDDefinition               `json:"requestID"`
	Type         TransactionTypeDefinition         `json:"type"`
	Alias        string                            `json:"alias"`
	MarginRate   DecimalNumberDefinition           `json:"marginRate"`
	RejectReason TransactionRejectReasonDefinition `json:"rejectReason"`
}

type TransferFundsTransactionDefinition struct {
	Id             TransactionIDDefinition   `json:"id"`
	Time           DateTimeDefinition        `json:"time"`
	UserID         int                       `json:"userID"`
	AccountID      AccountIDDefinition       `json:"accountID"`
	BatchID        TransactionIDDefinition   `json:"batchID"`
	RequestID      RequestIDDefinition       `json:"requestID"`
	Type           TransactionTypeDefinition `json:"type"`
	Amount         AccountUnitsDefinition    `json:"amount"`
	FundingReason  FundingReasonDefinition   `json:"fundingReason"`
	Comment        string                    `json:"comment"`
	AccountBalance AccountUnitsDefinition    `json:"accountBalance"`
}

type TransferFundsRejectTransactionDefinition struct {
	Id            TransactionIDDefinition           `json:"id"`
	Time          DateTimeDefinition                `json:"time"`
	UserID        int                               `json:"userID"`
	AccountID     AccountIDDefinition               `json:"accountID"`
	BatchID       TransactionIDDefinition           `json:"batchID"`
	RequestID     RequestIDDefinition               `json:"requestID"`
	Type          TransactionTypeDefinition         `json:"type"`
	Amount        AccountUnitsDefinition            `json:"amount"`
	FundingReason FundingReasonDefinition           `json:"fundingReason"`
	Comment       string                            `json:"comment"`
	RejectReason  TransactionRejectReasonDefinition `json:"rejectReason"`
}

type MarketOrderTransactionDefinition struct {
	Id                     TransactionIDDefinition                `json:"id"`
	Time                   DateTimeDefinition                     `json:"time"`
	UserID                 int                                    `json:"userID"`
	AccountID              AccountIDDefinition                    `json:"accountID"`
	BatchID                TransactionIDDefinition                `json:"batchID"`
	RequestID              RequestIDDefinition                    `json:"requestID"`
	Type                   TransactionTypeDefinition              `json:"type"`
	Instrument             InstrumentNameDefinition               `json:"instrument"`
	Units                  DecimalNumberDefinition                `json:"units"`
	TimeInForce            TimeInForceDefinition                  `json:"timeInForce"`
	PriceBound             PriceValueDefinition                   `json:"priceBound"`
	PositionFill           OrderPositionFillDefinition            `json:"positionFill"`
	TradeClose             MarketOrderTradeCloseDefinition        `json:"tradeClose"`
	LongPositionCloseout   MarketOrderPositionCloseoutDefinition  `json:"longPositionCloseout"`
	ShortPositionCloseout  MarketOrderPositionCloseoutDefinition  `json:"shortPositionCloseout"`
	MarginCloseout         MarketOrderMarginCloseoutDefinition    `json:"marginCloseout"`
	DelayedTradeClose      MarketOrderDelayedTradeCloseDefinition `json:"delayedTradeClose"`
	Reason                 MarketOrderReasonDefinition            `json:"reason"`
	ClientExtensions       ClientExtensionsDefinition             `json:"clientExtensions"`
	TakeProfitOnFill       TakeProfitDetailsDefinition            `json:"takeProfitOnFill"`
	StopLossOnFill         StopLossDetailsDefinition              `json:"stopLossOnFill"`
	TrailingStopLossOnFill TrailingStopLossDetailsDefinition      `json:"trailingStopLossOnFill"`
	TradeClientExtensions  ClientExtensionsDefinition             `json:"tradeClientExtensions"`
}

type MarketOrderRejectTransactionDefinition struct {
	Id                     TransactionIDDefinition                `json:"id"`
	Time                   DateTimeDefinition                     `json:"time"`
	UserID                 int                                    `json:"userID"`
	AccountID              AccountIDDefinition                    `json:"accountID"`
	BatchID                TransactionIDDefinition                `json:"batchID"`
	RequestID              RequestIDDefinition                    `json:"requestID"`
	Type                   TransactionTypeDefinition              `json:"type"`
	Instrument             InstrumentNameDefinition               `json:"instrument"`
	Units                  DecimalNumberDefinition                `json:"units"`
	TimeInForce            TimeInForceDefinition                  `json:"timeInForce"`
	PriceBound             PriceValueDefinition                   `json:"priceBound"`
	PositionFill           OrderPositionFillDefinition            `json:"positionFill"`
	TradeClose             MarketOrderTradeCloseDefinition        `json:"tradeClose"`
	LongPositionCloseout   MarketOrderPositionCloseoutDefinition  `json:"longPositionCloseout"`
	ShortPositionCloseout  MarketOrderPositionCloseoutDefinition  `json:"shortPositionCloseout"`
	MarginCloseout         MarketOrderMarginCloseoutDefinition    `json:"marginCloseout"`
	DelayedTradeClose      MarketOrderDelayedTradeCloseDefinition `json:"delayedTradeClose"`
	Reason                 MarketOrderReasonDefinition            `json:"reason"`
	ClientExtensions       ClientExtensionsDefinition             `json:"clientExtensions"`
	TakeProfitOnFill       TakeProfitDetailsDefinition            `json:"takeProfitOnFill"`
	StopLossOnFill         StopLossDetailsDefinition              `json:"stopLossOnFill"`
	TrailingStopLossOnFill TrailingStopLossDetailsDefinition      `json:"trailingStopLossOnFill"`
	TradeClientExtensions  ClientExtensionsDefinition             `json:"tradeClientExtensions"`
	RejectReason           TransactionRejectReasonDefinition      `json:"rejectReason"`
}

type FixedPriceOrderTransactionDefinition struct {
	Id                     TransactionIDDefinition           `json:"id"`
	Time                   DateTimeDefinition                `json:"time"`
	UserID                 int                               `json:"userID"`
	AccountID              AccountIDDefinition               `json:"accountID"`
	BatchID                TransactionIDDefinition           `json:"batchID"`
	RequestID              RequestIDDefinition               `json:"requestID"`
	Type                   TransactionTypeDefinition         `json:"type"`
	Instrument             InstrumentNameDefinition          `json:"instrument"`
	Units                  DecimalNumberDefinition           `json:"units"`
	Price                  PriceValueDefinition              `json:"price"`
	PositionFill           OrderPositionFillDefinition       `json:"positionFill"`
	TradeState             string                            `json:"tradeState"`
	Reason                 FixedPriceOrderReasonDefinition   `json:"reason"`
	ClientExtensions       ClientExtensionsDefinition        `json:"clientExtensions"`
	TakeProfitOnFill       TakeProfitDetailsDefinition       `json:"takeProfitOnFill"`
	StopLossOnFill         StopLossDetailsDefinition         `json:"stopLossOnFill"`
	TrailingStopLossOnFill TrailingStopLossDetailsDefinition `json:"trailingStopLossOnFill"`
	TradeClientExtensions  ClientExtensionsDefinition        `json:"tradeClientExtensions"`
}

type LimitOrderTransactionDefinition struct {
	Id                      TransactionIDDefinition           `json:"id"`
	Time                    DateTimeDefinition                `json:"time"`
	UserID                  int                               `json:"userID"`
	AccountID               AccountIDDefinition               `json:"accountID"`
	BatchID                 TransactionIDDefinition           `json:"batchID"`
	RequestID               RequestIDDefinition               `json:"requestID"`
	Type                    TransactionTypeDefinition         `json:"type"`
	Instrument              InstrumentNameDefinition          `json:"instrument"`
	Units                   DecimalNumberDefinition           `json:"units"`
	Price                   PriceValueDefinition              `json:"price"`
	TimeInForce             TimeInForceDefinition             `json:"timeInForce"`
	GtdTime                 DateTimeDefinition                `json:"gtdTime"`
	PositionFill            OrderPositionFillDefinition       `json:"positionFill"`
	TriggerCondition        OrderTriggerConditionDefinition   `json:"triggerCondition"`
	Reason                  LimitOrderReasonDefinition        `json:"reason"`
	ClientExtensions        ClientExtensionsDefinition        `json:"clientExtensions"`
	TakeProfitOnFill        TakeProfitDetailsDefinition       `json:"takeProfitOnFill"`
	StopLossOnFill          StopLossDetailsDefinition         `json:"stopLossOnFill"`
	TrailingStopLossOnFill  TrailingStopLossDetailsDefinition `json:"trailingStopLossOnFill"`
	TradeClientExtensions   ClientExtensionsDefinition        `json:"tradeClientExtensions"`
	ReplacesOrderID         OrderIDDefinition                 `json:"replacesOrderID"`
	CancellingTransactionID TransactionIDDefinition           `json:"cancellingTransactionID"`
}

type LimitOrderRejectTransactionDefinition struct {
	Id                      TransactionIDDefinition           `json:"id"`
	Time                    DateTimeDefinition                `json:"time"`
	UserID                  int                               `json:"userID"`
	AccountID               AccountIDDefinition               `json:"accountID"`
	BatchID                 TransactionIDDefinition           `json:"batchID"`
	RequestID               RequestIDDefinition               `json:"requestID"`
	Type                    TransactionTypeDefinition         `json:"type"`
	Instrument              InstrumentNameDefinition          `json:"instrument"`
	Units                   DecimalNumberDefinition           `json:"units"`
	Price                   PriceValueDefinition              `json:"price"`
	TimeInForce             TimeInForceDefinition             `json:"timeInForce"`
	GtdTime                 DateTimeDefinition                `json:"gtdTime"`
	PositionFill            OrderPositionFillDefinition       `json:"positionFill"`
	TriggerCondition        OrderTriggerConditionDefinition   `json:"triggerCondition"`
	Reason                  LimitOrderReasonDefinition        `json:"reason"`
	ClientExtensions        ClientExtensionsDefinition        `json:"clientExtensions"`
	TakeProfitOnFill        TakeProfitDetailsDefinition       `json:"takeProfitOnFill"`
	StopLossOnFill          StopLossDetailsDefinition         `json:"stopLossOnFill"`
	TrailingStopLossOnFill  TrailingStopLossDetailsDefinition `json:"trailingStopLossOnFill"`
	TradeClientExtensions   ClientExtensionsDefinition        `json:"tradeClientExtensions"`
	IntendedReplacesOrderID OrderIDDefinition                 `json:"intendedReplacesOrderID"`
	RejectReason            TransactionRejectReasonDefinition `json:"rejectReason"`
}

type StopOrderTransactionDefinition struct {
	Id                      TransactionIDDefinition           `json:"id"`
	Time                    DateTimeDefinition                `json:"time"`
	UserID                  int                               `json:"userID"`
	AccountID               AccountIDDefinition               `json:"accountID"`
	BatchID                 TransactionIDDefinition           `json:"batchID"`
	RequestID               RequestIDDefinition               `json:"requestID"`
	Type                    TransactionTypeDefinition         `json:"type"`
	Instrument              InstrumentNameDefinition          `json:"instrument"`
	Units                   DecimalNumberDefinition           `json:"units"`
	Price                   PriceValueDefinition              `json:"price"`
	PriceBound              PriceValueDefinition              `json:"priceBound"`
	TimeInForce             TimeInForceDefinition             `json:"timeInForce"`
	GtdTime                 DateTimeDefinition                `json:"gtdTime"`
	PositionFill            OrderPositionFillDefinition       `json:"positionFill"`
	TriggerCondition        OrderTriggerConditionDefinition   `json:"triggerCondition"`
	Reason                  StopOrderReasonDefinition         `json:"reason"`
	ClientExtensions        ClientExtensionsDefinition        `json:"clientExtensions"`
	TakeProfitOnFill        TakeProfitDetailsDefinition       `json:"takeProfitOnFill"`
	StopLossOnFill          StopLossDetailsDefinition         `json:"stopLossOnFill"`
	TrailingStopLossOnFill  TrailingStopLossDetailsDefinition `json:"trailingStopLossOnFill"`
	TradeClientExtensions   ClientExtensionsDefinition        `json:"tradeClientExtensions"`
	ReplacesOrderID         OrderIDDefinition                 `json:"replacesOrderID"`
	CancellingTransactionID TransactionIDDefinition           `json:"cancellingTransactionID"`
}

type StopOrderRejectTransactionDefinition struct {
	Id                      TransactionIDDefinition           `json:"id"`
	Time                    DateTimeDefinition                `json:"time"`
	UserID                  int                               `json:"userID"`
	AccountID               AccountIDDefinition               `json:"accountID"`
	BatchID                 TransactionIDDefinition           `json:"batchID"`
	RequestID               RequestIDDefinition               `json:"requestID"`
	Type                    TransactionTypeDefinition         `json:"type"`
	Instrument              InstrumentNameDefinition          `json:"instrument"`
	Units                   DecimalNumberDefinition           `json:"units"`
	Price                   PriceValueDefinition              `json:"price"`
	PriceBound              PriceValueDefinition              `json:"priceBound"`
	TimeInForce             TimeInForceDefinition             `json:"timeInForce"`
	GtdTime                 DateTimeDefinition                `json:"gtdTime"`
	PositionFill            OrderPositionFillDefinition       `json:"positionFill"`
	TriggerCondition        OrderTriggerConditionDefinition   `json:"triggerCondition"`
	Reason                  StopOrderReasonDefinition         `json:"reason"`
	ClientExtensions        ClientExtensionsDefinition        `json:"clientExtensions"`
	TakeProfitOnFill        TakeProfitDetailsDefinition       `json:"takeProfitOnFill"`
	StopLossOnFill          StopLossDetailsDefinition         `json:"stopLossOnFill"`
	TrailingStopLossOnFill  TrailingStopLossDetailsDefinition `json:"trailingStopLossOnFill"`
	TradeClientExtensions   ClientExtensionsDefinition        `json:"tradeClientExtensions"`
	IntendedReplacesOrderID OrderIDDefinition                 `json:"intendedReplacesOrderID"`
	RejectReason            TransactionRejectReasonDefinition `json:"rejectReason"`
}

type MarketIfTouchedOrderTransactionDefinition struct {
	Id                      TransactionIDDefinition              `json:"id"`
	Time                    DateTimeDefinition                   `json:"time"`
	UserID                  int                                  `json:"userID"`
	AccountID               AccountIDDefinition                  `json:"accountID"`
	BatchID                 TransactionIDDefinition              `json:"batchID"`
	RequestID               RequestIDDefinition                  `json:"requestID"`
	Type                    TransactionTypeDefinition            `json:"type"`
	Instrument              InstrumentNameDefinition             `json:"instrument"`
	Units                   DecimalNumberDefinition              `json:"units"`
	Price                   PriceValueDefinition                 `json:"price"`
	PriceBound              PriceValueDefinition                 `json:"priceBound"`
	TimeInForce             TimeInForceDefinition                `json:"timeInForce"`
	GtdTime                 DateTimeDefinition                   `json:"gtdTime"`
	PositionFill            OrderPositionFillDefinition          `json:"positionFill"`
	TriggerCondition        OrderTriggerConditionDefinition      `json:"triggerCondition"`
	Reason                  MarketIfTouchedOrderReasonDefinition `json:"reason"`
	ClientExtensions        ClientExtensionsDefinition           `json:"clientExtensions"`
	TakeProfitOnFill        TakeProfitDetailsDefinition          `json:"takeProfitOnFill"`
	StopLossOnFill          StopLossDetailsDefinition            `json:"stopLossOnFill"`
	TrailingStopLossOnFill  TrailingStopLossDetailsDefinition    `json:"trailingStopLossOnFill"`
	TradeClientExtensions   ClientExtensionsDefinition           `json:"tradeClientExtensions"`
	ReplacesOrderID         OrderIDDefinition                    `json:"replacesOrderID"`
	CancellingTransactionID TransactionIDDefinition              `json:"cancellingTransactionID"`
}

type MarketIfTouchedOrderRejectTransactionDefinition struct {
	Id                      TransactionIDDefinition              `json:"id"`
	Time                    DateTimeDefinition                   `json:"time"`
	UserID                  int                                  `json:"userID"`
	AccountID               AccountIDDefinition                  `json:"accountID"`
	BatchID                 TransactionIDDefinition              `json:"batchID"`
	RequestID               RequestIDDefinition                  `json:"requestID"`
	Type                    TransactionTypeDefinition            `json:"type"`
	Instrument              InstrumentNameDefinition             `json:"instrument"`
	Units                   DecimalNumberDefinition              `json:"units"`
	Price                   PriceValueDefinition                 `json:"price"`
	PriceBound              PriceValueDefinition                 `json:"priceBound"`
	TimeInForce             TimeInForceDefinition                `json:"timeInForce"`
	GtdTime                 DateTimeDefinition                   `json:"gtdTime"`
	PositionFill            OrderPositionFillDefinition          `json:"positionFill"`
	TriggerCondition        OrderTriggerConditionDefinition      `json:"triggerCondition"`
	Reason                  MarketIfTouchedOrderReasonDefinition `json:"reason"`
	ClientExtensions        ClientExtensionsDefinition           `json:"clientExtensions"`
	TakeProfitOnFill        TakeProfitDetailsDefinition          `json:"takeProfitOnFill"`
	StopLossOnFill          StopLossDetailsDefinition            `json:"stopLossOnFill"`
	TrailingStopLossOnFill  TrailingStopLossDetailsDefinition    `json:"trailingStopLossOnFill"`
	TradeClientExtensions   ClientExtensionsDefinition           `json:"tradeClientExtensions"`
	IntendedReplacesOrderID OrderIDDefinition                    `json:"intendedReplacesOrderID"`
	RejectReason            TransactionRejectReasonDefinition    `json:"rejectReason"`
}

type TakeProfitOrderTransactionDefinition struct {
	Id                      TransactionIDDefinition         `json:"id"`
	Time                    DateTimeDefinition              `json:"time"`
	UserID                  int                             `json:"userID"`
	AccountID               AccountIDDefinition             `json:"accountID"`
	BatchID                 TransactionIDDefinition         `json:"batchID"`
	RequestID               RequestIDDefinition             `json:"requestID"`
	Type                    TransactionTypeDefinition       `json:"type"`
	TradeID                 TradeIDDefinition               `json:"tradeID"`
	ClientTradeID           ClientIDDefinition              `json:"clientTradeID"`
	Price                   PriceValueDefinition            `json:"price"`
	TimeInForce             TimeInForceDefinition           `json:"timeInForce"`
	GtdTime                 DateTimeDefinition              `json:"gtdTime"`
	TriggerCondition        OrderTriggerConditionDefinition `json:"triggerCondition"`
	Reason                  TakeProfitOrderReasonDefinition `json:"reason"`
	ClientExtensions        ClientExtensionsDefinition      `json:"clientExtensions"`
	OrderFillTransactionID  TransactionIDDefinition         `json:"orderFillTransactionID"`
	ReplacesOrderID         OrderIDDefinition               `json:"replacesOrderID"`
	CancellingTransactionID TransactionIDDefinition         `json:"cancellingTransactionID"`
}

type TakeProfitOrderRejectTransactionDefinition struct {
	Id                      TransactionIDDefinition           `json:"id"`
	Time                    DateTimeDefinition                `json:"time"`
	UserID                  int                               `json:"userID"`
	AccountID               AccountIDDefinition               `json:"accountID"`
	BatchID                 TransactionIDDefinition           `json:"batchID"`
	RequestID               RequestIDDefinition               `json:"requestID"`
	Type                    TransactionTypeDefinition         `json:"type"`
	TradeID                 TradeIDDefinition                 `json:"tradeID"`
	ClientTradeID           ClientIDDefinition                `json:"clientTradeID"`
	Price                   PriceValueDefinition              `json:"price"`
	TimeInForce             TimeInForceDefinition             `json:"timeInForce"`
	GtdTime                 DateTimeDefinition                `json:"gtdTime"`
	TriggerCondition        OrderTriggerConditionDefinition   `json:"triggerCondition"`
	Reason                  TakeProfitOrderReasonDefinition   `json:"reason"`
	ClientExtensions        ClientExtensionsDefinition        `json:"clientExtensions"`
	OrderFillTransactionID  TransactionIDDefinition           `json:"orderFillTransactionID"`
	IntendedReplacesOrderID OrderIDDefinition                 `json:"intendedReplacesOrderID"`
	RejectReason            TransactionRejectReasonDefinition `json:"rejectReason"`
}

type StopLossOrderTransactionDefinition struct {
	Id                         TransactionIDDefinition         `json:"id"`
	Time                       DateTimeDefinition              `json:"time"`
	UserID                     int                             `json:"userID"`
	AccountID                  AccountIDDefinition             `json:"accountID"`
	BatchID                    TransactionIDDefinition         `json:"batchID"`
	RequestID                  RequestIDDefinition             `json:"requestID"`
	Type                       TransactionTypeDefinition       `json:"type"`
	TradeID                    TradeIDDefinition               `json:"tradeID"`
	ClientTradeID              ClientIDDefinition              `json:"clientTradeID"`
	Price                      PriceValueDefinition            `json:"price"`
	Distance                   DecimalNumberDefinition         `json:"distance"`
	TimeInForce                TimeInForceDefinition           `json:"timeInForce"`
	GtdTime                    DateTimeDefinition              `json:"gtdTime"`
	TriggerCondition           OrderTriggerConditionDefinition `json:"triggerCondition"`
	Guaranteed                 bool                            `json:"guaranteed"`
	GuaranteedExecutionPremium DecimalNumberDefinition         `json:"guaranteedExecutionPremium"`
	Reason                     StopLossOrderReasonDefinition   `json:"reason"`
	ClientExtensions           ClientExtensionsDefinition      `json:"clientExtensions"`
	OrderFillTransactionID     TransactionIDDefinition         `json:"orderFillTransactionID"`
	ReplacesOrderID            OrderIDDefinition               `json:"replacesOrderID"`
	CancellingTransactionID    TransactionIDDefinition         `json:"cancellingTransactionID"`
}

type StopLossOrderRejectTransactionDefinition struct {
	Id                      TransactionIDDefinition           `json:"id"`
	Time                    DateTimeDefinition                `json:"time"`
	UserID                  int                               `json:"userID"`
	AccountID               AccountIDDefinition               `json:"accountID"`
	BatchID                 TransactionIDDefinition           `json:"batchID"`
	RequestID               RequestIDDefinition               `json:"requestID"`
	Type                    TransactionTypeDefinition         `json:"type"`
	TradeID                 TradeIDDefinition                 `json:"tradeID"`
	ClientTradeID           ClientIDDefinition                `json:"clientTradeID"`
	Price                   PriceValueDefinition              `json:"price"`
	Distance                DecimalNumberDefinition           `json:"distance"`
	TimeInForce             TimeInForceDefinition             `json:"timeInForce"`
	GtdTime                 DateTimeDefinition                `json:"gtdTime"`
	TriggerCondition        OrderTriggerConditionDefinition   `json:"triggerCondition"`
	Guaranteed              bool                              `json:"guaranteed"`
	Reason                  StopLossOrderReasonDefinition     `json:"reason"`
	ClientExtensions        ClientExtensionsDefinition        `json:"clientExtensions"`
	OrderFillTransactionID  TransactionIDDefinition           `json:"orderFillTransactionID"`
	IntendedReplacesOrderID OrderIDDefinition                 `json:"intendedReplacesOrderID"`
	RejectReason            TransactionRejectReasonDefinition `json:"rejectReason"`
}

type TrailingStopLossOrderTransactionDefinition struct {
	Id                      TransactionIDDefinition               `json:"id"`
	Time                    DateTimeDefinition                    `json:"time"`
	UserID                  int                                   `json:"userID"`
	AccountID               AccountIDDefinition                   `json:"accountID"`
	BatchID                 TransactionIDDefinition               `json:"batchID"`
	RequestID               RequestIDDefinition                   `json:"requestID"`
	Type                    TransactionTypeDefinition             `json:"type"`
	TradeID                 TradeIDDefinition                     `json:"tradeID"`
	ClientTradeID           ClientIDDefinition                    `json:"clientTradeID"`
	Distance                DecimalNumberDefinition               `json:"distance"`
	TimeInForce             TimeInForceDefinition                 `json:"timeInForce"`
	GtdTime                 DateTimeDefinition                    `json:"gtdTime"`
	TriggerCondition        OrderTriggerConditionDefinition       `json:"triggerCondition"`
	Reason                  TrailingStopLossOrderReasonDefinition `json:"reason"`
	ClientExtensions        ClientExtensionsDefinition            `json:"clientExtensions"`
	OrderFillTransactionID  TransactionIDDefinition               `json:"orderFillTransactionID"`
	ReplacesOrderID         OrderIDDefinition                     `json:"replacesOrderID"`
	CancellingTransactionID TransactionIDDefinition               `json:"cancellingTransactionID"`
}

type TrailingStopLossOrderRejectTransactionDefinition struct {
	Id                      TransactionIDDefinition               `json:"id"`
	Time                    DateTimeDefinition                    `json:"time"`
	UserID                  int                                   `json:"userID"`
	AccountID               AccountIDDefinition                   `json:"accountID"`
	BatchID                 TransactionIDDefinition               `json:"batchID"`
	RequestID               RequestIDDefinition                   `json:"requestID"`
	Type                    TransactionTypeDefinition             `json:"type"`
	TradeID                 TradeIDDefinition                     `json:"tradeID"`
	ClientTradeID           ClientIDDefinition                    `json:"clientTradeID"`
	Distance                DecimalNumberDefinition               `json:"distance"`
	TimeInForce             TimeInForceDefinition                 `json:"timeInForce"`
	GtdTime                 DateTimeDefinition                    `json:"gtdTime"`
	TriggerCondition        OrderTriggerConditionDefinition       `json:"triggerCondition"`
	Reason                  TrailingStopLossOrderReasonDefinition `json:"reason"`
	ClientExtensions        ClientExtensionsDefinition            `json:"clientExtensions"`
	OrderFillTransactionID  TransactionIDDefinition               `json:"orderFillTransactionID"`
	IntendedReplacesOrderID OrderIDDefinition                     `json:"intendedReplacesOrderID"`
	RejectReason            TransactionRejectReasonDefinition     `json:"rejectReason"`
}

type OrderFillTransactionDefinition struct {
	Id                            TransactionIDDefinition   `json:"id"`
	Time                          DateTimeDefinition        `json:"time"`
	UserID                        int                       `json:"userID"`
	AccountID                     AccountIDDefinition       `json:"accountID"`
	BatchID                       TransactionIDDefinition   `json:"batchID"`
	RequestID                     RequestIDDefinition       `json:"requestID"`
	Type                          TransactionTypeDefinition `json:"type"`
	OrderID                       OrderIDDefinition         `json:"orderID"`
	ClientOrderID                 ClientIDDefinition        `json:"clientOrderID"`
	Instrument                    InstrumentNameDefinition  `json:"instrument"`
	Units                         DecimalNumberDefinition   `json:"units"`
	GainQuoteHomeConversionFactor DecimalNumberDefinition   `json:"gainQuoteHomeConversionFactor"`
	LossQuoteHomeConversionFactor DecimalNumberDefinition   `json:"lossQuoteHomeConversionFactor"`
	Price                         PriceValueDefinition      `json:"price"`
	FullPrice                     ClientPriceDefinition     `json:"fullPrice"`
	Reason                        OrderFillReasonDefinition `json:"reason"`
	Pl                            AccountUnitsDefinition    `json:"pl"`
	Financing                     AccountUnitsDefinition    `json:"financing"`
	Commission                    AccountUnitsDefinition    `json:"commission"`
	GuaranteedExecutionFee        AccountUnitsDefinition    `json:"guaranteedExecutionFee"`
	AccountBalance                AccountUnitsDefinition    `json:"accountBalance"`
	TradeOpened                   TradeOpenDefinition       `json:"tradeOpened"`
	TradesClosed                  []TradeReduceDefinition   `json:"tradesClosed"`
	TradeReduced                  TradeReduceDefinition     `json:"tradeReduced"`
	HalfSpreadCost                AccountUnitsDefinition    `json:"halfSpreadCost"`
}

type OrderCancelTransactionDefinition struct {
	Id                TransactionIDDefinition     `json:"id"`
	Time              DateTimeDefinition          `json:"time"`
	UserID            int                         `json:"userID"`
	AccountID         AccountIDDefinition         `json:"accountID"`
	BatchID           TransactionIDDefinition     `json:"batchID"`
	RequestID         RequestIDDefinition         `json:"requestID"`
	Type              TransactionTypeDefinition   `json:"type"`
	OrderID           OrderIDDefinition           `json:"orderID"`
	ClientOrderID     OrderIDDefinition           `json:"clientOrderID"`
	Reason            OrderCancelReasonDefinition `json:"reason"`
	ReplacedByOrderID OrderIDDefinition           `json:"replacedByOrderID"`
}

type OrderCancelRejectTransactionDefinition struct {
	Id            TransactionIDDefinition           `json:"id"`
	Time          DateTimeDefinition                `json:"time"`
	UserID        int                               `json:"userID"`
	AccountID     AccountIDDefinition               `json:"accountID"`
	BatchID       TransactionIDDefinition           `json:"batchID"`
	RequestID     RequestIDDefinition               `json:"requestID"`
	Type          TransactionTypeDefinition         `json:"type"`
	OrderID       OrderIDDefinition                 `json:"orderID"`
	ClientOrderID OrderIDDefinition                 `json:"clientOrderID"`
	RejectReason  TransactionRejectReasonDefinition `json:"rejectReason"`
}

type OrderClientExtensionsModifyTransactionDefinition struct {
	Id                          TransactionIDDefinition    `json:"id"`
	Time                        DateTimeDefinition         `json:"time"`
	UserID                      int                        `json:"userID"`
	AccountID                   AccountIDDefinition        `json:"accountID"`
	BatchID                     TransactionIDDefinition    `json:"batchID"`
	RequestID                   RequestIDDefinition        `json:"requestID"`
	Type                        TransactionTypeDefinition  `json:"type"`
	OrderID                     OrderIDDefinition          `json:"orderID"`
	ClientOrderID               ClientIDDefinition         `json:"clientOrderID"`
	ClientExtensionsModify      ClientExtensionsDefinition `json:"clientExtensionsModify"`
	TradeClientExtensionsModify ClientExtensionsDefinition `json:"tradeClientExtensionsModify"`
}

type OrderClientExtensionsModifyRejectTransactionDefinition struct {
	Id                          TransactionIDDefinition           `json:"id"`
	Time                        DateTimeDefinition                `json:"time"`
	UserID                      int                               `json:"userID"`
	AccountID                   AccountIDDefinition               `json:"accountID"`
	BatchID                     TransactionIDDefinition           `json:"batchID"`
	RequestID                   RequestIDDefinition               `json:"requestID"`
	Type                        TransactionTypeDefinition         `json:"type"`
	OrderID                     OrderIDDefinition                 `json:"orderID"`
	ClientOrderID               ClientIDDefinition                `json:"clientOrderID"`
	ClientExtensionsModify      ClientExtensionsDefinition        `json:"clientExtensionsModify"`
	TradeClientExtensionsModify ClientExtensionsDefinition        `json:"tradeClientExtensionsModify"`
	RejectReason                TransactionRejectReasonDefinition `json:"rejectReason"`
}

type TradeClientExtensionsModifyTransactionDefinition struct {
	Id                          TransactionIDDefinition    `json:"id"`
	Time                        DateTimeDefinition         `json:"time"`
	UserID                      int                        `json:"userID"`
	AccountID                   AccountIDDefinition        `json:"accountID"`
	BatchID                     TransactionIDDefinition    `json:"batchID"`
	RequestID                   RequestIDDefinition        `json:"requestID"`
	Type                        TransactionTypeDefinition  `json:"type"`
	TradeID                     TradeIDDefinition          `json:"tradeID"`
	ClientTradeID               ClientIDDefinition         `json:"clientTradeID"`
	TradeClientExtensionsModify ClientExtensionsDefinition `json:"tradeClientExtensionsModify"`
}

type TradeClientExtensionsModifyRejectTransactionDefinition struct {
	Id                          TransactionIDDefinition           `json:"id"`
	Time                        DateTimeDefinition                `json:"time"`
	UserID                      int                               `json:"userID"`
	AccountID                   AccountIDDefinition               `json:"accountID"`
	BatchID                     TransactionIDDefinition           `json:"batchID"`
	RequestID                   RequestIDDefinition               `json:"requestID"`
	Type                        TransactionTypeDefinition         `json:"type"`
	TradeID                     TradeIDDefinition                 `json:"tradeID"`
	ClientTradeID               ClientIDDefinition                `json:"clientTradeID"`
	TradeClientExtensionsModify ClientExtensionsDefinition        `json:"tradeClientExtensionsModify"`
	RejectReason                TransactionRejectReasonDefinition `json:"rejectReason"`
}

type MarginCallEnterTransactionDefinition struct {
	Id        TransactionIDDefinition   `json:"id"`
	Time      DateTimeDefinition        `json:"time"`
	UserID    int                       `json:"userID"`
	AccountID AccountIDDefinition       `json:"accountID"`
	BatchID   TransactionIDDefinition   `json:"batchID"`
	RequestID RequestIDDefinition       `json:"requestID"`
	Type      TransactionTypeDefinition `json:"type"`
}

type MarginCallExtendTransactionDefinition struct {
	Id              TransactionIDDefinition   `json:"id"`
	Time            DateTimeDefinition        `json:"time"`
	UserID          int                       `json:"userID"`
	AccountID       AccountIDDefinition       `json:"accountID"`
	BatchID         TransactionIDDefinition   `json:"batchID"`
	RequestID       RequestIDDefinition       `json:"requestID"`
	Type            TransactionTypeDefinition `json:"type"`
	ExtensionNumber int                       `json:"extensionNumber"`
}

type MarginCallExitTransactionDefinition struct {
	Id        TransactionIDDefinition   `json:"id"`
	Time      DateTimeDefinition        `json:"time"`
	UserID    int                       `json:"userID"`
	AccountID AccountIDDefinition       `json:"accountID"`
	BatchID   TransactionIDDefinition   `json:"batchID"`
	RequestID RequestIDDefinition       `json:"requestID"`
	Type      TransactionTypeDefinition `json:"type"`
}

type DelayedTradeClosureTransactionDefinition struct {
	Id        TransactionIDDefinition     `json:"id"`
	Time      DateTimeDefinition          `json:"time"`
	UserID    int                         `json:"userID"`
	AccountID AccountIDDefinition         `json:"accountID"`
	BatchID   TransactionIDDefinition     `json:"batchID"`
	RequestID RequestIDDefinition         `json:"requestID"`
	Type      TransactionTypeDefinition   `json:"type"`
	Reason    MarketOrderReasonDefinition `json:"reason"`
	TradeIDs  TradeIDDefinition           `json:"tradeIDs"`
}

type DailyFinancingTransactionDefinition struct {
	Id                   TransactionIDDefinition        `json:"id"`
	Time                 DateTimeDefinition             `json:"time"`
	UserID               int                            `json:"userID"`
	AccountID            AccountIDDefinition            `json:"accountID"`
	BatchID              TransactionIDDefinition        `json:"batchID"`
	RequestID            RequestIDDefinition            `json:"requestID"`
	Type                 TransactionTypeDefinition      `json:"type"`
	Financing            AccountUnitsDefinition         `json:"financing"`
	AccountBalance       AccountUnitsDefinition         `json:"accountBalance"`
	AccountFinancingMode AccountFinancingModeDefinition `json:"accountFinancingMode"`
	PositionFinancings   []PositionFinancingDefinition  `json:"positionFinancings"`
}

type ResetResettablePLTransactionDefinition struct {
	Id        TransactionIDDefinition   `json:"id"`
	Time      DateTimeDefinition        `json:"time"`
	UserID    int                       `json:"userID"`
	AccountID AccountIDDefinition       `json:"accountID"`
	BatchID   TransactionIDDefinition   `json:"batchID"`
	RequestID RequestIDDefinition       `json:"requestID"`
	Type      TransactionTypeDefinition `json:"type"`
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
	reason MarketOrderMarginCloseoutReasonDefinition `json:"reason"`
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

type TransactionHeartbeatDefinition struct {
	Type              string                  `json:"type"`
	LastTransactionID TransactionIDDefinition `json:"lastTransactionID"`
	Time              DateTimeDefinition      `json:"time"`
}

//
// Pricing Definitions
//

type PriceDefinition struct {
	Type                       string                               `json:"type"`
	Instrument                 InstrumentNameDefinition             `json:"instrument"`
	Time                       DateTimeDefinition                   `json:"time"`
	Status                     PriceStatusDefinition                `json:"status"`
	Tradeable                  bool                                 `json:"tradeable"`
	Bids                       []PriceBucketDefinition              `json:"bids"`
	Asks                       []PriceBucketDefinition              `json:"asks"`
	CloseoutBid                PriceValueDefinition                 `json:"closeoutBid"`
	CloseoutAsk                PriceValueDefinition                 `json:"closeoutAsk"`
	QuoteHomeConversionFactors QuoteHomeConversionFactorsDefinition `json:"quoteHomeConversionFactors"`
	UnitsAvailable             UnitsAvailableDefinition             `json:"unitsAvailable"`
}

type PriceValueDefinition string

type PriceBucketDefinition struct {
	Price     PriceValueDefinition `json:"price"`
	Liquidity int                  `json:"liquidity"`
}

type PriceStatusDefinition string

type QuoteHomeConversionFactorsDefinition struct {
	PositiveUnits DecimalNumberDefinition `json:"positiveUnits"`
	NegativeUnits DecimalNumberDefinition `json:"negativeUnits"`
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
