package dto

type GetBinanceOrderBookReq struct {
	Symbol string `query:"symbol" validate:"required,uppercase"`
	Limit  int    `query:"limit" validate:"gte=100,lte=5000"`
}

type GetBinanceRecentTradeListReq struct {
	Symbol string `query:"symbol" validate:"required,uppercase"`
	Limit  int    `query:"limit" validate:"gte=50,lte=1000"`
}