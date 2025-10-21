package entity

// https://developers.binance.com/docs/binance-spot-api-docs/rest-api/market-data-endpoints#order-book
type MarketOrderBookEntity struct {
	LastUpdateID int        
	Bids         [][]string
	Asks         [][]string
}