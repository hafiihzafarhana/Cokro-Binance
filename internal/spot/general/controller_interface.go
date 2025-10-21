package general

import "github.com/gofiber/fiber/v2"

type GeneralControllerInterface interface {
	CheckServerTime(ctx *fiber.Ctx) error
}