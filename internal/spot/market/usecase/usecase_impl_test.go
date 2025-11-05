package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/hafiihzafarhana/Cokro-Binance/domain/spot/market"
	"github.com/hafiihzafarhana/Cokro-Binance/domain/spot/market/entity"
	stringdata "github.com/hafiihzafarhana/Cokro-Binance/shared/utils/stringData"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockMarketRepository struct {
	mock.Mock
}

func (m *mockMarketRepository) GetOrderBook(ctx context.Context, req *market.GenericSymbolLimitParams) (*entity.MarketOrderBookEntity, error) {
	args := m.Called(ctx, req)
	return args.Get(0).(*entity.MarketOrderBookEntity), args.Error(1)
}

func (m *mockMarketRepository) GetRecentTradeList(ctx context.Context, req *market.GenericSymbolLimitParams) ([]*entity.MarketRecentTradeListEntity, error) {
	args := m.Called(ctx, req)
	return args.Get(0).([]*entity.MarketRecentTradeListEntity), args.Error(1)
}

func (m *mockMarketRepository) GetOldTradeLookup(ctx context.Context, req *market.GetOldTradeLookupParams) ([]*entity.MarketOldTradeLookupEntity, error) {
	args := m.Called(ctx, req)
	return args.Get(0).([]*entity.MarketOldTradeLookupEntity), args.Error(1)
}

func (m *mockMarketRepository) GetAgregateTradeList(ctx context.Context, req *market.GetAgregateTradeListParams) ([]*entity.MarketAgregateTradeListEntity, error) {
	args := m.Called(ctx, req)
	return args.Get(0).([]*entity.MarketAgregateTradeListEntity), args.Error(1)
}

func (m *mockMarketRepository) GetCandleStickData(ctx context.Context, req *market.GetCandleStickDataParams) ([]*entity.MarketCandleStickDataEntity, error) {
	args := m.Called(ctx, req)
	return args.Get(0).([]*entity.MarketCandleStickDataEntity), args.Error(1)
}

func (m *mockMarketRepository) GetCurrentAveragePrice(ctx context.Context, symbol string) (*entity.MarketCurrentAveragePriceEntity, error) {
	args := m.Called(ctx, symbol)
	return args.Get(0).(*entity.MarketCurrentAveragePriceEntity), args.Error(1)
}

func (m *mockMarketRepository) GetTickerPrice24hr(ctx context.Context, req *market.GetPriceChange24hrParams) ([]*entity.MarketTickerPrice24hrEntity, error) {
	args := m.Called(ctx, req)
	return args.Get(0).([]*entity.MarketTickerPrice24hrEntity), args.Error(1)
}

func (m *mockMarketRepository) GetTradingDayTicker(ctx context.Context, req *market.GetTradingDayTickerParams) ([]*entity.MarketTradingDayTickerEntity, error) {
	return nil, nil
}

func TestNewMarketUsecase(t *testing.T) {
	repo := new(mockMarketRepository)
	uc := NewMarketUsecase(repo)
	if uc == nil {
		t.Fatal("expected non-nil usecase")
	}
}

func TestGetBinanceOrderBook_Success(t *testing.T) {
	ctx := context.Background()
	req := &market.GenericSymbolLimitParams{
		Symbol: "BTCUSDT",
		Limit: 100,
	}

	expectedData := &entity.MarketOrderBookEntity{
		LastUpdateID: 1027024,
		Bids: [][]string{
			{"4.00000000", "431.00000000"},
			{"4.00000200", "12.00000000"},
		},
	}

	mockRepo := new(mockMarketRepository)
	mockRepo.On("GetOrderBook", ctx, req).Return(expectedData, nil)

	usecase := &MarketUsecaseImpl{
		repo: mockRepo,
	}

	result, err := usecase.GetBinanceOrderBook(ctx, req)

	assert.NoError(t, err)
	assert.Equal(t, expectedData, result)
	mockRepo.AssertExpectations(t)
}

func TestGetBinanceRecentTradeList_Success(t *testing.T) {
	ctx := context.Background()
	req := &market.GenericSymbolLimitParams{
		Symbol: "BTCUSDT",
		Limit: 100,
	}

	expectedData := []*entity.MarketRecentTradeListEntity{
		{
			Id: 28457,
			Price: "4.00000100",
			Qty: "12.00000000",
			QuoteQty: "48.000012",
			Time: 1499865549590,
			DateTime: "2006-01-02 15:04:05",
			IsBuyerMaker: true,
			IsBestMatch: true,
		},
	}

	mockRepo := new(mockMarketRepository)
	mockRepo.On("GetRecentTradeList", ctx, req).Return(expectedData, nil)

	usecase := &MarketUsecaseImpl{
		repo: mockRepo,
	}

	result, err := usecase.GetBinanceRecentTradeList(ctx, req)

	assert.NoError(t, err)
	assert.Equal(t, expectedData, result)
	mockRepo.AssertExpectations(t)
}

