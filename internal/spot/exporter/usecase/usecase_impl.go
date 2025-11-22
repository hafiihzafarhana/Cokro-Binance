package usecase

import (
	"context"

	"github.com/hafiihzafarhana/Cokro-Binance/domain/spot/exporter"
	"github.com/hafiihzafarhana/Cokro-Binance/domain/spot/market"
	"github.com/hafiihzafarhana/Cokro-Binance/domain/spot/market/entity"
)

type ExporterUsecaseImpl struct{
	repo market.MarketRepository
}

func NewExporterUsecase(repo market.MarketRepository) ExporterUseCaseInterface {
    return &ExporterUsecaseImpl{repo: repo}
}

func (s *ExporterUsecaseImpl) CsvBinanceCandleStickData(ctx context.Context, req *exporter.GetCandleStickDataParams) ([]*entity.MarketCandleStickDataEntity, error) {
	return nil, nil
}