package dto

import (
	"github.com/hafiihzafarhana/Cokro-Binance/entities"
)

type GetBinanceAggregateTradeRes struct {
	TradeId      int    `json:"tradeId"`
	Price        string `json:"price"`
	Quantity     string `json:"quantity"`
	FirstTradeId int    `json:"firstTradeId"`
	LastTradeId  int    `json:"lastTradeId"`
	Timestamp    int64  `json:"timestamp"`
	DateTime     string `json:"dateTime"`
	IsBuyerMaker bool   `json:"isBuyerMaker"`
	IsBestMatch  bool   `json:"isBestMatch"`
}

func GetBinanceAggregateTradeResponse(users []*entities.MarketAgregateTradeListEntity) []*GetBinanceAggregateTradeRes {
	res := make([]*GetBinanceAggregateTradeRes, len(users))

	for i, u := range users {
		res[i] = &GetBinanceAggregateTradeRes{
			TradeId:      u.TradeId,
			Price:        u.Price,
			Quantity:     u.Quantity,
			FirstTradeId: u.FirstTradeId,
			LastTradeId:  u.LastTradeId,
			Timestamp:    u.Timestamp,
			DateTime:     u.DateTime,
			IsBuyerMaker: u.IsBuyerMaker,
			IsBestMatch:  u.IsBestMatch,
		}
	}

	return res
}

