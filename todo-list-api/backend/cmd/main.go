package main

import (
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	usercontroller "github.com/imnzr/DevOps-Project-01/todo-list-api/backend/controller/user-controller"
	databases "github.com/imnzr/DevOps-Project-01/todo-list-api/backend/database"
	userrepository "github.com/imnzr/DevOps-Project-01/todo-list-api/backend/repository/user-repository"
	userservice "github.com/imnzr/DevOps-Project-01/todo-list-api/backend/service/user-service"
	"github.com/julienschmidt/httprouter"
)

func main() {
	db, err := databases.GetConnection()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	fmt.Println("database connection successfully")

	userRepository := userrepository.NewUserRepository()
	userService := userservice.NewUserService(userRepository, db)
	userController := usercontroller.NewUserController(userService)

	router := httprouter.New()

	// Mencari semua user
	router.GET("/api/users", userController.FindByAll)
	// Mencari user berdasarkan id
	router.GET("/api/user/:userId", userController.FindById)
	// Membuat user
	router.POST("/api/auth/signup", userController.Create)
	// Login user
	router.POST("/api/auth/signin", userController.Login)

	server := http.Server{
		Addr:    "localhost:8085",
		Handler: router,
	}

	err = server.ListenAndServe()
	if err != nil {
		panic(fmt.Errorf("error listen server: %w", err))
	}

}