func TestGetBinanceOldTradeLookup_Success(t *testing.T) {
	ctx := context.Background()
	req := &market.GetOldTradeLookupParams{
		FromId: 28457,
		GenericSymbolLimitParams: market.GenericSymbolLimitParams{
			Symbol: "BTCUSDT",
			Limit: 100,
		},
	}

	expectedData := []*entity.MarketOldTradeLookupEntity{
		{
			Id: 28457,
			Price: "4.00000100",
			Qty: "12.00000000",
			QuoteQty: "48.000012",
			Time: 1499865549590,
			DateTime: "2006-01-02 15:04:05",
			IsBuyerMaker: true,
			IsBestMatch: true,
		},
	}

	mockRepo := new(mockMarketRepository)
	mockRepo.On("GetOldTradeLookup", ctx, req).Return(expectedData, nil)

	usecase := &MarketUsecaseImpl{
		repo: mockRepo,
	}

	result, err := usecase.GetBinanceOldTradeLookup(ctx, req)

	assert.NoError(t, err)
	assert.Equal(t, expectedData, result)
	mockRepo.AssertExpectations(t)
}

	func TestGetBinanceAgregateTradeList_Success(t *testing.T) {
		ctx := context.Background()
		req := &market.GetAgregateTradeListParams{
			GenericSymbolLimitParams: market.GenericSymbolLimitParams{
				Symbol: "BTCUSDT",
				Limit: 100,
			},
			FromId: 28457,
			StartTime: "2025-01-01T00:00:00Z",
			EndTime:   "2025-01-01T01:00:00Z",
			TimeZone:  "UTC",
		}

		expected := []*entity.MarketAgregateTradeListEntity{
			{
				TradeId: 28457,
				Price: "0.01633102",
				Quantity: "4.70443515",
				FirstTradeId: 27781,
				LastTradeId: 27781,
				Timestamp: 1498793709153,
				IsBuyerMaker: true,
				IsBestMatch: true,
				DateTime: "2006-01-02 15:04:05",
			},
		}

		mockRepo := new(mockMarketRepository)
		mockRepo.On("GetAgregateTradeList", ctx, req).Return(expected, nil)
		usecase := &MarketUsecaseImpl{repo: mockRepo}

		result, err := usecase.GetBinanceAgregateTradeList(ctx, req)

		assert.NoError(t, err)
		assert.Equal(t, expected, result)
		mockRepo.AssertExpectations(t)
	}

func TestGetBinanceAgregateTradeList_DefaultTimeZone(t *testing.T) {
	ctx := context.Background()
	req := &market.GetAgregateTradeListParams{
		GenericSymbolLimitParams: market.GenericSymbolLimitParams{
			Symbol: "BTCUSDT",
			Limit:  100,
		},
		TimeZone: "",
	}

	expected := []*entity.MarketAgregateTradeListEntity{}

	mockRepo := new(mockMarketRepository)
	mockRepo.On("GetAgregateTradeList", ctx, mock.AnythingOfType("*market.GetAgregateTradeListParams")).Return(expected, nil)

	usecase := &MarketUsecaseImpl{repo: mockRepo}
	result, err := usecase.GetBinanceAgregateTradeList(ctx, req)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	mockRepo.AssertExpectations(t)
}

func TestGetBinanceAgregateTradeList_Error(t *testing.T) {
	ctx := context.Background()
	req := &market.GetAgregateTradeListParams{
		GenericSymbolLimitParams: market.GenericSymbolLimitParams{
			Symbol: "BTCUSDT",
			Limit:  100,
		},
		TimeZone: "UTC",
	}

	mockRepo := new(mockMarketRepository)
	var expected []*entity.MarketAgregateTradeListEntity
	mockRepo.On("GetAgregateTradeList", ctx, req).
		Return(expected, errors.New("failed to fetch"))

	usecase := &MarketUsecaseImpl{repo: mockRepo}
	result, err := usecase.GetBinanceAgregateTradeList(ctx, req)

	assert.Error(t, err)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}

