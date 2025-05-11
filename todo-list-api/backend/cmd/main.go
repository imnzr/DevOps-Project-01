package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	databases "github.com/imnzr/DevOps-Project-01/todo-list-api/backend/database"
)

func main() {
	db, err := databases.GetConnection()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	fmt.Println("database connection successfully")
}
