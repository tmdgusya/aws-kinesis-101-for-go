package writer

type StockPrice struct {
	TickerSymbol string
	Price        float32
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

func getRandomTrade() {

}
