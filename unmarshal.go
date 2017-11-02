package oanda

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
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


func (p Prices) UnmarshalPricing(instrument string) (float64, float64) {
        responseByte := p.SubscribeToPriceStream(instrument)

        err := json.Unmarshal(responseByte, &p)

				if err != nil {
                panic(err)
        }

				fmt.Println(p)
				fmt.Println("TEST3 TEST3 TEST3")
				fmt.Println(p.Asks)
        //totally fake return values
        return 1.232, 32.455
}
