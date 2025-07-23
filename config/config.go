package config

import (
	"fmt"

	"github.com/joho/godotenv"
)

func LoadEnvFile() {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Error loading file %v", err)
		return
	}
}
