package main

import (
	"todolist/controllers"
	"todolist/models"
)

func main() {
	models.InitDb()
	router := controllers.InitRouter("templates/**/*")
	router.Run(":8080")
}
