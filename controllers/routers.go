package controllers

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"time"
)

func InitRouter(templatePath string) *gin.Engine {
	// Creates a gin engine with default middleware:
	// logger and recovery (crash-free) middleware
	engine := gin.Default()
	engine.SetTrustedProxies(nil)
	engine.SetFuncMap(template.FuncMap{
		"formatAsDate": formatAsDate,
	})
	engine.LoadHTMLGlob(templatePath)

	engine.GET("/", index)
	engine.GET("/health", health)
	engine.GET("/todos/", GetTodos)
	engine.GET("/todos/:id", GetTodo)

	return engine
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
