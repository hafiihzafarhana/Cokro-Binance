package dto

type GenericSymbolLimitReq struct {
	Symbol string `query:"symbol" validate:"required,uppercase"`
	Limit  int `query:"limit" validate:"gte=100,lte=5000"`
}

type GetOldTradeLookupReq struct {
	GenericSymbolLimitReq
	FromId int `query:"fromId" validate:"omitempty,gte=0"`
}

type GetAgregateTradeListReq struct {
	GenericSymbolLimitReq
	FromId int `query:"fromId" validate:"omitempty,gte=0"`
	StartTime string `query:"startTime" validate:"omitempty"`
	EndTime   string `query:"endTime" validate:"omitempty"`
	TimeZone  string `query:"timeZone" validate:"omitempty"`
}

type GetCandleStickDataReq struct {
	GenericSymbolLimitReq
	Interval string `query:"interval" validate:"required,oneof=1m 3m 5m 15m 30m 1h 2h 4h 6h 8h 12h 1d 3d 1w 1M"`
	StartTime string `query:"startTime" validate:"omitempty"`
	EndTime   string `query:"endTime" validate:"omitempty"`
	TypeKline string `query:"typeKline" validate:"omitempty,oneof=uiklines klines"`
	TimeZone  string `query:"timeZone" validate:"omitempty"`
}

type GetPriceChange24hrParams struct {
	Symbol      string   `query:"symbol" validate:"required,uppercase"`
	Type       string   `query:"type" validate:"omitempty,oneof=FULL MINI"`
	SymbolStatus string  `query:"symbolStatus" validate:"omitempty,oneof=TRADING HALT BREAK"`
}