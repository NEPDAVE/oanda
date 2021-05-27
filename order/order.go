package order

import (
	"bytes"
	"encoding/json"
	"github.com/nepdave/oanda"
	"net/http"
	"time"
)

/*
Order Endpoints
*/

//Request represents the Order details sent to Oanda for new Orders
type Request struct {
	Order Order `json:"order"`
}
type StopLossOnFill struct {
	TimeInForce string `json:"timeInForce"`
	Price       string `json:"price"`
}
type TakeProfitOnFill struct {
	TimeInForce string `json:"timeInForce"`
	Price       string `json:"price"`
}
type Order struct {
	Price            string           `json:"price"`
	StopLossOnFill   StopLossOnFill   `json:"stopLossOnFill"`
	TakeProfitOnFill TakeProfitOnFill `json:"takeProfitOnFill"`
	TimeInForce      string           `json:"timeInForce"`
	Instrument       string           `json:"instrument"`
	Units            string           `json:"units"`
	Type             string           `json:"type"`
	PositionFill     string           `json:"positionFill"`
}

//Response represents the Order transaction details returned from Oanda for new Orders
type Response struct {
	LastTransactionID      string            `json:"lastTransactionID"`
	OrderCreateTransaction CreateTransaction `json:"orderCreateTransaction"`
	OrderFillTransaction   FillTransaction   `json:"orderFillTransaction"`
	RelatedTransactionIDs  []string          `json:"relatedTransactionIDs"`
}
type CreateTransaction struct {
	AccountID    string    `json:"accountID"`
	BatchID      string    `json:"batchID"`
	ID           string    `json:"id"`
	Instrument   string    `json:"instrument"`
	PositionFill string    `json:"positionFill"`
	Reason       string    `json:"reason"`
	Time         time.Time `json:"time"`
	TimeInForce  string    `json:"timeInForce"`
	Type         string    `json:"type"`
	Units        string    `json:"units"`
	UserID       int       `json:"userID"`
}
type TradeOpened struct {
	TradeID string `json:"tradeID"`
	Units   string `json:"units"`
}
type FillTransaction struct {
	AccountBalance string      `json:"accountBalance"`
	AccountID      string      `json:"accountID"`
	BatchID        string      `json:"batchID"`
	Financing      string      `json:"financing"`
	ID             string      `json:"id"`
	Instrument     string      `json:"instrument"`
	OrderID        string      `json:"orderID"`
	Pl             string      `json:"pl"`
	Price          string      `json:"price"`
	Reason         string      `json:"reason"`
	Time           time.Time   `json:"time"`
	TradeOpened    TradeOpened `json:"tradeOpened"`
	Type           string      `json:"type"`
	Units          string      `json:"units"`
	UserID         int         `json:"userID"`
}

func CreateOrder(request *Request) (*Response, error) {
	reqBytes, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	body := bytes.NewBuffer(reqBytes)

	reqArgs := &oanda.RequestArgs{
		Method: http.MethodPost,
		URL:    oanda.Host + "/accounts/" + oanda.AccountID + "/orders",
		Body:   body,
	}

	respBytes, err := oanda.MakeRequest(reqArgs)
	if err != nil {
		return nil, err
	}

	resp := &Response{}
	err = json.Unmarshal(respBytes, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
