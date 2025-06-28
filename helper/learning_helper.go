package helper

import (
	"net/http"
	"nexproject/database"
	"nexproject/model"
	"nexproject/model/request"
	"nexproject/model/response"

	"github.com/gin-gonic/gin"
)

// GET learning by skills
// GET learning by search
// GET learning details

func GetLearningBySkills(c *gin.Context) {
	var req request.LearningBySkillsRequestDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.BaseResponseDTO{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
		return
	}

	var learnings []model.Learning
	if err := database.GlobalDB.
		Joins("JOIN learning_skills ON learning_skills.learning_learning_id = learnings.learning_id").
		Where("learning_skills.skill_skill_id IN ?", req.SkillIDs).
		Preload("Skills").
		Find(&learnings).Error; err != nil {
		c.JSON(http.StatusInternalServerError, response.BaseResponseDTO{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to get learnings by skills: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.GetLearningResponseDTO{
		Learnings: learnings,
		BaseResponseDTO: response.BaseResponseDTO{
			StatusCode: http.StatusOK,
			Message:    "Success",
		},
	})
}

func SearchLearning(c *gin.Context) {
	var req request.LearningBySearchRequestDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.BaseResponseDTO{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
		return
	}

	var learnings []model.Learning
	if err := database.GlobalDB.
		Where("title ILIKE ? OR content ILIKE ?", "%"+req.Keyword+"%", "%"+req.Keyword+"%").
		Preload("Skills").
		Find(&learnings).Error; err != nil {
		c.JSON(http.StatusInternalServerError, response.BaseResponseDTO{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to search learnings: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.GetLearningResponseDTO{
		Learnings: learnings,
		BaseResponseDTO: response.BaseResponseDTO{
			StatusCode: http.StatusOK,
			Message:    "Success",
		},
	})
}

func GetLearningDetails(c *gin.Context) {
	var req request.LearningDetailRequestDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.BaseResponseDTO{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
		return
	}

	var learning model.Learning
	if err := database.GlobalDB.
		Preload("Skills").
		First(&learning, "learning_id = ?", req.LearningID).Error; err != nil {
		c.JSON(http.StatusNotFound, response.BaseResponseDTO{
			StatusCode: http.StatusNotFound,
			Message:    "Learning not found: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.GetLearningDetailResponseDTO{
		Data: learning,
		BaseResponseDTO: response.BaseResponseDTO{
			StatusCode: http.StatusOK,
			Message:    "Success",
		},
	})
}
