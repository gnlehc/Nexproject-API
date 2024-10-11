package main

import (
	"help/database"
	"help/database/migrate"
	"help/service"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	err := database.DatabaseConnection()
	if err != nil {
		log.Fatalln("Could not connect to database")
	}
	err = migrate.DBMigrate(database.GlobalDB)
	if err != nil {
		log.Fatalln("Failed to migrate database tables:", err)
	}
	gin.SetMode(gin.DebugMode)
	r := gin.Default()
	service.SetupRoutes(r)
	r.SetTrustedProxies(nil)
	r.Run(":8080")
}
