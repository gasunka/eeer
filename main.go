package main

import (
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func initDB() {
	dsn := "host=localhost user=postgres password=yourpassword dbname=postgres port=5432 sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)

	}
	if err := db.AutoMigrate(&Task{}); err != nil {
		log.Fatalf("Could not migrate:%v", err)
	}

}

type Task struct {
	ID        string         `gorm:"primaryKey" json:"id"`
	Task      string         `json:"task"`
	IsDone    bool           `json:"is_done"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func getTask(c echo.Context) error {
	var tasks []Task
	if result := db.Where("deleted_at IS NULL").Find(&tasks); result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "could not fetch tasks"})
	}
	return c.JSON(http.StatusOK, tasks)
}

func postTask(c echo.Context) error {
	var req struct {
		Task   string `json:"task"`
		IsDone bool   `json:"is_done"`
	}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}
	newTask := Task{
		ID:     uuid.New().String(),
		Task:   req.Task,
		IsDone: req.IsDone,
	}
	if result := db.Create(&newTask); result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "could not create task"})
	}

	return c.JSON(http.StatusCreated, newTask)
}

func patchTask(c echo.Context) error {
	taskID := c.Param("id")
	var req struct {
		Task   *string `json:"task"`
		IsDone *bool   `json:"is_done"`
	}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}
	var task Task
	if result := db.First(&task, "id = ? AND deleted_at IS NULL", taskID); result.Error != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "k not found"})
	}
	updates := make(map[string]interface{})
	if req.Task != nil {
		updates["task"] = *req.Task
	}
	if req.IsDone != nil {
		updates["is_done"] = *req.IsDone
	}
	if len(updates) > 0 {
		if result := db.Model(&task).Updates(updates); result.Error != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "uld not update task"})

		}

	}
	return c.JSON(http.StatusOK, task)
}

func deleteTask(c echo.Context) error {
	taskID := c.Param("id")
	result := db.Where("id=?", taskID).Delete(&Task{})
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "could not delete task"})
	}
	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "task not found"})
	}
	return c.NoContent(http.StatusNoContent)
}

func main() {

	initDB()
	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.GET("/task", getTask)
	e.POST("/task", postTask)
	e.PATCH("/task/:id", patchTask)
	e.DELETE("/task/:id", deleteTask)

	e.Start("localhost:8080")
}
