package oanda

import (
	"bytes"
	"encoding/json"
	"time"
)

/*
Order Endpoints
*/

//OrderPayload represents the Order details
type OrderPayload struct {
	Order Order `json:"order"`
}
type StopLossOnFill struct {
	TimeInForce string `json:"timeInForce"`
	Price       string `json:"price"`
}
type TakeProfitOnFill struct {
	Price string `json:"price"`
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

//OrderResponsePayload represents the Order transaction details sent from Oanda
type OrderResponsePayload struct {
	LastTransactionID      string                 `json:"lastTransactionID"`
	OrderCreateTransaction OrderCreateTransaction `json:"orderCreateTransaction"`
	OrderFillTransaction   OrderFillTransaction   `json:"orderFillTransaction"`
	RelatedTransactionIDs  []string               `json:"relatedTransactionIDs"`
}
type OrderCreateTransaction struct {
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
type OrderFillTransaction struct {
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

func CreateOrder(orderPayload *OrderPayload) (*OrderResponsePayload, error) {
	opBytes, err := json.Marshal(orderPayload)

	if err != nil {
		return nil, err
	}

	body := bytes.NewBuffer(opBytes)

	reqArgs := &ReqArgs{
		ReqMethod: "PUT",
		URL:       oandaHost + "/accounts/" + accountID + "/orders",
		Body:      body,
	}

	orderRespBytes, err := MakeRequest(reqArgs)

	if err != nil {
		return nil, err
	}

	orderResp := &OrderResponsePayload{}
	err = json.Unmarshal(orderRespBytes, orderResp)

	if err != nil {
		return nil, err
	}

	return orderResp, nil
}
