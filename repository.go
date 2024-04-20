package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

func getEnvDBUser() string {
	godotenv.Load(".env")
	return os.Getenv("USER")
}
func getEnvDB() string {
	godotenv.Load(".env")
	return os.Getenv("DB")
}
func getEnvDBPassword() string {
	godotenv.Load(".env")
	return os.Getenv("PASSWORD")
}

func getDB() *sql.DB {
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", getEnvDBUser(), getEnvDBPassword(), getEnvDB())
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
