package main

import (
	"github.com/firstneverrest/auth/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	v1 := app.Group("/v1", func(c *fiber.Ctx) error {
		c.Set("Version", "1.0.0")
		return c.Next()
	})

	v1.Get("/user/:id", handlers.UserProfile)
	v1.Post("/signup", handlers.Signup)
	v1.Post("/signin", handlers.Signin)
	v1.Post("/signout", handlers.Signout)
}
