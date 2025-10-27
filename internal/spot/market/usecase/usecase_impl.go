package usecase

import (
	"context"
	"time"

	"github.com/hafiihzafarhana/Cokro-Binance/domain/spot/market"
	"github.com/hafiihzafarhana/Cokro-Binance/domain/spot/market/entity"
	"github.com/hafiihzafarhana/Cokro-Binance/shared/utils/convert"
)

type MarketUsecaseImpl struct{
	repo market.MarketRepository
}

func NewMarketUsecase(repo market.MarketRepository) MarketUseCaseInterface {
    return &MarketUsecaseImpl{repo: repo}
}

func (s *MarketUsecaseImpl) GetBinanceOrderBook(ctx context.Context, req *market.GenericSymbolLimitParams) (*entity.MarketOrderBookEntity, error) {
	return s.repo.GetOrderBook(ctx, req)
}

func (s *MarketUsecaseImpl) GetBinanceRecentTradeList(ctx context.Context, req *market.GenericSymbolLimitParams) ([]*entity.MarketRecentTradeListEntity, error) {
	results, err := s.repo.GetRecentTradeList(ctx, req)
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (s *MarketUsecaseImpl) GetBinanceOldTradeLookup(ctx context.Context, req *market.GetOldTradeLookupParams) ([]*entity.MarketOldTradeLookupEntity, error) {
	results, err := s.repo.GetOldTradeLookup(ctx, req)
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (s *MarketUsecaseImpl) GetBinanceAgregateTradeList(ctx context.Context, req *market.GetAgregateTradeListParams) ([]*entity.MarketAgregateTradeListEntity, error) {
	timeZone := req.TimeZone
	if timeZone == "" {
		timeZone = "UTC"
	}
	req.StartTime = convert.TimeToUnixMilliString(convert.NormalizeTimeString(req.StartTime, req.TimeZone), req.TimeZone)
	req.EndTime = convert.TimeToUnixMilliString(convert.NormalizeTimeString(req.EndTime, req.TimeZone), req.TimeZone)
	results, err := s.repo.GetAgregateTradeList(ctx, req)
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (s *MarketUsecaseImpl) GetBinanceCandleStickData(ctx context.Context, req *market.GetCandleStickDataParams) ([]*entity.MarketCandleStickDataEntity, error) {
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