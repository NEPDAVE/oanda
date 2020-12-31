package oanda

import (
	"bytes"
	"encoding/json"
	"time"
)

/*
Trade Endpoints
*/

//TradePayload represent a single trade
type TradePayload struct {
	LastTransactionID string  `json:"lastTransactionID"`
	Trades            []Trade `json:"trade"`
}

//TradesPayload represents multiple trades
type TradesPayload struct {
	LastTransactionID string  `json:"lastTransactionID"`
	Trades            []Trade `json:"trades"`
}

type ClientExtensions struct {
	Comment string `json:"comment"`
	Tag     string `json:"tag"`
	ID      string `json:"id"`
}

type Trade struct {
	CurrentUnits     string           `json:"currentUnits"`
	Financing        string           `json:"financing"`
	ID               string           `json:"id"`
	InitialUnits     string           `json:"initialUnits"`
	Instrument       string           `json:"instrument"`
	OpenTime         time.Time        `json:"openTime"`
	Price            string           `json:"price"`
	RealizedPL       string           `json:"realizedPL"`
	State            string           `json:"state"`
	UnrealizedPL     string           `json:"unrealizedPL"`
	ClientExtensions ClientExtensions `json:"clientExtensions,omitempty"`
}

//CloseTradePayload represent the number of units to reduce a trade by
type CloseTradePayload struct {
	Units string
}

type DependentOrders struct {
	TakeProfit TakeProfit `json:"takeProfit"`
	StopLoss   StopLoss   `json:"stopLoss"`
}
type TakeProfit struct {
	TimeInForce string `json:"timeInForce"`
	Price       string `json:"price"`
}
type StopLoss struct {
	TimeInForce string `json:"timeInForce"`
	Price       string `json:"price"`
}

//GetOpenTrades returns all the open trades for an account
func GetOpenTrades() (*TradesPayload, error) {
	reqArgs := &ReqArgs{
		ReqMethod: "GET",
		URL:       OandaHost + "/accounts/" + accountID + "/openTrades",
	}

	openTradesBytes, err := MakeRequest(reqArgs)

	if err != nil {
		return nil, err
	}

	openTrades := &TradesPayload{}
	err = json.Unmarshal(openTradesBytes, openTrades)

	if err != nil {
		return nil, err
	}

	return openTrades, nil
}

//GetTrade gets the details of a specific trade in an account
func GetTrade(tradeSpecifier string) (*TradePayload, error) {
	reqArgs := &ReqArgs{
		ReqMethod: "GET",
		URL:       OandaHost + "/accounts/" + accountID + "/trades/" + tradeSpecifier,
	}

	tradeBytes, err := MakeRequest(reqArgs)

	if err != nil {
		return nil, err
	}

	trade := &TradePayload{}
	err = json.Unmarshal(tradeBytes, trade)

	if err != nil {
		return nil, err
	}

	return trade, nil
}

//CloseTrade partially or fully closes a specific open Trade in an Account
func CloseTrade(tradeSpecifier string, units string) (*TradePayload, error) {
	bodyBytes, err := json.Marshal(CloseTradePayload{Units: units})

	if err != nil {
		return nil, err
	}

	body := bytes.NewReader(bodyBytes)

	reqArgs := &ReqArgs{
		ReqMethod: "PUT",
		URL:       OandaHost + "/accounts/" + accountID + "/trades/" + tradeSpecifier + "/close",
		Body:      body,
	}

	tradeBytes, err := MakeRequest(reqArgs)

	if err != nil {
		return nil, err
	}

	trade := &TradePayload{}
	err = json.Unmarshal(tradeBytes, trade)

	if err != nil {
		return nil, err
	}

	return trade, nil
}

//ModifyDependentOrders replaces and/or cancels a trade's dependent orders - IE
//take profit, stop loss, and trailing stop loss orders.
func ModifyDependentOrders(stopLoss StopLoss, takeProfit TakeProfit, tradeSpecifier string) (*TradePayload, error) {
	dOBytes, err := json.Marshal(DependentOrders{StopLoss: stopLoss, TakeProfit: takeProfit})

	if err != nil {
		return nil, err
	}

	body := bytes.NewReader(dOBytes)

	reqArgs := &ReqArgs{
		ReqMethod: "PUT",
		URL:       OandaHost + "/accounts/" + accountID + "/trades/" + tradeSpecifier + "/orders",
		Body:      body,
	}

	tradeBytes, err := MakeRequest(reqArgs)

	if err != nil {
		return nil, err
	}

	trade := &TradePayload{}
	err = json.Unmarshal(tradeBytes, trade)

	if err != nil {
		return nil, err
	}

	return trade, nil
}