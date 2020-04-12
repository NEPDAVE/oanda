package oanda

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type OpenTrades struct {
	LastTransactionID string   `json:"lastTransactionID"`
	Trades            []Trades `json:"trades"`
}
type ClientExtensions struct {
	ID string `json:"id"`
}
type Trades struct {
	CurrentUnits     string           `json:"currentUnits"`
	Financing        string           `json:"financing"`
	ID               string           `json:"id"`
	InitialUnits     string           `json:"initialUnits"`
	Instrument       string           `json:"instrument"`
	OpenTime         time.Time        `json:"openTime"`
	Price            string           `json:"price"`
	RealizedPL       string           `json:"realizedPL"`
	State            string           `json:"state"`
	UnrealizedPL     string           `json:"unrealizedPL"`
	ClientExtensions ClientExtensions `json:"clientExtensions,omitempty"`
}

func NewOpenTrades() (*OpenTrades, error) {
	openTrades := &OpenTrades{}

	openTradesByte, err := GetOpenTrades()

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(openTradesByte, &openTrades)

	if err != nil {
		return nil, err
	}

	return openTrades, nil
}

func GetOpenTrades() ([]byte, error) {
	req, err := http.NewRequest("GET", oandaURL+"/accounts/"+accountID+"/openTrades", nil)

	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", bearer)
	req.Header.Add("Connection", "Keep-Alive")

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	openTradesByte, err := ioutil.ReadAll(resp.Body)

	fmt.Println(string(openTradesByte))

	if err != nil {
		return nil, err
	}

	return openTradesByte, nil

}
