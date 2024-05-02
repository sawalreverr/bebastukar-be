package main

import (
	"github.com/sawalreverr/bebastukar-be/config"
	"github.com/sawalreverr/bebastukar-be/internal/database"
	"github.com/sawalreverr/bebastukar-be/internal/server"
)

func main() {
	conf := config.GetConfig()
	db := database.NewMySQLDatabase(conf)
	database.AutoMigrate(db)

	app := server.NewEchoServer(conf, db)
	app.Start()
}
