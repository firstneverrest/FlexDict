package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	jwtware "github.com/gofiber/jwt/v2"
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

func JwtWare(path string, secret []byte, app *fiber.App) {
	app.Use(path, jwtware.New(jwtware.Config{
		SigningMethod: "HS256",
		SigningKey:    secret,
		SuccessHandler: func(c *fiber.Ctx) error {
			return c.Next()
		},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return fiber.ErrUnauthorized
		},
	}))
}
