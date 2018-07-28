package oanda

import (
	"encoding/json"
	"time"
)

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
		panic(err)
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
		panic(err)
	}

	return &p
}

//UnmarshalPricing unmarshals the Pricing data byte slice from Oanda
func (p Pricing) UnmarshalPricing(priceByte []byte) *Pricing {

	err := json.Unmarshal(priceByte, &p)

	if err != nil {
		panic(err)
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
func (c Candles) UnmarshalCandles(priceByte []byte) *Candles {

	err := json.Unmarshal(priceByte, &c)

	if err != nil {
		panic(err)
	}

	return &c
}

/*
***************************
orders
***************************
*/

type OrderSubmission struct {
	LastTransactionID      string `json:"lastTransactionID"`
	OrderCreateTransaction OrderCreateTransaction
}

type OrderCreateTransaction struct {
	AccountID             string         `json:"accountID"`
	BatchID               string         `json:"batchID"`
	ID                    string         `json:"id"`
	Instrument            string         `json:"instrument"`
	PositionFill          string         `json:"positionFill"`
	Price                 string         `json:"price"`
	Reason                string         `json:"reason"`
	StopLossOnFill        StopLossOnFill `json:"stopLossOnFill"`
	Time                  time.Time      `json:"time"`
	TimeInForce           string         `json:"timeInForce"`
	TriggerCondition      string         `json:"triggerCondition"`
	Type                  string         `json:"type"`
	Units                 string         `json:"units"`
	UserID                string         `json:"userID"`
	RelatedTransactionIDs []string       `json:"relatedTransactionIDs"`
}

//UnmarshalOrderSubmission unmarshals the returned data byte slice from Oanda
//that contains the order data
func (o OrderSubmission) UnmarshalOrderSubmission(ordersResponseByte []byte) *OrderSubmission {

	err := json.Unmarshal(ordersResponseByte, &o)

	if err != nil {
		panic(err)
	}

	return &o
}

/*

FIXME this is an example response for submitting an order
{
  "lastTransactionID": "6372",
  "orderCreateTransaction": {
    "accountID": "<ACCOUNT>",
    "batchID": "6372",
    "id": "6372",
    "instrument": "USD_CAD",
    "positionFill": "DEFAULT",
    "price": "1.50000",
    "reason": "CLIENT_ORDER",
    "stopLossOnFill": {
      "price": "1.70000",
      "timeInForce": "GTC"
    },
    "time": "2016-06-22T18:41:29.285982286Z",
    "timeInForce": "GTC",
    "triggerCondition": "TRIGGER_DEFAULT",
    "type": "LIMIT_ORDER",
    "units": "-1000",
    "userID": <USERID>
  },
  "relatedTransactionIDs": [
    "6372"
  ]
}


{"orderCreateTransaction":{
	"type":"MARKET_ORDER",
	"instrument":"EUR_USD",
	"units":"10",
	"timeInForce":"FOK",
	"positionFill":"DEFAULT",
	"takeProfitOnFill":{
		"price":"1.18372",
		"timeInForce":"GTC"
		},
	"stopLossOnFill":{
		"price":"1.16026",
		"timeInForce":"GTC"
		},
	"reason":"CLIENT_ORDER",
	"id":"7200",
	"userID":6395930,
	"accountID":"101-001-6395930-001",
	"batchID":"7200",
	"requestID":"78486291686830674",
	"time":"2018-07-27T19:49:18.256694465Z"
	},
	"orderFillTransaction":{
		"type":"ORDER_FILL",
		"orderID":"7200",
		"instrument":"EUR_USD",
		"units":"10",
		"price":"1.16623",
		"pl":"0.0000",
		"financing":"0.0000",
		"commission":"0.0000",
		"accountBalance":"100403.3955",
		"gainQuoteHomeConversionFactor":"1",
		"lossQuoteHomeConversionFactor":"1",
		"guaranteedExecutionFee":"0.0000",
		"halfSpreadCost":"0.0007",
		"reason":"MARKET_ORDER",
		"tradeOpened":{
			"price":"1.16623",
			"tradeID":"7201",
			"units":"10",
			"guaranteedExecutionFee":"0.0000",
			"halfSpreadCost":"0.0007",
			"initialMarginRequired":"0.2332"
			},
		"fullPrice":{
			"closeoutBid":"1.16594",
			"closeoutAsk":"1.16638",
			"timestamp":"2018-07-27T19:49:17.090610310Z",
			"bids":[{
				"price":"1.16609",
				"liquidity":"10000000"
				}],
			"asks":[{
				"price":"1.16623",
				"liquidity":"9999880"
				}]
				},
			"id":"7201",
			"userID":6395930,
			"accountID":"101-001-6395930-001",
			"batchID":"7200",
			"requestID":"78486291686830674",
			"time":"2018-07-27T19:49:18.256694465Z"
			},
			"relatedTransactionIDs":[
				"7200",
				"7201",
				"7202",
				"7203"
				],
			"lastTransactionID":"7203"}
*/
