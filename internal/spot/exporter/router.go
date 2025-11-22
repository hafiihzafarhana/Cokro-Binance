package exporter

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterGeneralRoutes(app *fiber.App, handler ExporterControllerInterface) {
	generalGroup := app.Group("/exporter")
	generalGroup.Get("/candlestick-data", handler.ExportBinanceCandleStickData)
}