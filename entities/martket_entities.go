package entities

// https://developers.binance.com/docs/binance-spot-api-docs/rest-api/market-data-endpoints#order-book
type MarketOrderBookEntity struct {
	LastUpdateID int        `json:"lastUpdateId"`
	Bids         [][]string `json:"bids"`
	Asks         [][]string `json:"asks"`
}

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

// https://developers.binance.com/docs/binance-spot-api-docs/rest-api/market-data-endpoints#old-trade-lookup
type MarketOldTradeLookupEntity struct {
	Id int `json:"id"`
	Price string `json:"price"`
	Qty string `json:"qty"`
	QuoteQty string `json:"quoteQty"`
	Time int64 `json:"time"`
	DateTime string `json:"dateTime"`
	IsBuyerMaker bool `json:"isBuyerMaker"`
	IsBestMatch bool `json:"isBestMatch"`
}

// https://developers.binance.com/docs/binance-spot-api-docs/rest-api/market-data-endpoints#compressedaggregate-trades-list
type MarketAgregateTradeListEntity struct {
	TradeId      int    `json:"a"`
	Price        string `json:"p"`
	Quantity     string `json:"q"`
	FirstTradeId int    `json:"f"`
	LastTradeId  int    `json:"l"`
	Timestamp    int64  `json:"T"`
	IsBuyerMaker bool   `json:"m"`
	IsBestMatch  bool   `json:"M"`
	DateTime     string `json:"-"`
}

// https://developers.binance.com/docs/binance-spot-api-docs/rest-api/market-data-endpoints#klinecandlestick-data
type MarketCandleStickDataEntity struct {
	OpenTime  int64  `json:"openTime"`
	OpenTimeStr string `json:"openTimeStr"`
	OpenPrice      string `json:"openPrice"`
	HighPrice      string `json:"highPrice"`
	LowPrice       string `json:"lowPrice"`
	ClosePrice     string `json:"closePrice"`
	Volume         string `json:"volume"`
	CloseTime      int64  `json:"closeTime"`
	CloseTimeStr  string `json:"closeTimeStr"`
	QuoteAssetVolume string `json:"quoteAssetVolume"`
	NumberOfTrades   int    `json:"numberOfTrades"`
	TakerBuyBaseAssetVolume string `json:"takerBuyBaseAssetVolume"`
	TakerBuyQuoteAssetVolume string `json:"takerBuyQuoteAssetVolume"`
	TypeKline string `json:"typeKline"` // klines or uiklines
}