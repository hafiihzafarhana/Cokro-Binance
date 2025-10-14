// @title Cokro Binance
// @version 1.0
// @description API server for market & trading data with Binance API
// @host localhost:8080
// @BasePath /
package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/hafiihzafarhana/Cokro-Binance/config"
	_ "github.com/hafiihzafarhana/Cokro-Binance/docs"
	cGeneral "github.com/hafiihzafarhana/Cokro-Binance/domain/general/controller"
	sGeneral "github.com/hafiihzafarhana/Cokro-Binance/domain/general/service"
	cMarket "github.com/hafiihzafarhana/Cokro-Binance/domain/market/controller"
	sMarket "github.com/hafiihzafarhana/Cokro-Binance/domain/market/service"
	"github.com/hafiihzafarhana/Cokro-Binance/middleware/logging"
	"github.com/hafiihzafarhana/Cokro-Binance/routes"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName:       "Cokro-Binance API",
		CaseSensitive: true,
	})

	app.Use(cors.New(cors.Config{
        AllowHeaders:     "Origin, Content-Type, Accept, Content-Length, Accept-Language, Accept-Encoding, Connection, Authorization",
		AllowOrigins:     "*",
		AllowCredentials: false,
		AllowMethods:     "GET, POST, HEAD, PUT, DELETE, PATCH, OPTIONS",
    }))

	var bootConfig = config.BootConfig()

	app.Use(logging.Logging())
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Ping Ping",
		})
	})
	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	generalService := sGeneral.NewGeneralService()
	generalController := cGeneral.NewGeneralController(generalService)

	marketService := sMarket.NewMarketService()
	marketController := cMarket.NewMarketController(marketService)

	routes.GeneralRoutes(app, generalController)
	routes.MarketRoutes(app, marketController)
	// routes.AccountRoutes(app)
	// routes.MarketRoutes(app)
	// routes.TradingRoutes(app)

	addr := fmt.Sprintf(":%d", bootConfig.AppPort)
	if err := app.Listen(addr).Error(); err != addr {
		panic("Appilaction failed to start")
	}
}