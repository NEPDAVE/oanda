package oanda

import (
	"encoding/json"
	"time"
)

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
