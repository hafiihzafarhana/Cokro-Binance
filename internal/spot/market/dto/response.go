package dto

type MarketAggregateTradeListResponse struct {
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

type MarketPriceChange24hrResponse struct {
	Symbol             string  `json:"symbol"`
	LastPrice          string  `json:"lastPrice"`
	OpenPrice          string  `json:"openPrice"`
	HighPrice          string  `json:"highPrice"`
	LowPrice           string  `json:"lowPrice"`
	Volume             string  `json:"volume"`
	QuoteVolume        string  `json:"quoteVolume"`
	OpenTime           int64   `json:"openTime"`
	CloseTime          int64   `json:"closeTime"`
	FirstId            int64   `json:"firstId"`
	LastId             int64   `json:"lastId"`
	Count              int64   `json:"count"`
	PriceChange        *string `json:"priceChange,omitempty"`
	PriceChangePercent *string `json:"priceChangePercent,omitempty"`
	WeightedAvgPrice   *string `json:"weightedAvgPrice,omitempty"`
	PrevClosePrice     *string `json:"prevClosePrice,omitempty"`
	LastQty            *string `json:"lastQty,omitempty"`
	BidPrice           *string `json:"bidPrice,omitempty"`
	BidQty             *string `json:"bidQty,omitempty"`
	AskPrice           *string `json:"askPrice,omitempty"`
	AskQty             *string `json:"askQty,omitempty"`
}

type MarketTradingDayTickerResponse struct {
	Symbol             string `json:"symbol"`
	OpenPrice          string `json:"openPrice"`
	HighPrice          string `json:"highPrice"`
	LowPrice           string `json:"lowPrice"`
	LastPrice          string `json:"lastPrice"`
	Volume             string `json:"volume"`
	QuoteVolume        string `json:"quoteVolume"`
	OpenTime           int64 `json:"openTime"`
	OpenTimeStr		   string `json:"openTimeStr"`
	CloseTime          int64 `json:"closeTime"`
	CloseTimeStr	   string `json:"closeTimeStr"`
	FirstId            int64 `json:"firstId"`
	LastId             int64 `json:"lastId"`
	Count              int64 `json:"count"`

	PriceChange        *string `json:"priceChange,omitempty"`
	PriceChangePercent *string `json:"priceChangePercent,omitempty"`
	WeightedAvgPrice   *string `json:"weightedAvgPrice,omitempty"`
}