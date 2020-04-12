package oanda

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

type OrderPayload struct {
	Order Order `json:"order"`
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

type StopLossOnFill struct {
	TimeInForce string `json:"timeInForce"`
	Price       string `json:"price"`
}

type TakeProfitOnFill struct {
	Price string `json:"price"`
}

//OrderResponse
type OrderResponse struct {
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

//TODO this should take a struct?
/*
	newOrder(struct)
	marshalOrder(struct) []byte
	createOrder([] byte)

*/

func NewOrder(op *OrderPayload) (*OrderResponse, error) {
	opByte, err := json.Marshal(op)

	if err != nil {
		return nil, err
	}

	orderByte, err := CreateOrder(opByte)

	if err != nil {
		return nil, err
	}

	orderResp := &OrderResponse{}

	err = json.Unmarshal(orderByte, orderResp)

	if err != nil {
		return nil, err
	}

	return orderResp, nil
}

func CreateOrder(orders []byte) ([]byte, error) {
	body := bytes.NewBuffer(orders)
	req, err := http.NewRequest("POST", oandaURL+"/accounts/"+accountID+"/orders", body)

	req.Header.Set("Authorization", bearer)
	req.Header.Set("content-type", "application/json")

	if err != nil {
		return []byte{}, err
	}

	resp, err := client.Do(req)

	if err != nil {
		return []byte{}, err
	}

	defer resp.Body.Close()

	createOrderByte, _ := ioutil.ReadAll(resp.Body)

	if err != nil {
		return []byte{}, err
	}

	return createOrderByte, err
}
