package entity

// https://developers.binance.com/docs/binance-spot-api-docs/rest-api/market-data-endpoints#recent-trades-list
type MarketRecentTradeListEntity struct {
	Id int `json:"id"`
	Price string `json:"price"`
	Qty string `json:"qty"`
	QuoteQty string `json:"quoteQty"`
	Time int64 `json:"time"`
	DateTime string `json:"dateTime"`
	IsBuyerMaker bool `json:"isBuyerMaker"`
	IsBestMatch bool `json:"isBestMatch"`
}