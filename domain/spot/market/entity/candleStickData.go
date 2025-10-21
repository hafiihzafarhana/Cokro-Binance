package entity

// https://developers.binance.com/docs/binance-spot-api-docs/rest-api/market-data-endpoints#klinecandlestick-data
type MarketCandleStickDataEntity struct {
	OpenTime  int64
	OpenTimeStr string
	OpenPrice      string
	HighPrice      string
	LowPrice       string
	ClosePrice     string
	Volume         string
	CloseTime      int64
	CloseTimeStr  string
	QuoteAssetVolume string
	NumberOfTrades   int
	TakerBuyBaseAssetVolume string
	TakerBuyQuoteAssetVolume string
	TypeKline string
}