package dto

import (
	"github.com/hafiihzafarhana/Cokro-Binance/domain/spot/market/entity"
	stringdata "github.com/hafiihzafarhana/Cokro-Binance/shared/utils/stringData"
)

func ToMarketPriceChange24hrResponse(e *entity.MarketTickerPrice24hrEntity) *MarketPriceChange24hrResponse {
	return &MarketPriceChange24hrResponse{
		Symbol:             e.Symbol,
		LastPrice:          e.LastPrice,
		OpenPrice:          e.OpenPrice,
		HighPrice:          e.HighPrice,
		LowPrice:           e.LowPrice,
		Volume:             e.Volume,
		QuoteVolume:        e.QuoteVolume,
		OpenTime:           e.OpenTime,
		CloseTime:          e.CloseTime,
		FirstId:            e.FirstId,
		LastId:             e.LastId,
		Count:              e.Count,
		PriceChange:        stringdata.ToNilIfEmpty(e.PriceChange),
		PriceChangePercent: stringdata.ToNilIfEmpty(e.PriceChangePercent),
		WeightedAvgPrice:   stringdata.ToNilIfEmpty(e.WeightedAvgPrice),
		PrevClosePrice:     stringdata.ToNilIfEmpty(e.PrevClosePrice),
		LastQty:            stringdata.ToNilIfEmpty(e.LastQty),
		BidPrice:           stringdata.ToNilIfEmpty(e.BidPrice),
		BidQty:             stringdata.ToNilIfEmpty(e.BidQty),
		AskPrice:           stringdata.ToNilIfEmpty(e.AskPrice),
		AskQty:             stringdata.ToNilIfEmpty(e.AskQty),
	}
}

func ToMarketTickerPrice24hrResponseList(list []*entity.MarketTickerPrice24hrEntity) []*MarketPriceChange24hrResponse {
	resp := make([]*MarketPriceChange24hrResponse, 0, len(list))
	for _, e := range list {
		resp = append(resp, ToMarketPriceChange24hrResponse(e))
	}
	return resp
}

func ToMarketTradingDayTickerResponse(e *entity.MarketTradingDayTickerEntity) *MarketTradingDayTickerResponse {
	return &MarketTradingDayTickerResponse{
		Symbol:             e.Symbol,
		LastPrice:          e.LastPrice,
		OpenPrice:          e.OpenPrice,
		HighPrice:          e.HighPrice,
		LowPrice:           e.LowPrice,
		Volume:             e.Volume,
		QuoteVolume:        e.QuoteVolume,
		OpenTime:           e.OpenTime,
		OpenTimeStr: 		e.OpenTimeStr,	
		CloseTime:          e.CloseTime,
		CloseTimeStr: 		e.CloseTimeStr,
		FirstId:            e.FirstId,
		LastId:             e.LastId,
		Count:              e.Count,
		PriceChange:        stringdata.ToNilIfEmpty(e.PriceChange),
		PriceChangePercent: stringdata.ToNilIfEmpty(e.PriceChangePercent),
		WeightedAvgPrice:   stringdata.ToNilIfEmpty(e.WeightedAvgPrice),
	}
}

func ToMarketTradingDayTickerResponseList(list []*entity.MarketTradingDayTickerEntity) []*MarketTradingDayTickerResponse {
	resp := make([]*MarketTradingDayTickerResponse, 0, len(list))
	for _, e := range list {
		resp = append(resp, ToMarketTradingDayTickerResponse(e))
	}
	return resp
}