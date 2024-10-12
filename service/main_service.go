package service

import (
	"help/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		public := api.Group("/public")
		{
			UserRoutes(public)
		}

		private := api.Group("/private")
		private.Use(middleware.JWTAuthMiddleware())
		{
			CVRoute(private)
			PortofolioRoute(private)
		}
	}

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to Loom API by BNCC Magnolia",
		})
	})
}
