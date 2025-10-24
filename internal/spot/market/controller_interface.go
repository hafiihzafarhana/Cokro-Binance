package market

import "github.com/gofiber/fiber/v2"

type MarketControllerInterface interface {
	CheckBinanceOrderBook(c *fiber.Ctx) error
	CheckBinanceRecentTradeList(c *fiber.Ctx) error
	CheckBinanceOldTradeLookup(c *fiber.Ctx) error
	CheckBinanceAgregateTradeList(c *fiber.Ctx) error
	CheckBinanceCandleStickData(c *fiber.Ctx) error
	CheckBinanceCurrentAveragePrice(c *fiber.Ctx) error
}