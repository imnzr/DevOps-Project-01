package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var DB *sql.DB

func DatabaseConnection() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	DBUser := os.Getenv("DBUser")
	DBPass := os.Getenv("DBPass")
	DBHost := os.Getenv("DBHost")
	DBName := os.Getenv("DBName")
	DBPort := os.Getenv("DBPort")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		DBUser,
		DBPass,
		DBHost,
		DBPort,
		DBName,
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Database Connection failed: ", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Database unreachable: ", err)
	}

	DB = db
	fmt.Println("Database connected successfully")
}
