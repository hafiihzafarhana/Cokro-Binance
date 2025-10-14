package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hafiihzafarhana/Cokro-Binance/domain/market"
	"github.com/hafiihzafarhana/Cokro-Binance/domain/market/dto"
	"github.com/hafiihzafarhana/Cokro-Binance/helper/response"
	"github.com/hafiihzafarhana/Cokro-Binance/helper/validator"
)

type MarketController struct {
	MarketService market.MarketServiceInterface
}

func NewMarketController(marketService market.MarketServiceInterface) market.MarketControllerInterface {
	return &MarketController{
		MarketService: marketService,
	}
}

func (m *MarketController) CheckOrderBook(c *fiber.Ctx) error {
	var q dto.GetBinanceOrderBookReq
	if err := c.QueryParser(&q); err != nil {
		return response.SendStatusBadRequest(c, "invalid query: "+err.Error())
	}
	if err := validator.ValidateStruct(q); err != nil {
		return response.SendStatusBadRequest(c, "error validating payload:"+err.Error())
	}
	data, err := m.MarketService.GetBinanceOrderBook(&q)
	if err != nil {
		return response.SendStatusInternalServerError(c, err.Error())
	}

	return response.SendStatusOkWithDataResponse(c, "Data Ok" , data)
}

func (m *MarketController) CheckRecentTradeList(c *fiber.Ctx) error {
	var q dto.GetBinanceRecentTradeListReq
	if err := c.QueryParser(&q); err != nil {
		return response.SendStatusBadRequest(c, "invalid query: "+err.Error())
	}
	if err := validator.ValidateStruct(q); err != nil {
		return response.SendStatusBadRequest(c, "error validating payload:"+err.Error())
	}
	data, err := m.MarketService.GetBinanceRecentTradeList(&q)
	if err != nil {
		return response.SendStatusInternalServerError(c, err.Error())
	}

	return response.SendStatusOkWithDataResponse(c, "Data Ok" , data)
}

