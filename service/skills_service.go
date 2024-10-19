package service

import (
	"loom/helper"

	"github.com/gin-gonic/gin"
)

func SkillsRoute(private *gin.RouterGroup) {
	private.GET("/get-all-skills", helper.GetAllSkills)
}
