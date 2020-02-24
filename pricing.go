package oanda

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Pricing struct {
	Prices []Prices `json:"prices"`
}

type Prices struct {
	Asks                       []Asks                     `json:"asks"`
	Bids                       []Bids                     `json:"bids"`
	CloseoutAsk                string                     `json:"closeoutAsk"`
	CloseoutBid                string                     `json:"closeoutBid"`
	Instrument                 string                     `json:"instrument"`
	QuoteHomeConversionFactors QuoteHomeConversionFactors `json:"quoteHomeConversionFactors"`
	Status                     string                     `json:"status"`
	Time                       time.Time                  `json:"time"`
	UnitsAvailable             UnitsAvailable             `json:"unitsAvailable"`
}

type Asks struct {
	Liquidity int    `json:"liquidity"`
	Price     string `json:"price"`
}

type Bids struct {
	Liquidity int    `json:"liquidity"`
	Price     string `json:"price"`
}

type QuoteHomeConversionFactors struct {
	NegativeUnits string `json:"negativeUnits"`
	PositiveUnits string `json:"positiveUnits"`
}

type Default struct {
	Long  string `json:"long"`
	Short string `json:"short"`
}

type OpenOnly struct {
	Long  string `json:"long"`
	Short string `json:"short"`
}

type ReduceFirst struct {
	Long  string `json:"long"`
	Short string `json:"short"`
}

type ReduceOnly struct {
	Long  string `json:"long"`
	Short string `json:"short"`
}

type UnitsAvailable struct {
	Default     Default     `json:"default"`
	OpenOnly    OpenOnly    `json:"openOnly"`
	ReduceFirst ReduceFirst `json:"reduceFirst"`
	ReduceOnly  ReduceOnly  `json:"reduceOnly"`
}

func NewPricing(instrument string) (*Pricing, error) {
	pricing := &Pricing{}

	pricesByte, err := pricing.GetPricing(instrument)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(pricesByte, &pricing)

	if err != nil {
		return nil, err
	}

	return pricing, nil

}

//GetPricing returns latest pricing data
func (p *Pricing) GetPricing(instruments ...string) ([]byte, error) {
	queryValues := url.Values{}
	instrumentsEncoded := strings.Join(instruments, ",")
	queryValues.Add("instruments", instrumentsEncoded)

	req, err := http.NewRequest("GET", oandaURL+"/accounts/"+accountID+
		"/pricing?"+queryValues.Encode(), nil)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", bearer)
	req.Header.Add("Connection", "Keep-Alive")

	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	pricesByte, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return pricesByte, nil
}
