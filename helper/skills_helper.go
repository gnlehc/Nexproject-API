package helper

import (
	"loom/database"
	"loom/model"
	"loom/model/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllSkills(c *gin.Context) {
	var skills []model.Skill
	database.GlobalDB.Find(&skills)
	c.JSON(200, gin.H{
		"skills": skills,
	})
}

func GetAllSkillsByJobID(c *gin.Context) {
	var job model.Job
	jobID := c.Param("job_id")

	db := database.GlobalDB

	if err := db.Preload("Skills").First(&job, "job_id = ?", jobID).Error; err != nil {
		c.JSON(http.StatusNotFound, response.BaseResponseDTO{
			Message:    "Job not found: " + err.Error(),
			StatusCode: http.StatusNotFound,
		})
		return
	}

	c.JSON(http.StatusOK, response.GetAllSkillResponseDTO{
		Skills: job.Skills,
		BaseResponse: response.BaseResponseDTO{
			StatusCode: http.StatusOK,
			Message:    "Success",
		},
	})
}
