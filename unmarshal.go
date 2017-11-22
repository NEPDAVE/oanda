package main

import (
	"encoding/json"
	"time"
)


type Prices struct {
	Type        string    `json:"type"`
	Asks        []Asks    `json:"asks"`
	Bids        []Bids    `json:"bids"`
	CloseOutAsk string    `json:"closeoutAsk"`
	CloseOutBid string    `json:"closeoutBid"`
	Instrument  string    `json:"instrument"`
	Status      string    `json:"tradeable"`
	Time        time.Time `json:"time"`
}

type Asks struct {
	Price     float64 `json:"price"`
	Liquidity string  `json:"liquidity"`
	}

type Bids struct {
	Price     float64 `json:"price"`
	Liquidity string  `json:"liquidity"`
}


func (p Prices) UnmarshalPricing(priceByte []byte) *Prices {

        err := json.Unmarshal(priceByte, &p)

				if err != nil {
                panic(err)
        }

        return &p
}
