package oanda

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

var oandaURL = os.Getenv("OANDA_URL")
var streamoandaURL = os.Getenv("STREAM_OANDA_URL")
var bearer = "Bearer " + os.Getenv("OANDA_TOKEN")
var accountID = os.Getenv("OANDA_ACCOUNT_ID")

/*
***************************
prices
***************************
*/

//StreamResult is sent over channel in StreamPricing func
type StreamResult struct {
	PriceByte []byte
	Error     error
}

//StreamPricing can stream multiple prices at once. opting not to for simplicity
func StreamPricing(instruments string, out chan StreamResult) {
	defer close(out)

	client := &http.Client{}
	queryValues := url.Values{}
	queryValues.Add("instruments", instruments)

	req, err := http.NewRequest("GET", streamoandaURL+"/accounts/"+accountID+
		"/pricing/stream?"+queryValues.Encode(), nil)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", bearer)

	if err != nil {
		out <- StreamResult{Error: err}
	}

	resp, err := client.Do(req)

	if err != nil {
		out <- StreamResult{Error: err}
	}

	defer resp.Body.Close()

	reader := bufio.NewReader(resp.Body)

	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			out <- StreamResult{Error: err}
		}
		out <- StreamResult{PriceByte: line, Error: err}
	}
}

//GetPricing does single request to API
func GetPricing(instruments ...string) ([]byte, error) {
	client := &http.Client{}
	queryValues := url.Values{}
	instrumentsEncoded := strings.Join(instruments, ",")
	queryValues.Add("instruments", instrumentsEncoded)

	req, err := http.NewRequest("GET", oandaURL+"/accounts/"+accountID+
		"/pricing?"+queryValues.Encode(), nil)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", bearer)

	if err != nil {
		return []byte{}, err
	}

	resp, err := client.Do(req)

	if err != nil {
		return []byte{}, err
	}

	defer resp.Body.Close()

	pricesByte, _ := ioutil.ReadAll(resp.Body)

	if err != nil {
		return []byte{}, err
	}

	return pricesByte, err
}

/*
***************************
history
***************************
*/

//GetCandles retrieves instrument history in the form of candle sticks
func GetCandles(instrument string, count string, granularity string) ([]byte, error) {
	client := &http.Client{}
	queryValues := url.Values{}
	queryValues.Add("instruments", instrument)
	queryValues.Add("count", count)
	queryValues.Add("granularity", granularity)

	req, err := http.NewRequest("GET", oandaURL+"/instruments"+"/"+instrument+
		"/candles?"+queryValues.Encode(), nil)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", bearer)

	if err != nil {
		return []byte{}, err
	}

	resp, err := client.Do(req)

	if err != nil {
		return []byte{}, err
	}

	defer resp.Body.Close()

	candlesByte, _ := ioutil.ReadAll(resp.Body)

	defer resp.Body.Close()

	if err != nil {
		return []byte{}, err
	}

	return candlesByte, err
}

/*
***************************
orders
***************************
*/

//SubmitOrder used to submit orders
func SubmitOrder(orders []byte) ([]byte, error) {
	body := bytes.NewBuffer(orders)
	client := &http.Client{}

	req, err := http.NewRequest("POST", oandaURL+"/accounts/"+accountID+"/orders", body)

	req.Header.Set("Authorization", bearer)
	req.Header.Set("content-type", "application/json")

	if err != nil {
		return []byte{}, err
	}

	resp, err := client.Do(req)

	if err != nil {
		return []byte{}, err
	}

	defer resp.Body.Close()

	pricesByte, _ := ioutil.ReadAll(resp.Body)

	if err != nil {
		return []byte{}, err
	}

	return pricesByte, err
}
