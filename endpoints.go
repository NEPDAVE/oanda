package oanda

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

<<<<<<< HEAD
//FIXME need to make the URLs environment variables
=======

>>>>>>> 4bdba3e1870da73dfcd92442215eb9ea85cfb719
var oandaUrl string = "https://api-fxpractice.oanda.com/v3"
var streamOandaUrl string = "https://stream-fxpractice.oanda.com/v3"
var bearer string = "Bearer " + os.Getenv("OANDA_TOKEN")
var accountId string = os.Getenv("OANDA_ACCOUNT_ID")

/*
***************************
prices
***************************
*/

<<<<<<< HEAD
//type sent over channel in StreamPricing func
=======
>>>>>>> 4bdba3e1870da73dfcd92442215eb9ea85cfb719
type StreamResult struct {
	PriceByte []byte
	Error     error
}

<<<<<<< HEAD
//possible to stream multiple prices at once. opting not to for simplicity
=======
>>>>>>> 4bdba3e1870da73dfcd92442215eb9ea85cfb719
func StreamPricing(instruments string, out chan StreamResult) {
	defer close(out)

	client := &http.Client{}
	queryValues := url.Values{}
	queryValues.Add("instruments", instruments)

	req, err := http.NewRequest("GET", streamOandaUrl+"/accounts/"+accountId+
		"/pricing/stream?"+queryValues.Encode(), nil)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", bearer)

	if err != nil {
		out <- StreamResult{Error: err}
	}

	resp, err := client.Do(req)
	defer resp.Body.Close()

	if err != nil {
		put <- StreamResult{Error: err}
	}

	reader := bufio.NewReader(resp.Body)

	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			out <- StreamResult{Error: err}
		}
		out <- StreamResult{PriceByte: line, Error: err}
	}
}

func GetPricing(instruments ...string) ([]byte, error) {
	client := &http.Client{}
	queryValues := url.Values{}
	instrumentsEncoded := strings.Join(instruments, ",")
	queryValues.Add("instruments", instrumentsEncoded)

	req, err := http.NewRequest("GET", oandaUrl+"/accounts/"+accountId+
		"/pricing?"+queryValues.Encode(), nil)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", bearer)

	if err != nil {
		return []byte{}, err
	}

	resp, err := client.Do(req)
	defer resp.Body.Close()

	pricesByte, _ := ioutil.ReadAll(resp.Body)
	status := strconv.Itoa(resp.StatusCode)

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

func GetCandles(instrument string, count string, granularity string) ([]byte, error) {
	client := &http.Client{}
	queryValues := url.Values{}
	queryValues.Add("instruments", instrument)
	queryValues.Add("count", count)
	queryValues.Add("granularity", granularity)

	req, err := http.NewRequest("GET", oandaUrl+"/instruments"+"/"+instrument+
		"/candles?"+queryValues.Encode(), nil)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", bearer)

	if err != nil {
		return []byte{}, err
	}

	resp, err := client.Do(req)
	defer resp.Body.Close()

	pricesByte, _ := ioutil.ReadAll(resp.Body)
	status := strconv.Itoa(resp.StatusCode)

	defer resp.Body.Close()

	if err != nil {
		return []byte{}, err
	}

	return pricesByte, err
}

/*
***************************
orders
***************************
*/

func SubmitOrder(orders []byte) ([]byte, error) {
	body := bytes.NewBuffer(orders)
	client := &http.Client{}

	req, err := http.NewRequest("POST", oandaUrl+"/accounts/"+accountId+"/orders", body)

	req.Header.Set("Authorization", bearer)
	req.Header.Set("content-type", "application/json")

	if err != nil {
		return []byte{}, err
	}

	resp, err := client.Do(req)
	defer resp.Body.Close()

	pricesByte, _ := ioutil.ReadAll(resp.Body)
	status := strconv.Itoa(resp.StatusCode)

	if err != nil {
		return []byte{}, err
	}

	return pricesByte, err
}
