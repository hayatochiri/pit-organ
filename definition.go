package pitOrgan

//
// Account Definitions
//

type DefinitionAccountID string

type DefinitionAccount struct {
	ID                          DefinitionAccountID                   `json:"id"`
	Alias                       string                                `json:"alias"`
	Currency                    DefinitionCurrency                    `json:"currency"`
	Balance                     DefinitionAccountUnits                `json:"balance"`
	CreatedByUserID             int                                   `json:"createdByUserID"`
	CreatedTime                 DefinitionDateTime                    `json:"createdTime"`
	GuaranteedStopLossOrderMode DefinitionGuaranteedStopLossOrderMode `json:"guaranteedStopLossOrderMode"`
	Pl                          DefinitionAccountUnits                `json:"pl"`
	ResettablePL                DefinitionAccountUnits                `json:"resettablePL"`
	ResettablePLTime            DefinitionDateTime                    `json:"resettablePLTime"`
	Financing                   DefinitionAccountUnits                `json:"financing"`
	Commission                  DefinitionAccountUnits                `json:"commission"`
	GuaranteedExecutionFees     DefinitionAccountUnits                `json:"guaranteedExecutionFees"`
	MarginRate                  DefinitionDecimalNumber               `json:"marginRate"`
	MarginCallEnterTime         DefinitionDateTime                    `json:"marginCallEnterTime"`
	MarginCallExtensionCount    int                                   `json:"marginCallExtensionCount"`
	LastMarginCallExtensionTime DefinitionDateTime                    `json:"lastMarginCallExtensionTime"`
	OpenTradeCount              int                                   `json:"openTradeCount"`
	OpenPositionCount           int                                   `json:"openPositionCount"`
	PendingOrderCount           int                                   `json:"pendingOrderCount"`
	HedgingEnabled              bool                                  `json:"hedgingEnabled"`
	UnrealizedPL                DefinitionAccountUnits                `json:"unrealizedPL"`
	NAV                         DefinitionAccountUnits                `json:"NAV"`
	MarginUsed                  DefinitionAccountUnits                `json:"marginUsed"`
	MarginAvailable             DefinitionAccountUnits                `json:"marginAvailable"`
	PositionValue               DefinitionAccountUnits                `json:"positionValue"`
	MarginCloseoutUnrealizedPL  DefinitionAccountUnits                `json:"marginCloseoutUnrealizedPL"`
	MarginCloseoutNAV           DefinitionAccountUnits                `json:"marginCloseoutNAV"`
	MarginCloseoutMarginUsed    DefinitionAccountUnits                `json:"marginCloseoutMarginUsed"`
	MarginCloseoutPercent       DefinitionDecimalNumber               `json:"marginCloseoutPercent"`
	MarginCloseoutPositionValue DefinitionDecimalNumber               `json:"marginCloseoutPositionValue"`
	WithdrawalLimit             DefinitionAccountUnits                `json:"withdrawalLimit"`
	MarginCallMarginUsed        DefinitionAccountUnits                `json:"marginCallMarginUsed"`
	MarginCallPercent           DefinitionDecimalNumber               `json:"marginCallPercent"`
	LastTransactionID           DefinitionTransactionID               `json:"lastTransactionID"`
	Trades                      []DefinitionTradeSummary              `json:"trades"`
	Positions                   []DefinitionPosition                  `json:"positions"`
	Orders                      []DefinitionOrder                     `json:"orders"`
}

type DefinitionAccountChangesState struct {
	UnrealizedPL                DefinitionAccountUnits              `json:"unrealizedPL"`
	NAV                         DefinitionAccountUnits              `json:"NAV"`
	MarginUsed                  DefinitionAccountUnits              `json:"marginUsed"`
	MarginAvailable             DefinitionAccountUnits              `json:"marginAvailable"`
	PositionValue               DefinitionAccountUnits              `json:"positionValue"`
	MarginCloseoutUnrealizedPL  DefinitionAccountUnits              `json:"marginCloseoutUnrealizedPL"`
	MarginCloseoutNAV           DefinitionAccountUnits              `json:"marginCloseoutNAV"`
	MarginCloseoutMarginUsed    DefinitionAccountUnits              `json:"marginCloseoutMarginUsed"`
	MarginCloseoutPercent       DefinitionDecimalNumber             `json:"marginCloseoutPercent"`
	MarginCloseoutPositionValue DefinitionDecimalNumber             `json:"marginCloseoutPositionValue"`
	WithdrawalLimit             DefinitionAccountUnits              `json:"withdrawalLimit"`
	MarginCallMarginUsed        DefinitionAccountUnits              `json:"marginCallMarginUsed"`
	MarginCallPercent           DefinitionDecimalNumber             `json:"marginCallPercent"`
	Orders                      []DefinitionDynamicOrderState       `json:"orders"`
	Trades                      []DefinitionCalculatedTradeState    `json:"trades"`
	Positions                   []DefinitionCalculatedPositionState `json:"positions"`
}

type DefinitionAccountProperties struct {
	ID           DefinitionAccountID `json:"id"`
	MT4AccountID int                 `json:"mt4AccountID"`
	Tags         []string            `json:"tags"`
}

type DefinitionGuaranteedStopLossOrderMode string

type DefinitionAccountSummary struct {
	Id                          DefinitionAccountID                   `json:"id"`
	Alias                       string                                `json:"alias"`
	Currency                    DefinitionCurrency                    `json:"currency"`
	Balance                     DefinitionAccountUnits                `json:"balance"`
	CreatedByUserID             int                                   `json:"createdByUserID"`
	CreatedTime                 DefinitionDateTime                    `json:"createdTime"`
	GuaranteedStopLossOrderMode DefinitionGuaranteedStopLossOrderMode `json:"guaranteedStopLossOrderMode"`
	Pl                          DefinitionAccountUnits                `json:"pl"`
	ResettablePL                DefinitionAccountUnits                `json:"resettablePL"`
	ResettablePLTime            DefinitionDateTime                    `json:"resettablePLTime"`
	Financing                   DefinitionAccountUnits                `json:"financing"`
	Commission                  DefinitionAccountUnits                `json:"commission"`
	GuaranteedExecutionFees     DefinitionAccountUnits                `json:"guaranteedExecutionFees"`
	MarginRate                  DefinitionDecimalNumber               `json:"marginRate"`
	MarginCallEnterTime         DefinitionDateTime                    `json:"marginCallEnterTime"`
	MarginCallExtensionCount    int                                   `json:"marginCallExtensionCount"`
	LastMarginCallExtensionTime DefinitionDateTime                    `json:"lastMarginCallExtensionTime"`
	OpenTradeCount              int                                   `json:"openTradeCount"`
	OpenPositionCount           int                                   `json:"openPositionCount"`
	PendingOrderCount           int                                   `json:"pendingOrderCount"`
	HedgingEnabled              bool                                  `json:"hedgingEnabled"`
	UnrealizedPL                DefinitionAccountUnits                `json:"unrealizedPL"`
	NAV                         DefinitionAccountUnits                `json:"NAV"`
	MarginUsed                  DefinitionAccountUnits                `json:"marginUsed"`
	MarginAvailable             DefinitionAccountUnits                `json:"marginAvailable"`
	PositionValue               DefinitionAccountUnits                `json:"positionValue"`
	MarginCloseoutUnrealizedPL  DefinitionAccountUnits                `json:"marginCloseoutUnrealizedPL"`
	MarginCloseoutNAV           DefinitionAccountUnits                `json:"marginCloseoutNAV"`
	MarginCloseoutMarginUsed    DefinitionAccountUnits                `json:"marginCloseoutMarginUsed"`
	MarginCloseoutPercent       DefinitionDecimalNumber               `json:"marginCloseoutPercent"`
	MarginCloseoutPositionValue DefinitionDecimalNumber               `json:"marginCloseoutPositionValue"`
	WithdrawalLimit             DefinitionAccountUnits                `json:"withdrawalLimit"`
	MarginCallMarginUsed        DefinitionAccountUnits                `json:"marginCallMarginUsed"`
	MarginCallPercent           DefinitionDecimalNumber               `json:"marginCallPercent"`
	LastTransactionID           DefinitionTransactionID               `json:"lastTransactionID"`
}

type DefinitionCalculatedAccountState struct {
	UnrealizedPL                DefinitionAccountUnits  `json:"unrealizedPL"`
	NAV                         DefinitionAccountUnits  `json:"NAV"`
	MarginUsed                  DefinitionAccountUnits  `json:"marginUsed"`
	MarginAvailable             DefinitionAccountUnits  `json:"marginAvailable"`
	PositionValue               DefinitionAccountUnits  `json:"positionValue"`
	MarginCloseoutUnrealizedPL  DefinitionAccountUnits  `json:"marginCloseoutUnrealizedPL"`
	MarginCloseoutNAV           DefinitionAccountUnits  `json:"marginCloseoutNAV"`
	MarginCloseoutMarginUsed    DefinitionAccountUnits  `json:"marginCloseoutMarginUsed"`
	MarginCloseoutPercent       DefinitionDecimalNumber `json:"marginCloseoutPercent"`
	MarginCloseoutPositionValue DefinitionDecimalNumber `json:"marginCloseoutPositionValue"`
	WithdrawalLimit             DefinitionAccountUnits  `json:"withdrawalLimit"`
	MarginCallMarginUsed        DefinitionAccountUnits  `json:"marginCallMarginUsed"`
	MarginCallPercent           DefinitionDecimalNumber `json:"marginCallPercent"`
}

type DefinitionAccountChanges struct {
	OrdersCreated   []DefinitionOrder        `json:"ordersCreated"`
	OrdersCancelled []DefinitionOrder        `json:"ordersCancelled"`
	OrdersFilled    []DefinitionOrder        `json:"ordersFilled"`
	OrdersTriggered []DefinitionOrder        `json:"ordersTriggered"`
	TradesOpened    []DefinitionTradeSummary `json:"tradesOpened"`
	TradesReduced   []DefinitionTradeSummary `json:"tradesReduced"`
	TradesClosed    []DefinitionTradeSummary `json:"tradesClosed"`
	Positions       []DefinitionPosition     `json:"positions"`
	Transactions    []DefinitionTransaction  `json:"transactions"`
}

type DefinitionAccountFinancingMode string

type DefinitionPositionAggregationMode string

//
// Instrument Definitions
//

type DefinitionCandlestickGranularity string

type DefinitionWeeklyAlignment string

type DefinitionCandlestick struct {
	Time     DefinitionDateTime        `json:"time"`
	Bid      DefinitionCandlestickData `json:"bid"`
	Ask      DefinitionCandlestickData `json:"ask"`
	Mid      DefinitionCandlestickData `json:"mid"`
	Volume   int                       `json:"volume"`
	Complete bool                      `json:"complete"`
}

type DefinitionCandlestickData struct {
	O DefinitionPriceValue `json:"o"`
	H DefinitionPriceValue `json:"h"`
	L DefinitionPriceValue `json:"l"`
	C DefinitionPriceValue `json:"c"`
}

type DefinitionOrderBook struct {
	Instrument  DefinitionInstrumentName    `json:"instrument"`
	Time        DefinitionDateTime          `json:"time"`
	Price       DefinitionPriceValue        `json:"price"`
	BucketWidth DefinitionPriceValue        `json:"bucketWidth"`
	Buckets     []DefinitionOrderBookBucket `json:"buckets"`
}

type DefinitionOrderBookBucket struct {
	Price             DefinitionPriceValue    `json:"price"`
	LongCountPercent  DefinitionDecimalNumber `json:"longCountPercent"`
	ShortCountPercent DefinitionDecimalNumber `json:"shortCountPercent"`
}

