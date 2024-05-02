package database

import (
	"log"

	"github.com/sawalreverr/bebastukar-be/internal/entity"
)

func AutoMigrate(db Database) {
	if err := db.GetDB().AutoMigrate(&entity.User{}); err != nil {
		log.Fatal("Database Migration Failed!")
	}

	log.Println("Database Migration Success")
}
