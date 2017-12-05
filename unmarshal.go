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

type Pricing struct {
	Prices []Prices  `json: "prices"`
	Time   time.Time `json: "time"`
}

type Prices struct {
	Type        string    `json: "type"`
	Asks        []Ask     `json: "asks"`
	Bids        []Bid     `json: "bids"`
	CloseOutAsk string    `json: "closeoutAsk"`
	CloseOutBid string    `json: "closeoutBid"`
	Instrument  string    `json: "instrument"`
	Status      string    `json: "status"`
	Time        time.Time `json: "time"`
}

type Ask struct {
	Price     string `json: "price"`
	Liquidity int64  `json: "liquidity"`
}

type Bid struct {
	Price     string `json: "price"`
	Liquidity int64  `json: "liquidity"`
}

func (p Pricing) UnmarshalPricing(priceByte []byte) *Pricing {

	err := json.Unmarshal(priceByte, &p)

	if err != nil {
		panic(err)
	}

	return &p
}

/*
{"instrument":"EUR_USD",
"granularity":"D",
"candles":[
{"complete":false,
"volume":6556,
"time":"2017-12-03T22:00:00.000000000Z",
"mid":{"o":"1.18650","h":"1.18761","l":"1.18560","c":"1.18700"}}]}

//{EUR_USD D [{false 37981 2017-12-04 22:00:00 +0000 UTC {   }}]}
{"instrument":"EUR_USD",
"granularity":"D",
"candles":[{
"complete":false,
"volume":37981,
"time":"2017-12-04T22:00:00.000000000Z",
"mid":{"o":"1.18642","h":"1.18768","l":"1.18006","c":"1.18044"}}]
}

*/

type Candles struct {
	Instrument  string   `json: "instrument"`
	Granularity string   `json: "granularity"`
	Candles     []Candle `json: "candles"`
}

type Candle struct {
	Complete bool                   `json: "complete"`
	Volume   int64                  `json: "volume"`
	Time     time.Time              `json: "time"`
	Mid      map[string]interface{} `json: "mid"`
}

func (c Candles) UnmarshalCandles(priceByte []byte) *Candles {

	err := json.Unmarshal(priceByte, &c)

	if err != nil {
		panic(err)
	}

	return &c
}
