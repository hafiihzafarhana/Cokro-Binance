package dto

type MarketAggregateTradeListRes struct {
	TradeId	  int64  `json:"trade_id"`
	Price     string `json:"price"`
	Quantity  string `json:"quantity"`
	FirstTradeId   int64  `json:"first_trade_id"`
	LastTradeId    int64  `json:"last_trade_id"`
	Timestamp int64  `json:"timestamp"`
	IsBuyerMaker bool `json:"is_buyer_maker"`
	IsBestMatch  bool `json:"is_best_match"`
	DateTime    string `json:"date_time"`
}