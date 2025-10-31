package entity

type MarketTickerPrice24hrEntity struct {
	Symbol             string
	LastPrice          string
	OpenPrice          string
	HighPrice          string
	LowPrice           string
	Volume             string
	QuoteVolume        string
	OpenTime           int64
	CloseTime          int64
	FirstId            int64
	LastId             int64
	Count              int64

	// FULL
	PriceChange        *string
	PriceChangePercent *string
	WeightedAvgPrice   *string
	PrevClosePrice     *string
	LastQty            *string
	BidPrice           *string
	BidQty             *string
	AskPrice           *string
	AskQty             *string
}
