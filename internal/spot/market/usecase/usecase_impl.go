package usecase

import (
	"context"
	"time"

	"github.com/hafiihzafarhana/Cokro-Binance/domain/spot/market"
	"github.com/hafiihzafarhana/Cokro-Binance/domain/spot/market/entity"
	"github.com/hafiihzafarhana/Cokro-Binance/internal/spot/market/dto"
	"github.com/hafiihzafarhana/Cokro-Binance/shared/utils/convert"
)

type MarketUsecaseImpl struct{
	repo market.MarketRepository
}

func NewMarketUsecase(repo market.MarketRepository) MarketUseCaseInterface {
    return &MarketUsecaseImpl{repo: repo}
}

func (s *MarketUsecaseImpl) GetBinanceOrderBook(ctx context.Context, req *dto.GenericSymbolLimitReq) (*entity.MarketOrderBookEntity, error) {
	params := &market.GenericSymbolLimitParams{
		Symbol: req.Symbol,
		Limit:  req.Limit,
	}

	return s.repo.GetOrderBook(ctx, params)
}

func (s *MarketUsecaseImpl) GetBinanceRecentTradeList(ctx context.Context, req *dto.GenericSymbolLimitReq) ([]*entity.MarketRecentTradeListEntity, error) {
	params := &market.GenericSymbolLimitParams{
		Symbol: req.Symbol,
		Limit:  req.Limit,
	}

	results, err := s.repo.GetRecentTradeList(ctx, params)
	if err != nil {
		return nil, err
	}

	for _, r := range results {
		r.DateTime = convert.UnixToDateTimeString(r.Time, "Asia/Jakarta")
	}

	return results, nil
}

func (s *MarketUsecaseImpl) GetBinanceOldTradeLookup(ctx context.Context, req *dto.GetOldTradeLookupReq) ([]*entity.MarketOldTradeLookupEntity, error) {
	params := &market.GetOldTradeLookupParams{
		FromId: req.FromId,
		GenericSymbolLimitParams: market.GenericSymbolLimitParams{
			Symbol: req.Symbol,
			Limit:  req.Limit,
		},
	}

	return s.repo.GetOldTradeLookup(ctx, params)
}

func (s *MarketUsecaseImpl) GetBinanceAgregateTradeList(ctx context.Context, req *dto.GetAgregateTradeListReq) ([]*entity.MarketAgregateTradeListEntity, error) {
	timeZone := req.TimeZone
	if timeZone == "" {
		timeZone = "UTC"
	}
	loc, err := time.LoadLocation(timeZone)
	if err != nil {
		loc = time.UTC
	}
	const layout = "2006-01-02T15:04:05"
	if req.StartTime != "" {
		if t, err := time.ParseInLocation(layout, req.StartTime, loc); err == nil {
			req.StartTime = t.UTC().Format(layout)
		}
	}
	if req.EndTime != "" {
		if t, err := time.ParseInLocation(layout, req.EndTime, loc); err == nil {
			req.EndTime = t.UTC().Format(layout)
		}
	}
	params := &market.GetAgregateTradeListParams{
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
		GenericSymbolLimitParams: market.GenericSymbolLimitParams{
			Symbol: req.Symbol,
			Limit:  req.Limit,
		},
	}

	return s.repo.GetAgregateTradeList(ctx, params)
}

func (s *MarketUsecaseImpl) GetBinanceCandleStickData(ctx context.Context, req *dto.GetCandleStickDataReq) ([]*entity.MarketCandleStickDataEntity, error) {
	timeZone := req.TimeZone
	if timeZone == "" {
		timeZone = "UTC"
	}
	loc, err := time.LoadLocation(timeZone)
	if err != nil {
		loc = time.UTC
	}
	const layout = "2006-01-02T15:04:05"
	if req.StartTime != "" {
		if t, err := time.ParseInLocation(layout, req.StartTime, loc); err == nil {
			req.StartTime = t.UTC().Format(layout)
		}
	}
	if req.EndTime != "" {
		if t, err := time.ParseInLocation(layout, req.EndTime, loc); err == nil {
			req.EndTime = t.UTC().Format(layout)
		}
	}
	params := &market.GetCandleStickDataParams{
		Interval:  req.Interval,
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
		TypeKline: req.TypeKline,
		TimeZone:  req.TimeZone,
		GenericSymbolLimitParams: market.GenericSymbolLimitParams{
			Symbol: req.Symbol,
			Limit:  req.Limit,
		},
	}

	return s.repo.GetCandleStickData(ctx, params)
}

func (s *MarketUsecaseImpl) GetBinanceCurrentAveragePrice(ctx context.Context, symbol string) (*entity.MarketCurrentAveragePriceEntity, error) {
	return s.repo.GetCurrentAveragePrice(ctx, symbol)
}