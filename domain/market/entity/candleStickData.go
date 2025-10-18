package entity

import "time"

// https://developers.binance.com/docs/binance-spot-api-docs/rest-api/market-data-endpoints#klinecandlestick-data
type MarketCandleStickDataEntity struct {
	OpenTime  int64  `json:"openTime"`
	OpenTimeStr string `json:"openTimeStr"`
	OpenPrice      string `json:"openPrice"`
	HighPrice      string `json:"highPrice"`
	LowPrice       string `json:"lowPrice"`
	ClosePrice     string `json:"closePrice"`
	Volume         string `json:"volume"`
	CloseTime      int64  `json:"closeTime"`
	CloseTimeStr  string `json:"closeTimeStr"`
	QuoteAssetVolume string `json:"quoteAssetVolume"`
	NumberOfTrades   int    `json:"numberOfTrades"`
	TakerBuyBaseAssetVolume string `json:"takerBuyBaseAssetVolume"`
	TakerBuyQuoteAssetVolume string `json:"takerBuyQuoteAssetVolume"`
	TypeKline string `json:"typeKline"` // klines or uiklines
}

func MapMarketCandleStickData(rawData [][]interface{}, typeKline string) []*MarketCandleStickDataEntity {
	var candles []*MarketCandleStickDataEntity

	for _, item := range rawData {
		if len(item) < 11 {
			continue
		}

		openTime, _ := item[0].(float64)
		openPrice, _ := item[1].(string)
		highPrice, _ := item[2].(string)
		lowPrice, _ := item[3].(string)
		closePrice, _ := item[4].(string)
		volume, _ := item[5].(string)
		closeTime, _ := item[6].(float64)
		quoteAssetVolume, _ := item[7].(string)
		numberOfTrades, _ := item[8].(float64)
		takerBuyBaseAssetVolume, _ := item[9].(string)
		takerBuyQuoteAssetVolume, _ := item[10].(string)

		candle := &MarketCandleStickDataEntity{
			OpenTime:                int64(openTime),
			OpenTimeStr:             time.UnixMilli(int64(openTime)).Local().Format("2006-01-02 15:04:05"),
			OpenPrice:               openPrice,
			HighPrice:               highPrice,
			LowPrice:                lowPrice,
			ClosePrice:              closePrice,
			Volume:                  volume,
			CloseTime:               int64(closeTime),
			CloseTimeStr:            time.UnixMilli(int64(closeTime)).Local().Format("2006-01-02 15:04:05"),
			QuoteAssetVolume:        quoteAssetVolume,
			NumberOfTrades:          int(numberOfTrades),
			TakerBuyBaseAssetVolume: takerBuyBaseAssetVolume,
			TakerBuyQuoteAssetVolume: takerBuyQuoteAssetVolume,
			TypeKline:               typeKline,
		}

		candles = append(candles, candle)
	}

	return candles
}