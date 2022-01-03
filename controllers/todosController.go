package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
	"todolist/models"
	"todolist/services"
)

func handleResponse(c *gin.Context, httpStatus int, templatePath string, data map[string]interface{}) {
	//default to text/html if no content-type specified
	if c.ContentType() == "" || strings.Contains(c.ContentType(), "text/html") {
		c.HTML(httpStatus, templatePath, data)
	} else if strings.Contains(c.ContentType(), "application/json") {
		c.JSON(httpStatus, data)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"errMsg": "Unknown content type",
		})
	}
}

func GetTodos(c *gin.Context) {
	todos := services.GetTodos()
	handleResponse(c, http.StatusOK, "todos/list", gin.H{
		"title": "All Todos",
		"count": len(*todos),
		"todos": todos,
	})
}

func GetTodo(c *gin.Context) {
	var todo models.Todo
	if err := c.ShouldBindUri(&todo); err != nil {
		switch err.(type) {
		case *strconv.NumError:
			handleResponse(c, http.StatusBadRequest, "todos/todo", gin.H{
				"errMsg": "Invalid id",
			})
		default:
			handleResponse(c, http.StatusBadRequest, "todos/todo", gin.H{
				"errMsg": "Unknown error",
			})
		}
		return
	}

	foundTodo := services.GetTodo(todo)
	if foundTodo == nil {
		handleResponse(c, http.StatusNotFound, "todos/todo", gin.H{
			"errMsg": "Record not found",
		})
		return
	}

	handleResponse(c, http.StatusOK, "todos/todo", gin.H{
		"id":          foundTodo.ID,
		"title":       foundTodo.Title,
		"description": foundTodo.Description,
		"dueDate":     foundTodo.DueDate,
	})
}
