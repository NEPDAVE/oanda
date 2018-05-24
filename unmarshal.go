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

//{"time":"2016-09-20T15:05:50.163791738Z","type":"HEARTBEAT"}
type Heartbeat struct {
	Time        time.Time `json:"time"`
	Type        string    `json:"type"`

}

func (h Heartbeat) UnmarshalHeartbeat(heartbeatByte []byte) *Heartbeat {

	err := json.Unmarshal(heartbeatByte, &h)

	if err != nil {
		panic(err)
	}

	return &h
}

type Pricing struct {
	Prices []Prices  `json:"prices"`
	Time   time.Time `json:"time"`
}

type Prices struct {
	Type        string    `json:"type"`
	Asks        []Ask     `json:"asks"`
	Bids        []Bid     `json:"bids"`
	CloseOutAsk string    `json:"closeoutAsk"`
	CloseOutBid string    `json:"closeoutBid"`
	Instrument  string    `json:"instrument"`
	Status      string    `json:"status"`
	Time        time.Time `json:"time"`
}

type Ask struct {
	Price     string `json:"price"`
	Liquidity int64  `json:"liquidity"`
}

type Bid struct {
	Price     string `json:"price"`
	Liquidity int64  `json:"liquidity"`
}

func (p Prices) UnmarshalPrices(priceByte []byte) *Prices {

	err := json.Unmarshal(priceByte, &p)

	if err != nil {
		panic(err)
	}

	return &p
}

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

type Candles struct {
	Instrument  string   `json:"instrument"`
	Granularity string   `json:"granularity"`
	Candles     []Candle `json:"candles"`
}

type Candle struct {
	Complete bool      `json:"complete"`
	Volume   int64     `json:"volume"`
	Time     time.Time `json:"time"`
	Mid      Mid       `json:"mid"`
}

type Mid struct {
	Open  string `json:"o"`
	High  string `json:"h"`
	Low   string `json:"l"`
	Close string `json:"c"`
}

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
*/
