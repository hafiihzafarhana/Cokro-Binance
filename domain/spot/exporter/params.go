package exporter

type GenericSymbolLimitParams struct {
	Symbol string
	Limit  int
}

type GetCandleStickDataParams struct {
	GenericSymbolLimitParams
	Interval string
	StartTime string
	EndTime   string
	TypeKline string
	TimeZone  string
}