type DefinitionPositionBook struct {
	Instrument  DefinitionInstrumentName       `json:"instrument"`
	Time        DefinitionDateTime             `json:"time"`
	Price       DefinitionPriceValue           `json:"price"`
	BucketWidth DefinitionPriceValue           `json:"bucketWidth"`
	Buckets     []DefinitionPositionBookBucket `json:"buckets"`
}

type DefinitionPositionBookBucket struct {
	Price             DefinitionPriceValue    `json:"price"`
	LongCountPercent  DefinitionDecimalNumber `json:"longCountPercent"`
	ShortCountPercent DefinitionDecimalNumber `json:"shortCountPercent"`
}

//
// Order Definitions
//

// Orders

type DefinitionOrder struct {
	Id               DefinitionOrderID          `json:"id"`
	CreateTime       DefinitionDateTime         `json:"createTime"`
	State            DefinitionOrderState       `json:"state"`
	ClientExtensions DefinitionClientExtensions `json:"clientExtensions"`
}

type DefinitionMarketOrder struct {
	Id                      DefinitionOrderID                      `json:"id"`
	CreateTime              DefinitionDateTime                     `json:"createTime"`
	State                   DefinitionOrderState                   `json:"state"`
	ClientExtensions        DefinitionClientExtensions             `json:"clientExtensions"`
	Type                    DefinitionOrderType                    `json:"type"`
	Instrument              DefinitionInstrumentName               `json:"instrument"`
	Units                   DefinitionDecimalNumber                `json:"units"`
	TimeInForce             DefinitionTimeInForce                  `json:"timeInForce"`
	PriceBound              DefinitionPriceValue                   `json:"priceBound"`
	PositionFill            DefinitionOrderPositionFill            `json:"positionFill"`
	TradeClose              DefinitionMarketOrderTradeClose        `json:"tradeClose"`
	LongPositionCloseout    DefinitionMarketOrderPositionCloseout  `json:"longPositionCloseout"`
	ShortPositionCloseout   DefinitionMarketOrderPositionCloseout  `json:"shortPositionCloseout"`
	MarginCloseout          DefinitionMarketOrderMarginCloseout    `json:"marginCloseout"`
	DelayedTradeClose       DefinitionMarketOrderDelayedTradeClose `json:"delayedTradeClose"`
	TakeProfitOnFill        DefinitionTakeProfitDetails            `json:"takeProfitOnFill"`
	StopLossOnFill          DefinitionStopLossDetails              `json:"stopLossOnFill"`
	TrailingStopLossOnFill  DefinitionTrailingStopLossDetails      `json:"trailingStopLossOnFill"`
	TradeClientExtensions   DefinitionClientExtensions             `json:"tradeClientExtensions"`
	FillingTransactionID    DefinitionTransactionID                `json:"fillingTransactionID"`
	FilledTime              DefinitionDateTime                     `json:"filledTime"`
	TradeOpenedID           DefinitionTradeID                      `json:"tradeOpenedID"`
	TradeReducedID          DefinitionTradeID                      `json:"tradeReducedID"`
	TradeClosedIDs          []DefinitionTradeID                    `json:"tradeClosedIDs"`
	CancellingTransactionID DefinitionTransactionID                `json:"cancellingTransactionID"`
	CancelledTime           DefinitionDateTime                     `json:"cancelledTime"`
}

type DefinitionFixedPriceOrder struct {
	Id                      DefinitionOrderID                 `json:"id"`
	CreateTime              DefinitionDateTime                `json:"createTime"`
	State                   DefinitionOrderState              `json:"state"`
	ClientExtensions        DefinitionClientExtensions        `json:"clientExtensions"`
	Type                    DefinitionOrderType               `json:"type"`
	Instrument              DefinitionInstrumentName          `json:"instrument"`
	Units                   DefinitionDecimalNumber           `json:"units"`
	Price                   DefinitionPriceValue              `json:"price"`
	PositionFill            DefinitionOrderPositionFill       `json:"positionFill"`
	TradeState              string                            `json:"tradeState"`
	TakeProfitOnFill        DefinitionTakeProfitDetails       `json:"takeProfitOnFill"`
	StopLossOnFill          DefinitionStopLossDetails         `json:"stopLossOnFill"`
	TrailingStopLossOnFill  DefinitionTrailingStopLossDetails `json:"trailingStopLossOnFill"`
	TradeClientExtensions   DefinitionClientExtensions        `json:"tradeClientExtensions"`
	FillingTransactionID    DefinitionTransactionID           `json:"fillingTransactionID"`
	FilledTime              DefinitionDateTime                `json:"filledTime"`
	TradeOpenedID           DefinitionTradeID                 `json:"tradeOpenedID"`
	TradeReducedID          DefinitionTradeID                 `json:"tradeReducedID"`
	TradeClosedIDs          []DefinitionTradeID               `json:"tradeClosedIDs"`
	CancellingTransactionID DefinitionTransactionID           `json:"cancellingTransactionID"`
	CancelledTime           DefinitionDateTime                `json:"cancelledTime"`
}

type DefinitionLimitOrder struct {
	Id                      DefinitionOrderID                 `json:"id"`
	CreateTime              DefinitionDateTime                `json:"createTime"`
	State                   DefinitionOrderState              `json:"state"`
	ClientExtensions        DefinitionClientExtensions        `json:"clientExtensions"`
	Type                    DefinitionOrderType               `json:"type"`
	Instrument              DefinitionInstrumentName          `json:"instrument"`
	Units                   DefinitionDecimalNumber           `json:"units"`
	Price                   DefinitionPriceValue              `json:"price"`
	TimeInForce             DefinitionTimeInForce             `json:"timeInForce"`
	GtdTime                 DefinitionDateTime                `json:"gtdTime"`
	PositionFill            DefinitionOrderPositionFill       `json:"positionFill"`
	TriggerCondition        DefinitionOrderTriggerCondition   `json:"triggerCondition"`
	TakeProfitOnFill        DefinitionTakeProfitDetails       `json:"takeProfitOnFill"`
	StopLossOnFill          DefinitionStopLossDetails         `json:"stopLossOnFill"`
	TrailingStopLossOnFill  DefinitionTrailingStopLossDetails `json:"trailingStopLossOnFill"`
	TradeClientExtensions   DefinitionClientExtensions        `json:"tradeClientExtensions"`
	FillingTransactionID    DefinitionTransactionID           `json:"fillingTransactionID"`
	FilledTime              DefinitionDateTime                `json:"filledTime"`
	TradeOpenedID           DefinitionTradeID                 `json:"tradeOpenedID"`
	TradeReducedID          DefinitionTradeID                 `json:"tradeReducedID"`
	TradeClosedIDs          []DefinitionTradeID               `json:"tradeClosedIDs"`
	CancellingTransactionID DefinitionTransactionID           `json:"cancellingTransactionID"`
	CancelledTime           DefinitionDateTime                `json:"cancelledTime"`
	ReplacesOrderID         DefinitionOrderID                 `json:"replacesOrderID"`
	ReplacedByOrderID       DefinitionOrderID                 `json:"replacedByOrderID"`
}

type DefinitionStopOrder struct {
	Id                      DefinitionOrderID                 `json:"id"`
	CreateTime              DefinitionDateTime                `json:"createTime"`
	State                   DefinitionOrderState              `json:"state"`
	ClientExtensions        DefinitionClientExtensions        `json:"clientExtensions"`
	Type                    DefinitionOrderType               `json:"type"`
	Instrument              DefinitionInstrumentName          `json:"instrument"`
	Units                   DefinitionDecimalNumber           `json:"units"`
	Price                   DefinitionPriceValue              `json:"price"`
	PriceBound              DefinitionPriceValue              `json:"priceBound"`
	TimeInForce             DefinitionTimeInForce             `json:"timeInForce"`
	GtdTime                 DefinitionDateTime                `json:"gtdTime"`
	PositionFill            DefinitionOrderPositionFill       `json:"positionFill"`
	TriggerCondition        DefinitionOrderTriggerCondition   `json:"triggerCondition"`
	TakeProfitOnFill        DefinitionTakeProfitDetails       `json:"takeProfitOnFill"`
	StopLossOnFill          DefinitionStopLossDetails         `json:"stopLossOnFill"`
	TrailingStopLossOnFill  DefinitionTrailingStopLossDetails `json:"trailingStopLossOnFill"`
	TradeClientExtensions   DefinitionClientExtensions        `json:"tradeClientExtensions"`
	FillingTransactionID    DefinitionTransactionID           `json:"fillingTransactionID"`
	FilledTime              DefinitionDateTime                `json:"filledTime"`
	TradeOpenedID           DefinitionTradeID                 `json:"tradeOpenedID"`
	TradeReducedID          DefinitionTradeID                 `json:"tradeReducedID"`
	TradeClosedIDs          []DefinitionTradeID               `json:"tradeClosedIDs"`
	CancellingTransactionID DefinitionTransactionID           `json:"cancellingTransactionID"`
	CancelledTime           DefinitionDateTime                `json:"cancelledTime"`
	ReplacesOrderID         DefinitionOrderID                 `json:"replacesOrderID"`
	ReplacedByOrderID       DefinitionOrderID                 `json:"replacedByOrderID"`
}

type DefinitionMarketIfTouchedOrder struct {
	Id                      DefinitionOrderID                 `json:"id"`
	CreateTime              DefinitionDateTime                `json:"createTime"`
	State                   DefinitionOrderState              `json:"state"`
	ClientExtensions        DefinitionClientExtensions        `json:"clientExtensions"`
	Type                    DefinitionOrderType               `json:"type"`
	Instrument              DefinitionInstrumentName          `json:"instrument"`
	Units                   DefinitionDecimalNumber           `json:"units"`
	Price                   DefinitionPriceValue              `json:"price"`
	PriceBound              DefinitionPriceValue              `json:"priceBound"`
	TimeInForce             DefinitionTimeInForce             `json:"timeInForce"`
	GtdTime                 DefinitionDateTime                `json:"gtdTime"`
	PositionFill            DefinitionOrderPositionFill       `json:"positionFill"`
	TriggerCondition        DefinitionOrderTriggerCondition   `json:"triggerCondition"`
	InitialMarketPrice      DefinitionPriceValue              `json:"initialMarketPrice"`
	TakeProfitOnFill        DefinitionTakeProfitDetails       `json:"takeProfitOnFill"`
	StopLossOnFill          DefinitionStopLossDetails         `json:"stopLossOnFill"`
	TrailingStopLossOnFill  DefinitionTrailingStopLossDetails `json:"trailingStopLossOnFill"`
	TradeClientExtensions   DefinitionClientExtensions        `json:"tradeClientExtensions"`
	FillingTransactionID    DefinitionTransactionID           `json:"fillingTransactionID"`
	FilledTime              DefinitionDateTime                `json:"filledTime"`
	TradeOpenedID           DefinitionTradeID                 `json:"tradeOpenedID"`
	TradeReducedID          DefinitionTradeID                 `json:"tradeReducedID"`
	TradeClosedIDs          []DefinitionTradeID               `json:"tradeClosedIDs"`
	CancellingTransactionID DefinitionTransactionID           `json:"cancellingTransactionID"`
	CancelledTime           DefinitionDateTime                `json:"cancelledTime"`
	ReplacesOrderID         DefinitionOrderID                 `json:"replacesOrderID"`
	ReplacedByOrderID       DefinitionOrderID                 `json:"replacedByOrderID"`
}

