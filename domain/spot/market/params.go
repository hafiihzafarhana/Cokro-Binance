package market

type GenericSymbolLimitParams struct {
	Symbol string
	Limit  int
}

type GetOldTradeLookupParams struct {
	GenericSymbolLimitParams
	FromId int
}

type GetAgregateTradeListParams struct {
	GenericSymbolLimitParams
	FromId int
	StartTime string
	EndTime   string
	TimeZone  string
}

type GetCandleStickDataParams struct {
	GenericSymbolLimitParams
	Interval string
	StartTime string
	EndTime   string
	TypeKline string
	TimeZone  string
}

type GetPriceChange24hrParams struct {
	Symbol string
	Type string
	SymbolStatus string
}