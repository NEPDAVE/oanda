package oanda

import (
	"encoding/json"
	"net/url"
	"strings"
	"time"
)

type PricingPayload struct {
	Prices []struct {
		Asks []struct {
			Liquidity int    `json:"liquidity"`
			Price     string `json:"price"`
		} `json:"asks"`
		Bids []struct {
			Liquidity int    `json:"liquidity"`
			Price     string `json:"price"`
		} `json:"bids"`
		CloseoutAsk                string `json:"closeoutAsk"`
		CloseoutBid                string `json:"closeoutBid"`
		Instrument                 string `json:"instrument"`
		QuoteHomeConversionFactors struct {
			NegativeUnits string `json:"negativeUnits"`
			PositiveUnits string `json:"positiveUnits"`
		} `json:"quoteHomeConversionFactors"`
		Status         string    `json:"status"`
		Time           time.Time `json:"time"`
		UnitsAvailable struct {
			Default struct {
				Long  string `json:"long"`
				Short string `json:"short"`
			} `json:"default"`
			OpenOnly struct {
				Long  string `json:"long"`
				Short string `json:"short"`
			} `json:"openOnly"`
			ReduceFirst struct {
				Long  string `json:"long"`
				Short string `json:"short"`
			} `json:"reduceFirst"`
			ReduceOnly struct {
				Long  string `json:"long"`
				Short string `json:"short"`
			} `json:"reduceOnly"`
		} `json:"unitsAvailable"`
	} `json:"prices"`
}

//GetPricing returns latest pricing data
func GetPricing(instruments []string) (*PricingPayload, error) {
	instrumentsString := strings.Join(instruments, ",")
	queryValues := url.Values{}
	queryValues.Add("instruments", instrumentsString)

	reqArgs := &ReqArgs{
		ReqMethod: "GET",
		URL:       oandaHost + "/accounts/" + accountID + "/pricing?" + queryValues.Encode(),
	}

	pricingBytes, err := MakeRequest(reqArgs)

	if err != nil {
		return nil, err
	}

	pricingPayload := &PricingPayload{}
	err = json.Unmarshal(pricingBytes, pricingPayload)

	if err != nil {
		return nil, err
	}

	return pricingPayload, nil
}
