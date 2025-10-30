package todoservice

import (
	"gorm.io/gorm"
	"time"
)

// Task представляет модель задачи в базе данных
type Task struct {
	ID        string         `json:"id" gorm:"primaryKey"`
	Task      string         `json:"task" gorm:"not null"`
	IsDone    bool           `json:"is_done" gorm:"default:false"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

func (Task) TableName() string {
	return "tasks"
}
