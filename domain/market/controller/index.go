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

// CheckOrderBook godoc
// @Summary Check order book
// @Description Check order book from Binance API
// @Tags Market
// @Accept json
// @Produce json
// @Param symbol query string true "Trading pair symbol (e.g. BTCUSDT)"
// @Param limit query int false "Number of entries to fetch (default: 100, min: 100, max: 5000)"
// @Router /market/order-book [get]
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

// CheckRecentTradeList godoc
// @Summary Check recent trade list
// @Description Check recent trade list from Binance API
// @Tags Market
// @Accept json
// @Produce json
// @Param symbol query string true "Trading pair symbol (e.g. BTCUSDT)"
// @Param limit query int false "Number of entries to fetch (min: 50, max: 1000)"
// @Router /market/recent-trade-lists [get]
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

// CheckOldTradeLookup godoc
// @Summary Check old trade lookup
// @Description Check old trade lookup from Binance API
// @Tags Market
// @Accept json
// @Produce json
// @Param symbol query string true "Trading pair symbol (e.g. BTCUSDT)"
// @Param limit query int false "Number of entries to fetch (min: 50, max: 1000)"
// @Param fromId query int false "Trade ID to fetch from (default: 0)"
// @Router /market/old-trade-lookup [get]
func (m *MarketController) CheckOldTradeLookup(c *fiber.Ctx) error {
	var q dto.GetBinanceOldTradeLookupReq
	if err := c.QueryParser(&q); err != nil {
		return response.SendStatusBadRequest(c, "invalid query: "+err.Error())
	}
	if err := validator.ValidateStruct(q); err != nil {
		return response.SendStatusBadRequest(c, "error validating payload:"+err.Error())
	}
	data, err := m.MarketService.GetBinanceOldTradeLookup(&q)
	if err != nil {
		return response.SendStatusInternalServerError(c, err.Error())
	}

	return response.SendStatusOkWithDataResponse(c, "Data Ok" , data)
}

// CheckAgregateTradeList godoc
// @Summary Check agregate trade list
// @Description Check agregate trade list from Binance API
// @Tags Market
// @Accept json
// @Produce json
// @Param symbol query string true "Trading pair symbol (e.g. BTCUSDT)"
// @Param limit query int false "Number of entries to fetch (min: 50, max: 1000)"
// @Param fromId query int false "Trade ID to fetch from (default: 0)"
// @Param startTime query string false "Start time in RFC3339 format (e.g. 2025-10-14T00:00:00Z)"
// @Param endTime query string false "End time in RFC3339 format (e.g. 2025-10-14T12:00:00Z)"
// @Router /market/agregate-trade-lists [get]
func (m *MarketController) CheckAgregateTradeList(c *fiber.Ctx) error {
	var q dto.GetBinanceAgregateTradeListReq
	if err := c.QueryParser(&q); err != nil {
		return response.SendStatusBadRequest(c, "invalid query: "+err.Error())
	}
	if err := validator.ValidateStruct(q); err != nil {
		return response.SendStatusBadRequest(c, "error validating payload:"+err.Error())
	}
	data, err := m.MarketService.GetAgregateTradeList(&q)
	if err != nil {
		return response.SendStatusInternalServerError(c, err.Error())
	}
	return response.SendStatusOkWithDataResponse(c, "Data Ok" , dto.GetBinanceAggregateTradeResponse(data))
}

