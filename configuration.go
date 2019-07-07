package oanda

import (
	"log"
	"os"
	"net/http"
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
func OandaInit(oandaLogger *log.Logger) {
	client = &http.Client{}
	logger = oandaLogger
	oandaURL = os.Getenv("OANDA_URL")
	streamoandaURL = os.Getenv("STREAM_OANDA_URL")
	bearer = "Bearer " + os.Getenv("OANDA_TOKEN")
	accountID = os.Getenv("OANDA_ACCOUNT_ID")

}
