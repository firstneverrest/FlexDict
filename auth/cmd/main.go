package main

import (
	"fmt"
	"log"
	"os"

	"database/sql"

	// "github.com/firstneverrest/auth/internal/handlers"
	"github.com/firstneverrest/auth/internal/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

var DB *sql.DB

func main() {
	app := fiber.New()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")

	// mysql
	db, err := sql.Open("mysql", "root:root@tcp(mysql:3306)/flexdict")
	if err != nil {
		log.Fatal(err)
	}

	models.DB = db
	fmt.Println("Connected to database")

	defer db.Close()

	// routes
	Routes(app)

	// middleware
	CorsHandler(app)
	// JwtWare("/user/:id/vocab", handlers.GetJWTSecret(), app)
	Logger(app)

	app.Listen(":" + port)
}
