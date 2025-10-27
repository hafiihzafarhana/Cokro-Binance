package market

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hafiihzafarhana/Cokro-Binance/domain/spot/market"
	"github.com/hafiihzafarhana/Cokro-Binance/internal/spot/market/dto"
	"github.com/hafiihzafarhana/Cokro-Binance/internal/spot/market/usecase"
	"github.com/hafiihzafarhana/Cokro-Binance/shared/response"
	"github.com/hafiihzafarhana/Cokro-Binance/shared/utils/validator"
)

type MarketController struct {
	usecase usecase.MarketUseCaseInterface
}

func NewMarketController(usecase usecase.MarketUseCaseInterface) MarketControllerInterface {
	return &MarketController{
		usecase: usecase,
	}
}

// CheckBinanceOrderBook godoc
// @Summary Check order book
// @Description Check order book from Binance API
// @Tags Market
// @Accept json
// @Produce json
// @Param symbol query string true "Trading pair symbol (e.g. BTCUSDT)"
// @Param limit query int false "Number of entries to fetch (default: 100, min: 100, max: 5000)"
// @Router /market/order-book [get]
func (m *MarketController) CheckBinanceOrderBook(c *fiber.Ctx) error {
	var q dto.GenericSymbolLimitReq
	if err := c.QueryParser(&q); err != nil {
		return response.SendStatusBadRequest(c, "invalid query: "+err.Error())
	}
	if err := validator.ValidateStruct(q); err != nil {
		return response.SendStatusBadRequest(c, "error validating payload:"+err.Error())
	}
	params := market.GenericSymbolLimitParams{
		Symbol: q.Symbol,
		Limit:  q.Limit,
	}
	data, err := m.usecase.GetBinanceOrderBook(c.UserContext(), &params)
	if err != nil {
		return response.SendStatusInternalServerError(c, err.Error())
	}

	return response.SendStatusOkWithDataResponse(c, "Data Ok" , data)
}

