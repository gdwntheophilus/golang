package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func LoadConfig() string {
	// fmt.Println("Loading Config")
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(err)
	}
	return os.Getenv("MONGO_DB_URL")
}
