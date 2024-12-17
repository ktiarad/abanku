package infra

import (
	"abanku/configs"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func ConnectDB(configs *configs.DBConfig) *sql.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		configs.DBHost,
		configs.DBPort,
		configs.DBUser,
		configs.DBPass,
		configs.DBName)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	db.SetMaxOpenConns(configs.DbMaxConn)

	return db
}
