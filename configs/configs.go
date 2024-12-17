package configs

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type (
	DBConfig struct {
		DBHost    string
		DBPort    string
		DBUser    string
		DBPass    string
		DBName    string
		DbMaxConn int
	}

	configs struct {
		DBConfig
	}
)

func InitConfigs() *configs {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	return &configs{
		DBConfig: DBConfig{
			DBHost: os.Getenv("DB_HOST"),
			DBPort: os.Getenv("DB_PORT"),
			DBUser: os.Getenv("DB_USER"),
			DBPass: os.Getenv("DB_PASS"),
			DBName: os.Getenv("DB_NAME"),
			DbMaxConn: func() int {
				maxconn, err := strconv.Atoi(os.Getenv("DB_MAXCONN"))
				if err != nil {
					panic(err)
				}
				return maxconn
			}(),
		},
	}
}
