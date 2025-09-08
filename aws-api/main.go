package main

import (
	"backend/internal/todo/controller"
	"backend/internal/todo/repository"
	"backend/internal/todo/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// CORS
	config := corsConfig()
	router.Use(cors.New(config))

	// Factories
	todoRepository := repository.NewTodoRepository()
	todoService := service.NewTodoService(todoRepository)
	todoController := controller.NewTodoController(todoService)

	// Routes
	router.POST("/todos", todoController.CreateTodo)
	router.GET("/todos", todoController.GetAllTodos)

	// Run
	router.Run(":8080")
}

// CORS Config
func corsConfig() cors.Config {
	return cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		AllowCredentials: true,
	}
}
