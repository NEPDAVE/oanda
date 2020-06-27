package oanda

import (
	"bufio"
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type PricingPayload struct {
	Prices []Prices `json:"prices"`
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

//GetPricing returns latest pricing data - pricing data is returned for each
//instrument passed in array
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

//StreamPayload is sent over channel in StreamPricing func
type StreamPayload struct {
	PricingPayload *PricingPayload
	Error          error
}

//StreamPricing can stream multiple prices at once
func StreamPricing(instruments []string) chan StreamPayload {
	out := make(chan StreamPayload)
	defer close(out)

	client := &http.Client{}
	instrumentsString := strings.Join(instruments, ",")
	queryValues := url.Values{}
	queryValues.Add("instruments", instrumentsString)

	req, err := http.NewRequest("GET", streamOandaURL+"/accounts/"+accountID+
		"/pricing/stream?"+queryValues.Encode(), nil)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", bearer)

	if err != nil {
		out <- StreamPayload{Error: err}
	}

	resp, err := client.Do(req)

	if err != nil {
		out <- StreamPayload{Error: err}
	}

	defer resp.Body.Close()

	reader := bufio.NewReader(resp.Body)

	for {
		pricingBytes, err := reader.ReadBytes('\n')
		if err != nil {
			out <- StreamPayload{Error: err}
		}

		pricingPayload := &PricingPayload{}
		err = json.Unmarshal(pricingBytes, pricingPayload)

		if err != nil {
			out <- StreamPayload{Error: err}
		}
		out <- StreamPayload{PricingPayload: pricingPayload, Error: err}
	}
}
