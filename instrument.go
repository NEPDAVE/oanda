package oanda

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

//IC represents the JSON returned by /v3/instruments/{instrument}/candles
type IC struct {
	Candles     []Candles `json:"candles"`
	Granularity string    `json:"granularity"`
	Instrument  string    `json:"instrument"`
}
type Mid struct {
	C string `json:"c"`
	H string `json:"h"`
	L string `json:"l"`
	O string `json:"o"`
}
type Candles struct {
	Complete bool      `json:"complete"`
	Mid      Mid       `json:"mid"`
	Time     time.Time `json:"time"`
	Volume   int       `json:"volume"`
}

func NewIC(instrument string, count string, granularity string) (*IC, error) {
	ic := &IC{}

	icByte, err := ic.GetCandles(instrument, count, granularity)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(icByte, &ic)

	if err != nil {
		return nil, err
	}

	return ic, nil
}

//GetCandles returns historical instrument candle data
func (i *IC) GetCandles(instrument string, count string, granularity string) ([]byte, error) {
	queryValues := url.Values{}
	queryValues.Add("instruments", instrument)
	queryValues.Add("count", count)
	queryValues.Add("granularity", granularity)

	req, err := http.NewRequest("GET", oandaURL+"/instruments"+"/"+instrument+
		"/candles?"+queryValues.Encode(), nil)

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

	icByte, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return icByte, err
}
