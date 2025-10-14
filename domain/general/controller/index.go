package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hafiihzafarhana/Cokro-Binance/domain/general"
	"github.com/hafiihzafarhana/Cokro-Binance/helper/response"
)

type GeneralController struct {
	GeneralService general.GeneralServiceInterface
}

func NewGeneralController(generalService general.GeneralServiceInterface) general.GeneralControllerInterface {
	return &GeneralController{
		GeneralService: generalService,
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
	data, err := g.GeneralService.GetBinanceServerTime()
	if err != nil {
		return response.SendStatusInternalServerError(c, err.Error())
	}

	return response.SendStatusOkWithDataResponse(c, "Data Ok" , data)
}