package main

import (
	"backend/internal/todo/cache"
	"backend/internal/todo/controller"
	"backend/internal/todo/repository"
	"backend/internal/todo/service"
	"context"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

func main() {
	redisUri := os.Getenv("REDIS_URI")
	godotenv.Load()

	// Redis Client
	redisClient := redis.NewClient(&redis.Options{
		Addr: redisUri,
		DB:   0,
	})
	stats := redisClient.Ping(context.Background())
	if stats.Err() != nil {
		panic(stats.Err())
	}

	// Redis Cache
	redisCache := cache.NewRedisCache(redisClient, "allTodos", 10*time.Second)

	// Router
	router := gin.Default()

	// CORS
	config := corsConfig()
	router.Use(cors.New(config))

	// Factories
	todoRepository := repository.NewTodoRepository()
	todoService := service.NewTodoService(todoRepository, redisCache)
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
