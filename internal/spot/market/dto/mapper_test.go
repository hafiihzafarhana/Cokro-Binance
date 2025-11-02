package dto

import (
	"testing"

	"github.com/hafiihzafarhana/Cokro-Binance/domain/spot/market/entity"
	stringdata "github.com/hafiihzafarhana/Cokro-Binance/shared/utils/stringData"
	"github.com/stretchr/testify/assert"
)

func TestToMarketPriceChange24hrResponse(t *testing.T) {
	// Arrange
	priceChange := stringdata.StringPtr("500.00")
	priceChangePercent := stringdata.StringPtr("1.25")
	weightedAvgPrice := stringdata.StringPtr("40050.00")
	prevClosePrice := stringdata.StringPtr("39950.00")
	lastQty := stringdata.StringPtr("0.1")
	bidPrice := stringdata.StringPtr("40490.00")
	bidQty := stringdata.StringPtr("0.5")
	askPrice := stringdata.StringPtr("40510.00")
	askQty := stringdata.StringPtr("0.3")

	entityInput := &entity.MarketTickerPrice24hrEntity{
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

		PriceChange:        priceChange,
		PriceChangePercent: priceChangePercent,
		WeightedAvgPrice:   weightedAvgPrice,
		PrevClosePrice:     prevClosePrice,
		LastQty:            lastQty,
		BidPrice:           bidPrice,
		BidQty:             bidQty,
		AskPrice:           askPrice,
		AskQty:             askQty,
	}

	// Act
	result := ToMarketPriceChange24hrResponse(entityInput)

	// Assert
	assert.Equal(t, entityInput.Symbol, result.Symbol)
	assert.Equal(t, entityInput.LastPrice, result.LastPrice)
	assert.Equal(t, entityInput.OpenPrice, result.OpenPrice)
	assert.Equal(t, entityInput.HighPrice, result.HighPrice)
	assert.Equal(t, entityInput.LowPrice, result.LowPrice)
	assert.Equal(t, entityInput.Volume, result.Volume)
	assert.Equal(t, entityInput.QuoteVolume, result.QuoteVolume)
	assert.Equal(t, entityInput.OpenTime, result.OpenTime)
	assert.Equal(t, entityInput.CloseTime, result.CloseTime)
	assert.Equal(t, entityInput.FirstId, result.FirstId)
	assert.Equal(t, entityInput.LastId, result.LastId)
	assert.Equal(t, entityInput.Count, result.Count)

	assert.Equal(t, *priceChange, *result.PriceChange)
	assert.Equal(t, *priceChangePercent, *result.PriceChangePercent)
	assert.Equal(t, *weightedAvgPrice, *result.WeightedAvgPrice)
	assert.Equal(t, *prevClosePrice, *result.PrevClosePrice)
	assert.Equal(t, *lastQty, *result.LastQty)
	assert.Equal(t, *bidPrice, *result.BidPrice)
	assert.Equal(t, *bidQty, *result.BidQty)
	assert.Equal(t, *askPrice, *result.AskPrice)
	assert.Equal(t, *askQty, *result.AskQty)
}

func TestToMarketTickerPrice24hrResponseList(t *testing.T) {
	// Arrange
	priceChange := stringdata.StringPtr("500.00")
	priceChangePercent := stringdata.StringPtr("1.25")

	inputList := []*entity.MarketTickerPrice24hrEntity{
		{
			Symbol:             "BTCUSDT",
			LastPrice:          "40500.00",
			OpenPrice:          "40000.00",
			HighPrice:          "40600.00",
			LowPrice:           "39800.00",
			Volume:             "1200.123",
			QuoteVolume:        "48000000.50",
			OpenTime:           1700000000000,
			CloseTime:          1700086400000,
			FirstId:            1000,
			LastId:             2000,
			Count:              1001,
			PriceChange:        priceChange,
			PriceChangePercent: priceChangePercent,
		},
		{
			Symbol:    "ETHUSDT",
			LastPrice: "2100.00",
			OpenPrice: "2000.00",
			HighPrice: "2150.00",
			LowPrice:  "1980.00",
			Volume:    "500.55",
			Count:     250,
		},
	}

	// Act
	result := ToMarketTickerPrice24hrResponseList(inputList)

	// Assert
	assert.Len(t, result, len(inputList), "result slice length should match input length")

	// Check first item mapping
	first := result[0]
	assert.Equal(t, "BTCUSDT", first.Symbol)
	assert.Equal(t, "40500.00", first.LastPrice)
	assert.Equal(t, "40000.00", first.OpenPrice)
	assert.Equal(t, "40600.00", first.HighPrice)
	assert.Equal(t, "39800.00", first.LowPrice)
	assert.Equal(t, "1200.123", first.Volume)
	assert.Equal(t, "48000000.50", first.QuoteVolume)
	assert.Equal(t, int64(1000), first.FirstId)
	assert.Equal(t, int64(2000), first.LastId)
	assert.Equal(t, int64(1001), first.Count)
	assert.Equal(t, *priceChange, *first.PriceChange)
	assert.Equal(t, *priceChangePercent, *first.PriceChangePercent)

	// Check second item mapping
	second := result[1]
	assert.Equal(t, "ETHUSDT", second.Symbol)
	assert.Equal(t, "2100.00", second.LastPrice)
	assert.Equal(t, "2000.00", second.OpenPrice)
	assert.Equal(t, "2150.00", second.HighPrice)
	assert.Equal(t, "1980.00", second.LowPrice)
	assert.Equal(t, "500.55", second.Volume)
	assert.Equal(t, int64(250), second.Count)
	assert.Nil(t, second.PriceChange) // karena field pointer nil
}

func TestToMarketTickerPrice24hrResponseList_EmptyInput(t *testing.T) {
	// Arrange
	inputList := []*entity.MarketTickerPrice24hrEntity{}

	// Act
	result := ToMarketTickerPrice24hrResponseList(inputList)

	// Assert
	assert.NotNil(t, result)
	assert.Empty(t, result, "result should be empty slice, not nil")
}