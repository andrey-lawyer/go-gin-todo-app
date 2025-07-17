package main

import (
	"github.com/andrey-lawyer/go-gin-todo-app/config"
	"github.com/andrey-lawyer/go-gin-todo-app/handlers"
	"github.com/andrey-lawyer/go-gin-todo-app/infrastructure"
	"github.com/andrey-lawyer/go-gin-todo-app/middleware"
	"github.com/andrey-lawyer/go-gin-todo-app/repositories"
	"github.com/andrey-lawyer/go-gin-todo-app/services"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	infrastructure.InitMongo()
	db := infrastructure.MongoClient.Database("todo_db")

	userRepo := repositories.NewUserRepository(db)
	authService := services.NewAuthService(userRepo)
	authHandler := handlers.NewAuthHandler(authService)

	taskRepo := repositories.NewTaskRepository(db)
	taskService := services.NewTaskService(taskRepo)
	taskHandler := handlers.NewTaskHandler(taskService)

	router := gin.Default()

	// Подключаем recovery middleware
	router.Use(middleware.RecoveryWithLog())

	// cookie session
	store := cookie.NewStore([]byte(config.SessionSecret))
	router.Use(sessions.Sessions("auth-session", store))

	router.POST("/register", authHandler.Register)
	router.POST("/login", authHandler.Login)

	auth := router.Group("/")
	auth.Use(middleware.AuthRequired)
	{
		auth.GET("/tasks", taskHandler.GetTasks)
		auth.POST("/tasks", taskHandler.CreateTask)
	}

	router.Run(":" + config.Port)
}
