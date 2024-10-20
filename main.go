package main

import (
	"log"
	"loom/database"
	"loom/database/migrate"
	"loom/service"

	"github.com/gin-contrib/cors"
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

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	service.SetupRoutes(r, database.GlobalDB)
	r.SetTrustedProxies(nil)
	r.Run(":8080")
}
