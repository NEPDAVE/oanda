package oanda

import (
	"encoding/json"
	"log"
	"time"
)

/*
***************************
errors
***************************
*/

//ErrorCode captures an Oanda error code returned by their API
type ErrorCode struct {
	Code int `json:"code"`
}

//UnmarshalErrorCode used by StreamPricing
func (e ErrorCode) UnmarshalErrorCode(errorByte []byte) *ErrorCode {

	err := json.Unmarshal(errorByte, &e)

	if err != nil {
		log.Println(err)
	}


	return &e
}

/*
***************************
prices
***************************
*/

//Heartbeat is returned from the Oanda streaming endpoint
type Heartbeat struct {
	Time time.Time `json:"time"`
	Type string    `json:"type"`
}

//UnmarshalHeartbeat is a method of Heartbeat
func (h Heartbeat) UnmarshalHeartbeat(heartbeatByte []byte) *Heartbeat {

	err := json.Unmarshal(heartbeatByte, &h)

	if err != nil {
		log.Println(ErrorCode{}.UnmarshalErrorCode(heartbeatByte))
	}

	return &h
}

//Pricing is returned from the Oanda pricing endpoint
type Pricing struct {
	Prices []Prices  `json:"prices"`
	Time   time.Time `json:"time"`
}

//Prices is embedded within each Pricing struct and  is returned object from
//the Oanda streaming endpoint
type Prices struct {
	Type        string    `json:"type"`
	Bids        []Bid     `json:"bids"`
	Asks        []Ask     `jons:"asks"`
	CloseOutAsk string    `json:"closeoutAsk"`
	CloseOutBid string    `json:"closeoutBid"`
	Instrument  string    `json:"instrument"`
	Status      string    `json:"status"`
	Time        time.Time `json:"time"`
}

//Ask represents one element in the Asks list of a Prices Struct
type Ask struct {
	Price     string `json:"price"`
	Liquidity int64  `json:"liquidity"`
}

//Bid represents one element in the Bids list of a Prices Struct
type Bid struct {
	Price     string `json:"price"`
	Liquidity int64  `json:"liquidity"`
}

//UnmarshalPrices used by StreamPricing
func (p Prices) UnmarshalPrices(priceByte []byte) *Prices {

	err := json.Unmarshal(priceByte, &p)

	if err != nil {
		log.Println(ErrorCode{}.UnmarshalErrorCode(priceByte))
	}

	return &p
}

//UnmarshalPricing unmarshals the Pricing data byte slice from Oanda
func (p Pricing) UnmarshalPricing(priceByte []byte) *Pricing {

	err := json.Unmarshal(priceByte, &p)

	if err != nil {
		log.Println(ErrorCode{}.UnmarshalErrorCode(priceByte))
	}

	return &p
}

/*
***************************
history
***************************
*/

//Candles represents the data structure returned by Oanda when requesting
//multiple Candles
type Candles struct {
	Instrument  string   `json:"instrument"`
	Granularity string   `json:"granularity"`
	Candles     []Candle `json:"candles"`
}

//Candle represents a single data point in an instrument's pricing history
type Candle struct {
	Complete bool      `json:"complete"`
	Volume   int64     `json:"volume"`
	Time     time.Time `json:"time"`
	Mid      Mid       `json:"mid"`
}

//Mid represents the actual quotes/prices in a Candle
type Mid struct {
	Open  string `json:"o"`
	High  string `json:"h"`
	Low   string `json:"l"`
	Close string `json:"c"`
}

//UnmarshalCandles unmarshals History data byte slice from Oanda
func (c Candles) UnmarshalCandles(candlesByte []byte) *Candles {

	err := json.Unmarshal(candlesByte, &c)

	if err != nil {
		log.Println(ErrorCode{}.UnmarshalErrorCode(candlesByte))
	}

	return &c
}

/*
***************************
orders
***************************
*/

//OrderCreateTransaction represents the data structure returned by oanda after
//submiting an order
type OrderCreateTransaction struct {
	OrderCreateTransaction OrderCreateTransactionData `json:"orderCreateTransaction"`
	OrderFillTransaction   OrderFillTransactionData   `json:"orderFillTransaction"`
}

//OrderCreateTransactionData represents a data structure embedded in
//OrderCreateTransaction
type OrderCreateTransactionData struct {
	Type             string           `json:"type"`
	Instrument       string           `json:"instrument"`
	Units            string           `json:"units"`
	TimeInForce      string           `json:"timeInForce"`
	PositionFill     string           `json:"positionFill"`
	TakeProfitOnFill TakeProfitOnFill `json:"takeProfitOnFill"` //see orders.go
	StopLossOnFill   StopLossOnFill   `json:"stopLossOnFill"`   //see orders.go
	Reason           string           `json:"reason"`
	ID               string           `json:"id"`
	UserID           int              `json:"userID"`
	AccountID        string           `json:"accountID"`
	BatchID          string           `json:"batchID"`
	RequestID        string           `json:"requestID"`
	Time             time.Time        `json:"time"`
}

