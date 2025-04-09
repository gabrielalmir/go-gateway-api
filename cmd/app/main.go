package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gabrielalmir/go-gateway-api/internal/repository"
	"github.com/gabrielalmir/go-gateway-api/internal/service"
	"github.com/gabrielalmir/go-gateway-api/internal/web/server"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		getEnv("DB_HOST"),
		getEnv("DB_PORT"),
		getEnv("DB_USER"),
		getEnv("DB_PASSWORD"),
		getEnv("DB_NAME"),
	)

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("Error pinging the database: %v", err)
	}

	accountRepository := repository.NewAccountRepository(db)
	accountService := service.NewAccountService(accountRepository)

	srv := server.NewServer(accountService, getEnv("PORT"))
	if err := srv.Start(); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

func getEnv(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Fatalf("Environment variable %s not set", key)
	}
	return value
}
