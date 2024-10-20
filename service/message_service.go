package service

import (
	"loom/helper"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func MessageRoutes(private *gin.RouterGroup, db *gorm.DB) {
	private.GET("/ws", func(c *gin.Context) {
		helper.HandleWebSocket(c, db)
	})
	go helper.HandleMessages(db)

}
