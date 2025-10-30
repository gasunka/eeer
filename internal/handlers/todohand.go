package handlers

import (
	"net/http"
	"todobackend/internal/todoservice"

	"github.com/labstack/echo/v4"
)

type TodoHandler struct {
	service *todoservice.TodoService
}

func NewTodoHandler(service *todoservice.TodoService) *TodoHandler {
	return &TodoHandler{service: service}
}

func (h *TodoHandler) GetTasks(c echo.Context) error {
	tasks, err := h.service.GetAllTasks()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "could not fetch tasks"})
	}
	return c.JSON(http.StatusOK, tasks)
}

func (h *TodoHandler) CreateTask(c echo.Context) error {
	var req struct {
		Task   string `json:"task"`
		IsDone bool   `json:"is_done"`
	}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	task, err := h.service.CreateTask(req.Task, req.IsDone)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "could not create task"})
	}

	return c.JSON(http.StatusCreated, task)
}

func (h *TodoHandler) UpdateTask(c echo.Context) error {
	taskID := c.Param("id")

	var req struct {
		Task   *string `json:"task"`
		IsDone *bool   `json:"is_done"`
	}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	task, err := h.service.UpdateTask(taskID, req.Task, req.IsDone)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, task)
}

func (h *TodoHandler) DeleteTask(c echo.Context) error {
	taskID := c.Param("id")

	if err := h.service.DeleteTask(taskID); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "could not delete task"})
	}

	return c.NoContent(http.StatusNoContent)
}
