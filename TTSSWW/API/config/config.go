package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func DBURL() string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	DBHost := os.Getenv("SUPABASE_HOST")
	DBUser := os.Getenv("SUPABASE_USER")
	DBPassword := os.Getenv("SUPABASE_PASSWORD")
	DBPort := os.Getenv("SUPABASE_PORT")
	DBName := os.Getenv("SUPABASE_DB")

	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable&statement_cache_mode=describe", DBUser, DBPassword, DBHost, DBPort, DBName)
}
