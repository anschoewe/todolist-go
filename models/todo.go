package models

import (
	"time"
)

type TodoState int

const (
	Active TodoState = iota
	Archive
	Deleted
)

type Todo struct {
	//gorm.Model
	ID          uint `json:"id" uri:"id" binding:"required" gorm:"primarykey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time `gorm:"index"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"dueDate"`
	State       TodoState `json:"state"`
}
