package usecase

import (
	"context"

	"github.com/hafiihzafarhana/Cokro-Binance/domain/spot/exporter"
	"github.com/hafiihzafarhana/Cokro-Binance/domain/spot/market/entity"
)

type ExporterUseCaseInterface interface {
	CsvBinanceCandleStickData(ctx context.Context, data *exporter.GetCandleStickDataParams) ([]*entity.MarketCandleStickDataEntity, error)
}