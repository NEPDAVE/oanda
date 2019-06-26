package oanda

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

var (
	oandaURL       string
	streamoandaURL string
	bearer         string
	accountID      string
	client         *http.Client
	logger         *log.Logger
)

//OandaInit populates the global variables using using the evironment variables
func OandaInit(logger *log.Logger) {
	client = &http.Client{}
	logger = logger
	oandaURL = os.Getenv("OANDA_URL")
	streamoandaURL = os.Getenv("STREAM_OANDA_URL")
	bearer = "Bearer " + os.Getenv("OANDA_TOKEN")
	accountID = os.Getenv("OANDA_ACCOUNT_ID")

}

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
	queryValues := url.Values{}
	instrumentsEncoded := strings.Join(instruments, ",")
	queryValues.Add("instruments", instrumentsEncoded)

	req, err := http.NewRequest("GET", oandaURL+"/accounts/"+accountID+
		"/pricing?"+queryValues.Encode(), nil)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", bearer)
	req.Header.Add("Connection", "Keep-Alive")

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
	req, err := http.NewRequest("POST", oandaURL+"/accounts/"+accountID+"/orders", body)

	req.Header.Set("Authorization", bearer)
	req.Header.Set("content-type", "application/json")

	if err != nil {
		logger.Println(err)
		return []byte{}, err
	}

	resp, err := client.Do(req)

	if err != nil {
		logger.Println(err)
		return []byte{}, err
	}

	defer resp.Body.Close()

	createOrderByte, _ := ioutil.ReadAll(resp.Body)

	if err != nil {
		logger.Println(err)
		return []byte{}, err
	}

	return createOrderByte, err
}

/*
curl \
  -X PUT \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer 9fd32dee7bac39d8af58cd654b193b61-f6c942e3a94280431256657ffe9d9a70" \
  "https://api-fxpractice.oanda.com/v3/accounts/101-001-6395930-001/orders/6372/cancel"
*/

//CancelOrder used to submit orders
func CancelOrder(OrderID string) ([]byte, error) {
	client := &http.Client{}

	req, err := http.NewRequest("PUT", oandaURL+"/accounts/"+accountID+
		"/orders/"+OrderID+"/cancel", nil)

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

	cancelOrderByte, _ := ioutil.ReadAll(resp.Body)

	if err != nil {
		return []byte{}, err
	}

	return cancelOrderByte, err
}

//curl \
//  -H "Content-Type: application/json" \
//  -H "Authorization: Bearer 9fd32dee7bac39d8af58cd654b193b61-f6c942e3a94280431256657ffe9d9a70" \
//  "https://api-fxpractice.oanda.com/v3/accounts/101-001-6395930-001/orders/6372"
// my req  = https://api-fxpractice.oanda.com/v3/accounts/101-001-6395930-001/orders/10650

//GetOrder gets information on single order
func GetOrder(orderID string) ([]byte, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", oandaURL+"/accounts/"+accountID+
		"/orders/"+orderID, nil)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", bearer)

	if err != nil {
		return []byte{}, err
	}

	resp, err := client.Do(req)

	fmt.Println("GET ORDER STATUS CODE:")
	fmt.Println(resp.StatusCode)

	if err != nil {
		return []byte{}, err
	}

	defer resp.Body.Close()

	getOrderByte, _ := ioutil.ReadAll(resp.Body)

	if err != nil {
		return []byte{}, err
	}

	return getOrderByte, err
}

/*
***************************
trades
***************************
*/

