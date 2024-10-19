package helper

import (
	"loom/database"
	"loom/model"

	"github.com/gin-gonic/gin"
)

func GetAllSkills(c *gin.Context) {
	var skills []model.Skill
	database.GlobalDB.Find(&skills)
	c.JSON(200, gin.H{
		"skills": skills,
	})
}
