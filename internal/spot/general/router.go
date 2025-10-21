package general

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterGeneralRoutes(app *fiber.App, handler GeneralControllerInterface) {
	generalGroup := app.Group("/general")
	generalGroup.Get("/server-time", handler.CheckServerTime)
}