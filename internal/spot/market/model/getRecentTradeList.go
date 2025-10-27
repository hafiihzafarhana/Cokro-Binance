package model

type GetRecentTradeListModel struct {
	Id int `json:"id"`
	Price string `json:"price"`
	Qty string `json:"qty"`
	QuoteQty string `json:"quoteQty"`
	Time int64 `json:"time"`
	IsBuyerMaker bool `json:"isBuyerMaker"`
	IsBestMatch bool `json:"isBestMatch"`
}