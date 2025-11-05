package entity

type MarketTradingDayTickerEntity struct {
	Symbol             string
	OpenPrice          string
	HighPrice          string
	LowPrice           string
	LastPrice          string
	Volume             string
	QuoteVolume        string
	OpenTime           int64
	OpenTimeStr        string
	CloseTime          int64
	CloseTimeStr       string
	FirstId            int64
	LastId             int64
	Count              int64

	PriceChange        *string
	PriceChangePercent *string
	WeightedAvgPrice   *string
}