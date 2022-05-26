package handlers

import (
	"github.com/gofiber/fiber/v2"
)

type user struct {
	Username string `json:"username"`
	Password string `json:"password"`
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