type DefinitionTakeProfitOrder struct {
	Id                      DefinitionOrderID               `json:"id"`
	CreateTime              DefinitionDateTime              `json:"createTime"`
	State                   DefinitionOrderState            `json:"state"`
	ClientExtensions        DefinitionClientExtensions      `json:"clientExtensions"`
	Type                    DefinitionOrderType             `json:"type"`
	TradeID                 DefinitionTradeID               `json:"tradeID"`
	ClientTradeID           DefinitionClientID              `json:"clientTradeID"`
	Price                   DefinitionPriceValue            `json:"price"`
	TimeInForce             DefinitionTimeInForce           `json:"timeInForce"`
	GtdTime                 DefinitionDateTime              `json:"gtdTime"`
	TriggerCondition        DefinitionOrderTriggerCondition `json:"triggerCondition"`
	FillingTransactionID    DefinitionTransactionID         `json:"fillingTransactionID"`
	FilledTime              DefinitionDateTime              `json:"filledTime"`
	TradeOpenedID           DefinitionTradeID               `json:"tradeOpenedID"`
	TradeReducedID          DefinitionTradeID               `json:"tradeReducedID"`
	TradeClosedIDs          []DefinitionTradeID             `json:"tradeClosedIDs"`
	CancellingTransactionID DefinitionTransactionID         `json:"cancellingTransactionID"`
	CancelledTime           DefinitionDateTime              `json:"cancelledTime"`
	ReplacesOrderID         DefinitionOrderID               `json:"replacesOrderID"`
	ReplacedByOrderID       DefinitionOrderID               `json:"replacedByOrderID"`
}

type DefinitionStopLossOrder struct {
	Id                         DefinitionOrderID               `json:"id"`
	CreateTime                 DefinitionDateTime              `json:"createTime"`
	State                      DefinitionOrderState            `json:"state"`
	ClientExtensions           DefinitionClientExtensions      `json:"clientExtensions"`
	Type                       DefinitionOrderType             `json:"type"`
	GuaranteedExecutionPremium DefinitionDecimalNumber         `json:"guaranteedExecutionPremium"`
	TradeID                    DefinitionTradeID               `json:"tradeID"`
	ClientTradeID              DefinitionClientID              `json:"clientTradeID"`
	Price                      DefinitionPriceValue            `json:"price"`
	Distance                   DefinitionDecimalNumber         `json:"distance"`
	TimeInForce                DefinitionTimeInForce           `json:"timeInForce"`
	GtdTime                    DefinitionDateTime              `json:"gtdTime"`
	TriggerCondition           DefinitionOrderTriggerCondition `json:"triggerCondition"`
	Guaranteed                 bool                            `json:"guaranteed"`
	FillingTransactionID       DefinitionTransactionID         `json:"fillingTransactionID"`
	FilledTime                 DefinitionDateTime              `json:"filledTime"`
	TradeOpenedID              DefinitionTradeID               `json:"tradeOpenedID"`
	TradeReducedID             DefinitionTradeID               `json:"tradeReducedID"`
	TradeClosedIDs             []DefinitionTradeID             `json:"tradeClosedIDs"`
	CancellingTransactionID    DefinitionTransactionID         `json:"cancellingTransactionID"`
	CancelledTime              DefinitionDateTime              `json:"cancelledTime"`
	ReplacesOrderID            DefinitionOrderID               `json:"replacesOrderID"`
	ReplacedByOrderID          DefinitionOrderID               `json:"replacedByOrderID"`
}

type DefinitionTrailingStopLossOrder struct {
	Id                      DefinitionOrderID               `json:"id"`
	CreateTime              DefinitionDateTime              `json:"createTime"`
	State                   DefinitionOrderState            `json:"state"`
	ClientExtensions        DefinitionClientExtensions      `json:"clientExtensions"`
	Type                    DefinitionOrderType             `json:"type"`
	TradeID                 DefinitionTradeID               `json:"tradeID"`
	ClientTradeID           DefinitionClientID              `json:"clientTradeID"`
	Distance                DefinitionDecimalNumber         `json:"distance"`
	TimeInForce             DefinitionTimeInForce           `json:"timeInForce"`
	GtdTime                 DefinitionDateTime              `json:"gtdTime"`
	TriggerCondition        DefinitionOrderTriggerCondition `json:"triggerCondition"`
	TrailingStopValue       DefinitionPriceValue            `json:"trailingStopValue"`
	FillingTransactionID    DefinitionTransactionID         `json:"fillingTransactionID"`
	FilledTime              DefinitionDateTime              `json:"filledTime"`
	TradeOpenedID           DefinitionTradeID               `json:"tradeOpenedID"`
	TradeReducedID          DefinitionTradeID               `json:"tradeReducedID"`
	TradeClosedIDs          []DefinitionTradeID             `json:"tradeClosedIDs"`
	CancellingTransactionID DefinitionTransactionID         `json:"cancellingTransactionID"`
	CancelledTime           DefinitionDateTime              `json:"cancelledTime"`
	ReplacesOrderID         DefinitionOrderID               `json:"replacesOrderID"`
	ReplacedByOrderID       DefinitionOrderID               `json:"replacedByOrderID"`
}

// Order Requests

// type DefinitionOrderRequest
// TODO: Implemented by: MarketOrderRequest, LimitOrderRequest, StopOrderRequest, MarketIfTouchedOrderRequest, TakeProfitOrderRequest, StopLossOrderRequest, TrailingStopLossOrderRequest

type DefinitionMarketOrderRequest struct {
	Type                   DefinitionOrderType               `json:"type"`
	Instrument             DefinitionInstrumentName          `json:"instrument"`
	Units                  DefinitionDecimalNumber           `json:"units"`
	TimeInForce            DefinitionTimeInForce             `json:"timeInForce"`
	PriceBound             DefinitionPriceValue              `json:"priceBound"`
	PositionFill           DefinitionOrderPositionFill       `json:"positionFill"`
	ClientExtensions       DefinitionClientExtensions        `json:"clientExtensions"`
	TakeProfitOnFill       DefinitionTakeProfitDetails       `json:"takeProfitOnFill"`
	StopLossOnFill         DefinitionStopLossDetails         `json:"stopLossOnFill"`
	TrailingStopLossOnFill DefinitionTrailingStopLossDetails `json:"trailingStopLossOnFill"`
	TradeClientExtensions  DefinitionClientExtensions        `json:"tradeClientExtensions"`
}

type DefinitionLimitOrderRequest struct {
	Type                   DefinitionOrderType               `json:"type"`
	Instrument             DefinitionInstrumentName          `json:"instrument"`
	Units                  DefinitionDecimalNumber           `json:"units"`
	Price                  DefinitionPriceValue              `json:"price"`
	TimeInForce            DefinitionTimeInForce             `json:"timeInForce"`
	GtdTime                DefinitionDateTime                `json:"gtdTime"`
	PositionFill           DefinitionOrderPositionFill       `json:"positionFill"`
	TriggerCondition       DefinitionOrderTriggerCondition   `json:"triggerCondition"`
	ClientExtensions       DefinitionClientExtensions        `json:"clientExtensions"`
	TakeProfitOnFill       DefinitionTakeProfitDetails       `json:"takeProfitOnFill"`
	StopLossOnFill         DefinitionStopLossDetails         `json:"stopLossOnFill"`
	TrailingStopLossOnFill DefinitionTrailingStopLossDetails `json:"trailingStopLossOnFill"`
	TradeClientExtensions  DefinitionClientExtensions        `json:"tradeClientExtensions"`
}

type DefinitionStopOrderRequest struct {
	Type                   DefinitionOrderType               `json:"type"`
	Instrument             DefinitionInstrumentName          `json:"instrument"`
	Units                  DefinitionDecimalNumber           `json:"units"`
	Price                  DefinitionPriceValue              `json:"price"`
	PriceBound             DefinitionPriceValue              `json:"priceBound"`
	TimeInForce            DefinitionTimeInForce             `json:"timeInForce"`
	GtdTime                DefinitionDateTime                `json:"gtdTime"`
	PositionFill           DefinitionOrderPositionFill       `json:"positionFill"`
	TriggerCondition       DefinitionOrderTriggerCondition   `json:"triggerCondition"`
	ClientExtensions       DefinitionClientExtensions        `json:"clientExtensions"`
	TakeProfitOnFill       DefinitionTakeProfitDetails       `json:"takeProfitOnFill"`
	StopLossOnFill         DefinitionStopLossDetails         `json:"stopLossOnFill"`
	TrailingStopLossOnFill DefinitionTrailingStopLossDetails `json:"trailingStopLossOnFill"`
	TradeClientExtensions  DefinitionClientExtensions        `json:"tradeClientExtensions"`
}

type DefinitionMarketIfTouchedOrderRequest struct {
	Type                   DefinitionOrderType               `json:"type"`
	Instrument             DefinitionInstrumentName          `json:"instrument"`
	Units                  DefinitionDecimalNumber           `json:"units"`
	Price                  DefinitionPriceValue              `json:"price"`
	PriceBound             DefinitionPriceValue              `json:"priceBound"`
	TimeInForce            DefinitionTimeInForce             `json:"timeInForce"`
	GtdTime                DefinitionDateTime                `json:"gtdTime"`
	PositionFill           DefinitionOrderPositionFill       `json:"positionFill"`
	TriggerCondition       DefinitionOrderTriggerCondition   `json:"triggerCondition"`
	ClientExtensions       DefinitionClientExtensions        `json:"clientExtensions"`
	TakeProfitOnFill       DefinitionTakeProfitDetails       `json:"takeProfitOnFill"`
	StopLossOnFill         DefinitionStopLossDetails         `json:"stopLossOnFill"`
	TrailingStopLossOnFill DefinitionTrailingStopLossDetails `json:"trailingStopLossOnFill"`
	TradeClientExtensions  DefinitionClientExtensions        `json:"tradeClientExtensions"`
}

type DefinitionTakeProfitOrderRequest struct {
	Type             DefinitionOrderType             `json:"type"`
	TradeID          DefinitionTradeID               `json:"tradeID"`
	ClientTradeID    DefinitionClientID              `json:"clientTradeID"`
	Price            DefinitionPriceValue            `json:"price"`
	TimeInForce      DefinitionTimeInForce           `json:"timeInForce"`
	GtdTime          DefinitionDateTime              `json:"gtdTime"`
	TriggerCondition DefinitionOrderTriggerCondition `json:"triggerCondition"`
	ClientExtensions DefinitionClientExtensions      `json:"clientExtensions"`
}

type DefinitionStopLossOrderRequest struct {
	Type             DefinitionOrderType             `json:"type"`
	TradeID          DefinitionTradeID               `json:"tradeID"`
	ClientTradeID    DefinitionClientID              `json:"clientTradeID"`
	Price            DefinitionPriceValue            `json:"price"`
	Distance         DefinitionDecimalNumber         `json:"distance"`
	TimeInForce      DefinitionTimeInForce           `json:"timeInForce"`
	GtdTime          DefinitionDateTime              `json:"gtdTime"`
	TriggerCondition DefinitionOrderTriggerCondition `json:"triggerCondition"`
	Guaranteed       bool                            `json:"guaranteed"`
	ClientExtensions DefinitionClientExtensions      `json:"clientExtensions"`
}

type DefinitionTrailingStopLossOrderRequest struct {
	Type             DefinitionOrderType             `json:"type"`
	TradeID          DefinitionTradeID               `json:"tradeID"`
	ClientTradeID    DefinitionClientID              `json:"clientTradeID"`
	Distance         DefinitionDecimalNumber         `json:"distance"`
	TimeInForce      DefinitionTimeInForce           `json:"timeInForce"`
	GtdTime          DefinitionDateTime              `json:"gtdTime"`
	TriggerCondition DefinitionOrderTriggerCondition `json:"triggerCondition"`
	ClientExtensions DefinitionClientExtensions      `json:"clientExtensions"`
}

