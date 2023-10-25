package writer

import (
	"kinesis-for-go/model"
	"math"
	"math/rand"
)

var (
	MAX_DEVIATION          = 0.2
	MAX_QUANTITY     int64 = 10000
	PROBABILITY_SELL       = 0.4
)

type StockPrice struct {
	TickerSymbol string
	Price        float64
}

func createStockPrices() []StockPrice {
	return []StockPrice{
		{TickerSymbol: "AAPL", Price: 119.72},
		{TickerSymbol: "XOM", Price: 91.56},
		{TickerSymbol: "GOOG", Price: 527.83},
		{TickerSymbol: "BRK.A", Price: 223999.88},
		{TickerSymbol: "MSFT", Price: 42.36},
		{TickerSymbol: "WFC", Price: 54.21},
		{TickerSymbol: "JNJ", Price: 99.78},
		{TickerSymbol: "WMT", Price: 85.91},
	}
}

// GetRandomTrade do randomly buying or selling stock
func GetRandomTrade() model.StockTrade {
	prices := createStockPrices()
	stockPrice := prices[rand.Intn(len(prices))]
	deviation := (rand.Float64() - 0.5) * 2.0 * MAX_DEVIATION
	price := stockPrice.Price * (1 + deviation)
	price = math.Round(price*100) / 100.0

	tradeType := "BUY"

	if rand.Float64() < PROBABILITY_SELL {
		tradeType = "SELL"
	}

	return model.StockTrade{
		TickerSymbol: stockPrice.TickerSymbol,
		TradeType:    tradeType,
		Price:        price,
		Quantity:     rand.Int63n(MAX_QUANTITY) + 1,
		Id:           1,
	}
}
