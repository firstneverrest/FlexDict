package models

import (
	"database/sql"
)

var DB *sql.DB

type User struct {
	Id       string
	Username string
	Password string
}