// Order-related Definitions

type DefinitionOrderID string

type DefinitionOrderType string

type DefinitionCancellableOrderType string

type DefinitionOrderState string

type DefinitionOrderStateFilter string

type DefinitionOrderIdentifier struct {
	OrderID       DefinitionOrderID  `json:"orderID"`
	ClientOrderID DefinitionClientID `json:"clientOrderID"`
}

type DefinitionOrderSpecifier string

type DefinitionTimeInForce string

type DefinitionOrderPositionFill string

type DefinitionOrderTriggerCondition string

type DefinitionDynamicOrderState struct {
	Id                     DefinitionOrderID    `json:"id"`
	TrailingStopValue      DefinitionPriceValue `json:"trailingStopValue"`
	TriggerDistance        DefinitionPriceValue `json:"triggerDistance"`
	IsTriggerDistanceExact bool                 `json:"isTriggerDistanceExact"`
}

type DefinitionUnitsAvailableDetails struct {
	Long  DefinitionDecimalNumber `json:"long"`
	Short DefinitionDecimalNumber `json:"short"`
}

type DefinitionUnitsAvailable struct {
	Default     DefinitionUnitsAvailableDetails `json:"default"`
	ReduceFirst DefinitionUnitsAvailableDetails `json:"reduceFirst"`
	ReduceOnly  DefinitionUnitsAvailableDetails `json:"reduceOnly"`
	OpenOnly    DefinitionUnitsAvailableDetails `json:"openOnly"`
}

type DefinitionGuaranteedStopLossOrderEntryData struct {
	MinimumDistance  DefinitionDecimalNumber                           `json:"minimumDistance"`
	Premium          DefinitionDecimalNumber                           `json:"premium"`
	LevelRestriction DefinitionGuaranteedStopLossOrderLevelRestriction `json:"levelRestriction"`
}

//
// Trade Definitions
//

type DefinitionTradeID string

type DefinitionTradeState string

type DefinitionTradeStateFilter string

type DefinitionTradeSpecifier string

type DefinitionTrade struct {
	Id                    DefinitionTradeID               `json:"id"`
	Instrument            DefinitionInstrumentName        `json:"instrument"`
	Price                 DefinitionPriceValue            `json:"price"`
	OpenTime              DefinitionDateTime              `json:"openTime"`
	State                 DefinitionTradeState            `json:"state"`
	InitialUnits          DefinitionDecimalNumber         `json:"initialUnits"`
	InitialMarginRequired DefinitionAccountUnits          `json:"initialMarginRequired"`
	CurrentUnits          DefinitionDecimalNumber         `json:"currentUnits"`
	RealizedPL            DefinitionAccountUnits          `json:"realizedPL"`
	UnrealizedPL          DefinitionAccountUnits          `json:"unrealizedPL"`
	MarginUsed            DefinitionAccountUnits          `json:"marginUsed"`
	AverageClosePrice     DefinitionPriceValue            `json:"averageClosePrice"`
	ClosingTransactionIDs []DefinitionTransactionID       `json:"closingTransactionIDs"`
	Financing             DefinitionAccountUnits          `json:"financing"`
	CloseTime             DefinitionDateTime              `json:"closeTime"`
	ClientExtensions      DefinitionClientExtensions      `json:"clientExtensions"`
	TakeProfitOrder       DefinitionTakeProfitOrder       `json:"takeProfitOrder"`
	StopLossOrder         DefinitionStopLossOrder         `json:"stopLossOrder"`
	TrailingStopLossOrder DefinitionTrailingStopLossOrder `json:"trailingStopLossOrder"`
}

type DefinitionTradeSummary struct {
	Id                      DefinitionTradeID          `json:"id"`
	Instrument              DefinitionInstrumentName   `json:"instrument"`
	Price                   DefinitionPriceValue       `json:"price"`
	OpenTime                DefinitionDateTime         `json:"openTime"`
	State                   DefinitionTradeState       `json:"state"`
	InitialUnits            DefinitionDecimalNumber    `json:"initialUnits"`
	InitialMarginRequired   DefinitionAccountUnits     `json:"initialMarginRequired"`
	CurrentUnits            DefinitionDecimalNumber    `json:"currentUnits"`
	RealizedPL              DefinitionAccountUnits     `json:"realizedPL"`
	UnrealizedPL            DefinitionAccountUnits     `json:"unrealizedPL"`
	MarginUsed              DefinitionAccountUnits     `json:"marginUsed"`
	AverageClosePrice       DefinitionPriceValue       `json:"averageClosePrice"`
	ClosingTransactionIDs   []DefinitionTransactionID  `json:"closingTransactionIDs"`
	Financing               DefinitionAccountUnits     `json:"financing"`
	CloseTime               DefinitionDateTime         `json:"closeTime"`
	ClientExtensions        DefinitionClientExtensions `json:"clientExtensions"`
	TakeProfitOrderID       DefinitionOrderID          `json:"takeProfitOrderID"`
	StopLossOrderID         DefinitionOrderID          `json:"stopLossOrderID"`
	TrailingStopLossOrderID DefinitionOrderID          `json:"trailingStopLossOrderID"`
}

type DefinitionCalculatedTradeState struct {
	Id           DefinitionTradeID      `json:"id"`
	UnrealizedPL DefinitionAccountUnits `json:"unrealizedPL"`
	MarginUsed   DefinitionAccountUnits `json:"marginUsed"`
}

type DefinitionTradePL string

//
// Position Definitions
//

type DefinitionPosition struct {
	Instrument              DefinitionInstrumentName `json:"instrument"`
	Pl                      DefinitionAccountUnits   `json:"pl"`
	UnrealizedPL            DefinitionAccountUnits   `json:"unrealizedPL"`
	MarginUsed              DefinitionAccountUnits   `json:"marginUsed"`
	ResettablePL            DefinitionAccountUnits   `json:"resettablePL"`
	Financing               DefinitionAccountUnits   `json:"financing"`
	Commission              DefinitionAccountUnits   `json:"commission"`
	GuaranteedExecutionFees DefinitionAccountUnits   `json:"guaranteedExecutionFees"`
	Long                    DefinitionPositionSide   `json:"long"`
	Short                   DefinitionPositionSide   `json:"short"`
}

type DefinitionPositionSide struct {
	Units                   DefinitionDecimalNumber `json:"units"`
	AveragePrice            DefinitionPriceValue    `json:"averagePrice"`
	TradeIDs                []DefinitionTradeID     `json:"tradeIDs"`
	Pl                      DefinitionAccountUnits  `json:"pl"`
	UnrealizedPL            DefinitionAccountUnits  `json:"unrealizedPL"`
	ResettablePL            DefinitionAccountUnits  `json:"resettablePL"`
	Financing               DefinitionAccountUnits  `json:"financing"`
	GuaranteedExecutionFees DefinitionAccountUnits  `json:"guaranteedExecutionFees"`
}

type DefinitionCalculatedPositionState struct {
	Instrument        DefinitionInstrumentName `json:"instrument"`
	NetUnrealizedPL   DefinitionAccountUnits   `json:"netUnrealizedPL"`
	LongUnrealizedPL  DefinitionAccountUnits   `json:"longUnrealizedPL"`
	ShortUnrealizedPL DefinitionAccountUnits   `json:"shortUnrealizedPL"`
	MarginUsed        DefinitionAccountUnits   `json:"marginUsed"`
}

//
// Transaction Definitions
//

// Transactions

// TODO: Implemented by: OrderFillTransaction, OrderCancelTransaction, OrderCancelRejectTransaction, OrderClientExtensionsModifyTransaction, OrderClientExtensionsModifyRejectTransaction, CreateTransaction, CloseTransaction, ReopenTransaction, ClientConfigureTransaction, ClientConfigureRejectTransaction, TransferFundsTransaction, TransferFundsRejectTransaction, MarketOrderTransaction, MarketOrderRejectTransaction, FixedPriceOrderTransaction, LimitOrderTransaction, LimitOrderRejectTransaction, StopOrderTransaction, StopOrderRejectTransaction, MarketIfTouchedOrderTransaction, MarketIfTouchedOrderRejectTransaction, TakeProfitOrderTransaction, TakeProfitOrderRejectTransaction, StopLossOrderTransaction, StopLossOrderRejectTransaction, TrailingStopLossOrderTransaction, TrailingStopLossOrderRejectTransaction, TradeClientExtensionsModifyTransaction, TradeClientExtensionsModifyRejectTransaction, MarginCallEnterTransaction, MarginCallExtendTransaction, MarginCallExitTransaction, DelayedTradeClosureTransaction, DailyFinancingTransaction, ResetResettablePLTransaction
type DefinitionTransaction struct {
	Id        DefinitionTransactionID `json:"id"`
	Time      DefinitionDateTime      `json:"time"`
	UserID    int                     `json:"userID"`
	AccountID DefinitionAccountID     `json:"accountID"`
	BatchID   DefinitionTransactionID `json:"batchID"`
	RequestID DefinitionRequestID     `json:"requestID"`
}

type DefinitionCreateTransaction struct {
	Id            DefinitionTransactionID   `json:"id"`
	Time          DefinitionDateTime        `json:"time"`
	UserID        int                       `json:"userID"`
	AccountID     DefinitionAccountID       `json:"accountID"`
	BatchID       DefinitionTransactionID   `json:"batchID"`
	RequestID     DefinitionRequestID       `json:"requestID"`
	Type          DefinitionTransactionType `json:"type"`
	DivisionID    int                       `json:"divisionID"`
	SiteID        int                       `json:"siteID"`
	AccountUserID int                       `json:"accountUserID"`
	AccountNumber int                       `json:"accountNumber"`
	HomeCurrency  DefinitionCurrency        `json:"homeCurrency"`
}

type DefinitionCloseTransaction struct {
	Id        DefinitionTransactionID   `json:"id"`
	Time      DefinitionDateTime        `json:"time"`
	UserID    int                       `json:"userID"`
	AccountID DefinitionAccountID       `json:"accountID"`
	BatchID   DefinitionTransactionID   `json:"batchID"`
	RequestID DefinitionRequestID       `json:"requestID"`
	Type      DefinitionTransactionType `json:"type"`
}

type DefinitionReopenTransaction struct {
	Id        DefinitionTransactionID   `json:"id"`
	Time      DefinitionDateTime        `json:"time"`
	UserID    int                       `json:"userID"`
	AccountID DefinitionAccountID       `json:"accountID"`
	BatchID   DefinitionTransactionID   `json:"batchID"`
	RequestID DefinitionRequestID       `json:"requestID"`
	Type      DefinitionTransactionType `json:"type"`
}

type DefinitionClientConfigureTransaction struct {
	Id         DefinitionTransactionID   `json:"id"`
	Time       DefinitionDateTime        `json:"time"`
	UserID     int                       `json:"userID"`
	AccountID  DefinitionAccountID       `json:"accountID"`
	BatchID    DefinitionTransactionID   `json:"batchID"`
	RequestID  DefinitionRequestID       `json:"requestID"`
	Type       DefinitionTransactionType `json:"type"`
	Alias      string                    `json:"alias"`
	MarginRate DefinitionDecimalNumber   `json:"marginRate"`
}

