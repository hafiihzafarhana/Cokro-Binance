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
		GenericSymbolLimitParams: exporter.GenericSymbolLimitParams{
			Symbol: q.Symbol,
			Limit:  q.Limit,
		},
	}
	data, err := m.usecase.CsvBinanceCandleStickData(c.UserContext(), &params)
	if err != nil {
		return response.SendStatusInternalServerError(c, err.Error())
	}

	return response.SendStatusOkWithDataResponse(c, "Data Ok" , data)
}