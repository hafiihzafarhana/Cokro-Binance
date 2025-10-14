package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hafiihzafarhana/Cokro-Binance/domain/general"
	"github.com/hafiihzafarhana/Cokro-Binance/domain/market"
)

func GeneralRoutes(app *fiber.App, handler general.GeneralControllerInterface) {
	generalGroup := app.Group("/general")
	generalGroup.Get("/server-time", handler.CheckServerTime)
}

func AccountRoutes(app *fiber.App) {
	accountGroup := app.Group("/account")
	accountGroup.Get("/balance", )
}
func MarketRoutes(app *fiber.App, handler market.MarketControllerInterface) {
	marketGroup := app.Group("/market")
	marketGroup.Get("/order-book", handler.CheckOrderBook)
	marketGroup.Get("/recent-trade-lists", handler.CheckRecentTradeList)
	marketGroup.Get("/old-trade-lookup", handler.CheckOldTradeLookup)
	marketGroup.Get("/agregate-trade-lists", handler.CheckAgregateTradeList)
	marketGroup.Get("/candlestick-data", handler.CheckCandleStickData) // klines and uiklines
}
func TradingRoutes(app *fiber.App) {
	tradingGroup := app.Group("/trading")
	tradingGroup.Post("/order", )
}