type DefinitionClientConfigureRejectTransaction struct {
	Id           DefinitionTransactionID           `json:"id"`
	Time         DefinitionDateTime                `json:"time"`
	UserID       int                               `json:"userID"`
	AccountID    DefinitionAccountID               `json:"accountID"`
	BatchID      DefinitionTransactionID           `json:"batchID"`
	RequestID    DefinitionRequestID               `json:"requestID"`
	Type         DefinitionTransactionType         `json:"type"`
	Alias        string                            `json:"alias"`
	MarginRate   DefinitionDecimalNumber           `json:"marginRate"`
	RejectReason DefinitionTransactionRejectReason `json:"rejectReason"`
}

type DefinitionTransferFundsTransaction struct {
	Id             DefinitionTransactionID   `json:"id"`
	Time           DefinitionDateTime        `json:"time"`
	UserID         int                       `json:"userID"`
	AccountID      DefinitionAccountID       `json:"accountID"`
	BatchID        DefinitionTransactionID   `json:"batchID"`
	RequestID      DefinitionRequestID       `json:"requestID"`
	Type           DefinitionTransactionType `json:"type"`
	Amount         DefinitionAccountUnits    `json:"amount"`
	FundingReason  DefinitionFundingReason   `json:"fundingReason"`
	Comment        string                    `json:"comment"`
	AccountBalance DefinitionAccountUnits    `json:"accountBalance"`
}

type DefinitionTransferFundsRejectTransaction struct {
	Id            DefinitionTransactionID           `json:"id"`
	Time          DefinitionDateTime                `json:"time"`
	UserID        int                               `json:"userID"`
	AccountID     DefinitionAccountID               `json:"accountID"`
	BatchID       DefinitionTransactionID           `json:"batchID"`
	RequestID     DefinitionRequestID               `json:"requestID"`
	Type          DefinitionTransactionType         `json:"type"`
	Amount        DefinitionAccountUnits            `json:"amount"`
	FundingReason DefinitionFundingReason           `json:"fundingReason"`
	Comment       string                            `json:"comment"`
	RejectReason  DefinitionTransactionRejectReason `json:"rejectReason"`
}

type DefinitionMarketOrderTransaction struct {
	Id                     DefinitionTransactionID                `json:"id"`
	Time                   DefinitionDateTime                     `json:"time"`
	UserID                 int                                    `json:"userID"`
	AccountID              DefinitionAccountID                    `json:"accountID"`
	BatchID                DefinitionTransactionID                `json:"batchID"`
	RequestID              DefinitionRequestID                    `json:"requestID"`
	Type                   DefinitionTransactionType              `json:"type"`
	Instrument             DefinitionInstrumentName               `json:"instrument"`
	Units                  DefinitionDecimalNumber                `json:"units"`
	TimeInForce            DefinitionTimeInForce                  `json:"timeInForce"`
	PriceBound             DefinitionPriceValue                   `json:"priceBound"`
	PositionFill           DefinitionOrderPositionFill            `json:"positionFill"`
	TradeClose             DefinitionMarketOrderTradeClose        `json:"tradeClose"`
	LongPositionCloseout   DefinitionMarketOrderPositionCloseout  `json:"longPositionCloseout"`
	ShortPositionCloseout  DefinitionMarketOrderPositionCloseout  `json:"shortPositionCloseout"`
	MarginCloseout         DefinitionMarketOrderMarginCloseout    `json:"marginCloseout"`
	DelayedTradeClose      DefinitionMarketOrderDelayedTradeClose `json:"delayedTradeClose"`
	Reason                 DefinitionMarketOrderReason            `json:"reason"`
	ClientExtensions       DefinitionClientExtensions             `json:"clientExtensions"`
	TakeProfitOnFill       DefinitionTakeProfitDetails            `json:"takeProfitOnFill"`
	StopLossOnFill         DefinitionStopLossDetails              `json:"stopLossOnFill"`
	TrailingStopLossOnFill DefinitionTrailingStopLossDetails      `json:"trailingStopLossOnFill"`
	TradeClientExtensions  DefinitionClientExtensions             `json:"tradeClientExtensions"`
}

type DefinitionMarketOrderRejectTransaction struct {
	Id                     DefinitionTransactionID                `json:"id"`
	Time                   DefinitionDateTime                     `json:"time"`
	UserID                 int                                    `json:"userID"`
	AccountID              DefinitionAccountID                    `json:"accountID"`
	BatchID                DefinitionTransactionID                `json:"batchID"`
	RequestID              DefinitionRequestID                    `json:"requestID"`
	Type                   DefinitionTransactionType              `json:"type"`
	Instrument             DefinitionInstrumentName               `json:"instrument"`
	Units                  DefinitionDecimalNumber                `json:"units"`
	TimeInForce            DefinitionTimeInForce                  `json:"timeInForce"`
	PriceBound             DefinitionPriceValue                   `json:"priceBound"`
	PositionFill           DefinitionOrderPositionFill            `json:"positionFill"`
	TradeClose             DefinitionMarketOrderTradeClose        `json:"tradeClose"`
	LongPositionCloseout   DefinitionMarketOrderPositionCloseout  `json:"longPositionCloseout"`
	ShortPositionCloseout  DefinitionMarketOrderPositionCloseout  `json:"shortPositionCloseout"`
	MarginCloseout         DefinitionMarketOrderMarginCloseout    `json:"marginCloseout"`
	DelayedTradeClose      DefinitionMarketOrderDelayedTradeClose `json:"delayedTradeClose"`
	Reason                 DefinitionMarketOrderReason            `json:"reason"`
	ClientExtensions       DefinitionClientExtensions             `json:"clientExtensions"`
	TakeProfitOnFill       DefinitionTakeProfitDetails            `json:"takeProfitOnFill"`
	StopLossOnFill         DefinitionStopLossDetails              `json:"stopLossOnFill"`
	TrailingStopLossOnFill DefinitionTrailingStopLossDetails      `json:"trailingStopLossOnFill"`
	TradeClientExtensions  DefinitionClientExtensions             `json:"tradeClientExtensions"`
	RejectReason           DefinitionTransactionRejectReason      `json:"rejectReason"`
}

type DefinitionFixedPriceOrderTransaction struct {
	Id                     DefinitionTransactionID           `json:"id"`
	Time                   DefinitionDateTime                `json:"time"`
	UserID                 int                               `json:"userID"`
	AccountID              DefinitionAccountID               `json:"accountID"`
	BatchID                DefinitionTransactionID           `json:"batchID"`
	RequestID              DefinitionRequestID               `json:"requestID"`
	Type                   DefinitionTransactionType         `json:"type"`
	Instrument             DefinitionInstrumentName          `json:"instrument"`
	Units                  DefinitionDecimalNumber           `json:"units"`
	Price                  DefinitionPriceValue              `json:"price"`
	PositionFill           DefinitionOrderPositionFill       `json:"positionFill"`
	TradeState             string                            `json:"tradeState"`
	Reason                 DefinitionFixedPriceOrderReason   `json:"reason"`
	ClientExtensions       DefinitionClientExtensions        `json:"clientExtensions"`
	TakeProfitOnFill       DefinitionTakeProfitDetails       `json:"takeProfitOnFill"`
	StopLossOnFill         DefinitionStopLossDetails         `json:"stopLossOnFill"`
	TrailingStopLossOnFill DefinitionTrailingStopLossDetails `json:"trailingStopLossOnFill"`
	TradeClientExtensions  DefinitionClientExtensions        `json:"tradeClientExtensions"`
}

type DefinitionLimitOrderTransaction struct {
	Id                      DefinitionTransactionID           `json:"id"`
	Time                    DefinitionDateTime                `json:"time"`
	UserID                  int                               `json:"userID"`
	AccountID               DefinitionAccountID               `json:"accountID"`
	BatchID                 DefinitionTransactionID           `json:"batchID"`
	RequestID               DefinitionRequestID               `json:"requestID"`
	Type                    DefinitionTransactionType         `json:"type"`
	Instrument              DefinitionInstrumentName          `json:"instrument"`
	Units                   DefinitionDecimalNumber           `json:"units"`
	Price                   DefinitionPriceValue              `json:"price"`
	TimeInForce             DefinitionTimeInForce             `json:"timeInForce"`
	GtdTime                 DefinitionDateTime                `json:"gtdTime"`
	PositionFill            DefinitionOrderPositionFill       `json:"positionFill"`
	TriggerCondition        DefinitionOrderTriggerCondition   `json:"triggerCondition"`
	Reason                  DefinitionLimitOrderReason        `json:"reason"`
	ClientExtensions        DefinitionClientExtensions        `json:"clientExtensions"`
	TakeProfitOnFill        DefinitionTakeProfitDetails       `json:"takeProfitOnFill"`
	StopLossOnFill          DefinitionStopLossDetails         `json:"stopLossOnFill"`
	TrailingStopLossOnFill  DefinitionTrailingStopLossDetails `json:"trailingStopLossOnFill"`
	TradeClientExtensions   DefinitionClientExtensions        `json:"tradeClientExtensions"`
	ReplacesOrderID         DefinitionOrderID                 `json:"replacesOrderID"`
	CancellingTransactionID DefinitionTransactionID           `json:"cancellingTransactionID"`
}

type DefinitionLimitOrderRejectTransaction struct {
	Id                      DefinitionTransactionID           `json:"id"`
	Time                    DefinitionDateTime                `json:"time"`
	UserID                  int                               `json:"userID"`
	AccountID               DefinitionAccountID               `json:"accountID"`
	BatchID                 DefinitionTransactionID           `json:"batchID"`
	RequestID               DefinitionRequestID               `json:"requestID"`
	Type                    DefinitionTransactionType         `json:"type"`
	Instrument              DefinitionInstrumentName          `json:"instrument"`
	Units                   DefinitionDecimalNumber           `json:"units"`
	Price                   DefinitionPriceValue              `json:"price"`
	TimeInForce             DefinitionTimeInForce             `json:"timeInForce"`
	GtdTime                 DefinitionDateTime                `json:"gtdTime"`
	PositionFill            DefinitionOrderPositionFill       `json:"positionFill"`
	TriggerCondition        DefinitionOrderTriggerCondition   `json:"triggerCondition"`
	Reason                  DefinitionLimitOrderReason        `json:"reason"`
	ClientExtensions        DefinitionClientExtensions        `json:"clientExtensions"`
	TakeProfitOnFill        DefinitionTakeProfitDetails       `json:"takeProfitOnFill"`
	StopLossOnFill          DefinitionStopLossDetails         `json:"stopLossOnFill"`
	TrailingStopLossOnFill  DefinitionTrailingStopLossDetails `json:"trailingStopLossOnFill"`
	TradeClientExtensions   DefinitionClientExtensions        `json:"tradeClientExtensions"`
	IntendedReplacesOrderID DefinitionOrderID                 `json:"intendedReplacesOrderID"`
	RejectReason            DefinitionTransactionRejectReason `json:"rejectReason"`
}