// //ClosePositions closes all positions for instrument
// func ClosePositions(instrument string) ([]byte, error) {
// 	close := Close{LongUnits: "ALL", ShortUnits: "ALL"}
// 	longAndShort := MarshalClosePositions(close)
// 	body := bytes.NewBuffer(longAndShort)
// 	client := &http.Client{}
//
// 	req, err := http.NewRequest("PUT", oandaURL+"/accounts/"+accountID+
// 		"/positions/"+instrument+"/close", body)
//
// 	fmt.Println(req)
//
// 	req.Header.Set("Authorization", bearer)
// 	req.Header.Set("content-type", "application/json")
//
// 	if err != nil {
// 		return []byte{}, err
// 	}
//
// 	resp, err := client.Do(req)
//
// 	if err != nil {
// 		return []byte{}, err
// 	}
//
// 	defer resp.Body.Close()
//
// 	positionsResponseByte, _ := ioutil.ReadAll(resp.Body)
//
// 	if err != nil {
// 		return []byte{}, err
// 	}
//
// 	return positionsResponseByte, err
// }

//curl: Set dependent Orders for Trade 6397
// body=$(cat << EOF
// {
//   "takeProfit": {
//     "timeInForce": "GTC",
//     "price": "0.5"
//   },
//   "stopLoss": {
//     "timeInForce": "GTC",
//     "price": "2.5"
//   }
// }
// EOF
// )
//
// curl \
//   -X PUT \
//   -H "Content-Type: application/json" \
//   -H "Authorization: Bearer 9fd32dee7bac39d8af58cd654b193b61-f6c942e3a94280431256657ffe9d9a70" \
//   -d "$body" \
//   "https://api-fxtrade.oanda.com/v3/accounts/101-001-6395930-001/trades/6397/orders"

/*
***************************
position
***************************
*/

/*
curl: Get EUR_USD Position details for Account
curl \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer 9fd32dee7bac39d8af58cd654b193b61-f6c942e3a94280431256657ffe9d9a70" \
  "https://api-fxtrade.oanda.com/v3/accounts/101-001-6395930-001/positions/EUR_USD"
*/

//GetPositionDetails gets the position details for a single instrument
func GetPositionDetails(instrument string) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", oandaURL+"/accounts/"+accountID+
		"/positions/"+instrument, nil)

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

	positionResponseByte, _ := ioutil.ReadAll(resp.Body)

	if err != nil {
		return []byte{}, err
	}

	return positionResponseByte, err
}

//Close contains the number of longUnits and shortUnits to close for the instrument
type Close struct {
	LongUnits  string `json:"longUnits"`
	ShortUnits string `json:"shortUnits"`
}

//CloseLongPositions closes all positions for instrument
func CloseLongPositions(instrument string) ([]byte, error) {
	closeAllLongUnits := []byte(`{"longUnits": "ALL" }`)
	body := bytes.NewBuffer(closeAllLongUnits)
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

//CloseShortPositions closes all positions for instrument
func CloseShortPositions(instrument string) ([]byte, error) {
	closeAllShortUnits := []byte(`{"shortUnits": "ALL" }`)
	body := bytes.NewBuffer(closeAllShortUnits)
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

// body=$(cat << EOF
// {
//   "clientExtensions": {
//     "comment": "New comment for my limit order"
//   }
// }
// EOF
// )
//
// curl \
//   -X PUT \
//   -H "Content-Type: application/json" \
//   -H "Authorization: Bearer 9fd32dee7bac39d8af58cd654b193b61-f6c942e3a94280431256657ffe9d9a70" \
//   -d "$body" \
//   "https://api-fxtrade.oanda.com/v3/accounts/101-001-6395930-001/orders/6372/clientExtensions"

/*
	***************************
	account
	***************************
*/

/*
	curl \
	  -H "Content-Type: application/json" \
	  -H "Authorization: Bearer 9fd32dee7bac39d8af58cd654b193b61-f6c942e3a94280431256657ffe9d9a70" \
	  "https://api-fxtrade.oanda.com/v3/accounts/101-001-6395930-001/summary"
*/

//GetAccountSummary gets the position details for a single instrument
func GetAccountSummary() ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", oandaURL+"/accounts/"+accountID+
		"/summary", nil)

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

	accountSummaryResponseByte, _ := ioutil.ReadAll(resp.Body)

	if err != nil {
		return []byte{}, err
	}

	return accountSummaryResponseByte, err
}
