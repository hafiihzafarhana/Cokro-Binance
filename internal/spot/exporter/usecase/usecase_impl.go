package usecase

import (
	"context"

	"github.com/hafiihzafarhana/Cokro-Binance/domain/spot/exporter"
	"github.com/hafiihzafarhana/Cokro-Binance/domain/spot/market"
	"github.com/hafiihzafarhana/Cokro-Binance/internal/spot/exporter/mapper"
	"github.com/hafiihzafarhana/Cokro-Binance/internal/spot/exporter/utils"
	"github.com/hafiihzafarhana/Cokro-Binance/shared/utils/convert"
)

type ExporterUsecaseImpl struct{
	repo market.MarketRepository
}

func NewExporterUsecase(repo market.MarketRepository) ExporterUseCaseInterface {
    return &ExporterUsecaseImpl{repo: repo}
}

func (s *ExporterUsecaseImpl) CsvBinanceCandleStickData(ctx context.Context, req *exporter.GetCandleStickDataParams) ([]byte, error) {
	req.TimeZone = convert.NormalizeTimeZone(req.TimeZone)
	req.StartTime = convert.TimeToUnixMilliString(convert.NormalizeTimeString(req.StartTime, req.TimeZone), req.TimeZone)
	req.EndTime = convert.TimeToUnixMilliString(convert.NormalizeTimeString(req.EndTime, req.TimeZone), req.TimeZone)
	mp := mapper.ToMarketParams(req)
	candles, err := s.repo.GetCandleStickData(ctx, mp)
	if err != nil {
        return nil, err
    }

	rows := mapper.ToMarketCsvRows(candles)
	csvBytes, err := utils.GenerateCsv(rows)
    if err != nil {
        return nil, err
    }
	return csvBytes, nil
}