package main

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

type Task struct {
	ID     string `json:"id"`
	Text   string `json:"text"`
	Status string `json:"status"`
}

type TaskUpdateRequest struct {
	Text   string `json:"text"`
	Status string `json:"status"`
}

var tasks []Task

func getTask(c echo.Context) error {
	return c.JSON(http.StatusOK, tasks)
}

func postTask(c echo.Context) error {
	var req TaskUpdateRequest
	if err := json.NewDecoder(c.Request().Body).Decode(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid task"})
	}
	newTask := Task{
		ID:     uuid.New().String(),
		Text:   req.Text,
		Status: req.Status,
	}
	tasks = append(tasks, newTask)

	return c.JSON(http.StatusOK, tasks)
}

func patchTask(c echo.Context) error {
	taskID := c.Param("id")
	var req TaskUpdateRequest
	if err := json.NewDecoder(c.Request().Body).Decode(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid task"})
	}
	for i, task := range tasks {
		if task.ID == taskID {
			if req.Text != "" {
				tasks[i].Text = req.Text
			}
			if req.Status != "" {
				tasks[i].Status = req.Status
			}

			return c.JSON(http.StatusOK, tasks[i])
		}

	}
	return c.JSON(http.StatusNotFound, map[string]string{"error": "task not found"})
}

func deleteTask(c echo.Context) error {
	taskID := c.Param("id")
	for i, task := range tasks {
		if task.ID == taskID {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return c.NoContent(http.StatusNoContent)
		}
	}
	return c.JSON(http.StatusNotFound, map[string]string{"error": "task id not found"})
}

func main() {
	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.GET("/task", getTask)
	e.POST("/task", postTask)
	e.PATCH("/task/:id", patchTask)
	e.DELETE("/task/:id", deleteTask)

	e.Start("localhost:8080")
}
