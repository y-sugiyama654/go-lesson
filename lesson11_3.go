package main

import (
	"fmt"

	"github.com/markcheno/go-quote"
	"github.com/markcheno/go-talib"
)

func main() {
	spy, _ := quote.NewQuoteFromYahoo("spy", "2018-04-01", "2019-01-01", quote.Daily, true)
	fmt.Print(spy.CSV())

	// RSIの算出
	rsi2 := talib.Rsi(spy.Close, 2)
	fmt.Println(rsi2)

	// EMAの算出
	mva := talib.Ema(spy.Close, 14)
	fmt.Println(mva)
}
