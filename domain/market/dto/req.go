package dto

import "time"

type GetBinanceOrderBookReq struct {
	Symbol string `query:"symbol" validate:"required,uppercase"`
	Limit  int    `query:"limit" validate:"gte=100,lte=5000"`
}

type GetBinanceRecentTradeListReq struct {
	Symbol string `query:"symbol" validate:"required,uppercase"`
	Limit  int    `query:"limit" validate:"gte=50,lte=1000"`
}

type GetBinanceOldTradeLookupReq struct {
	Symbol string `query:"symbol" validate:"required,uppercase"`
	Limit  int    `query:"limit" validate:"gte=50,lte=1000"`
	FromId int    `query:"fromId" validate:"omitempty,gte=0"`
}

type GetBinanceAgregateTradeListReq struct {
	Symbol string `query:"symbol" validate:"required,uppercase"`
	Limit  int    `query:"limit" validate:"gte=1,lte=1000"`
	FromId int    `query:"fromId" validate:"omitempty,gte=0"`
	StartTime time.Time `query:"startTime" validate:"omitempty"`
	EndTime   time.Time `query:"endTime" validate:"omitempty"`
}