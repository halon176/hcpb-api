package configs

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

var (
	DB_USER string
	DB_PASS string
	DB_NAME string
	DB_HOST string
	DB_PORT string

	API_KEY string

	LOG_LEVEL    string = "info"
	SERVICE_PORT string
)

func init() {
	godotenv.Load()

	DB_USER = os.Getenv("DB_USER")
	DB_PASS = os.Getenv("DB_PASS")
	DB_NAME = os.Getenv("DB_NAME")
	DB_HOST = os.Getenv("DB_HOST")
	DB_PORT = os.Getenv("DB_PORT")

	API_KEY = os.Getenv("API_KEY")
	SERVICE_PORT = os.Getenv("SERVICE_PORT")
	if SERVICE_PORT == "" {
		SERVICE_PORT = "7777"
	}

	LOG_LEVEL = os.Getenv("LOG_LEVEL")

	if LOG_LEVEL == "" {
		LOG_LEVEL = "debug"
	}
	LOG_LEVEL = strings.ToLower(LOG_LEVEL)

	switch LOG_LEVEL {
	case "debug", "info", "warn", "error":
		break
	default:
		log.Fatalf("LOG_LEVEL %s is not valid", LOG_LEVEL)
	}

	initLogger(LOG_LEVEL)

}
