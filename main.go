package main

import (
	"gin-api/app/config"
	"gin-api/app/route"
	"gin-api/database"
)

func main() {
	config := config.LoadConfig()
	// Connect database
	db := database.NewConnection(config)

	// Load router
	route.LoadRouter(config, db)

}
