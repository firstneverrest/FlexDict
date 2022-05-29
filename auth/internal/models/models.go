package models

import (
	"database/sql"
)

var DB *sql.DB

type User struct {
	Id       string `db:"id" json:"id"`
	Username string `db:"username" json:"username"`
	Password string `db:"password" json:"password"`
}

type Vocabulary struct {
	Id      string `db:"id" json:"id"`
	Title   string `db:"title" json:"title"`
	Meaning string `db:"meaning" json:"meaning"`
}
