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


var oandaUrl string = "https://api-fxpractice.oanda.com/v3"
var streamOandaUrl string = "https://stream-fxpractice.oanda.com/v3"
var bearer string = "Bearer " + os.Getenv("OANDA_TOKEN")
var accountId string = os.Getenv("OANDA_ACCOUNT_ID")

//callback function for printing out network requests
func LogComms(req *http.Request, pricesByte []byte, statusCode int, err error) {
	log.Printf("Request: %p\n", req)
	log.Printf("Response: %s\n", string(pricesByte))
	log.Printf("Status Code: %d\n", statusCode)
	log.Printf("GetPricing Response Error: %s\n", err)
}

/*
***************************
prices
***************************
*/

type StreamResult struct {
	PriceByte []byte
	Error     error
}

func StreamPricing(instruments string, out chan StreamResult) {
	client := &http.Client{}
	queryValues := url.Values{}
	queryValues.Add("instruments", instruments)

	req, err := http.NewRequest("GET", streamOandaUrl+"/accounts/"+accountId+
		"/pricing/stream?"+queryValues.Encode(), nil)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", bearer)

	if err != nil {
		log.Printf("StreamPricing error building request: %s\n", err)
	}

	resp, err := client.Do(req)

	if err != nil {
		log.Printf("StreamPricing error making request: %s\n", err)
	}

	defer resp.Body.Close()
	defer close(out)

	reader := bufio.NewReader(resp.Body)
	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			log.Printf("StreamPricing error reading response byteSlice: %s\n", err)
		}
		out <- StreamResult{PriceByte: line, Error: err}
	}
	log.Printf("closing streamPricing channel")
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
		return []byte{}, errors.New("error building request")
	}

	resp, err := client.Do(req)
	pricesByte, _ := ioutil.ReadAll(resp.Body)
	status := strconv.Itoa(resp.StatusCode)

	defer resp.Body.Close()

	if err != nil {
		LogComms(req, pricesByte, resp.StatusCode, err)
		errorMessage := fmt.Sprintf("error making request - status code: %s", status)
		return []byte{}, errors.New(errorMessage)
	}

	return pricesByte, nil

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
		return []byte{}, errors.New("error building request")
	}

	resp, err := client.Do(req)
	pricesByte, _ := ioutil.ReadAll(resp.Body)
	status := strconv.Itoa(resp.StatusCode)

	defer resp.Body.Close()

	if err != nil {
		LogComms(req, pricesByte, resp.StatusCode, err)
		errorMessage := fmt.Sprintf("error making request - status code: %s", status)
		return []byte{}, errors.New(errorMessage)
	}

	return pricesByte, nil

}

/*
***************************
orders
***************************
*/

func SubmitOrder(orders []byte) ([]byte, error) {
	body := bytes.NewBuffer(orders)
	//FIXME these should be moved somewhere or removed
	fmt.Println(body)
	fmt.Println()
	client := &http.Client{}

	req, err := http.NewRequest("POST", oandaUrl+"/accounts/"+accountId+"/orders", body)

	req.Header.Set("Authorization", bearer)
	req.Header.Set("content-type", "application/json")

	if err != nil {
		return []byte{}, errors.New("error building request")
	}

	resp, err := client.Do(req)
	pricesByte, _ := ioutil.ReadAll(resp.Body)
	status := strconv.Itoa(resp.StatusCode)

	defer resp.Body.Close()

	if err != nil {
		LogComms(req, pricesByte, resp.StatusCode, err)
		errorMessage := fmt.Sprintf("error making request - status code: %s", status)
		return []byte{}, errors.New(errorMessage)
	}

	return pricesByte, nil

}
