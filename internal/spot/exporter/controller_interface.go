package exporter

import "github.com/gofiber/fiber/v2"

type ExporterControllerInterface interface {
	ExportBinanceCandleStickData(c *fiber.Ctx) error
}