type DefinitionStopOrderTransaction struct {
	Id                      DefinitionTransactionID           `json:"id"`
	Time                    DefinitionDateTime                `json:"time"`
	UserID                  int                               `json:"userID"`
	AccountID               DefinitionAccountID               `json:"accountID"`
	BatchID                 DefinitionTransactionID           `json:"batchID"`
	RequestID               DefinitionRequestID               `json:"requestID"`
	Type                    DefinitionTransactionType         `json:"type"`
	Instrument              DefinitionInstrumentName          `json:"instrument"`
	Units                   DefinitionDecimalNumber           `json:"units"`
	Price                   DefinitionPriceValue              `json:"price"`
	PriceBound              DefinitionPriceValue              `json:"priceBound"`
	TimeInForce             DefinitionTimeInForce             `json:"timeInForce"`
	GtdTime                 DefinitionDateTime                `json:"gtdTime"`
	PositionFill            DefinitionOrderPositionFill       `json:"positionFill"`
	TriggerCondition        DefinitionOrderTriggerCondition   `json:"triggerCondition"`
	Reason                  DefinitionStopOrderReason         `json:"reason"`
	ClientExtensions        DefinitionClientExtensions        `json:"clientExtensions"`
	TakeProfitOnFill        DefinitionTakeProfitDetails       `json:"takeProfitOnFill"`
	StopLossOnFill          DefinitionStopLossDetails         `json:"stopLossOnFill"`
	TrailingStopLossOnFill  DefinitionTrailingStopLossDetails `json:"trailingStopLossOnFill"`
	TradeClientExtensions   DefinitionClientExtensions        `json:"tradeClientExtensions"`
	ReplacesOrderID         DefinitionOrderID                 `json:"replacesOrderID"`
	CancellingTransactionID DefinitionTransactionID           `json:"cancellingTransactionID"`
}

type DefinitionStopOrderRejectTransaction struct {
	Id                      DefinitionTransactionID           `json:"id"`
	Time                    DefinitionDateTime                `json:"time"`
	UserID                  int                               `json:"userID"`
	AccountID               DefinitionAccountID               `json:"accountID"`
	BatchID                 DefinitionTransactionID           `json:"batchID"`
	RequestID               DefinitionRequestID               `json:"requestID"`
	Type                    DefinitionTransactionType         `json:"type"`
	Instrument              DefinitionInstrumentName          `json:"instrument"`
	Units                   DefinitionDecimalNumber           `json:"units"`
	Price                   DefinitionPriceValue              `json:"price"`
	PriceBound              DefinitionPriceValue              `json:"priceBound"`
	TimeInForce             DefinitionTimeInForce             `json:"timeInForce"`
	GtdTime                 DefinitionDateTime                `json:"gtdTime"`
	PositionFill            DefinitionOrderPositionFill       `json:"positionFill"`
	TriggerCondition        DefinitionOrderTriggerCondition   `json:"triggerCondition"`
	Reason                  DefinitionStopOrderReason         `json:"reason"`
	ClientExtensions        DefinitionClientExtensions        `json:"clientExtensions"`
	TakeProfitOnFill        DefinitionTakeProfitDetails       `json:"takeProfitOnFill"`
	StopLossOnFill          DefinitionStopLossDetails         `json:"stopLossOnFill"`
	TrailingStopLossOnFill  DefinitionTrailingStopLossDetails `json:"trailingStopLossOnFill"`
	TradeClientExtensions   DefinitionClientExtensions        `json:"tradeClientExtensions"`
	IntendedReplacesOrderID DefinitionOrderID                 `json:"intendedReplacesOrderID"`
	RejectReason            DefinitionTransactionRejectReason `json:"rejectReason"`
}

type DefinitionMarketIfTouchedOrderTransaction struct {
	Id                      DefinitionTransactionID              `json:"id"`
	Time                    DefinitionDateTime                   `json:"time"`
	UserID                  int                                  `json:"userID"`
	AccountID               DefinitionAccountID                  `json:"accountID"`
	BatchID                 DefinitionTransactionID              `json:"batchID"`
	RequestID               DefinitionRequestID                  `json:"requestID"`
	Type                    DefinitionTransactionType            `json:"type"`
	Instrument              DefinitionInstrumentName             `json:"instrument"`
	Units                   DefinitionDecimalNumber              `json:"units"`
	Price                   DefinitionPriceValue                 `json:"price"`
	PriceBound              DefinitionPriceValue                 `json:"priceBound"`
	TimeInForce             DefinitionTimeInForce                `json:"timeInForce"`
	GtdTime                 DefinitionDateTime                   `json:"gtdTime"`
	PositionFill            DefinitionOrderPositionFill          `json:"positionFill"`
	TriggerCondition        DefinitionOrderTriggerCondition      `json:"triggerCondition"`
	Reason                  DefinitionMarketIfTouchedOrderReason `json:"reason"`
	ClientExtensions        DefinitionClientExtensions           `json:"clientExtensions"`
	TakeProfitOnFill        DefinitionTakeProfitDetails          `json:"takeProfitOnFill"`
	StopLossOnFill          DefinitionStopLossDetails            `json:"stopLossOnFill"`
	TrailingStopLossOnFill  DefinitionTrailingStopLossDetails    `json:"trailingStopLossOnFill"`
	TradeClientExtensions   DefinitionClientExtensions           `json:"tradeClientExtensions"`
	ReplacesOrderID         DefinitionOrderID                    `json:"replacesOrderID"`
	CancellingTransactionID DefinitionTransactionID              `json:"cancellingTransactionID"`
}

type DefinitionMarketIfTouchedOrderRejectTransaction struct {
	Id                      DefinitionTransactionID              `json:"id"`
	Time                    DefinitionDateTime                   `json:"time"`
	UserID                  int                                  `json:"userID"`
	AccountID               DefinitionAccountID                  `json:"accountID"`
	BatchID                 DefinitionTransactionID              `json:"batchID"`
	RequestID               DefinitionRequestID                  `json:"requestID"`
	Type                    DefinitionTransactionType            `json:"type"`
	Instrument              DefinitionInstrumentName             `json:"instrument"`
	Units                   DefinitionDecimalNumber              `json:"units"`
	Price                   DefinitionPriceValue                 `json:"price"`
	PriceBound              DefinitionPriceValue                 `json:"priceBound"`
	TimeInForce             DefinitionTimeInForce                `json:"timeInForce"`
	GtdTime                 DefinitionDateTime                   `json:"gtdTime"`
	PositionFill            DefinitionOrderPositionFill          `json:"positionFill"`
	TriggerCondition        DefinitionOrderTriggerCondition      `json:"triggerCondition"`
	Reason                  DefinitionMarketIfTouchedOrderReason `json:"reason"`
	ClientExtensions        DefinitionClientExtensions           `json:"clientExtensions"`
	TakeProfitOnFill        DefinitionTakeProfitDetails          `json:"takeProfitOnFill"`
	StopLossOnFill          DefinitionStopLossDetails            `json:"stopLossOnFill"`
	TrailingStopLossOnFill  DefinitionTrailingStopLossDetails    `json:"trailingStopLossOnFill"`
	TradeClientExtensions   DefinitionClientExtensions           `json:"tradeClientExtensions"`
	IntendedReplacesOrderID DefinitionOrderID                    `json:"intendedReplacesOrderID"`
	RejectReason            DefinitionTransactionRejectReason    `json:"rejectReason"`
}

type DefinitionTakeProfitOrderTransaction struct {
	Id                      DefinitionTransactionID         `json:"id"`
	Time                    DefinitionDateTime              `json:"time"`
	UserID                  int                             `json:"userID"`
	AccountID               DefinitionAccountID             `json:"accountID"`
	BatchID                 DefinitionTransactionID         `json:"batchID"`
	RequestID               DefinitionRequestID             `json:"requestID"`
	Type                    DefinitionTransactionType       `json:"type"`
	TradeID                 DefinitionTradeID               `json:"tradeID"`
	ClientTradeID           DefinitionClientID              `json:"clientTradeID"`
	Price                   DefinitionPriceValue            `json:"price"`
	TimeInForce             DefinitionTimeInForce           `json:"timeInForce"`
	GtdTime                 DefinitionDateTime              `json:"gtdTime"`
	TriggerCondition        DefinitionOrderTriggerCondition `json:"triggerCondition"`
	Reason                  DefinitionTakeProfitOrderReason `json:"reason"`
	ClientExtensions        DefinitionClientExtensions      `json:"clientExtensions"`
	OrderFillTransactionID  DefinitionTransactionID         `json:"orderFillTransactionID"`
	ReplacesOrderID         DefinitionOrderID               `json:"replacesOrderID"`
	CancellingTransactionID DefinitionTransactionID         `json:"cancellingTransactionID"`
}

type DefinitionTakeProfitOrderRejectTransaction struct {
	Id                      DefinitionTransactionID           `json:"id"`
	Time                    DefinitionDateTime                `json:"time"`
	UserID                  int                               `json:"userID"`
	AccountID               DefinitionAccountID               `json:"accountID"`
	BatchID                 DefinitionTransactionID           `json:"batchID"`
	RequestID               DefinitionRequestID               `json:"requestID"`
	Type                    DefinitionTransactionType         `json:"type"`
	TradeID                 DefinitionTradeID                 `json:"tradeID"`
	ClientTradeID           DefinitionClientID                `json:"clientTradeID"`
	Price                   DefinitionPriceValue              `json:"price"`
	TimeInForce             DefinitionTimeInForce             `json:"timeInForce"`
	GtdTime                 DefinitionDateTime                `json:"gtdTime"`
	TriggerCondition        DefinitionOrderTriggerCondition   `json:"triggerCondition"`
	Reason                  DefinitionTakeProfitOrderReason   `json:"reason"`
	ClientExtensions        DefinitionClientExtensions        `json:"clientExtensions"`
	OrderFillTransactionID  DefinitionTransactionID           `json:"orderFillTransactionID"`
	IntendedReplacesOrderID DefinitionOrderID                 `json:"intendedReplacesOrderID"`
	RejectReason            DefinitionTransactionRejectReason `json:"rejectReason"`
}

type DefinitionStopLossOrderTransaction struct {
	Id                         DefinitionTransactionID         `json:"id"`
	Time                       DefinitionDateTime              `json:"time"`
	UserID                     int                             `json:"userID"`
	AccountID                  DefinitionAccountID             `json:"accountID"`
	BatchID                    DefinitionTransactionID         `json:"batchID"`
	RequestID                  DefinitionRequestID             `json:"requestID"`
	Type                       DefinitionTransactionType       `json:"type"`
	TradeID                    DefinitionTradeID               `json:"tradeID"`
	ClientTradeID              DefinitionClientID              `json:"clientTradeID"`
	Price                      DefinitionPriceValue            `json:"price"`
	Distance                   DefinitionDecimalNumber         `json:"distance"`
	TimeInForce                DefinitionTimeInForce           `json:"timeInForce"`
	GtdTime                    DefinitionDateTime              `json:"gtdTime"`
	TriggerCondition           DefinitionOrderTriggerCondition `json:"triggerCondition"`
	Guaranteed                 bool                            `json:"guaranteed"`
	GuaranteedExecutionPremium DefinitionDecimalNumber         `json:"guaranteedExecutionPremium"`
	Reason                     DefinitionStopLossOrderReason   `json:"reason"`
	ClientExtensions           DefinitionClientExtensions      `json:"clientExtensions"`
	OrderFillTransactionID     DefinitionTransactionID         `json:"orderFillTransactionID"`
	ReplacesOrderID            DefinitionOrderID               `json:"replacesOrderID"`
	CancellingTransactionID    DefinitionTransactionID         `json:"cancellingTransactionID"`
}

