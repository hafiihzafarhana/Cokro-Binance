package general

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/hafiihzafarhana/Cokro-Binance/internal/spot/general/usecase"
	"github.com/hafiihzafarhana/Cokro-Binance/shared/response"
)

type GeneralController struct {
	usecase usecase.MarketUsecaseInterface
}

func NewGeneralController(usecase usecase.MarketUsecaseInterface) GeneralControllerInterface {
	return &GeneralController{
		usecase: usecase,
	}
}

// CheckServerTime godoc
// @Summary Check server time
// @Description Check server time from Binance API
// @Tags General
// @Accept json
// @Produce json
// @Router /general/server-time [get]
func (g *GeneralController) CheckServerTime(c *fiber.Ctx) error {
	data, err := g.usecase.GetServerTime(context.Background())
	if err != nil {
		return response.SendStatusInternalServerError(c, err.Error())
	}

	return response.SendStatusOkWithDataResponse(c, "Data Ok" , data)
}
