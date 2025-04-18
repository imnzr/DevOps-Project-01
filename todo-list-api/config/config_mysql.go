package config

import (
	"database/sql"

	"github.com/imnzr/DevOps-Project-01/todo-list-api/config"
)

func DatabaseConnection(cfg *config.Config) {
	result, err := sql.Open("mysql", "imznr:2165")
}