type DefinitionStopLossOrderRejectTransaction struct {
	Id                      DefinitionTransactionID           `json:"id"`
	Time                    DefinitionDateTime                `json:"time"`
	UserID                  int                               `json:"userID"`
	AccountID               DefinitionAccountID               `json:"accountID"`
	BatchID                 DefinitionTransactionID           `json:"batchID"`
	RequestID               DefinitionRequestID               `json:"requestID"`
	Type                    DefinitionTransactionType         `json:"type"`
	TradeID                 DefinitionTradeID                 `json:"tradeID"`
	ClientTradeID           DefinitionClientID                `json:"clientTradeID"`
	Price                   DefinitionPriceValue              `json:"price"`
	Distance                DefinitionDecimalNumber           `json:"distance"`
	TimeInForce             DefinitionTimeInForce             `json:"timeInForce"`
	GtdTime                 DefinitionDateTime                `json:"gtdTime"`
	TriggerCondition        DefinitionOrderTriggerCondition   `json:"triggerCondition"`
	Guaranteed              bool                              `json:"guaranteed"`
	Reason                  DefinitionStopLossOrderReason     `json:"reason"`
	ClientExtensions        DefinitionClientExtensions        `json:"clientExtensions"`
	OrderFillTransactionID  DefinitionTransactionID           `json:"orderFillTransactionID"`
	IntendedReplacesOrderID DefinitionOrderID                 `json:"intendedReplacesOrderID"`
	RejectReason            DefinitionTransactionRejectReason `json:"rejectReason"`
}

type DefinitionTrailingStopLossOrderTransaction struct {
	Id                      DefinitionTransactionID               `json:"id"`
	Time                    DefinitionDateTime                    `json:"time"`
	UserID                  int                                   `json:"userID"`
	AccountID               DefinitionAccountID                   `json:"accountID"`
	BatchID                 DefinitionTransactionID               `json:"batchID"`
	RequestID               DefinitionRequestID                   `json:"requestID"`
	Type                    DefinitionTransactionType             `json:"type"`
	TradeID                 DefinitionTradeID                     `json:"tradeID"`
	ClientTradeID           DefinitionClientID                    `json:"clientTradeID"`
	Distance                DefinitionDecimalNumber               `json:"distance"`
	TimeInForce             DefinitionTimeInForce                 `json:"timeInForce"`
	GtdTime                 DefinitionDateTime                    `json:"gtdTime"`
	TriggerCondition        DefinitionOrderTriggerCondition       `json:"triggerCondition"`
	Reason                  DefinitionTrailingStopLossOrderReason `json:"reason"`
	ClientExtensions        DefinitionClientExtensions            `json:"clientExtensions"`
	OrderFillTransactionID  DefinitionTransactionID               `json:"orderFillTransactionID"`
	ReplacesOrderID         DefinitionOrderID                     `json:"replacesOrderID"`
	CancellingTransactionID DefinitionTransactionID               `json:"cancellingTransactionID"`
}

type DefinitionTrailingStopLossOrderRejectTransaction struct {
	Id                      DefinitionTransactionID               `json:"id"`
	Time                    DefinitionDateTime                    `json:"time"`
	UserID                  int                                   `json:"userID"`
	AccountID               DefinitionAccountID                   `json:"accountID"`
	BatchID                 DefinitionTransactionID               `json:"batchID"`
	RequestID               DefinitionRequestID                   `json:"requestID"`
	Type                    DefinitionTransactionType             `json:"type"`
	TradeID                 DefinitionTradeID                     `json:"tradeID"`
	ClientTradeID           DefinitionClientID                    `json:"clientTradeID"`
	Distance                DefinitionDecimalNumber               `json:"distance"`
	TimeInForce             DefinitionTimeInForce                 `json:"timeInForce"`
	GtdTime                 DefinitionDateTime                    `json:"gtdTime"`
	TriggerCondition        DefinitionOrderTriggerCondition       `json:"triggerCondition"`
	Reason                  DefinitionTrailingStopLossOrderReason `json:"reason"`
	ClientExtensions        DefinitionClientExtensions            `json:"clientExtensions"`
	OrderFillTransactionID  DefinitionTransactionID               `json:"orderFillTransactionID"`
	IntendedReplacesOrderID DefinitionOrderID                     `json:"intendedReplacesOrderID"`
	RejectReason            DefinitionTransactionRejectReason     `json:"rejectReason"`
}

type DefinitionOrderFillTransaction struct {
	Id                            DefinitionTransactionID   `json:"id"`
	Time                          DefinitionDateTime        `json:"time"`
	UserID                        int                       `json:"userID"`
	AccountID                     DefinitionAccountID       `json:"accountID"`
	BatchID                       DefinitionTransactionID   `json:"batchID"`
	RequestID                     DefinitionRequestID       `json:"requestID"`
	Type                          DefinitionTransactionType `json:"type"`
	OrderID                       DefinitionOrderID         `json:"orderID"`
	ClientOrderID                 DefinitionClientID        `json:"clientOrderID"`
	Instrument                    DefinitionInstrumentName  `json:"instrument"`
	Units                         DefinitionDecimalNumber   `json:"units"`
	GainQuoteHomeConversionFactor DefinitionDecimalNumber   `json:"gainQuoteHomeConversionFactor"`
	LossQuoteHomeConversionFactor DefinitionDecimalNumber   `json:"lossQuoteHomeConversionFactor"`
	Price                         DefinitionPriceValue      `json:"price"`
	FullPrice                     DefinitionClientPrice     `json:"fullPrice"`
	Reason                        DefinitionOrderFillReason `json:"reason"`
	Pl                            DefinitionAccountUnits    `json:"pl"`
	Financing                     DefinitionAccountUnits    `json:"financing"`
	Commission                    DefinitionAccountUnits    `json:"commission"`
	GuaranteedExecutionFee        DefinitionAccountUnits    `json:"guaranteedExecutionFee"`
	AccountBalance                DefinitionAccountUnits    `json:"accountBalance"`
	TradeOpened                   DefinitionTradeOpen       `json:"tradeOpened"`
	TradesClosed                  []DefinitionTradeReduce   `json:"tradesClosed"`
	TradeReduced                  DefinitionTradeReduce     `json:"tradeReduced"`
	HalfSpreadCost                DefinitionAccountUnits    `json:"halfSpreadCost"`
}

type DefinitionOrderCancelTransaction struct {
	Id                DefinitionTransactionID     `json:"id"`
	Time              DefinitionDateTime          `json:"time"`
	UserID            int                         `json:"userID"`
	AccountID         DefinitionAccountID         `json:"accountID"`
	BatchID           DefinitionTransactionID     `json:"batchID"`
	RequestID         DefinitionRequestID         `json:"requestID"`
	Type              DefinitionTransactionType   `json:"type"`
	OrderID           DefinitionOrderID           `json:"orderID"`
	ClientOrderID     DefinitionOrderID           `json:"clientOrderID"`
	Reason            DefinitionOrderCancelReason `json:"reason"`
	ReplacedByOrderID DefinitionOrderID           `json:"replacedByOrderID"`
}

type DefinitionOrderCancelRejectTransaction struct {
	Id            DefinitionTransactionID           `json:"id"`
	Time          DefinitionDateTime                `json:"time"`
	UserID        int                               `json:"userID"`
	AccountID     DefinitionAccountID               `json:"accountID"`
	BatchID       DefinitionTransactionID           `json:"batchID"`
	RequestID     DefinitionRequestID               `json:"requestID"`
	Type          DefinitionTransactionType         `json:"type"`
	OrderID       DefinitionOrderID                 `json:"orderID"`
	ClientOrderID DefinitionOrderID                 `json:"clientOrderID"`
	RejectReason  DefinitionTransactionRejectReason `json:"rejectReason"`
}

type DefinitionOrderClientExtensionsModifyTransaction struct {
	Id                          DefinitionTransactionID    `json:"id"`
	Time                        DefinitionDateTime         `json:"time"`
	UserID                      int                        `json:"userID"`
	AccountID                   DefinitionAccountID        `json:"accountID"`
	BatchID                     DefinitionTransactionID    `json:"batchID"`
	RequestID                   DefinitionRequestID        `json:"requestID"`
	Type                        DefinitionTransactionType  `json:"type"`
	OrderID                     DefinitionOrderID          `json:"orderID"`
	ClientOrderID               DefinitionClientID         `json:"clientOrderID"`
	ClientExtensionsModify      DefinitionClientExtensions `json:"clientExtensionsModify"`
	TradeClientExtensionsModify DefinitionClientExtensions `json:"tradeClientExtensionsModify"`
}

type DefinitionOrderClientExtensionsModifyRejectTransaction struct {
	Id                          DefinitionTransactionID           `json:"id"`
	Time                        DefinitionDateTime                `json:"time"`
	UserID                      int                               `json:"userID"`
	AccountID                   DefinitionAccountID               `json:"accountID"`
	BatchID                     DefinitionTransactionID           `json:"batchID"`
	RequestID                   DefinitionRequestID               `json:"requestID"`
	Type                        DefinitionTransactionType         `json:"type"`
	OrderID                     DefinitionOrderID                 `json:"orderID"`
	ClientOrderID               DefinitionClientID                `json:"clientOrderID"`
	ClientExtensionsModify      DefinitionClientExtensions        `json:"clientExtensionsModify"`
	TradeClientExtensionsModify DefinitionClientExtensions        `json:"tradeClientExtensionsModify"`
	RejectReason                DefinitionTransactionRejectReason `json:"rejectReason"`
}

type DefinitionTradeClientExtensionsModifyTransaction struct {
	Id                          DefinitionTransactionID    `json:"id"`
	Time                        DefinitionDateTime         `json:"time"`
	UserID                      int                        `json:"userID"`
	AccountID                   DefinitionAccountID        `json:"accountID"`
	BatchID                     DefinitionTransactionID    `json:"batchID"`
	RequestID                   DefinitionRequestID        `json:"requestID"`
	Type                        DefinitionTransactionType  `json:"type"`
	TradeID                     DefinitionTradeID          `json:"tradeID"`
	ClientTradeID               DefinitionClientID         `json:"clientTradeID"`
	TradeClientExtensionsModify DefinitionClientExtensions `json:"tradeClientExtensionsModify"`
}

type DefinitionTradeClientExtensionsModifyRejectTransaction struct {
	Id                          DefinitionTransactionID           `json:"id"`
	Time                        DefinitionDateTime                `json:"time"`
	UserID                      int                               `json:"userID"`
	AccountID                   DefinitionAccountID               `json:"accountID"`
	BatchID                     DefinitionTransactionID           `json:"batchID"`
	RequestID                   DefinitionRequestID               `json:"requestID"`
	Type                        DefinitionTransactionType         `json:"type"`
	TradeID                     DefinitionTradeID                 `json:"tradeID"`
	ClientTradeID               DefinitionClientID                `json:"clientTradeID"`
	TradeClientExtensionsModify DefinitionClientExtensions        `json:"tradeClientExtensionsModify"`
	RejectReason                DefinitionTransactionRejectReason `json:"rejectReason"`
}

type DefinitionMarginCallEnterTransaction struct {
	Id        DefinitionTransactionID   `json:"id"`
	Time      DefinitionDateTime        `json:"time"`
	UserID    int                       `json:"userID"`
	AccountID DefinitionAccountID       `json:"accountID"`
	BatchID   DefinitionTransactionID   `json:"batchID"`
	RequestID DefinitionRequestID       `json:"requestID"`
	Type      DefinitionTransactionType `json:"type"`
}

type DefinitionMarginCallExtendTransaction struct {
	Id              DefinitionTransactionID   `json:"id"`
	Time            DefinitionDateTime        `json:"time"`
	UserID          int                       `json:"userID"`
	AccountID       DefinitionAccountID       `json:"accountID"`
	BatchID         DefinitionTransactionID   `json:"batchID"`
	RequestID       DefinitionRequestID       `json:"requestID"`
	Type            DefinitionTransactionType `json:"type"`
	ExtensionNumber int                       `json:"extensionNumber"`
}

