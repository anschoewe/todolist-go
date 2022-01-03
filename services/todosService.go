package services

import (
	"todolist/models"
)

func GetTodos() *[]models.Todo {
	var todos []models.Todo
	models.DB.Find(&todos)
	return &todos
}

func GetTodo(wantedTodo models.Todo) *models.Todo {
	var foundTodo models.Todo
	if err := models.DB.Where("id = ?", wantedTodo.ID).First(&foundTodo).Error; err != nil {
		return nil
	}

	return &foundTodo
}
