package database

import (
	"log"

	"github.com/sawalreverr/bebastukar-be/internal/entity"
)

func AutoMigrate(db Database) {
	if err := db.GetDB().AutoMigrate(
		&entity.User{},
		&entity.Discussions{},
		&entity.DiscussionImages{},
		&entity.DiscussionComments{},
		&entity.DiscussionReplyComments{},
	); err != nil {
		log.Fatal("Database Migration Failed!")
	}

	log.Println("Database Migration Success")
}
