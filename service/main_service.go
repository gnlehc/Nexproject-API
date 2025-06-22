package service

import (
	"net/http"
	"nexproject/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(router *gin.Engine, db *gorm.DB) {
	api := router.Group("/api")
	{
		public := api.Group("/public")
		{
			AuthUserRoutes(public)
		}

		private := api.Group("/private")
		private.Use(middleware.JWTAuthMiddleware())
		{
			CVRoute(private)
			PortofolioRoute(private)
			JobRoutes(private)
			UserRoutes(private)
			SkillsRoute(private)
			ApplicationRoutes(private)
			ProjectRoutes(private)
		}
	}

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to Nexproject API by BNCC Magnolia",
		})
	})
}