type DefinitionMarginCallExitTransaction struct {
	Id        DefinitionTransactionID   `json:"id"`
	Time      DefinitionDateTime        `json:"time"`
	UserID    int                       `json:"userID"`
	AccountID DefinitionAccountID       `json:"accountID"`
	BatchID   DefinitionTransactionID   `json:"batchID"`
	RequestID DefinitionRequestID       `json:"requestID"`
	Type      DefinitionTransactionType `json:"type"`
}

type DefinitionDelayedTradeClosureTransaction struct {
	Id        DefinitionTransactionID     `json:"id"`
	Time      DefinitionDateTime          `json:"time"`
	UserID    int                         `json:"userID"`
	AccountID DefinitionAccountID         `json:"accountID"`
	BatchID   DefinitionTransactionID     `json:"batchID"`
	RequestID DefinitionRequestID         `json:"requestID"`
	Type      DefinitionTransactionType   `json:"type"`
	Reason    DefinitionMarketOrderReason `json:"reason"`
	TradeIDs  DefinitionTradeID           `json:"tradeIDs"`
}

type DefinitionDailyFinancingTransaction struct {
	Id                   DefinitionTransactionID        `json:"id"`
	Time                 DefinitionDateTime             `json:"time"`
	UserID               int                            `json:"userID"`
	AccountID            DefinitionAccountID            `json:"accountID"`
	BatchID              DefinitionTransactionID        `json:"batchID"`
	RequestID            DefinitionRequestID            `json:"requestID"`
	Type                 DefinitionTransactionType      `json:"type"`
	Financing            DefinitionAccountUnits         `json:"financing"`
	AccountBalance       DefinitionAccountUnits         `json:"accountBalance"`
	AccountFinancingMode DefinitionAccountFinancingMode `json:"accountFinancingMode"`
	PositionFinancings   []DefinitionPositionFinancing  `json:"positionFinancings"`
}

type DefinitionResetResettablePLTransaction struct {
	Id        DefinitionTransactionID   `json:"id"`
	Time      DefinitionDateTime        `json:"time"`
	UserID    int                       `json:"userID"`
	AccountID DefinitionAccountID       `json:"accountID"`
	BatchID   DefinitionTransactionID   `json:"batchID"`
	RequestID DefinitionRequestID       `json:"requestID"`
	Type      DefinitionTransactionType `json:"type"`
}

// Transaction-related Definitions

type DefinitionTransactionID string

type DefinitionTransactionType string

type DefinitionFundingReason string

type DefinitionMarketOrderReason string

type DefinitionFixedPriceOrderReason string

type DefinitionLimitOrderReason string

type DefinitionStopOrderReason string

type DefinitionMarketIfTouchedOrderReason string

type DefinitionTakeProfitOrderReason string

type DefinitionStopLossOrderReason string

type DefinitionTrailingStopLossOrderReason string

type DefinitionOrderFillReason string

type DefinitionOrderCancelReason string

type DefinitionClientID string

type DefinitionClientTag string

type DefinitionClientComment string

type DefinitionClientExtensions struct {
	Id      DefinitionClientID      `json:"id"`
	Tag     DefinitionClientTag     `json:"tag"`
	Comment DefinitionClientComment `json:"comment"`
}

type DefinitionTakeProfitDetails struct {
	Price            DefinitionPriceValue       `json:"price"`
	TimeInForce      DefinitionTimeInForce      `json:"timeInForce"`
	GtdTime          DefinitionDateTime         `json:"gtdTime"`
	ClientExtensions DefinitionClientExtensions `json:"clientExtensions"`
}

type DefinitionStopLossDetails struct {
	Price            DefinitionPriceValue       `json:"price"`
	Distance         DefinitionDecimalNumber    `json:"distance"`
	TimeInForce      DefinitionTimeInForce      `json:"timeInForce"`
	GtdTime          DefinitionDateTime         `json:"gtdTime"`
	ClientExtensions DefinitionClientExtensions `json:"clientExtensions"`
	Guaranteed       bool                       `json:"guaranteed"`
}

type DefinitionTrailingStopLossDetails struct {
	Distance         DefinitionDecimalNumber    `json:"distance"`
	TimeInForce      DefinitionTimeInForce      `json:"timeInForce"`
	GtdTime          DefinitionDateTime         `json:"gtdTime"`
	ClientExtensions DefinitionClientExtensions `json:"clientExtensions"`
}

type DefinitionTradeOpen struct {
	TradeID                DefinitionTradeID          `json:"tradeID"`
	Units                  DefinitionDecimalNumber    `json:"units"`
	Price                  DefinitionPriceValue       `json:"price"`
	GuaranteedExecutionFee DefinitionAccountUnits     `json:"guaranteedExecutionFee"`
	ClientExtensions       DefinitionClientExtensions `json:"clientExtensions"`
	HalfSpreadCost         DefinitionAccountUnits     `json:"halfSpreadCost"`
	InitialMarginRequired  DefinitionAccountUnits     `json:"initialMarginRequired"`
}

type DefinitionTradeReduce struct {
	TradeID                DefinitionTradeID       `json:"tradeID"`
	Units                  DefinitionDecimalNumber `json:"units"`
	Price                  DefinitionPriceValue    `json:"price"`
	RealizedPL             DefinitionAccountUnits  `json:"realizedPL"`
	Financing              DefinitionAccountUnits  `json:"financing"`
	GuaranteedExecutionFee DefinitionAccountUnits  `json:"guaranteedExecutionFee"`
	HalfSpreadCost         DefinitionAccountUnits  `json:"halfSpreadCost"`
}

type DefinitionMarketOrderTradeClose struct {
	TradeID       DefinitionTradeID `json:"tradeID"`
	ClientTradeID string            `json:"clientTradeID"`
	Units         string            `json:"units"`
}

type DefinitionMarketOrderMarginCloseout struct {
	reason DefinitionMarketOrderMarginCloseoutReason `json:"reason"`
}

type DefinitionMarketOrderMarginCloseoutReason string

type DefinitionMarketOrderDelayedTradeClose struct {
	TradeID             DefinitionTradeID       `json:"tradeID"`
	ClientTradeID       DefinitionTradeID       `json:"clientTradeID"`
	SourceTransactionID DefinitionTransactionID `json:"sourceTransactionID"`
}

type DefinitionMarketOrderPositionCloseout struct {
	Instrument DefinitionInstrumentName `json:"instrument"`
	Units      string                   `json:"units"`
}

type DefinitionLiquidityRegenerationSchedule struct {
	Steps []DefinitionLiquidityRegenerationScheduleStep `json:"steps"`
}

type DefinitionLiquidityRegenerationScheduleStep struct {
	Timestamp        DefinitionDateTime      `json:"timestamp"`
	BidLiquidityUsed DefinitionDecimalNumber `json:"bidLiquidityUsed"`
	AskLiquidityUsed DefinitionDecimalNumber `json:"askLiquidityUsed"`
}

type DefinitionOpenTradeFinancing struct {
	TradeID   DefinitionTradeID      `json:"tradeID"`
	Financing DefinitionAccountUnits `json:"financing"`
}

type DefinitionPositionFinancing struct {
	Instrument          DefinitionInstrumentName       `json:"instrument"`
	Financing           DefinitionAccountUnits         `json:"financing"`
	OpenTradeFinancings []DefinitionOpenTradeFinancing `json:"openTradeFinancings"`
}

type DefinitionRequestID string

type DefinitionTransactionRejectReason string

type DefinitionTransactionFilter string

type DefinitionTransactionHeartbeat struct {
	Type              string                  `json:"type"`
	LastTransactionID DefinitionTransactionID `json:"lastTransactionID"`
	Time              DefinitionDateTime      `json:"time"`
}

//
// Pricing Definitions
//

type DefinitionPrice struct {
	Type                       string                               `json:"type"`
	Instrument                 DefinitionInstrumentName             `json:"instrument"`
	Time                       DefinitionDateTime                   `json:"time"`
	Status                     DefinitionPriceStatus                `json:"status"`
	Tradeable                  bool                                 `json:"tradeable"`
	Bids                       []DefinitionPriceBucket              `json:"bids"`
	Asks                       []DefinitionPriceBucket              `json:"asks"`
	CloseoutBid                DefinitionPriceValue                 `json:"closeoutBid"`
	CloseoutAsk                DefinitionPriceValue                 `json:"closeoutAsk"`
	QuoteHomeConversionFactors DefinitionQuoteHomeConversionFactors `json:"quoteHomeConversionFactors"`
	UnitsAvailable             DefinitionUnitsAvailable             `json:"unitsAvailable"`
}

type DefinitionPriceValue string

type DefinitionPriceBucket struct {
	Price     DefinitionPriceValue `json:"price"`
	Liquidity int                  `json:"liquidity"`
}

type DefinitionPriceStatus string

type DefinitionQuoteHomeConversionFactors struct {
	PositiveUnits DefinitionDecimalNumber `json:"positiveUnits"`
	NegativeUnits DefinitionDecimalNumber `json:"negativeUnits"`
}

type DefinitionHomeConversions struct {
	Currency      DefinitionCurrency      `json:"currency"`
	AccountGain   DefinitionDecimalNumber `json:"accountGain"`
	AccountLoss   DefinitionDecimalNumber `json:"accountLoss"`
	PositionValue DefinitionDecimalNumber `json:"positionValue"`
}

type DefinitionClientPrice struct {
	Bids        []DefinitionPriceBucket `json:"bids"`
	Asks        []DefinitionPriceBucket `json:"asks"`
	CloseoutBid DefinitionPriceValue    `json:"closeoutBid"`
	CloseoutAsk DefinitionPriceValue    `json:"closeoutAsk"`
	Timestamp   DefinitionDateTime      `json:"timestamp"`
}

type DefinitionPricingHeartbeat struct {
	Type string             `json:"type"`
	Time DefinitionDateTime `json:"time"`
}

//
// Primitives Definitions
//

type DefinitionDecimalNumber string

type DefinitionAccountUnits string

type DefinitionCurrency string

type DefinitionInstrumentName string

type DefinitionInstrumentType string

type DefinitionInstrument struct {
	Name                        DefinitionInstrumentName       `json:"name"`
	Type                        DefinitionInstrumentType       `json:"type"`
	DisplayName                 string                         `json:"displayName"`
	PipLocation                 int                            `json:"pipLocation"`
	DisplayPrecision            int                            `json:"displayPrecision"`
	TradeUnitsPrecision         int                            `json:"tradeUnitsPrecision"`
	MinimumTradeSize            DefinitionDecimalNumber        `json:"minimumTradeSize"`
	MaximumTrailingStopDistance DefinitionDecimalNumber        `json:"maximumTrailingStopDistance"`
	MinimumTrailingStopDistance DefinitionDecimalNumber        `json:"minimumTrailingStopDistance"`
	MaximumPositionSize         DefinitionDecimalNumber        `json:"maximumPositionSize"`
	MaximumOrderUnits           DefinitionDecimalNumber        `json:"maximumOrderUnits"`
	MarginRate                  DefinitionDecimalNumber        `json:"marginRate"`
	Commission                  DefinitionInstrumentCommission `json:"commission"`
}

type DefinitionDateTime string

type DefinitionAcceptDatetimeFormat string

type DefinitionInstrumentCommission struct {
	Commission        DefinitionDecimalNumber `json:"commission"`
	UnitsTraded       DefinitionDecimalNumber `json:"unitsTraded"`
	MinimumCommission DefinitionDecimalNumber `json:"minimumCommission"`
}

type DefinitionGuaranteedStopLossOrderLevelRestriction struct {
	Volume     DefinitionDecimalNumber `json:"volume"`
	PriceRange DefinitionDecimalNumber `json:"priceRange"`
}

type DefinitionDirection string
