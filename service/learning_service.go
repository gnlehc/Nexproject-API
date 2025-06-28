package service

import (
	"nexproject/helper"

	"github.com/gin-gonic/gin"
)

func LearningRoutes(private *gin.RouterGroup) {
	private.POST("/learnings/skills", helper.GetLearningBySkills)
	private.POST("/learnings/search", helper.SearchLearning)
	private.POST("/learnings/details", helper.GetLearningDetails)
}
