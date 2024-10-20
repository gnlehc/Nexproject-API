package service

import (
	"loom/helper"

	"github.com/gin-gonic/gin"
)

func JobRoutes(private *gin.RouterGroup) {
	private.POST("/post-a-job", helper.PostJob)
	private.GET("/get-all-job", helper.GetAllJobs)
	private.POST("/apply-job", helper.ApplyJob)
	private.POST("/get-job-posted", helper.GetAllJobsPostedBySME)
	private.GET("/jobs/:job_id", helper.GetJobByID)
}