func TestGetBinanceCandleStickData_Success(t *testing.T) {
	ctx := context.Background()
	req := &market.GetCandleStickDataParams{
		GenericSymbolLimitParams: market.GenericSymbolLimitParams{
			Symbol: "BTCUSDT",
			Limit: 100,
		},
		StartTime: "2025-01-01T00:00:00Z",
		EndTime:   "2025-01-01T01:00:00Z",
		TimeZone:  "UTC",
		Interval: "5m",
		TypeKline: "uiKlines",

	}

	expected := []*entity.MarketCandleStickDataEntity{
		{
			OpenTime: 1499040000000,
			OpenTimeStr: "2025-10-31 15:04:05",
			OpenPrice: "0.01634790",
			HighPrice: "0.80000000",
			LowPrice: "0.01575800",
			ClosePrice: "0.01577100",
			Volume: "148976.11427815",
			CloseTime: 1499644799999,
			CloseTimeStr: "2025-10-31 20:04:05",
			QuoteAssetVolume: "2434.19055334",
			NumberOfTrades: 308,
			TakerBuyBaseAssetVolume: "1756.87402397",
			TakerBuyQuoteAssetVolume: "28.46694368",
			TypeKline: "uiKlines",
		},
	}

	mockRepo := new(mockMarketRepository)
	mockRepo.On("GetCandleStickData", ctx, req).Return(expected, nil)
	usecase := &MarketUsecaseImpl{repo: mockRepo}

	result, err := usecase.GetBinanceCandleStickData(ctx, req)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	mockRepo.AssertExpectations(t)
}

func TestGetBinanceCandleStickData_DefaultTimeZone(t *testing.T) {
	ctx := context.Background()
	req := &market.GetCandleStickDataParams{
		GenericSymbolLimitParams: market.GenericSymbolLimitParams{
			Symbol: "BTCUSDT",
			Limit:  100,
		},
		TimeZone: "",
	}

	expected := []*entity.MarketCandleStickDataEntity{}

	mockRepo := new(mockMarketRepository)
	mockRepo.On("GetCandleStickData", ctx, mock.AnythingOfType("*market.GetCandleStickDataParams")).Return(expected, nil)

	usecase := &MarketUsecaseImpl{repo: mockRepo}
	result, err := usecase.GetBinanceCandleStickData(ctx, req)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	mockRepo.AssertExpectations(t)
}

func TestGetBinanceCurrentAveragePrice_Success(t *testing.T) {
	ctx := context.Background()
	symbol := "BTCUSDT"

	expectedData := &entity.MarketCurrentAveragePriceEntity{
		IntervalInMinute:  5,
		Price 		  : "9.35751834",
		CloseTime 	  : 1694061154503,
		CloseTimeStr:  "2025-10-31 20:04:05",
	}

	mockRepo := new(mockMarketRepository)
	mockRepo.On("GetCurrentAveragePrice", ctx, symbol).Return(expectedData, nil)

	usecase := &MarketUsecaseImpl{
		repo: mockRepo,
	}

	result, err := usecase.GetBinanceCurrentAveragePrice(ctx, symbol)

	assert.NoError(t, err)
	assert.Equal(t, expectedData, result)
	mockRepo.AssertExpectations(t)
}

func TestGetBinancePriceChange24hr_Success(t *testing.T) {
	ctx := context.Background()
	req := &market.GetPriceChange24hrParams{
		Symbol: "BTCUSDT",
		Type: "FULL",
		SymbolStatus: "TRADING",
	}

	expectedData := []*entity.MarketTickerPrice24hrEntity{
		{
			Symbol:      "BTCUSDT",
			LastPrice:   "68000.12",
			OpenPrice:   "67000.00",
			HighPrice:   "68500.00",
			LowPrice:    "66500.00",
			Volume:      "1234.56",
			QuoteVolume: "84000000.00",
			OpenTime:    1730000000000,
			CloseTime:   1730086400000,
			FirstId:     100,
			LastId:      200,
			Count:       101,

			PriceChange:        stringdata.StringPtr("1000.12"),
			PriceChangePercent: stringdata.StringPtr("1.23"),
			WeightedAvgPrice:   stringdata.StringPtr("67500.50"),
			PrevClosePrice:     stringdata.StringPtr("67000.00"),
			LastQty:            stringdata.StringPtr("0.25"),
			BidPrice:           stringdata.StringPtr("67990.00"),
			BidQty:             stringdata.StringPtr("0.15"),
			AskPrice:           stringdata.StringPtr("68010.00"),
			AskQty:             stringdata.StringPtr("0.20"),
		},
	}

	mockRepo := new(mockMarketRepository)
	mockRepo.On("GetTickerPrice24hr", ctx, req).Return(expectedData, nil)

	usecase := &MarketUsecaseImpl{
		repo: mockRepo,
	}

	result, err := usecase.GetBinanceTickerPrice24hr(ctx, req)

	assert.NoError(t, err)
	assert.Equal(t, expectedData, result)
	mockRepo.AssertExpectations(t)
}