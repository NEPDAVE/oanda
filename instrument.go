package oanda

import (
	"encoding/json"
	"net/url"
	"time"
)

//InstrumentHistory represents the JSON returned by /v3/instruments/{instrument}/candles
type InstrumentHistory struct {
	Candles     []Candle `json:"candles"`
	Granularity string    `json:"granularity"`
	Instrument  string    `json:"instrument"`
}

type Candle struct {
	Complete bool      `json:"complete"`
	Mid      Mid       `json:"mid"`
	Time     time.Time `json:"time"`
	Volume   int       `json:"volume"`
}

type Mid struct {
	C string `json:"c"`
	H string `json:"h"`
	L string `json:"l"`
	O string `json:"o"`
}

func GetCandles(instrument string, count string, granularity string) (*InstrumentHistory, error) {
	queryValues := url.Values{}
	queryValues.Add("instruments", instrument)
	queryValues.Add("count", count)
	queryValues.Add("granularity", granularity)
	queryValues.Add("alignmentTimezone", "America/New_York")

	reqArgs := &ReqArgs{
		ReqMethod: "GET",
		URL:       OandaHost + "/instruments/" + instrument + "/candles?" + queryValues.Encode(),
	}

	candlesBytes, err := MakeRequest(reqArgs)

	if err != nil {
		return nil, err
	}

	instrumentHistory := &InstrumentHistory{}
	err = json.Unmarshal(candlesBytes, instrumentHistory)

	if err != nil {
		return nil, err
	}

	return instrumentHistory, nil
}
