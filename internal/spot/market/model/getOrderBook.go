package model

type GetOrderBookModel struct {
	LastUpdateID int      `json:"lastUpdateId"`
	Bids         [][]string `json:"bids"`
	Asks         [][]string `json:"asks"`
}
