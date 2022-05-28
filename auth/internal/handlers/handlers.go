package handlers

import (
	"strconv"

	"github.com/firstneverrest/auth/internal/models"
	"github.com/firstneverrest/auth/internal/utils"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       string `db:"id" json:"id"`
	Username string `db:"username" json:"username"`
	Password string `db:"password" json:"password"`
}

type SignupRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SigninRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Signup(c *fiber.Ctx) error {
	request := SignupRequest{}
	err := c.BodyParser(&request)
	if err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, "Invalid request body")
	}

	if request.Username == "" || request.Password == "" {
		return fiber.NewError(fiber.StatusUnprocessableEntity, "Username and password are required")
	}

	password, err := bcrypt.GenerateFromPassword([]byte(request.Password), 10)
	if err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, err.Error())
	}

	query := "INSERT INTO Users (Id, Username, Password) VALUES (?, ?, ?)"
	randomId := utils.RandStringBytes(13)

	result, err := models.DB.Exec(query, randomId, request.Username, string(password))
	if err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, err.Error())
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, err.Error())
	}

	user := User{
		Id:       strconv.Itoa(int(id)),
		Username: request.Username,
		Password: string(password),
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}

func Signin(c *fiber.Ctx) error {
	request := SigninRequest{}
	err := c.BodyParser(&request)
	if err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, "Invalid request body")
	}

	if request.Username == "" || request.Password == "" {
		return fiber.NewError(fiber.StatusUnprocessableEntity, "Username and password are required")
	}

	user := User{}
	query := "SELECT Id, Username, Password FROM Users WHERE Username = ?"
	err = models.DB.QueryRow(query, request.Username).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Invalid username or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Invalid username or password")
	}

	return c.SendStatus(fiber.StatusOK)
}

func Signout(c *fiber.Ctx) error {
	return c.SendString("Signout")
}

func UserProfile(c *fiber.Ctx) error {
	id := c.Params("id")
	return c.SendString("UserProfile: " + id)
}
