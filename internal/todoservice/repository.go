package todoservice

import "gorm.io/gorm"

type TodoRepository interface {
	GetAll() ([]Task, error)
	GetByID(id string) (*Task, error)
	Create(task *Task) error
	Update(task *Task) error
	Delete(id string) error
}

type todoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) TodoRepository {
	return &todoRepository{db: db}
}

func (r *todoRepository) GetAll() ([]Task, error) {
	var tasks []Task
	result := r.db.Where("deleted_at IS NULL").Find(&tasks)
	return tasks, result.Error
}

func (r *todoRepository) GetByID(id string) (*Task, error) {
	var task Task
	result := r.db.Where("id = ? AND deleted_at IS NULL", id).First(&task)
	if result.Error != nil {
		return nil, result.Error
	}
	return &task, nil
}

func (r *todoRepository) Create(task *Task) error {
	result := r.db.Create(task)
	return result.Error
}

func (r *todoRepository) Update(task *Task) error {
	result := r.db.Save(task)
	return result.Error
}

func (r *todoRepository) Delete(id string) error {
	result := r.db.Where("id = ?", id).Delete(&Task{})
	return result.Error
}
