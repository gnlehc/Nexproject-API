package service

import (
	"nexproject/helper"

	"github.com/gin-gonic/gin"
)

func ProjectRoutes(private *gin.RouterGroup) {
	private.GET("/get-all-project-by-smeid", helper.GetAllProjectBySMEID)
	private.GET("/get-project-detail", helper.GetProjectByID)
	private.POST("/add-projects", helper.AddProject)
}
