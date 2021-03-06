package handlers

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/firstneverrest/auth/internal/models"
	"github.com/firstneverrest/auth/internal/utils"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type SignupRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SigninRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AddVocabularyRequest struct {
	Title   string `json:"title"`
	Meaning string `json:"meaning"`
}

const jwtSecret = "mySecret"

func Signup(c *fiber.Ctx) error {
	request := SignupRequest{}
	err := c.BodyParser(&request)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	if request.Username == "" || request.Password == "" {
		return fiber.NewError(fiber.StatusBadRequest, "Username and password are required")
	}

	password, err := bcrypt.GenerateFromPassword([]byte(request.Password), 10)
	if err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, err.Error())
	}

	// check if username already exists
	user := models.User{}
	query := "SELECT Id, Username, Password FROM Users WHERE Username = ?"
	err = models.DB.QueryRow(query, request.Username).Scan(&user.Id, &user.Username, &user.Password)
	if err == nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, "Username already exists")
	}

	// add user
	query = "INSERT INTO Users (Id, Username, Password) VALUES (?, ?, ?)"
	randomId := utils.RandStringBytes(13)

	_, err = models.DB.Exec(query, randomId, request.Username, string(password))
	if err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, err.Error())
	}

	// create new table for users
	query = "CREATE TABLE " + randomId + "_Vocabulary (Id NVARCHAR(13) NOT NULL, Title NVARCHAR(50), Meaning NVARCHAR(300), PRIMARY KEY (Id))"
	_, err = models.DB.Query(query)
	if err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"msg": "User created successfully",
	})
}

func Signin(c *fiber.Ctx) error {
	request := SigninRequest{}
	err := c.BodyParser(&request)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	if request.Username == "" || request.Password == "" {
		return fiber.NewError(fiber.StatusBadRequest, "Username and password are required")
	}

	user := models.User{}
	query := "SELECT Id, Username, Password FROM Users WHERE Username = ?"
	err = models.DB.QueryRow(query, request.Username).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Invalid username or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Invalid username or password")
	}

	// jwt
	claims := jwt.MapClaims{}
	claims["id"] = user.Id
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	claims["iat"] = time.Now().Unix()
	claims["issuer"] = "flexdict"

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := jwtToken.SignedString([]byte(jwtSecret))
	if err != nil {
		return fiber.ErrInternalServerError
	}

	return c.JSON(fiber.Map{
		"jwt":      token,
		"username": user.Username,
	})
}

func GetVocabulary(c *fiber.Ctx) error {
	jwtUserToken := c.Get("Authorization")[7:]

	if jwtUserToken == "" {
		return fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
	}

	// parse JWT to claims
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(jwtUserToken, claims, func(token *jwt.Token) (interface{}, error) {
		return GetJWTSecret(), nil
	})
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
	}

	// get user id
	userId := claims["id"].(string)

	vocabulary := make([]*models.Vocabulary, 0)
	query := "SELECT Id, Title, Meaning FROM " + userId + "_Vocabulary"
	rows, err := models.DB.Query(query)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, err.Error())
	}

	for rows.Next() {
		vocab := new(models.Vocabulary)
		err = rows.Scan(&vocab.Id, &vocab.Title, &vocab.Meaning)
		if err != nil {
			return fiber.NewError(fiber.StatusUnauthorized, err.Error())
		}
		vocabulary = append(vocabulary, vocab)
	}

	return c.Status(fiber.StatusOK).JSON(vocabulary)
}

func AddVocabulary(c *fiber.Ctx) error {
	jwtUserToken := c.Get("Authorization")[7:]

	if jwtUserToken == "" {
		return fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
	}

	// parse JWT to claims
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(jwtUserToken, claims, func(token *jwt.Token) (interface{}, error) {
		return GetJWTSecret(), nil
	})
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
	}

	// get user id
	userId := claims["id"].(string)
	fmt.Println(userId)

	request := AddVocabularyRequest{}
	err = c.BodyParser(&request)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	if request.Title == "" {
		return fiber.NewError(fiber.StatusBadRequest, "Title and Meaning are required")
	}

	// add vocab to db
	randomId := utils.RandStringBytes(13)
	fmt.Println("passed")
	query := "INSERT INTO " + userId + "_Vocabulary (Id, Title, Meaning) VALUES (?, ?, ?)"
	_, err = models.DB.Query(query, randomId, request.Title, request.Meaning)
	if err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"msg": "Vocabulary added successfully",
	})
}

func GetJWTSecret() []byte {
	return []byte(jwtSecret)
}
