package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET, POST, PUT, DELETE",
	}))

	// add logging
	app.Use(logger.New(logger.Config{
		TimeZone: "Asia/Bangkok",
	}))

	// Group
	v1 := app.Group("/v1", func(c *fiber.Ctx) error {
		c.Set("Version", "1.0.0")
		return c.Next()
	})
	v1.Get("/user/:id", UserProfile)
	v1.Post("/signup", Signup)
	v1.Post("/signin", Signin)
	v1.Post("/signout", Signout)

	v1.Get("/env", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"Version":    c.Get("Version"),
			"BaseURL":    c.BaseURL(),
			"Hostname":   c.Hostname(),
			"IP":         c.IP(),
			"IPs":        c.IPs(),
			"Protocol":   c.Protocol(),
			"Subdomains": c.Subdomains(),
		})
	})

	app.Listen(":8000")
}

func Signup(c *fiber.Ctx) error {
	return c.SendString("Signup")
}

func Signin(c *fiber.Ctx) error {
	return c.SendString("Signin")
}

func Signout(c *fiber.Ctx) error {
	return c.SendString("Signout")
}

func UserProfile(c *fiber.Ctx) error {
	id := c.Params("id")
	return c.SendString("UserProfile: " + id)
}
