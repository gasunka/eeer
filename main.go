package main

import (
	"fmt"
	"github.com/Knetic/govaluate"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

type Calculation struct {
	ID         string `json:"id"`
	Expression string `json:"expression"`
	Result     string `json:"result"`
}

type CalculationRequest struct {
	Expression string `json:"expression"`
}

var calculations = []Calculation{}

var task string

type TaskUpdateRequest struct {
	Task string `json:"task"`
}

func updateTask(c echo.Context) error {
	var req TaskUpdateRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid task"})
	}
	task = req.Task
	return c.JSON(http.StatusOK, map[string]string{"status": "update"})
}
func calculateExpression(expression string) (string, error) {
	expr, err := govaluate.NewEvaluableExpression(expression)
	if err != nil {
		return "", err
	}
	result, err := expr.Evaluate(nil)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", result), err

}
func getCalculations(c echo.Context) error {
	return c.JSON(http.StatusOK, calculations)
}
func postCalculations(c echo.Context) error {
	var req CalculationRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid operation"})

	}

	result, err := calculateExpression(req.Expression)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid expression"})
	}
	calc := Calculation{
		ID:         uuid.NewString(),
		Expression: req.Expression,
		Result:     result,
	}
	calculations = append(calculations, calc)
	return c.JSON(http.StatusCreated, calc)
}
func getTask(c echo.Context) error {
	return c.String(http.StatusOK, "hello"+task)
}
func main() {
	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.GET("/calculations", getCalculations)
	e.POST("/calculations", postCalculations)

	e.PUT("/task", updateTask)
	e.GET("/task", getTask)

	e.Start("localhost:8080")
}
