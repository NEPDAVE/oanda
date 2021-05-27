package oanda

import (
	"io"
	"io/ioutil"
	"net/http"
)

var client = http.Client{}

//ReqArgs represents the Request Arguments passed to MakeRequest to hit the correct Oanda endpoint
type RequestArgs struct {
	Method string
	URL    string
	Body   io.Reader
}

//MakeRequest takes a RequestArgs as an argument and uses it to hit the correct Oanda endpoint
func MakeRequest(ra *RequestArgs) ([]byte, error) {
	req, err := http.NewRequest(ra.Method, ra.URL, ra.Body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", Bearer)
	req.Header.Add("connection", "keep-alive")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return respBytes, nil
}
