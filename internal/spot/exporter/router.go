package exporter

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterExporterRoutes(app *fiber.App, handler ExporterControllerInterface) {
	exporterGroup := app.Group("/exporter")
	exporterGroup.Get("/candlestick-data", handler.ExportBinanceCandleStickData)
}