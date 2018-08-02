package oanda

import (
	"bufio"
	"bytes"
	"fmt"
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
		fmt.Println(string(line))
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

//CreateOrder used to submit orders
func CreateOrder(orders []byte) ([]byte, error) {
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

	createOrderByte, _ := ioutil.ReadAll(resp.Body)

	if err != nil {
		return []byte{}, err
	}

	return createOrderByte, err
}


//curl \
//  -H "Content-Type: application/json" \
//  -H "Authorization: Bearer 9fd32dee7bac39d8af58cd654b193b61-f6c942e3a94280431256657ffe9d9a70" \
//  "https://api-fxpractice.oanda.com/v3/accounts/101-001-6395930-001/orders/6372"

//CheckOrder gets information on single order
func CheckOrder(orderID string) ([]byte, error) {
	client := &http.Client{}

  req, err := http.NewRequest("GET", oandaURL+"/accounts/"+accountID+
		"/orders/"+orderID, nil)

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

	checkOrderByte, _ := ioutil.ReadAll(resp.Body)

	if err != nil {
		return []byte{}, err
	}

	return checkOrderByte, err
}

/*
***************************
position
***************************
*/

//Close contains the number of longUnits and shortUnits to close for the instrument
type Close struct {
	LongUnits  string `json:"longUnits"`
	ShortUnits string `json:"shortUnits"`
}

//ClosePositions closes all positions for instrument
func ClosePositions(instrument string) ([]byte, error) {
	close := Close{LongUnits: "ALL", ShortUnits: "ALL"}
	fmt.Println("###########################")
	fmt.Println(close)
	fmt.Println("###########################")

	longAndShort := MarshalClosePositions(close)
	body := bytes.NewBuffer(longAndShort)
	client := &http.Client{}

	req, err := http.NewRequest("PUT", oandaURL+"/accounts/"+accountID+
		"/positions/"+instrument+"/close", body)

	fmt.Println(req)

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

	positionsResponseByte, _ := ioutil.ReadAll(resp.Body)

	if err != nil {
		return []byte{}, err
	}

	return positionsResponseByte, err
}
