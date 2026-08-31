package db

import (
	"database/sql"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
)

var DB *sql.DB

func InitDB() {
	var err error

	godotenv.Load()

	connStr := os.Getenv("DATABASE_URL")

	DB, err = sql.Open("pgx", connStr)
	if err != nil {
		panic("could not connect database")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
}
