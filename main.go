package main

import (
	"github.com/anikets08/blog-backend/database"
	"github.com/anikets08/blog-backend/router"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	database.InitDB()
	router.ApiV1(app)
	app.Listen(":3000")
}
