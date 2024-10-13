package service

import (
	"loom/helper"

	"github.com/gin-gonic/gin"
)

func CVRoute(private *gin.RouterGroup) {
	private.POST("/upload-cv", helper.UploadCV)
}
