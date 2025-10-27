package model

type GetAgregateTradeListModel struct {
	TradeId      int64  `json:"a"`
	Price        string `json:"p"`
	Quantity     string `json:"q"`
	FirstTradeId int64 `json:"f"`
	LastTradeId  int64 `json:"l"`
	Timestamp    int64 `json:"T"`
	IsBuyerMaker bool `json:"m"`
	IsBestMatch  bool `json:"M"`
}