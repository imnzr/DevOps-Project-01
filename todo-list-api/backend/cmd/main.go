package main

import (
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	todocontroller "github.com/imnzr/DevOps-Project-01/todo-list-api/backend/controller/todo-controller"
	usercontroller "github.com/imnzr/DevOps-Project-01/todo-list-api/backend/controller/user-controller"
	databases "github.com/imnzr/DevOps-Project-01/todo-list-api/backend/database"
	todorepository "github.com/imnzr/DevOps-Project-01/todo-list-api/backend/repository/todo-repository"
	userrepository "github.com/imnzr/DevOps-Project-01/todo-list-api/backend/repository/user-repository"
	todoservice "github.com/imnzr/DevOps-Project-01/todo-list-api/backend/service/todo-service"
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

	// user
	userRepository := userrepository.NewUserRepository()
	userService := userservice.NewUserService(userRepository, db)
	userController := usercontroller.NewUserController(userService)

	// todo
	todoRepository := todorepository.NewTodoRepository()
	todoService := todoservice.NewTodoService(todoRepository, db)
	todoController := todocontroller.NewTodoController(todoService)

	router := httprouter.New()

	/*
		USER END POINT
	*/

	// Mencari semua user
	router.GET("/api/users", userController.FindByAll)
	// Mencari user berdasarkan id
	router.GET("/api/user/:userId", userController.FindById)
	// Membuat user
	router.POST("/api/auth/signup", userController.Create)
	// Login user
	router.POST("/api/auth/signin", userController.Login)

	// Membuat todo
	router.POST("/todo/create", todoController.Create)
	// Menampilkan semua todo
	router.GET("/api/todos", todoController.FindByAll)
	// Mengubah title dari todo
	router.PUT("/todo/update-title/:todoId", todoController.UpdateByTitle)
	// Mengubah description dari todo
	router.PUT("/todo/update-description/:todoId", todoController.UpdateByDescription)

	server := http.Server{
		Addr:    "localhost:8085",
		Handler: router,
	}

	/*
		USER END POINT END
	*/

	/*
		TODO END POINT

	*/

	err = server.ListenAndServe()
	if err != nil {
		panic(fmt.Errorf("error listen server: %w", err))
	}

}
