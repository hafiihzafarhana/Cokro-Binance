package mapper

import (
	"reflect"
	"testing"

	"github.com/hafiihzafarhana/Cokro-Binance/domain/spot/market/entity"
	"github.com/hafiihzafarhana/Cokro-Binance/internal/spot/market/model"
	"github.com/hafiihzafarhana/Cokro-Binance/shared/utils/convert"
	stringdata "github.com/hafiihzafarhana/Cokro-Binance/shared/utils/stringData"
	"github.com/stretchr/testify/assert"
)

func TestToMarketOrderBookEntity(t *testing.T){
	input := model.GetOrderBookModel{
		LastUpdateID: 12345,
		Bids: [][]string{
			{"50000.00", "0.5"},
			{"49900.00", "1.2"},
		},
		Asks: [][]string{
			{"50100.00", "0.4"},
			{"50200.00", "1.1"},
		},
	}

	expected := &entity.MarketOrderBookEntity{
		LastUpdateID: 12345,
		Bids: [][]string{
			{"50000.00", "0.5"},
			{"49900.00", "1.2"},
		},
		Asks: [][]string{
			{"50100.00", "0.4"},
			{"50200.00", "1.1"},
		},
	}

	result := ToMarketOrderBookEntity(input)

	if result.LastUpdateID != expected.LastUpdateID {
		t.Errorf("expected LastUpdateID %d, got %d", expected.LastUpdateID, result.LastUpdateID)
	}

	// reflect.DeepEqual for compare 2 slice
	if !reflect.DeepEqual(result.Bids, expected.Bids) {
		t.Errorf("expected Bids %v, got %v", expected.Bids, result.Bids)
	}

	if !reflect.DeepEqual(result.Asks, expected.Asks) {
		t.Errorf("expected Asks %v, got %v", expected.Asks, result.Asks)
	}
}

func TestToMarketRecentTradeListEntity(t *testing.T){
	input := model.GetRecentTradeListModel{
		Id: 28457,
		Price: "4.00000100",
		Qty: "12.00000000",
		QuoteQty: "48.000012",
		Time: 1499865549590,
		IsBuyerMaker: true,
		IsBestMatch: true,
	}

	expected := &entity.MarketRecentTradeListEntity {
		Id: 28457,
		Price: "4.00000100",
		Qty: "12.00000000",
		QuoteQty: "48.000012",
		Time: 1499865549590,
		IsBuyerMaker: true,
		IsBestMatch: true,
		DateTime: convert.UnixToDateTimeString(input.Time, "Asia/Jakarta"),
	}

	result := ToMarketRecentTradeListEntity(input)
	assert.Equal(t, expected.Id, result.Id)
	assert.Equal(t, expected.Price, result.Price)
	assert.Equal(t, expected.Qty, result.Qty)
	assert.Equal(t, expected.QuoteQty, result.QuoteQty)
	assert.Equal(t, expected.Time, result.Time)
	assert.Equal(t, expected.IsBuyerMaker, result.IsBuyerMaker)
	assert.Equal(t, expected.IsBestMatch, result.IsBestMatch)
	assert.Equal(t, expected.DateTime, result.DateTime)
}

func TestToMarketOldTradeLookupEntity(t *testing.T){
	input := model.GetOldTradeLookupModel{
		Id: 28457,
		Price: "4.00000100",
		Qty: "12.00000000",
		QuoteQty: "48.000012",
		Time: 1499865549590,
		IsBuyerMaker: true,
		IsBestMatch: true,
	}

	expected := &entity.MarketOldTradeLookupEntity {
		Id: 28457,
		Price: "4.00000100",
		Qty: "12.00000000",
		QuoteQty: "48.000012",
		Time: 1499865549590,
		IsBuyerMaker: true,
		IsBestMatch: true,
		DateTime: convert.UnixToDateTimeString(input.Time, "Asia/Jakarta"),
	}

	result := ToMarketOldTradeLookupEntity(input)
	assert.Equal(t, expected.Id, result.Id)
	assert.Equal(t, expected.Price, result.Price)
	assert.Equal(t, expected.Qty, result.Qty)
	assert.Equal(t, expected.QuoteQty, result.QuoteQty)
	assert.Equal(t, expected.Time, result.Time)
	assert.Equal(t, expected.IsBuyerMaker, result.IsBuyerMaker)
	assert.Equal(t, expected.IsBestMatch, result.IsBestMatch)
	assert.Equal(t, expected.DateTime, result.DateTime)
}

