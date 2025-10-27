package entity

// https://developers.binance.com/docs/binance-spot-api-docs/rest-api/market-data-endpoints#compressedaggregate-trades-list
type MarketAgregateTradeListEntity struct {
	TradeId      int64    
	Price        string
	Quantity     string
	FirstTradeId int64
	LastTradeId  int64
	Timestamp    int64
	IsBuyerMaker bool
	IsBestMatch  bool
	DateTime     string
}