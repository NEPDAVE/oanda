package oanda

import (
	//"log"
	//"strconv"
	//"errors"
	"fmt"
)

//golang playground testing
//https://play.golang.org/p/AC_eqpsQApa

/*
***************************
orders
***************************
*/

//Orders represents entire order(s) object when submiting and order to Oanda
type Orders struct {
	OrderData Order `json:"order"`
}

//Order represents single order to Oanda
type Order struct {
	Price            string           `json:"prices"`
	StopLossOnFill   StopLossOnFill   `json:"stopLossOnFill"`
	TakeProfitOnFill TakeProfitOnFill `json:"takeProfitOnFill"`
	TimeInForce      string           `json:"timeInForce"`
	Instrument       string           `json:"instrument"`
	Units            string           `json:"units"`
	Type             string           `json:"type"`
	PositionFill     string           `json:"positionFill"`
}

//StopLossOnFill represents stop loss parameters for an Order
type StopLossOnFill struct {
	TimeInForce string `json:"timeInForce"`
	Price       string `json:"price"`
}

//TakeProfitOnFill represents take profit parameters for an Order
type TakeProfitOnFill struct {
	TimeInForce string `json:"timeInForce"`
	Price       string `json:"price"`
}

//MarketBuyOrder builds struct needed for marshaling data into a []byte
//FIXME should not be setting SL/TP in this package
func (o Orders) MarketBuyOrder(bid float64, ask float64, instrument string, units string) Orders {
	//tp/sl ratio is 3 to 1
	stopLossPrice := fmt.Sprintf("%.5f", bid-(ask*.005))
	stopLossOnFill := StopLossOnFill{TimeInForce: "GTC", Price: stopLossPrice}

	takeProfitPrice := fmt.Sprintf("%.5f", ask+(ask*.015))
	takeProfitOnFill := TakeProfitOnFill{TimeInForce: "GTC", Price: takeProfitPrice}

	o.OrderData = Order{
		StopLossOnFill:   stopLossOnFill,
		TakeProfitOnFill: takeProfitOnFill,
		TimeInForce:      "FOK",
		Instrument:       instrument,
		Type:             "MARKET",
		PositionFill:     "DEFAULT"}

	return o
}

//MarketSellOrder builds struct needed for marshaling data into a []byte
//FIXME should not be setting SL/TP in this package
func (o Orders) MarketSellOrder(bid float64, ask float64, instrument string, units string) Orders {
	//tp/sl ratio is 3 to 1
	stopLossPrice := fmt.Sprintf("%.5f", ask+(bid*.005))
	stopLossOnFill := StopLossOnFill{TimeInForce: "GTC", Price: stopLossPrice}

	takeProfitPrice := fmt.Sprintf("%.5f", bid-(ask*.015))
	takeProfitOnFill := TakeProfitOnFill{TimeInForce: "GTC", Price: takeProfitPrice}

	o.OrderData = Order{
		StopLossOnFill:   stopLossOnFill,
		TakeProfitOnFill: takeProfitOnFill,
		TimeInForce:      "FOK",
		Instrument:       instrument,
		Type:             "MARKET",
		PositionFill:     "DEFAULT"}

	return o
}
