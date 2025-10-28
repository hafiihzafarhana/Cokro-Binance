package mapper

import (
	"github.com/hafiihzafarhana/Cokro-Binance/domain/spot/market/entity"
	"github.com/hafiihzafarhana/Cokro-Binance/internal/spot/market/model"
	"github.com/hafiihzafarhana/Cokro-Binance/shared/utils/convert"
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