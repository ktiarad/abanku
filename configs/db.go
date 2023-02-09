package configs

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	DB_HOST = "<your_config>"
	DB_PORT = "<your_config>"
	DB_USER = "<your_config>"
	DB_PASS = "<your_config>"
	DB_NAME = "<your_config>"
)

func ConnectDB() *sql.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", DB_HOST, DB_PORT, DB_USER, DB_PASS, DB_NAME)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	log.Default().Println("Connection db success")

	return db
}
