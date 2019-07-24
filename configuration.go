package oanda

import (
	"log"
	"net/http"
	"os"
)

var (
	oandaURL       string
	streamoandaURL string
	bearer         string
	accountID      string
	logger         *log.Logger
	client         = &http.Client{}
)

//OandaInit populates the global variables using using the evironment variables
func OandaInit(oandaLogger *log.Logger) {
	logger = oandaLogger
	oandaURL = os.Getenv("OANDA_URL")
	streamoandaURL = os.Getenv("STREAM_OANDA_URL")
	bearer = "Bearer " + os.Getenv("OANDA_TOKEN")
	accountID = os.Getenv("OANDA_ACCOUNT_ID")

}
