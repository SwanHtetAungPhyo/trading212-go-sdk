package trading212

import "time"

// AccountInfo represents account metadata information
type AccountInfo struct {
	Currency string `json:"currencyCode"`
	ID       int64  `json:"id"`
}

// AccountCash represents account cash balance information
type AccountCash struct {
	Free     float64 `json:"free"`
	Invested float64 `json:"invested"`
	Result   float64 `json:"result"`
	Total    float64 `json:"total"`
}

// AccountSummary represents combined account information (custom struct for convenience)
type AccountSummary struct {
	AccountInfo
	Cash AccountCash `json:"cash"`
}

// Order represents an order
type Order struct {
	CreatedAt      time.Time     `json:"createdAt"`
	Currency       string        `json:"currency"`
	ExtendedHours  bool          `json:"extendedHours"`
	FilledQuantity float64       `json:"filledQuantity"`
	FilledValue    float64       `json:"filledValue"`
	ID             int64         `json:"id"`
	InitiatedFrom  string        `json:"initiatedFrom"`
	Instrument     Instrument    `json:"instrument"`
	LimitPrice     *float64      `json:"limitPrice,omitempty"`
	Quantity       float64       `json:"quantity"`
	Side           OrderSide     `json:"side"`
	Status         OrderStatus   `json:"status"`
	StopPrice      *float64      `json:"stopPrice,omitempty"`
	Strategy       OrderStrategy `json:"strategy"`
	Ticker         string        `json:"ticker"`
	TimeInForce    TimeValidity  `json:"timeInForce"`
	Type           OrderType     `json:"type"`
	Value          float64       `json:"value"`
}

// OrderSide represents order side
type OrderSide string

const (
	OrderSideBuy  OrderSide = "BUY"
	OrderSideSell OrderSide = "SELL"
)

// OrderStatus represents order status
type OrderStatus string

const (
	OrderStatusLocal          OrderStatus = "LOCAL"
	OrderStatusUnconfirmed    OrderStatus = "UNCONFIRMED"
	OrderStatusConfirmed      OrderStatus = "CONFIRMED"
	OrderStatusNew            OrderStatus = "NEW"
	OrderStatusCancelling     OrderStatus = "CANCELLING"
	OrderStatusCancelled      OrderStatus = "CANCELLED"
	OrderStatusPartiallyFilled OrderStatus = "PARTIALLY_FILLED"
	OrderStatusFilled         OrderStatus = "FILLED"
	OrderStatusRejected       OrderStatus = "REJECTED"
	OrderStatusReplacing      OrderStatus = "REPLACING"
	OrderStatusReplaced       OrderStatus = "REPLACED"
)

// OrderStrategy represents order strategy
type OrderStrategy string

const (
	OrderStrategyQuantity OrderStrategy = "QUANTITY"
	OrderStrategyValue    OrderStrategy = "VALUE"
)

// OrderType represents order type
type OrderType string

const (
	OrderTypeLimit     OrderType = "LIMIT"
	OrderTypeStop      OrderType = "STOP"
	OrderTypeMarket    OrderType = "MARKET"
	OrderTypeStopLimit OrderType = "STOP_LIMIT"
)

// TimeValidity represents time validity
type TimeValidity string

const (
	TimeValidityDay             TimeValidity = "DAY"
	TimeValidityGoodTillCancel  TimeValidity = "GOOD_TILL_CANCEL"
)

// Instrument represents instrument information
type Instrument struct {
	Currency string `json:"currency"`
	ISIN     string `json:"isin"`
	Name     string `json:"name"`
	Ticker   string `json:"ticker"`
}

// Position represents a position
type Position struct {
	AveragePrice     float64   `json:"averagePrice"`
	CurrentPrice     float64   `json:"currentPrice"`
	Frontend         string    `json:"frontend"`
	FxPpl            float64   `json:"fxPpl"`
	InitialFillDate  time.Time `json:"initialFillDate"`
	MaxBuy           float64   `json:"maxBuy"`
	MaxSell          float64   `json:"maxSell"`
	PieQuantity      float64   `json:"pieQuantity"`
	Ppl              float64   `json:"ppl"`
	Quantity         float64   `json:"quantity"`
	Ticker           string    `json:"ticker"`
}

// PositionWalletImpact represents position wallet impact
type PositionWalletImpact struct {
	Currency             string  `json:"currency"`
	CurrentValue         float64 `json:"currentValue"`
	FxImpact             float64 `json:"fxImpact"`
	TotalCost            float64 `json:"totalCost"`
	UnrealizedProfitLoss float64 `json:"unrealizedProfitLoss"`
}

// TradableInstrument represents a tradable instrument
type TradableInstrument struct {
	AddedOn            time.Time      `json:"addedOn"`
	CurrencyCode       string         `json:"currencyCode"`
	ExtendedHours      bool           `json:"extendedHours"`
	ISIN               string         `json:"isin"`
	MaxOpenQuantity    float64        `json:"maxOpenQuantity"`
	Name               string         `json:"name"`
	ShortName          string         `json:"shortName"`
	Ticker             string         `json:"ticker"`
	Type               InstrumentType `json:"type"`
	WorkingScheduleID  int64          `json:"workingScheduleId"`
}

// InstrumentType represents instrument type
type InstrumentType string

const (
	InstrumentTypeCryptocurrency InstrumentType = "CRYPTOCURRENCY"
	InstrumentTypeETF            InstrumentType = "ETF"
	InstrumentTypeForex          InstrumentType = "FOREX"
	InstrumentTypeFutures        InstrumentType = "FUTURES"
	InstrumentTypeIndex          InstrumentType = "INDEX"
	InstrumentTypeStock          InstrumentType = "STOCK"
	InstrumentTypeWarrant        InstrumentType = "WARRANT"
	InstrumentTypeCrypto         InstrumentType = "CRYPTO"
	InstrumentTypeCVR            InstrumentType = "CVR"
	InstrumentTypeCorpact        InstrumentType = "CORPACT"
)

// Exchange represents an exchange
type Exchange struct {
	ID               int64             `json:"id"`
	Name             string            `json:"name"`
	WorkingSchedules []WorkingSchedule `json:"workingSchedules"`
}

// WorkingSchedule represents working schedule
type WorkingSchedule struct {
	ID         int64       `json:"id"`
	TimeEvents []TimeEvent `json:"timeEvents"`
}

// TimeEvent represents a time event
type TimeEvent struct {
	Date time.Time     `json:"date"`
	Type TimeEventType `json:"type"`
}

// TimeEventType represents time event type
type TimeEventType string

const (
	TimeEventTypeOpen             TimeEventType = "OPEN"
	TimeEventTypeClose            TimeEventType = "CLOSE"
	TimeEventTypeBreakStart       TimeEventType = "BREAK_START"
	TimeEventTypeBreakEnd         TimeEventType = "BREAK_END"
	TimeEventTypePreMarketOpen    TimeEventType = "PRE_MARKET_OPEN"
	TimeEventTypeAfterHoursOpen   TimeEventType = "AFTER_HOURS_OPEN"
	TimeEventTypeAfterHoursClose  TimeEventType = "AFTER_HOURS_CLOSE"
	TimeEventTypeOvernightOpen    TimeEventType = "OVERNIGHT_OPEN"
)