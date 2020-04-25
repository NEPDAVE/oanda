package oanda

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

/*
Trade Endpoints
*/

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

//TODO check to make sure url values are encoded properly
//URL values should be encoded and ready to go before passing to MakeRequest
type Units struct {
	Units string `json:"units"`
}

func GetOpenTrades() (*OpenTrades, error) {
	reqArgs := &ReqArgs{
		ReqMethod: "GET",
		URL:       oandaURL + "/accounts/" + accountID + "/openTrades",
	}

	openTradesByte, err := MakeRequest(reqArgs)

	if err != nil {
		return nil, err
	}

	openTrades := &OpenTrades{}
	err = json.Unmarshal(openTradesByte, &openTrades)

	if err != nil {
		return nil, err
	}

	return openTrades, nil
}

//FIXME this should return a the trade you made, not all open trades
//func NewTrade() (*OpenTrades, error) {
//}

//GetTrade gets the details of a specific trade in an account
func GetTrade(tradeSpecifier string) ([]byte, error) {
	req, err := http.NewRequest("GET", oandaURL+"/accounts/"+accountID+"/trades/"+tradeSpecifier, nil)

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

	tradeByte, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return tradeByte, nil

}

//CloseTrade partially or fully closes a specific open Trade in an Account
func CloseTrade(tradeSpecifier string, units string) ([]byte, error) {
	b, err := json.Marshal(Units{Units: units})

	if err != nil {
		return nil, err
	}

	body := bytes.NewReader(b)

	req, err := http.NewRequest("PUT", oandaURL+"/accounts/"+accountID+"/trades/"+tradeSpecifier+"/close", body)

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

	closeTradeByte, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return closeTradeByte, nil
}
