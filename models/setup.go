package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"time"
)

var DB *gorm.DB

func InitDb() {
	// use in-memory db
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// Migrate the schema
	db.AutoMigrate(&Todo{})

	DB = db

	//seed some data
	db.Create(&Todo{
		Title:       "Pay bills",
		Description: "Gas, Electricity, Sewage",
		DueDate:     time.Date(2022, 01, 01, 9, 0, 0, 0, time.UTC),
		State:       Active,
	})
	db.Create(&Todo{
		Title:       "Wash floor",
		Description: "mop",
		DueDate:     time.Now().Add(time.Hour * 48),
		State:       Active,
	})
	db.Create(&Todo{
		Title:       "Donate clothes",
		Description: "2T and crib",
		DueDate:     time.Now().Add(time.Hour * 48),
		State:       Active,
	})
}

func ClearDb() {
	if DB == nil {
		return
	}

	DB.Exec("DELETE FROM Todos")
}
