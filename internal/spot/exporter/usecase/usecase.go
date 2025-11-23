package usecase

import (
	"context"

	"github.com/hafiihzafarhana/Cokro-Binance/domain/spot/exporter"
)

type ExporterUseCaseInterface interface {
	CsvBinanceCandleStickData(ctx context.Context, data *exporter.GetCandleStickDataParams) ([]byte, error)
}