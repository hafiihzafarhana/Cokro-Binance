package dto

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
	StartTime string `query:"startTime" validate:"omitempty"`
	EndTime   string `query:"endTime" validate:"omitempty"`
}

type GetBinanceCandleStickDataReq struct {
	Symbol  string    `query:"symbol" validate:"required,uppercase"`
	Interval string    `query:"interval" validate:"required,oneof=1m 3m 5m 15m 30m 1h 2h 4h 6h 8h 12h 1d 3d 1w 1M"`
	StartTime string `query:"startTime" validate:"omitempty"`
	EndTime   string `query:"endTime" validate:"omitempty"`
	TimeZone  string    `query:"timeZone" validate:"omitempty"`
	Limit   int       `query:"limit" validate:"gte=1,lte=1000"`
	TypeKline string    `query:"typeKline" validate:"omitempty,oneof=uiklines klines"`
}