func TestToMarketAgregateTradeListEntity(t *testing.T){
	input := model.GetAgregateTradeListModel{
		TradeId: 26129,
		Price: "0.01633102",
		Quantity: "4.70443515",
		Timestamp: 1498793709153,
		IsBuyerMaker: true,
		IsBestMatch:  true,
		FirstTradeId: 27781,
		LastTradeId: 27781,
	}

	expected := &entity.MarketAgregateTradeListEntity {
		TradeId: 26129,
		Price: "0.01633102",
		Quantity: "4.70443515",
		Timestamp: 1498793709153,
		IsBuyerMaker: true,
		IsBestMatch:  true,
		FirstTradeId: 27781,
		LastTradeId: 27781,
		DateTime: convert.UnixToDateTimeString(input.Timestamp, "Asia/Jakarta"),
	}

	result := ToMarketAgregateTradeListEntity(input)
	assert.Equal(t, expected.TradeId, result.TradeId)
	assert.Equal(t, expected.Price, result.Price)
	assert.Equal(t, expected.Quantity, result.Quantity)
	assert.Equal(t, expected.Timestamp, result.Timestamp)
	assert.Equal(t, expected.FirstTradeId, result.FirstTradeId)
	assert.Equal(t, expected.LastTradeId, result.LastTradeId)
	assert.Equal(t, expected.IsBuyerMaker, result.IsBuyerMaker)
	assert.Equal(t, expected.IsBestMatch, result.IsBestMatch)
	assert.Equal(t, expected.DateTime, result.DateTime)
}

func TestToMarketCandleStickDataEntity(t *testing.T){
	openTime := float64(1609459200000) // 2021-01-01 00:00:00 UTC
	closeTime := float64(1609462800000) // 2021-01-01 01:00:00 UTC
	input := []interface{}{
		openTime,           // 0 OpenTime
		"40000.00",         // 1 OpenPrice
		"40500.00",         // 2 HighPrice
		"39500.00",         // 3 LowPrice
		"40200.00",         // 4 ClosePrice
		"123.45",           // 5 Volume
		closeTime,          // 6 CloseTime
		"5000000.00",       // 7 QuoteAssetVolume
		float64(150),       // 8 NumberOfTrades
		"50.00",            // 9 TakerBuyBaseAssetVolume
		"2000000.00",       // 10 TakerBuyQuoteAssetVolume
	}
	endpointType := "1h"

	expected := &entity.MarketCandleStickDataEntity{
		OpenTime:                 int64(openTime),
		OpenTimeStr:              convert.UnixToDateTimeString(int64(openTime), "Asia/Jakarta"),
		OpenPrice:                "40000.00",
		HighPrice:                "40500.00",
		LowPrice:                 "39500.00",
		ClosePrice:               "40200.00",
		Volume:                   "123.45",
		CloseTime:                int64(closeTime),
		CloseTimeStr:             convert.UnixToDateTimeString(int64(closeTime), "Asia/Jakarta"),
		QuoteAssetVolume:         "5000000.00",
		NumberOfTrades:           150,
		TakerBuyBaseAssetVolume:  "50.00",
		TakerBuyQuoteAssetVolume: "2000000.00",
		TypeKline:                "1h",
	}

	result := ToMarketCandleStickDataEntity(input, endpointType)
	assert.Equal(t, expected.OpenTime, result.OpenTime)
	assert.Equal(t, expected.OpenTimeStr, result.OpenTimeStr)
	assert.Equal(t, expected.OpenPrice, result.OpenPrice)
	assert.Equal(t, expected.HighPrice, result.HighPrice)
	assert.Equal(t, expected.LowPrice, result.LowPrice)
	assert.Equal(t, expected.ClosePrice, result.ClosePrice)
	assert.Equal(t, expected.Volume, result.Volume)
	assert.Equal(t, expected.CloseTime, result.CloseTime)
	assert.Equal(t, expected.CloseTimeStr, result.CloseTimeStr)
	assert.Equal(t, expected.QuoteAssetVolume, result.QuoteAssetVolume)
	assert.Equal(t, expected.NumberOfTrades, result.NumberOfTrades)
	assert.Equal(t, expected.TakerBuyBaseAssetVolume, result.TakerBuyBaseAssetVolume)
	assert.Equal(t, expected.TakerBuyQuoteAssetVolume, result.TakerBuyQuoteAssetVolume)
	assert.Equal(t, expected.TypeKline, result.TypeKline)
}

