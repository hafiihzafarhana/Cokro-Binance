package usecase

import (
	"context"

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
	req.TimeZone = convert.NormalizeTimeZone(timeZone)
	req.StartTime = convert.TimeToUnixMilliString(convert.NormalizeTimeString(req.StartTime, timeZone), timeZone)
	req.EndTime = convert.TimeToUnixMilliString(convert.NormalizeTimeString(req.EndTime, timeZone), timeZone)
	return s.repo.GetCandleStickData(ctx, req)
}

func (s *MarketUsecaseImpl) GetBinanceCurrentAveragePrice(ctx context.Context, symbol string) (*entity.MarketCurrentAveragePriceEntity, error) {
	return s.repo.GetCurrentAveragePrice(ctx, symbol)
}