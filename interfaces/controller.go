package interfaces

import "context"

type Controller interface {
	GetTodo(c *context.Context)
	GetTodos(c *context.Context)
}

type todoController struct {
}

//func NewTodoController() *Controller {
//	return &todoController{}
//}
