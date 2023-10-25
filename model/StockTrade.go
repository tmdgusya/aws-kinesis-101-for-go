package model

import "encoding/json"

type StockTrade struct {
	TickerSymbol string
	TradeType    string
	Price        float32
	Quantity     int64
	Id           int64
}

func (s *StockTrade) ToJsonAsBytes() []byte {
	bytes, err := json.Marshal(s)

	if err != nil {
		return nil
	}

	return bytes
}

func FromJsonAsByte(bytes []byte) *StockTrade {
	s := new(StockTrade)
	if err := json.Unmarshal(bytes, s); err != nil {
		return nil
	}
	return s
}
