package entity

// https://developers.binance.com/docs/binance-spot-api-docs/rest-api/market-data-endpoints#recent-trades-list
type MarketRecentTradeListEntity struct {
	Id int
	Price string
	Qty string
	QuoteQty string
	Time int64
	DateTime string
	IsBuyerMaker bool
	IsBestMatch bool
}