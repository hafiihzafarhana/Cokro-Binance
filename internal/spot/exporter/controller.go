package exporter

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hafiihzafarhana/Cokro-Binance/domain/spot/exporter"
	"github.com/hafiihzafarhana/Cokro-Binance/internal/spot/exporter/dto"
	"github.com/hafiihzafarhana/Cokro-Binance/internal/spot/exporter/usecase"
	"github.com/hafiihzafarhana/Cokro-Binance/shared/response"
	"github.com/hafiihzafarhana/Cokro-Binance/shared/utils/validator"
)

type ExporterController struct {
	usecase usecase.ExporterUseCaseInterface
}

func NewExporterController(usecase usecase.ExporterUseCaseInterface) ExporterControllerInterface {
	return &ExporterController{
		usecase: usecase,
	}
}

// ExportBinanceCandleStickData godoc
// @Summary Export candle stick data
// @Description Export candle stick data from Binance API
// @Tags Exporter
// @Accept json
// @Produce json
// @Param symbol query string true "Trading pair symbol (e.g. BTCUSDT)"
// @Param limit query int false "Number of entries to fetch (min: 100, max: 1000)"
// @Param startTime query string false "Start time in RFC3339 format (e.g. 2025-10-14T00:00:00Z)"
// @Param endTime query string false "End time in RFC3339 format (e.g. 2025-10-14T12:00:00Z)"
// @Param interval query string true "Candle stick interval (e.g. 1m, 5m, 1h, 1d)" Enums(1m,3m,5m,15m,30m,1h,2h,4h,6h,12h,1d,3d,1w,1M)
// @Router /exporter/candlestick-data [get]
func (m *ExporterController) ExportBinanceCandleStickData(c *fiber.Ctx) error {
	var q dto.GetCandleStickDataReq
	if err := c.QueryParser(&q); err != nil {
		return response.SendStatusBadRequest(c, "invalid query: "+err.Error())
	}
	if err := validator.ValidateStruct(q); err != nil {
		return response.SendStatusBadRequest(c, "error validating payload:"+err.Error())
	}
	params := exporter.GetCandleStickDataParams{
		StartTime: q.StartTime,
		EndTime:   q.EndTime,
		Interval:  q.Interval,
		TypeKline: "uiklines",
		TimeZone: "UTC",
		GenericSymbolLimitParams: exporter.GenericSymbolLimitParams{
			Symbol: q.Symbol,
			Limit:  q.Limit,
		},
	}
	data, err := m.usecase.CsvBinanceCandleStickData(c.UserContext(), &params)
	if err != nil {
		return response.SendStatusInternalServerError(c, err.Error())
	}

	filename := "binance_candlestick_" + q.Symbol + "_" + q.StartTime+ "_" + q.EndTime + ".csv"
	c.Set("Content-Type", "text/csv")
	c.Set("Content-Disposition", `attachment; filename="`+filename+`"`)

	return c.Send(data)
}