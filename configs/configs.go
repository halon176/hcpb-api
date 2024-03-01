package configs

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	DB_USER string
	DB_PASS string
	DB_NAME string
	DB_HOST string
	DB_PORT string
)

func init() {
	godotenv.Load()

	DB_USER = os.Getenv("DB_USER")
	DB_PASS = os.Getenv("DB_PASS")
	DB_NAME = os.Getenv("DB_NAME")
	DB_HOST = os.Getenv("DB_HOST")
	DB_PORT = os.Getenv("DB_PORT")

}
