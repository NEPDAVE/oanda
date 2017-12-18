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
