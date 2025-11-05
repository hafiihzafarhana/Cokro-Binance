package market

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

type mockMarketController struct{}

func (m *mockMarketController) CheckBinanceOrderBook(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).SendString("order-book ok")
}
func (m *mockMarketController) CheckBinanceRecentTradeList(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).SendString("recent-trade ok")
}
func (m *mockMarketController) CheckBinanceOldTradeLookup(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).SendString("old-trade ok")
}
func (m *mockMarketController) CheckBinanceAgregateTradeList(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).SendString("agregate ok")
}
func (m *mockMarketController) CheckBinanceCandleStickData(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).SendString("candlestick ok")
}
func (m *mockMarketController) CheckBinanceCurrentAveragePrice(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).SendString("avg-price ok")
}
func (m *mockMarketController) CheckBinancePriceChange24hr(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).SendString("price-change ok")
}
func (m *mockMarketController) CheckBinanceTradingDayTicker(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).SendString("day-ticker ok")
}

func TestRegisterMarketRoutes(t *testing.T) {
	app := fiber.New()
	handler := &mockMarketController{}

	RegisterMarketRoutes(app, handler)

	tests := []struct {
		route string
		want  string
	}{
		{"/market/order-book", "order-book ok"},
		{"/market/recent-trade-lists", "recent-trade ok"},
		{"/market/old-trade-lookup", "old-trade ok"},
		{"/market/agregate-trade-lists", "agregate ok"},
		{"/market/candlestick-data", "candlestick ok"},
		{"/market/current-average-price", "avg-price ok"},
		{"/market/price-change-24hr", "price-change ok"},
		{"/market/day-ticker", "day-ticker ok"},
	}

	for _, tt := range tests {
		req := httptest.NewRequest(http.MethodGet, tt.route, nil)
		resp, _ := app.Test(req, -1)

		assert.Equal(t, http.StatusOK, resp.StatusCode)
		body := make([]byte, resp.ContentLength)
		resp.Body.Read(body)
		assert.Contains(t, string(body), tt.want)
		resp.Body.Close()
	}
}