//OrderFillTransactionData represents a data structure embedded in
//OrderCreateTransaction
type OrderFillTransactionData struct {
	Type                          string          `json:"type"`
	OrderID                       string          `json:"orderID"`
	Instrument                    string          `json:"instrument"`
	Units                         string          `json:"units"`
	Price                         string          `json:"price"`
	PL                            string          `json:"pl"`
	Financing                     string          `json:"financing"`
	Commission                    string          `json:"commission"`
	AccountBalance                string          `json:"accountBalance"`
	GainQuoteHomeConversionFactor string          `json:"gainQuoteHomeConversionFactor"`
	LossQuoteHomeConversionFactor string          `json:"lossQuoteHomeConversionFactor"`
	HalfSpreadCost                string          `json:"halfSpreadCost"`
	Reason                        string          `json:"reason"`
	TradeOpened                   TradeOpenedData `json:"tradeOpened"`
	FullPrice                     FullPrice       `json:"fullPrice"`
	RelatedTransactionIDs         []string        `json:"relatedTransactionIDs"`
	LastTransactionID             string          `json:"lastTransactionID"`
}

//TradeOpenedData represents the data structure embedded in OrderFillTransactionData
type TradeOpenedData struct {
	Price                  string `json:"price"`
	TradeID                string `json:"tradeID"`
	Units                  string `json:"units"`
	GuaranteedExecutionFee string `json:"guaranteedExecutionFee"`
	HalfSpreadCost         string `json:"halfSpreadCost"`
	InitialMarginRequired  string `json:"initialMarginRequired"`
	LastTransactionID      string `json:"lastTransactionID"`
}

//FullPrice represents the data structure embedded in OrderFillTransactionData
type FullPrice struct {
	CloseoutBid string         `json:"closeoutBid"`
	CloseoutAsk string         `json:"closeoutAsk"`
	Time        time.Time      `json:"timestamp"`
	Bids        []FullPriceBid `json:"bids"`
	Asks        []FullPriceAsk `json:"asks"`
	ID          string         `json:"id"`
	UserID      string         `json:"userID"`
	AccountID   string         `json:"accountID"`
	BatchID     string         `json:"batchID"`
}

//FullPriceBid represents one element in the Bids list of a Prices Struct
//this differs from Bid which has an int for Liquidity
type FullPriceBid struct {
	Price     string `json:"price"`
	Liquidity string `json:"liquidity"`
}

//FullPriceAsk represents one element in the Asks list of a Prices Struct
//this differs from Ask which has an int for Liquidity
type FullPriceAsk struct {
	Price     string `json:"price"`
	Liquidity string `json:"liquidity"`
}

//UnmarshalOrderCreateTransaction unmarshals the returned data byte slice from Oanda
//that contains the order data
func (o OrderCreateTransaction) UnmarshalOrderCreateTransaction(
	ordersResponseByte []byte) *OrderCreateTransaction {

	err := json.Unmarshal(ordersResponseByte, &o)

	if err != nil {
		log.Println(ErrorCode{}.UnmarshalErrorCode(ordersResponseByte))
	}

	return &o
}

//Order represents the data structure returned by Oanda after calling
//fxtech.GetOrder()
type Order struct {
	OrderData         OrderData `json:"order"`
	LastTransactionID string    `json:"lastTransactionID"`
}

//OrderData represents the data associated with an order
type OrderData struct {
	ID               string           `json:"id"`
	CreateTime       time.Time        `json:"createTime"`
	Type             string           `json:"type"`
	Instrument       string           `json:"instrument"`
	Units            string           `json:"units"`
	TimeInForce      string           `json:"timeInForce"`
	StopLossOnFill   StopLossOnFill   `json:"stopLossOnFill"`
	TakeProfitOnFill TakeProfitOnFill `json:"takeProfitOnFill"`
	Price            string           `json:"price"`
	TriggerCondition string           `json:"triggerCondition"`
	PartialFill      string           `json:"partialFill"`
	PositionFill     string           `json:"positionFill"`
	State            string           `json:"state"`
}

//UnmarshalOrder unmarshals the returned data byte slice from Oanda
//after calling the fxtech.GetOrderStatus func
func (o Order) UnmarshalOrder(
	getOrderByte []byte) *Order {

	err := json.Unmarshal(getOrderByte, &o)

	if err != nil {
		log.Println(ErrorCode{}.UnmarshalErrorCode(getOrderByte))
	}

	return &o
}

type OrderCancelTransaction struct {
	OrderCancelTransactionData OrderCancelTransactionData `json:"orderCancelTransaction"`
	RelatedTransactionIDs      []string                   `json:"relatedTransactionIDs"`
	LastTransactionID          string                     `json:"lastTransactionID"`
}

