package configs

import (
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

var (
	DB_USER string
	DB_PASS string
	DB_NAME string
	DB_HOST string
	DB_PORT string

	DB_MAX_CONNS           int32
	DB_MIN_CONNS           int32
	DB_MAX_CONN_LIFETIME   time.Duration
	DB_MAX_CONN_IDLE_TIME  time.Duration
	DB_HEALTH_CHECK_PERIOD time.Duration

	API_KEY string

	LOG_LEVEL    string = "info"
	SERVICE_PORT string
)

func GetEnvInt32(key string, defaultVal int32) int32 {
	if val, exists := os.LookupEnv(key); exists {
		if i, err := strconv.Atoi(val); err == nil {
			return int32(i)
		}
	}
	return defaultVal
}

func GetEnvDurationMinutes(key string, defaultVal time.Duration) time.Duration {
	if val, exists := os.LookupEnv(key); exists {
		if i, err := strconv.Atoi(val); err == nil {
			return time.Duration(i) * time.Minute
		}
	}
	return defaultVal
}

func init() {
	godotenv.Load()

	DB_USER = os.Getenv("DB_USER")
	DB_PASS = os.Getenv("DB_PASS")
	DB_NAME = os.Getenv("DB_NAME")
	DB_HOST = os.Getenv("DB_HOST")
	DB_PORT = os.Getenv("DB_PORT")

	DB_MAX_CONNS = GetEnvInt32("DB_MAX_CONNS", 10)
	DB_MIN_CONNS = GetEnvInt32("DB_MIN_CONNS", 2)

	DB_MAX_CONN_LIFETIME = GetEnvDurationMinutes("DB_MAX_CONN_LIFETIME", 30*time.Minute)
	DB_MAX_CONN_IDLE_TIME = GetEnvDurationMinutes("DB_MAX_CONN_IDLE_TIME", 5*time.Minute)
	DB_HEALTH_CHECK_PERIOD = GetEnvDurationMinutes("DB_HEALTH_CHECK_PERIOD", 1*time.Minute)

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
