package main

import (
	"fmt"
)

func main(){
  pricesByte := GetPricing("EUR_USD")
  fmt.Println("MAIN FUNCTION LINE 9")
  fmt.Println(pricesByte)
  prices := Prices{}.UnmarshalPricing(pricesByte)
  fmt.Println("MAIN FUNCTION LINE 10")
  fmt.Println(prices.Status)
}
