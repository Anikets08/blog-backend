package router

import (
	"github.com/gofiber/fiber/v2"
)

func ApiV1(app *fiber.App) {
	api := app.Group("/api/v1")
	api.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})
}
