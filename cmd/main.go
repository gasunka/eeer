package main

import (
	"todobackend/internal/db"
	"todobackend/internal/handlers"
	"todobackend/internal/todoservice"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Инициализация БД
	db.Init()

	// Инициализация сервисов
	todoRepo := todoservice.NewTodoRepository(db.GetDB())
	todoSvc := todoservice.NewTodoService(todoRepo)
	todoHandler := handlers.NewTodoHandler(todoSvc)

	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	// Маршруты
	e.GET("/task", todoHandler.GetTasks)
	e.POST("/task", todoHandler.CreateTask)
	e.PATCH("/task/:id", todoHandler.UpdateTask)
	e.DELETE("/task/:id", todoHandler.DeleteTask)

	e.Start("localhost:8080")
}