// CheckBinanceRecentTradeList godoc
// @Summary Check recent trade list
// @Description Check recent trade list from Binance API
// @Tags Market
// @Accept json
// @Produce json
// @Param symbol query string true "Trading pair symbol (e.g. BTCUSDT)"
// @Param limit query int false "Number of entries to fetch (min: 100, max: 1000)"
// @Router /market/recent-trade-lists [get]
func (m *MarketController) CheckBinanceRecentTradeList(c *fiber.Ctx) error {
	var q dto.GenericSymbolLimitReq
	if err := c.QueryParser(&q); err != nil {
		return response.SendStatusBadRequest(c, "invalid query: "+err.Error())
	}
	if err := validator.ValidateStruct(q); err != nil {
		return response.SendStatusBadRequest(c, "error validating payload:"+err.Error())
	}
	params := market.GenericSymbolLimitParams{
		Symbol: q.Symbol,
		Limit:  q.Limit,
	}
	data, err := m.usecase.GetBinanceRecentTradeList(c.UserContext(), &params)
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
// @Param limit query int false "Number of entries to fetch (min: 100, max: 1000)"
// @Param fromId query int false "Trade ID to fetch from (default: 0)"
// @Router /market/old-trade-lookup [get]
func (m *MarketController) CheckBinanceOldTradeLookup(c *fiber.Ctx) error {
	var q dto.GetOldTradeLookupReq
	if err := c.QueryParser(&q); err != nil {
		return response.SendStatusBadRequest(c, "invalid query: "+err.Error())
	}
	if err := validator.ValidateStruct(q); err != nil {
		return response.SendStatusBadRequest(c, "error validating payload:"+err.Error())
	}
	params := market.GetOldTradeLookupParams{
		FromId: q.FromId,
		GenericSymbolLimitParams: market.GenericSymbolLimitParams{
			Symbol: q.Symbol,
			Limit:  q.Limit,
		},
	}
	data, err := m.usecase.GetBinanceOldTradeLookup(c.UserContext(), &params)
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
// @Param limit query int false "Number of entries to fetch (min: 100, max: 1000)"
// @Param fromId query int false "Trade ID to fetch from (default: 0)"
// @Param startTime query string false "Start time in RFC3339 format (e.g. 2025-10-14T00:00:00Z)"
// @Param endTime query string false "End time in RFC3339 format (e.g. 2025-10-14T12:00:00Z)"
// @Param timeZone query string false "Time zone (default: UTC)"
// @Router /market/agregate-trade-lists [get]
func (m *MarketController) CheckBinanceAgregateTradeList(c *fiber.Ctx) error {
	var q dto.GetAgregateTradeListReq
	if err := c.QueryParser(&q); err != nil {
		return response.SendStatusBadRequest(c, "invalid query: "+err.Error())
	}
	if err := validator.ValidateStruct(q); err != nil {
		return response.SendStatusBadRequest(c, "error validating payload:"+err.Error())
	}
	params := market.GetAgregateTradeListParams{
		FromId: q.FromId,
		StartTime: q.StartTime,
		EndTime:   q.EndTime,
		TimeZone:  q.TimeZone,
		GenericSymbolLimitParams: market.GenericSymbolLimitParams{
			Symbol: q.Symbol,
			Limit:  q.Limit,
		},
	}
	data, err := m.usecase.GetBinanceAgregateTradeList(c.UserContext(), &params)
	if err != nil {
		return response.SendStatusInternalServerError(c, err.Error())
	}
	return response.SendStatusOkWithDataResponse(c, "Data Ok" , data)
}

// CheckCandleStickData godoc
// @Summary Check candle stick data
// @Description Check candle stick data from Binance API
// @Tags Market
// @Accept json
// @Produce json
// @Param symbol query string true "Trading pair symbol (e.g. BTCUSDT)"
// @Param limit query int false "Number of entries to fetch (min: 100, max: 1000)"
// @Param startTime query string false "Start time in RFC3339 format (e.g. 2025-10-14T00:00:00Z)"
// @Param endTime query string false "End time in RFC3339 format (e.g. 2025-10-14T12:00:00Z)"
// @Param interval query string true "Candle stick interval (e.g. 1m, 5m, 1h, 1d)" Enums(1m,3m,5m,15m,30m,1h,2h,4h,6h,12h,1d,3d,1w,1M)
// @Param timeZone query string false "Time zone (default: UTC)"
// @Param typeKline query string false "Type of kline (default: klines)" Enums(klines,uiklines)
// @Router /market/candlestick-data [get]
func (m *MarketController) CheckBinanceCandleStickData(c *fiber.Ctx) error {
	var q dto.GetCandleStickDataReq
	if err := c.QueryParser(&q); err != nil {
		return response.SendStatusBadRequest(c, "invalid query: "+err.Error())
	}
	if err := validator.ValidateStruct(q); err != nil {
		return response.SendStatusBadRequest(c, "error validating payload:"+err.Error())
	}
	params := market.GetCandleStickDataParams{
		StartTime: q.StartTime,
		EndTime:   q.EndTime,
		Interval:  q.Interval,
		TypeKline: q.TypeKline,
		TimeZone:  q.TimeZone,
		GenericSymbolLimitParams: market.GenericSymbolLimitParams{
			Symbol: q.Symbol,
			Limit:  q.Limit,
		},
	}
	data, err := m.usecase.GetBinanceCandleStickData(c.UserContext(), &params)
	if err != nil {
		return response.SendStatusInternalServerError(c, err.Error())
	}
	return response.SendStatusOkWithDataResponse(c, "Data Ok" , data)
}

// CheckCurrentAveragePrice godoc
// @Summary Check current average price
// @Description Check current average price from Binance API
// @Tags Market
// @Accept json
// @Produce json
// @Param symbol query string true "Trading pair symbol (e.g. BTCUSDT)"
// @Router /market/current-average-price [get]
func (m *MarketController) CheckBinanceCurrentAveragePrice(c *fiber.Ctx) error {
	symbol := c.Query("symbol")
	if symbol == "" {
		return response.SendStatusBadRequest(c, "symbol query parameter is required")
	}

	data, err := m.usecase.GetBinanceCurrentAveragePrice(c.UserContext(), symbol)
	if err != nil {
		return response.SendStatusInternalServerError(c, err.Error())
	}

	return response.SendStatusOkWithDataResponse(c, "Data Ok", data)
}