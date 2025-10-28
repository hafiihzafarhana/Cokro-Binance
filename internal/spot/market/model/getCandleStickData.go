package model

type GetCandleStickDataModel struct {
	OpenTime  int64 `json:"openTime"`
	OpenPrice      string `json:"openPrice"`
	HighPrice      string `json:"highPrice"`
	LowPrice       string `json:"lowPrice"`
	ClosePrice     string `json:"closePrice"`
	Volume         string `json:"volume"`
	CloseTime      int64  `json:"closeTime"`
	QuoteAssetVolume string `json:"quoteAssetVolume"`
	NumberOfTrades   int    `json:"numberOfTrades"`
	TakerBuyBaseAssetVolume string `json:"takerBuyBaseAssetVolume"`
	TakerBuyQuoteAssetVolume string `json:"takerBuyQuoteAssetVolume"`
}