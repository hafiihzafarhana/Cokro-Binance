package market

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterMarketRoutes(app *fiber.App, handler MarketControllerInterface) {
	marketGroup := app.Group("/market")
	marketGroup.Get("/order-book", handler.CheckBinanceOrderBook)
	marketGroup.Get("/recent-trade-lists", handler.CheckBinanceRecentTradeList)
	marketGroup.Get("/old-trade-lookup", handler.CheckBinanceOldTradeLookup)
	marketGroup.Get("/agregate-trade-lists", handler.CheckBinanceAgregateTradeList)
	marketGroup.Get("/candlestick-data", handler.CheckBinanceCandleStickData) // klines and uiklines
	marketGroup.Get("/current-average-price", handler.CheckBinanceCurrentAveragePrice)
	marketGroup.Get("/price-change-24hr", handler.CheckBinancePriceChange24hr)
}