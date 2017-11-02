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

var oandaUrl string = "https://api-fxtrade.oanda.com/v3"
var bearer string = "Bearer " + os.Getenv("OANDA_TOKEN")
var accountId string = os.Getenv("OANDA_ACCOUNT_ID")

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

func (p Prices) GetPricing(instrument string) []byte {
        client := &http.Client{}
        queryValues := url.Values{}

        req, err := http.NewRequest("GET", oandaUrl+"/accounts/"+accountId+"/pricing"+queryValues.Encode(), nil)
				queryValues.Add("instruments", instrument)
				req.Header.Add("Content-Type", "application/json")
				req.Header.Add("Authorization", bearer)


				fmt.Println("TEST TEST TEST")
				fmt.Println(req.Header)
				fmt.Println(queryValues)

				resp, err := client.Do(req)

        if err != nil {
                fmt.Println(err)
                os.Exit(1)
        }

        defer resp.Body.Close()
        byte, _ := ioutil.ReadAll(resp.Body)
        return byte
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
