package trade

import (
	"bytes"
	"encoding/json"
	"time"

	"github.com/nepdave/oanda"
)

/*
Trade Endpoints
*/

type Response struct {
	LastTransactionID          string                    `json:"lastTransactionID"`
	Trades                     []Trade                   `json:"trades"`
	RelatedTransactionIDs      []string                  `json:"relatedTransactionIDs"`
	StopLossOrderTransaction   DependentOrderTransaction `json:"stopLossOrderTransaction"`
	TakeProfitOrderTransaction DependentOrderTransaction `json:"takeProfitOrderTransaction"`
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

type ClientExtensions struct {
	Comment string `json:"comment"`
	Tag     string `json:"tag"`
	ID      string `json:"id"`
}

type DependentOrderTransaction struct {
	AccountID        string    `json:"accountID"`
	BatchID          string    `json:"batchID"`
	ClientTradeID    string    `json:"clientTradeID"`
	ID               string    `json:"id"`
	Price            string    `json:"price"`
	Reason           string    `json:"reason"`
	Time             time.Time `json:"time"`
	TimeInForce      string    `json:"timeInForce"`
	TradeID          string    `json:"tradeID"`
	TriggerCondition string    `json:"triggerCondition"`
	Type             string    `json:"type"`
	UserID           string    `json:"userID"`
}

type Close struct {
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

//GetTrades returns a list of trades for an account
func GetTrades() (*Response, error) {
	reqArgs := &oanda.RequestArgs{
		Method: "GET",
		URL:    oanda.Host + "/accounts/" + oanda.AccountID + "/trades",
	}

	openTradesBytes, err := oanda.MakeRequest(reqArgs)

	if err != nil {
		return nil, err
	}

	trades := &Response{}
	err = json.Unmarshal(openTradesBytes, trades)

	if err != nil {
		return nil, err
	}

	return trades, nil
}

//GetOpenTrades returns a list of open trades for an account
func GetOpenTrades() (*Response, error) {
	reqArgs := &oanda.RequestArgs{
		Method: "GET",
		URL:    oanda.Host + "/accounts/" + oanda.AccountID + "/openTrades",
	}

	openTradesBytes, err := oanda.MakeRequest(reqArgs)

	if err != nil {
		return nil, err
	}

	openTrades := &Response{}
	err = json.Unmarshal(openTradesBytes, openTrades)

	if err != nil {
		return nil, err
	}

	return openTrades, nil
}

//GetTrade returns the details of a specific trade in an account
func GetTrade(tradeSpecifier string) (*Response, error) {
	reqArgs := &oanda.RequestArgs{
		Method: "GET",
		URL:    oanda.Host + "/accounts/" + oanda.AccountID + "/trades/" + tradeSpecifier,
	}

	tradeBytes, err := oanda.MakeRequest(reqArgs)

	if err != nil {
		return nil, err
	}

	trade := &Response{}
	err = json.Unmarshal(tradeBytes, trade)

	if err != nil {
		return nil, err
	}

	return trade, nil
}

//PutClose partially or fully closes a specific open Trade in an Account
func PutClose(tradeSpecifier string, units string) (*Response, error) {
	bodyBytes, err := json.Marshal(Close{Units: units})
	if err != nil {
		return nil, err
	}

	body := bytes.NewReader(bodyBytes)

	reqArgs := &oanda.RequestArgs{
		Method: "PUT",
		URL:    oanda.Host + "/accounts/" + oanda.AccountID + "/trades/" + tradeSpecifier + "/close",
		Body:   body,
	}

	tradeBytes, err := oanda.MakeRequest(reqArgs)
	if err != nil {
		return nil, err
	}

	trade := &Response{}
	err = json.Unmarshal(tradeBytes, trade)

	if err != nil {
		return nil, err
	}

	return trade, nil
}

func PutClientExtensions(tradeSpecifier string, comment string, tag string, id string) (*Response, error) {
	bodyBytes, err := json.Marshal(ClientExtensions{
		Comment: comment,
		Tag:     tag,
		ID:      id,
	})
	if err != nil {
		return nil, err
	}

	body := bytes.NewReader(bodyBytes)

	reqArgs := &oanda.RequestArgs{
		Method: "PUT",
		URL:    oanda.Host + "/accounts/" + oanda.AccountID + "/trades/" + tradeSpecifier + "/clientExtensions",
		Body:   body,
	}

	tradeBytes, err := oanda.MakeRequest(reqArgs)
	if err != nil {
		return nil, err
	}

	trade := &Response{}
	err = json.Unmarshal(tradeBytes, trade)
	if err != nil {
		return nil, err
	}

	return trade, nil
}

//PutOrders replaces and/or cancels a trade's dependent orders - IE
//take profit, stop loss, and trailing stop loss orders.
func PutOrders(tradeSpecifier string, stopLoss StopLoss, takeProfit TakeProfit) (*Response, error) {
	dOBytes, err := json.Marshal(DependentOrders{
		StopLoss:   stopLoss,
		TakeProfit: takeProfit,
	})
	if err != nil {
		return nil, err
	}

	body := bytes.NewReader(dOBytes)

	reqArgs := &oanda.RequestArgs{
		Method: "PUT",
		URL:    oanda.Host + "/accounts/" + oanda.AccountID + "/trades/" + tradeSpecifier + "/orders",
		Body:   body,
	}

	tradeBytes, err := oanda.MakeRequest(reqArgs)
	if err != nil {
		return nil, err
	}

	trade := &Response{}
	err = json.Unmarshal(tradeBytes, trade)
	if err != nil {
		return nil, err
	}

	return trade, nil
}
