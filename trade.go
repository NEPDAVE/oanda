package oanda

import (
	"bytes"
	"encoding/json"
	"time"
)

/*
Trade Endpoints
*/

//TradesPayload represents mutltiple trades
type TradesPayload struct {
	LastTransactionID string `json:"lastTransactionID"`
	Trades            []struct {
		CurrentUnits string    `json:"currentUnits"`
		Financing    string    `json:"financing"`
		ID           string    `json:"id"`
		InitialUnits string    `json:"initialUnits"`
		Instrument   string    `json:"instrument"`
		OpenTime     time.Time `json:"openTime"`
		Price        string    `json:"price"`
		RealizedPL   string    `json:"realizedPL"`
		State        string    `json:"state"`
		UnrealizedPL string    `json:"unrealizedPL"`
	} `json:"trades"`
}

//TradePayload represent a single trade
type TradePayload struct {
	LastTransactionID string `json:"lastTransactionID"`
	Trades            struct {
		CurrentUnits string    `json:"currentUnits"`
		Financing    string    `json:"financing"`
		ID           string    `json:"id"`
		InitialUnits string    `json:"initialUnits"`
		Instrument   string    `json:"instrument"`
		OpenTime     time.Time `json:"openTime"`
		Price        string    `json:"price"`
		RealizedPL   string    `json:"realizedPL"`
		State        string    `json:"state"`
		UnrealizedPL string    `json:"unrealizedPL"`
	} `json:"trade"`
}

//CloseTradePayload represents the number of Units to reduce a trade by
type CloseTradePayload struct {
	Units string
}

type ModifiedTrade struct {
	OrderCreateTransaction struct {
		Type         string `json:"type"`
		Instrument   string `json:"instrument"`
		Units        string `json:"units"`
		TimeInForce  string `json:"timeInForce"`
		PositionFill string `json:"positionFill"`
		Reason       string `json:"reason"`
		TradeClose   struct {
			Units   string `json:"units"`
			TradeID string `json:"tradeID"`
		} `json:"tradeClose"`
		ID        string    `json:"id"`
		UserID    int       `json:"userID"`
		AccountID string    `json:"accountID"`
		BatchID   string    `json:"batchID"`
		RequestID string    `json:"requestID"`
		Time      time.Time `json:"time"`
	} `json:"orderCreateTransaction"`
	OrderFillTransaction struct {
		Type           string `json:"type"`
		Instrument     string `json:"instrument"`
		Units          string `json:"units"`
		Price          string `json:"price"`
		FullPrice      string `json:"fullPrice"`
		PL             string `json:"pl"`
		Financing      string `json:"financing"`
		Commission     string `json:"commission"`
		AccountBalance string `json:"accountBalance"`
		TradeOpened    string `json:"tradeOpened"`
		TimeInForce    string `json:"timeInForce"`
		PositionFill   string `json:"positionFill"`
		Reason         string `json:"reason"`
		TradesClosed   []struct {
			TradeID    string `json:"tradeID"`
			Units      string `json:"units"`
			RealizedPL string `json:"realizedPL"`
			Financing  string `json:"financing"`
		} `json:"tradesClosed"`
		TradeReduced struct {
			TradeID    string `json:"tradeID"`
			Units      string `json:"units"`
			RealizedPL string `json:"realizedPL"`
			Financing  string `json:"financing"`
		} `json:"tradeReduced"`
		ID            string    `json:"id"`
		UserID        int       `json:"userID"`
		AccountID     string    `json:"accountID"`
		BatchID       string    `json:"batchID"`
		RequestID     string    `json:"requestID"`
		OrderID       string    `json:"orderId"`
		ClientOrderID string    `json:"clientOrderId"`
		Time          time.Time `json:"time"`
	} `json:"orderFillTransaction"`
	OrderCancelTransaction struct {
		Type      string    `json:"type"`
		OrderID   string    `json:"orderID"`
		Reason    string    `json:"reason"`
		ID        string    `json:"id"`
		UserID    int       `json:"userID"`
		AccountID string    `json:"accountID"`
		BatchID   string    `json:"batchID"`
		RequestID string    `json:"requestID"`
		Time      time.Time `json:"time"`
	} `json:"orderCancelTransaction"`
	RelatedTransactionIDs []string `json:"relatedTransactionIDs"`
	LastTransactionID     string   `json:"lastTransactionID"`
}

//GetOpenTrades returns all the open trades for an account
func GetOpenTrades() (*TradesPayload, error) {
	reqArgs := &ReqArgs{
		ReqMethod: "GET",
		URL:       oandaHost + "/accounts/" + accountID + "/openTrades",
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
		URL:       oandaHost + "/accounts/" + accountID + "/trades/" + tradeSpecifier,
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
		URL:       oandaHost + "/accounts/" + accountID + "/trades/" + tradeSpecifier + "/close",
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
