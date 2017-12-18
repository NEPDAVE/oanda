package oanda

import (
	"encoding/json"
	"time"
)

/*
***************************
order
***************************
*/

type OrderBody struct {
	Order Order `json:"order"`
}

type Order struct {
	Price            float64          `json:"prices"`
	StopLossOnFill   StopLossOnFill   `json:"stopLossOnFill"`
	TakeProfitOnFill TakeProfitOnFill `json:"takeProfitOnFill"`
	TimeInForce      string           `json:"timeInForce"`
	Instrument       string           `json:"instrument"`
	Units            int              `json:"units"`
	Type             string           `json:"type"`
	PositionFill     string           `json:"positionFill"`
}

type StopLossOnFill struct {
	TimeInForce string  `json:"timeInForce"`
	Price       float64 `json:"price"`
}

type TakeProfitOnFill struct {
	TimeInForce string  `json:"timeInForce"`
	Price       float64 `json:"price"`
}
