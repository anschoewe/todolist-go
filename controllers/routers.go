package controllers

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"time"
)

func InitRouter(templatePath string) *gin.Engine {
	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()
	router.SetFuncMap(template.FuncMap{
		"formatAsDate": formatAsDate,
	})
	router.LoadHTMLGlob(templatePath)

	router.GET("/", index)
	router.GET("/health", health)
	router.GET("/todos/", GetTodos)
	router.GET("/todos/:id", GetTodo)

	return router
}

func formatAsDate(t time.Time) string {
	return t.Format("Jan 2, 2006")
}

func health(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.Header("Content-Encoding", "utf-8")
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func index(c *gin.Context) {
	c.HTML(http.StatusOK, "index/index", gin.H{
		"title": "Andrew's Todos",
	})
}
