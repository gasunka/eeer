package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"todobackend/internal/db"
	"todobackend/internal/handlers"
	"todobackend/internal/todoservice"
)

func main() {
	// Восстановление после паники
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Recovered from panic: %v", r)
		}
	}()

	// Инициализация БД (ОДИН раз!)
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
