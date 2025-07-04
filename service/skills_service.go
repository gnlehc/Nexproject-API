package service

import (
	"nexproject/helper"

	"github.com/gin-gonic/gin"
)

func SkillsRoute(private *gin.RouterGroup) {
	private.GET("/get-all-skills", helper.GetAllSkills)
	private.GET("/jobs/:job_id/skills", helper.GetAllSkillsByJobID)

}
