// @title Binance API Service
// @version 1.0
// @description API to interact with Binance public endpoints
// @host localhost:8080
// @BasePath /
package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/hafiihzafarhana/Cokro-Binance/config"
	_ "github.com/hafiihzafarhana/Cokro-Binance/docs"
	"github.com/hafiihzafarhana/Cokro-Binance/internal/spot/general/bootstrap"
	"github.com/hafiihzafarhana/Cokro-Binance/shared/middleware/logging"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName:       "Cokro-Binance API",
		CaseSensitive: true,
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
		AllowHeaders:     "Origin, Content-Type, Accept, Content-Length, Accept-Language, Accept-Encoding, Connection, Authorization",
		AllowCredentials: false,
	}))

	app.Use(logging.Logging())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Ping Ping",
		})
	})

	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	bootConfig := config.BootConfig()

	bootstrap.SetupSpotGeneralModule(app)

	addr := fmt.Sprintf(":%d", bootConfig.AppPort)
	if err := app.Listen(addr); err != nil {
		log.Fatalf("Failed to start application: %v", err)
	}
}