func TestToMarketCurrentAveragePrice(t *testing.T){
	input := model.GetCurrentAveragePriceModel{
		Price: "0.01633102",
		IntervalInMinute: 5,
		CloseTime: 1499865549590,
	}

	expected := &entity.MarketCurrentAveragePriceEntity {
		IntervalInMinute: 5,
		Price: "0.01633102",
		CloseTime: 1499865549590,
		CloseTimeStr: convert.UnixToDateTimeString(input.CloseTime, "Asia/Jakarta"),
	}

	result := ToMarketCurrentAveragePrice(input)
	assert.Equal(t, expected.CloseTime, result.CloseTime)
	assert.Equal(t, expected.CloseTimeStr, result.CloseTimeStr)
	assert.Equal(t, expected.IntervalInMinute, result.IntervalInMinute)
	assert.Equal(t, expected.Price, result.Price)
}

func TestToMarketTickerPrice24hrEntity(t *testing.T) {
	input := model.GetTickerPrice24hrModel{
		Symbol:            "BTCUSDT",
		PriceChange:       "500.00",
		PriceChangePercent:"1.25",
		WeightedAvgPrice:  "40050.00",
		PrevClosePrice:    "39950.00",
		LastPrice:         "40500.00",
		LastQty:           "0.1",
		BidPrice:          "40490.00",
		BidQty:            "0.5",
		AskPrice:          "40510.00",
		AskQty:            "0.3",
		OpenPrice:         "40000.00",
		HighPrice:         "40600.00",
		LowPrice:          "39800.00",
		Volume:            "1200.123",
		QuoteVolume:       "48000000.50",
		OpenTime:          1700000000000,
		CloseTime:         1700086400000,
		FirstId:           1000,
		LastId:            2000,
		Count:             1001,
	}

	expected := &entity.MarketTickerPrice24hrEntity{
		Symbol:      "BTCUSDT",
		LastPrice:   "40500.00",
		OpenPrice:   "40000.00",
		HighPrice:   "40600.00",
		LowPrice:    "39800.00",
		Volume:      "1200.123",
		QuoteVolume: "48000000.50",
		OpenTime:    1700000000000,
		CloseTime:   1700086400000,
		FirstId:     1000,
		LastId:      2000,
		Count:       1001,

		PriceChange:        stringdata.StringPtr("500.00"),
		PriceChangePercent: stringdata.StringPtr("1.25"),
		WeightedAvgPrice:   stringdata.StringPtr("40050.00"),
		PrevClosePrice:     stringdata.StringPtr("39950.00"),
		LastQty:            stringdata.StringPtr("0.1"),
		BidPrice:           stringdata.StringPtr("40490.00"),
		BidQty:             stringdata.StringPtr("0.5"),
		AskPrice:           stringdata.StringPtr("40510.00"),
		AskQty:             stringdata.StringPtr("0.3"),
	}

	result := ToMarketTickerPrice24hrEntity(input)

	assert.Equal(t, expected.Symbol, result.Symbol)
	assert.Equal(t, expected.LastPrice, result.LastPrice)
	assert.Equal(t, expected.OpenPrice, result.OpenPrice)
	assert.Equal(t, expected.HighPrice, result.HighPrice)
	assert.Equal(t, expected.LowPrice, result.LowPrice)
	assert.Equal(t, expected.Volume, result.Volume)
	assert.Equal(t, expected.QuoteVolume, result.QuoteVolume)
	assert.Equal(t, expected.OpenTime, result.OpenTime)
	assert.Equal(t, expected.CloseTime, result.CloseTime)
	assert.Equal(t, expected.FirstId, result.FirstId)
	assert.Equal(t, expected.LastId, result.LastId)
	assert.Equal(t, expected.Count, result.Count)

	// Pointer checks
	assert.Equal(t, *expected.PriceChange, *result.PriceChange)
	assert.Equal(t, *expected.PriceChangePercent, *result.PriceChangePercent)
	assert.Equal(t, *expected.WeightedAvgPrice, *result.WeightedAvgPrice)
	assert.Equal(t, *expected.PrevClosePrice, *result.PrevClosePrice)
	assert.Equal(t, *expected.LastQty, *result.LastQty)
	assert.Equal(t, *expected.BidPrice, *result.BidPrice)
	assert.Equal(t, *expected.BidQty, *result.BidQty)
	assert.Equal(t, *expected.AskPrice, *result.AskPrice)
	assert.Equal(t, *expected.AskQty, *result.AskQty)
}