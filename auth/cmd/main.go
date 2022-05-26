package main

import (
	"github.com/gofiber/fiber/v2"
)

const portNumber = "8000"

func main() {
	app := fiber.New()

	// routes
	Routes(app)

	// middleware
	CorsHandler(app)
	Logger(app)

	app.Listen(":" + portNumber)
}
