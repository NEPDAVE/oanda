package oanda

import (
	//"log"
	//"strconv"
	//"errors"
	"fmt"
)

//golang playground testing
//https://play.golang.org/p/AC_eqpsQApa

//MarketBuyOrder builds struct needed for marshaling data into a []byte
//FIXME stopLoss and takeProfit are hard coded? make this ratio 3 to 1
func MarketBuyOrder(bid float64, ask float64, instrument string, units int) Orders {
	stopLossPrice := fmt.Sprintf("%.6f", bid-(ask*.005))
	takeProfitPrice := fmt.Sprintf("%.6f", ask+(ask*.01))
	stopLoss := StopLossOnFill{TimeInForce: "GTC", Price: stopLossPrice}
	takeProfit := TakeProfitOnFill{TimeInForce: "GTC", Price: takeProfitPrice}
	orderData := Order{
		StopLossOnFill:   stopLoss,
		TakeProfitOnFill: takeProfit,
		TimeInForce:      "FOK",
		Instrument:       instrument,
		Type:             "MARKET",
		PositionFill:     "DEFAULT"}
	order := Orders{Order: orderData}

	return order
}

//MarketSellOrder builds struct needed for marshaling data into a []byte
//FIXME stopLoss and takeProfit are hard coded? make this ratio 3 to 1
func MarketSellOrder(bid float64, ask float64, instrument string, units int) Orders {
	stopLossPrice := fmt.Sprintf("%.6f", ask+(bid*.005))
	takeProfitPrice := fmt.Sprintf("%.6f", bid-(ask*.01))
	stopLoss := StopLossOnFill{TimeInForce: "GTC", Price: stopLossPrice}
	takeProfit := TakeProfitOnFill{TimeInForce: "GTC", Price: takeProfitPrice}
	orderData := Order{
		StopLossOnFill:   stopLoss,
		TakeProfitOnFill: takeProfit,
		TimeInForce:      "FOK",
		Instrument:       instrument,
		Type:             "MARKET",
		PositionFill:     "DEFAULT"}
	order := Orders{Order: orderData}

	return order
}
