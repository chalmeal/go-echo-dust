package env

import (
	"os"

	"github.com/joho/godotenv"
)

// DB接続情報
func DbInfo() (string, string) {

	if os.Getenv("APP_ENV") == "local" {
		godotenv.Load("env/.env.local")
		return os.Getenv("DB_DRIVER_LOCAL"), os.Getenv("DB_URL_LOCAL")
	} else {
		return os.Getenv("DB_DRIVER"), os.Getenv("DB_URL")
	}

}
