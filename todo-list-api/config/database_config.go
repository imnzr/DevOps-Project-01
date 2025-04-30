package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetConnection() (*sql.DB, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading env")
	}

	DBUser := os.Getenv("DBUser")
	DBPassword := os.Getenv("DBPassword")
	DBHost := os.Getenv("DBHost")
	DBPort := os.Getenv("DBPort")
	DBName := os.Getenv("DBName")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		DBUser, DBPassword, DBHost, DBPort, DBName,
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open db: %w", err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping db: %w", err)
	}

	return db, nil
}
