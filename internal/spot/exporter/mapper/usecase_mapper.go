package mapper

import (
	"github.com/hafiihzafarhana/Cokro-Binance/domain/spot/exporter"
	"github.com/hafiihzafarhana/Cokro-Binance/domain/spot/market"
	"github.com/hafiihzafarhana/Cokro-Binance/domain/spot/market/entity"
)

func ToMarketParams(e *exporter.GetCandleStickDataParams) *market.GetCandleStickDataParams {
	return &market.GetCandleStickDataParams{
		Interval: e.Interval,
		GenericSymbolLimitParams: market.GenericSymbolLimitParams{
			Symbol: e.Symbol,
			Limit: e.Limit,
		},
		StartTime: e.StartTime,
		EndTime: e.EndTime,
		TypeKline: e.TypeKline,
		TimeZone: e.TimeZone,
	}
}

func ToMarketCsvRows(datas []*entity.MarketCandleStickDataEntity) [][]string {
	rows := make([][]string, 0, len(datas))

	rows = append(rows, []string{
        "OpenTime",
        "Open",
        "High",
        "Low",
        "Close",
        "Volume",
        "CloseTime",
    })

	for _, d := range datas {
		row := []string{
			d.OpenTimeStr,
			d.OpenPrice,
			d.HighPrice,
			d.LowPrice,
			d.ClosePrice,
			d.Volume,
			d.CloseTimeStr,
		}
		rows = append(rows, row)
	}

	return rows
}