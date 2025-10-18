package entity

// https://developers.binance.com/docs/binance-spot-api-docs/rest-api/market-data-endpoints#order-book
type MarketOrderBookEntity struct {
	LastUpdateID int        `json:"lastUpdateId"`
	Bids         [][]string `json:"bids"`
	Asks         [][]string `json:"asks"`
}