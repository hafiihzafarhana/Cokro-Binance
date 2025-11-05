package mapper

import (
	"fmt"

	"github.com/hafiihzafarhana/Cokro-Binance/domain/spot/market/entity"
	"github.com/hafiihzafarhana/Cokro-Binance/internal/spot/market/model"
	"github.com/hafiihzafarhana/Cokro-Binance/shared/utils/convert"
	stringdata "github.com/hafiihzafarhana/Cokro-Binance/shared/utils/stringData"
)

func ToMarketOrderBookEntity(e model.GetOrderBookModel) *entity.MarketOrderBookEntity {
	return &entity.MarketOrderBookEntity{
		LastUpdateID: e.LastUpdateID,
		Bids:         e.Bids,
		Asks:         e.Asks,
	}
}

func ToMarketRecentTradeListEntity(e model.GetRecentTradeListModel) *entity.MarketRecentTradeListEntity {
	return &entity.MarketRecentTradeListEntity{
		Id:           e.Id,
		Price:        e.Price,
		Qty:          e.Qty,
		QuoteQty:     e.QuoteQty,
		Time:         e.Time,
		DateTime:     convert.UnixToDateTimeString(e.Time, "Asia/Jakarta"),
		IsBuyerMaker: e.IsBuyerMaker,
		IsBestMatch:  e.IsBestMatch,
	}
}

func ToMarketOldTradeLookupEntity(e model.GetOldTradeLookupModel) *entity.MarketOldTradeLookupEntity {
	return &entity.MarketOldTradeLookupEntity{
		Id: e.Id,
		Price: e.Price,
		Qty: e.Qty,
		QuoteQty: e.QuoteQty,
		Time: e.Time,
		DateTime: convert.UnixToDateTimeString(e.Time, "Asia/Jakarta"),
		IsBuyerMaker: e.IsBuyerMaker,
		IsBestMatch:  e.IsBestMatch,
	}
}

func ToMarketAgregateTradeListEntity(e model.GetAgregateTradeListModel) *entity.MarketAgregateTradeListEntity {
	return &entity.MarketAgregateTradeListEntity{
		TradeId: e.TradeId,
		Price: e.Price,
		Quantity: e.Quantity,
		Timestamp: e.Timestamp,
		DateTime: convert.UnixToDateTimeString(e.Timestamp, "Asia/Jakarta"),
		IsBuyerMaker: e.IsBuyerMaker,
		IsBestMatch:  e.IsBestMatch,
		FirstTradeId: e.FirstTradeId,
		LastTradeId: e.LastTradeId,
	}
}

func ToMarketCandleStickDataEntity(e []interface{}, endpointType string) *entity.MarketCandleStickDataEntity {
	return &entity.MarketCandleStickDataEntity{
		OpenTime:                 int64(e[0].(float64)),
		OpenTimeStr:              convert.UnixToDateTimeString(int64(e[0].(float64)), "Asia/Jakarta"),
		OpenPrice:                e[1].(string),
		HighPrice:                e[2].(string),
		LowPrice:                 e[3].(string),
		ClosePrice:               e[4].(string),
		Volume:                   e[5].(string),
		CloseTime:                int64(e[6].(float64)),
		CloseTimeStr:             convert.UnixToDateTimeString(int64(e[6].(float64)), "Asia/Jakarta"),
		QuoteAssetVolume:         e[7].(string),
		NumberOfTrades:           int(e[8].(float64)),
		TakerBuyBaseAssetVolume:  e[9].(string),
		TakerBuyQuoteAssetVolume: e[10].(string),
		TypeKline:                endpointType,
	}
}

func ToMarketCurrentAveragePrice(e model.GetCurrentAveragePriceModel) *entity.MarketCurrentAveragePriceEntity {
	return &entity.MarketCurrentAveragePriceEntity{
		IntervalInMinute: e.IntervalInMinute,
		Price: e.Price,
		CloseTime: e.CloseTime,
		CloseTimeStr: convert.UnixToDateTimeString(e.CloseTime, "Asia/Jakarta"),
	}
}

func ToMarketTickerPrice24hrEntity(m model.GetTickerPrice24hrModel) *entity.MarketTickerPrice24hrEntity {
	return &entity.MarketTickerPrice24hrEntity{
		Symbol:      m.Symbol,
		LastPrice:   m.LastPrice,
		OpenPrice:   m.OpenPrice,
		HighPrice:   m.HighPrice,
		LowPrice:    m.LowPrice,
		Volume:      m.Volume,
		QuoteVolume: m.QuoteVolume,
		OpenTime:    m.OpenTime,
		CloseTime:   m.CloseTime,
		FirstId:     m.FirstId,
		LastId:      m.LastId,
		Count:       m.Count,

		PriceChange:        stringdata.StringPtr(m.PriceChange),
		PriceChangePercent: stringdata.StringPtr(m.PriceChangePercent),
		WeightedAvgPrice:   stringdata.StringPtr(m.WeightedAvgPrice),
		PrevClosePrice:     stringdata.StringPtr(m.PrevClosePrice),
		LastQty:            stringdata.StringPtr(m.LastQty),
		BidPrice:           stringdata.StringPtr(m.BidPrice),
		BidQty:             stringdata.StringPtr(m.BidQty),
		AskPrice:           stringdata.StringPtr(m.AskPrice),
		AskQty:             stringdata.StringPtr(m.AskQty),
	}
}

func ToMarketTradingDayTickerEntity(m model.GetTradingDayTickerModel) *entity.MarketTradingDayTickerEntity {
	fmt.Println(convert.UnixToDateTimeString(m.CloseTime, "Asia/Jakarta"),)
	return &entity.MarketTradingDayTickerEntity{
		Symbol:      m.Symbol,
		LastPrice:   m.LastPrice,
		OpenPrice:   m.OpenPrice,
		HighPrice:   m.HighPrice,
		LowPrice:    m.LowPrice,
		Volume:      m.Volume,
		QuoteVolume: m.QuoteVolume,
		OpenTime:    m.OpenTime,
		OpenTimeStr: convert.UnixToDateTimeString(m.OpenTime, "Asia/Jakarta"),
		CloseTime:   m.CloseTime,
		CloseTimeStr: convert.UnixToDateTimeString(m.CloseTime, "Asia/Jakarta"),
		FirstId:     m.FirstId,
		LastId:      m.LastId,
		Count:       m.Count,

		PriceChange:        stringdata.StringPtr(m.PriceChange),
		PriceChangePercent: stringdata.StringPtr(m.PriceChangePercent),
		WeightedAvgPrice:   stringdata.StringPtr(m.WeightedAvgPrice),
	}
}