type OrderCancelTransactionData struct {
	Type      string    `json:"type"`
	OrderID   string    `json:"orderID"`
	Reason    string    `json:"reason"`
	ID        string    `json:"id"`
	AccountID string    `json:"accountID"`
	UserID    int       `json:"userID"`
	BatchID   string    `json:"batchID"`
	RequestID string    `json:"requestID"`
	Time      time.Time `json:"time"`
}

//UnmarshalOrderCancelTransaction unmarshals the returned data byte slice from
//Oanada after calling the fxtech.CancelOrder func
func (o OrderCancelTransaction) UnmarshalOrderCancelTransaction(
	cancelOrderByte []byte) *OrderCancelTransaction {

	err := json.Unmarshal(cancelOrderByte, &o)

	if err != nil {
		log.Println(ErrorCode{}.UnmarshalErrorCode(cancelOrderByte))
	}

	return &o
}

/*
***************************
positions
***************************
*/

type Position struct {
	PositionDetails   PositionDetails `json:"position"`
	LastTransactionID string          `json:"lastTransactionID"`
}

type PositionDetails struct {
	Instrument              string              `json:"instrument"`
	Long                    PositionSideDetails `json:"long"`
	Short                   PositionSideDetails `json:"short"`
	PL                      string              `json:"pl"`
	ResettablePL            string              `json:"resettablePL"`
	Financing               string              `json:"financing"`
	Commission              string              `json:"commission"`
	GuaranteedExecutionFees string              `json:"guaranteedExecutionFees"`
	UnrealizedPL            string              `json:"unrealizedPL"`
	MarginUsed              string              `json:"marginUsed"`
}

type PositionSideDetails struct {
	AveragePrice            string   `json:"averagePrice"`
	Units                   string   `json:"units"`
	PL                      string   `json:"pl"`
	ResettablePL            string   `json:"resettablePL"`
	Financing               string   `json:"financing"`
	GuaranteedExecutionFees string   `json:"guaranteedExecutionFees"`
	UnrealizedPL            string   `json:"unrealizedPL"`
	TradeIDs                []string `json:"tradeIDs"`
}

//UnmarshalPosition unmarshals Position data byte slice from Oanda
func (p Position) UnmarshalPosition(positionByte []byte) *Position {

	err := json.Unmarshal(positionByte, &p)

	if err != nil {
		log.Println(ErrorCode{}.UnmarshalErrorCode(positionByte))
	}

	return &p
}

/*
***************************
account
***************************
*/

type AccountSummary struct {
	AccountSummaryDetails AccountSummaryDetails `json:"account"`
	LastTransactionID     string                `json:"lastTransactionID"`
}

type AccountSummaryDetails struct {
	GuaranteedStopLossOrderMode string    `json:"guaranteedStopLossOrderMode"`
	ID                          string    `json:"id"`
	CreatedTime                 string    `json:"createdTime"`
	Currency                    string    `json:"currency"`
	CreateByUserID              int       `json:"createdByUserID"`
	Alias                       string    `json:"Primary"`
	MarginRate                  string    `json:"marginRate"`
	HedgingEnabled              string    `json:"hedgingEnabled"`
	LastTransactionID           string    `json:"lastTransactionID"`
	Balance                     string    `json:"balance"`
	OpenTradeCount              int       `json:"openTradeCount"`
	OpenPositionCount           int       `json:"openPositionCount"`
	PendingOrderCount           int       `json:"pendingOrderCount"`
	PL                          string    `json:"pl"`
	ResettablePL                string    `json:"resettablePL"`
	ResettablePLTime            time.Time `json:"resettablePLTime"`
	Financing                   string    `json:"financing"`
	Commission                  string    `json:"commission"`
	GuaranteedExecutionFees     string    `json:"guaranteedExecutionFees"`
	UnrealizedPL                string    `json:"unrealizedPL"`
	NAV                         string    `json:"NAV"` //The net asset value of the Account. Equal to Account balance + unrealizedPL
	MarginUsed                  string    `json:"marginUsed"`
	MarginAvailable             string    `json:"marginAvailable"`
	PositionValue               string    `json:"positionValue"`
	MarginCloseoutUnrealizedPL  string    `json:"marginCloseoutUnrealizedPL"`
	MarginCloseoutNAV           string    `json:"marginCloseoutNAV"`
	MarginCloseoutMarginUsed    string    `json:"marginCloseoutMarginUsed"`
	MarginCloseoutPositionValue string    `json:"marginCloseoutPositionValue"`
	MarginCloseoutPercent       string    `json:"marginCloseoutPercent"`
	WithdrawalLimit             string    `json:"withdrawalLimit"`
	MarginCallMarginUsed        string    `json:"MarginCallMarginUsed"`
	MarginCallPercent           string    `json:"marginCallPercent"`
}

//UnmarshalAccount unmarshals Account data byte slice from Oanda
func (a AccountSummary) UnmarshalAccountSummary(accountByte []byte) *AccountSummary {

	err := json.Unmarshal(accountByte, &a)

	if err != nil {
		log.Println(ErrorCode{}.UnmarshalErrorCode(accountByte))
	}

	return &a
}
