package service

import (
	"nexproject/helper"

	"github.com/gin-gonic/gin"
)

func ApplicationRoutes(private *gin.RouterGroup) {
	private.POST("/update-application-status", helper.UpdateApplicationStatus)
	private.POST("/get-all-applicants-by-job-id", helper.GetAllAppIDByJobID)
	private.POST("/get-all-jobs-talent-applied", helper.GetAllAppIDByTalentID)
	private.GET("/get-all-applications", helper.GetAllApplications)
}
