package todoservice

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

type TodoService struct {
	repo TodoRepository
}

func NewTodoService(repo TodoRepository) *TodoService {
	return &TodoService{repo: repo}
}

func (s *TodoService) GetAllTasks() ([]Task, error) {
	return s.repo.GetAll()
}

func (s *TodoService) CreateTask(taskText string, isDone bool) (*Task, error) {
	task := &Task{
		ID:        uuid.New().String(),
		Task:      taskText,
		IsDone:    isDone,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := s.repo.Create(task); err != nil {
		return nil, err
	}
	return task, nil
}

func (s *TodoService) UpdateTask(id string, taskText *string, isDone *bool) (*Task, error) {
	task, err := s.repo.GetByID(id)
	if err != nil {
		return nil, errors.New("task not found")
	}

	if taskText != nil {
		task.Task = *taskText
	}
	if isDone != nil {
		task.IsDone = *isDone
	}
	task.UpdatedAt = time.Now()

	if err := s.repo.Update(task); err != nil {
		return nil, errors.New("could not update task")
	}

	return task, nil
}

func (s *TodoService) DeleteTask(id string) error {
	return s.repo.Delete(id)
}
