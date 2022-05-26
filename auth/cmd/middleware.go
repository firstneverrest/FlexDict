package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func CorsHandler(app *fiber.App) {
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET, POST, PUT, DELETE",
	}))
}

func Logger(app *fiber.App) {
	app.Use(logger.New(logger.Config{
		TimeZone: "Asia/Bangkok",
	}))
}
