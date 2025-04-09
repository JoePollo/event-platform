package main

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func createEventTable(db *gorm.DB) {
	type EventTable struct {
		ID              int
		EventName       string
		Location        string
		DateTimeOfEvent time.Time
	}
	err := db.AutoMigrate(&EventTable{})
	if err != nil {
		panic(err)
	}
}

func main() {
	dsn := "host=localhost user=eventuser password=eventpass dbname=event_service port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("Failed to connect to database due to error:\n %v.", err))
	}
	createEventTable(db)
}
