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