package entity

// https://developers.binance.com/docs/binance-spot-api-docs/rest-api/market-data-endpoints#compressedaggregate-trades-list
type MarketAgregateTradeListEntity struct {
	TradeId      int    
	Price        string
	Quantity     string
	FirstTradeId int
	LastTradeId  int
	Timestamp    int64
	IsBuyerMaker bool
	IsBestMatch  bool
	DateTime     string
}