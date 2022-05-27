package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	app := fiber.New()

	// routes
	Routes(app)

	// middleware
	CorsHandler(app)
	Logger(app)

	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	port := os.Getenv("PORT")

	app.Listen(":" + port)
}
