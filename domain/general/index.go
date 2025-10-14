package general

import "github.com/gofiber/fiber/v2"

type GeneralServiceInterface interface {
	GetBinanceServerTime() (map[string]interface{}, error)
}

type GeneralControllerInterface interface {
	CheckServerTime(c *fiber.Ctx) error
}