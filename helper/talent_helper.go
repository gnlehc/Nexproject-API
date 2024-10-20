package helper

import (
	"errors"
	"fmt"
	"loom/database"
	"loom/model"
	"loom/model/request"
	"loom/model/response"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func LoginTalent(c *gin.Context) {
	var req request.TalentLoginRequestDTO
	var talent model.Talent

	if err := c.ShouldBindJSON(&req); err != nil {
		res := response.TalentLoginResponseDTO{StatusCode: 400, Message: "Invalid Request"}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	if err := database.GlobalDB.Where("Email = ?", strings.ToLower(req.Email)).First(&talent).Error; err != nil {
		res := response.TalentLoginResponseDTO{StatusCode: 401, Message: "Credentials not matched"}
		c.JSON(http.StatusUnauthorized, res)
		return
	}

	if !CheckPasswordHash(req.Password, talent.Password) {
		res := response.TalentLoginResponseDTO{StatusCode: 401, Message: "Invalid Password"}
		c.JSON(http.StatusUnauthorized, res)
		return
	}

	token, err := GenerateJWT(talent.TalentID, "Talent")
	if err != nil {
		res := response.TalentLoginResponseDTO{StatusCode: 500, Message: "Failed to generate token"}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	res := response.TalentLoginResponseDTO{
		StatusCode: 200,
		Message:    "Login Success",
		Token:      token,
		Data: response.TalentLoginData{
			TalentID: talent.TalentID,
			Email:    talent.Email,
		},
	}
	c.JSON(http.StatusOK, res)
}

func RegisterTalent(c *gin.Context) {
	var req request.TalentRegisterRequestDTO
	var talent model.Talent

	if err := c.ShouldBindJSON(&req); err != nil {
		res := response.BaseResponseDTO{StatusCode: 400, Message: "Invalid Request"}
		c.JSON(http.StatusBadRequest, res)
		return
	}
	req.Email = strings.ToLower(req.Email)
	if err := database.GlobalDB.Where("Email = ?", req.Email).First(&talent).Error; err == nil {
		res := response.BaseResponseDTO{StatusCode: 409, Message: "Account already exists"}
		c.JSON(http.StatusConflict, res)
		return
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		res := response.BaseResponseDTO{StatusCode: 500, Message: "Database error"}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	hashedPassword, err := HashPassword(req.Password)
	if err != nil {
		res := response.BaseResponseDTO{StatusCode: 500, Message: "Failed to hash password"}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	talent = model.Talent{
		TalentID:     uuid.New(),
		Email:        req.Email,
		Password:     hashedPassword,
		FullName:     req.FullName,
		ActiveStatus: true,
		HireCount:    0,
		PhoneNumber:  req.PhoneNumber,
	}

	if err := database.GlobalDB.Create(&talent).Error; err != nil {
		res := response.BaseResponseDTO{StatusCode: 500, Message: "Failed to create account"}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	res := response.BaseResponseDTO{
		StatusCode: 200,
		Message:    "Register Success",
	}
	c.JSON(http.StatusCreated, res)
}

func GetAllTalents(c *gin.Context) {
	role, err := GetUserRole(c)
	if err != nil || role != "SME" {
		c.JSON(http.StatusForbidden, response.BaseResponseDTO{
			Message:    "You are not authorized to post a job",
			StatusCode: http.StatusForbidden,
		})
		return
	}
	var talents []model.Talent
	if err := database.GlobalDB.Find(&talents).Error; err != nil {
		c.JSON(http.StatusInternalServerError, response.BaseResponseDTO{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to retrieve Talents",
		})
		return
	}
	if len(talents) == 0 {
		c.JSON(http.StatusNotFound, response.BaseResponseDTO{
			StatusCode: http.StatusNotFound,
			Message:    "No Talent found",
		})
		return
	}
	c.JSON(http.StatusOK, response.GetAllTalentResponseDTO{
		Talents: talents,
	})
}

func GetAllTalentByAppID(c *gin.Context) {
	role, err := GetUserRole(c)
	if err != nil || strings.TrimSpace(strings.ToLower(role)) != "sme" {
		c.JSON(http.StatusForbidden, response.BaseResponseDTO{
			Message:    "Unauthorized",
			StatusCode: http.StatusForbidden,
		})
		return
	}

	var requestBody request.GetAllTalentByAppIDRequestDTO
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		res := response.BaseResponseDTO{StatusCode: http.StatusBadRequest, Message: "Invalid Request"}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	appID := requestBody.AppID
	fmt.Println("Received AppID:", appID)

	var application model.TrApplication
	if err := database.GlobalDB.Where("app_id = ?", appID).First(&application).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, response.BaseResponseDTO{
				StatusCode: http.StatusNotFound,
				Message:    "Application not found",
			})
		} else {
			c.JSON(http.StatusInternalServerError, response.BaseResponseDTO{
				StatusCode: http.StatusInternalServerError,
				Message:    "Failed to check application: " + err.Error(),
			})
		}
		return
	}

	var talents []model.Talent
	if err := database.GlobalDB.Where("talent_id IN (?)", application.TalentID).Find(&talents).Error; err != nil {
		c.JSON(http.StatusInternalServerError, response.BaseResponseDTO{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to retrieve Talents: " + err.Error(),
		})
		return
	}

	if len(talents) == 0 {
		c.JSON(http.StatusNotFound, response.BaseResponseDTO{
			StatusCode: http.StatusNotFound,
			Message:    "No Talent found",
		})
		return
	}

	c.JSON(http.StatusOK, response.GetAllTalentResponseDTO{
		Talents: talents,
		BaseResponse: response.BaseResponseDTO{
			StatusCode: http.StatusOK,
			Message:    "Success",
		},
	})
}

func GetTalentDetail(c *gin.Context) {
	db := database.GlobalDB
	talentID := c.Query("talent_id")
	if talentID == "" {
		c.JSON(http.StatusBadRequest, response.BaseResponseDTO{
			Message:    "Talent_ID is required",
			StatusCode: http.StatusBadRequest,
		})
		return
	}
	var talentDetail model.Talent

	if err := db.Where("talent_id", talentID).Find(&talentDetail).Error; err != nil {
		c.JSON(http.StatusInternalServerError, response.BaseResponseDTO{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to get Talent Detail",
		})
		return
	}
	c.JSON(http.StatusOK, response.TalentDetailResponseDTO{
		Data: talentDetail,
		BaseResponse: response.BaseResponseDTO{
			StatusCode: http.StatusOK,
			Message:    "Success",
		},
	})
}

func EditTalentDetail(c *gin.Context) {
	db := database.GlobalDB
	talentID := c.Query("talent_id")
	if talentID == "" {
		c.JSON(http.StatusBadRequest, response.BaseResponseDTO{
			Message:    "Talent_ID is required",
			StatusCode: http.StatusBadRequest,
		})
		return
	}

	var talent model.Talent
	if err := db.Where("talent_id", talentID).Find(&talent).Error; err != nil {
		c.JSON(http.StatusInternalServerError, response.BaseResponseDTO{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to get Talent Detail",
		})
		return
	}

	var req request.EditTalentRequestDTO

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.BaseResponseDTO{
			Message:    "Invalid Request",
			StatusCode: http.StatusBadRequest,
		})
		return
	}

	if req.FullName != "" {
		talent.FullName = req.FullName
	}
	if req.PhoneNumber != "" {
		talent.PhoneNumber = req.PhoneNumber
	}
	if req.Email != "" {
		talent.Email = req.Email
	}
	if req.Bio != "" {
		talent.Bio = req.Bio
	}
	if req.Location != "" {
		talent.Location = req.Location
	}
	if req.CV != "" {
		talent.CV = req.CV
	}

	if err := db.Save(&talent).Error; err != nil {
		c.JSON(http.StatusInternalServerError, response.BaseResponseDTO{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to update Talent Detail",
		})
		return
	}

	c.JSON(http.StatusOK, response.EditTalentDetailResponseDTO{
		Data: talent,
		BaseResponse: response.BaseResponseDTO{
			StatusCode: http.StatusOK,
			Message:    "Success",
		},
	})
}

func GetAllApplicantsByJobID(c *gin.Context) {
	role, err := GetUserRole(c)
	if err != nil || strings.TrimSpace(strings.ToLower(role)) != "sme" {
		c.JSON(http.StatusForbidden, response.BaseResponseDTO{
			Message:    "Unauthorized",
			StatusCode: http.StatusForbidden,
		})
		return
	}

	var requestBody request.GetAllTalentByAppAndJobIDRequestDTO
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		res := response.BaseResponseDTO{StatusCode: http.StatusBadRequest, Message: "Invalid Request"}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	var application model.TrApplication
	if err := database.GlobalDB.Where("app_id = ? AND job_id = ?", requestBody.AppID, requestBody.JobID).First(&application).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, response.BaseResponseDTO{
				StatusCode: http.StatusNotFound,
				Message:    "Application not found",
			})
		} else {
			c.JSON(http.StatusInternalServerError, response.BaseResponseDTO{
				StatusCode: http.StatusInternalServerError,
				Message:    "Failed to check application: " + err.Error(),
			})
		}
		return
	}

	var talents []model.Talent
	if err := database.GlobalDB.Where("talent_id IN (?)", application.TalentID).Find(&talents).Error; err != nil {
		c.JSON(http.StatusInternalServerError, response.BaseResponseDTO{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to retrieve Talents: " + err.Error(),
		})
		return
	}

	if len(talents) == 0 {
		c.JSON(http.StatusNotFound, response.BaseResponseDTO{
			StatusCode: http.StatusNotFound,
			Message:    "No Talent found",
		})
		return
	}

	c.JSON(http.StatusOK, response.GetAllTalentResponseDTO{
		Talents: talents,
		BaseResponse: response.BaseResponseDTO{
			StatusCode: http.StatusOK,
			Message:    "Success",
		},
	})
}

func GetTalentSkills(c *gin.Context) {
	var req request.GetTalentSkillsRequestDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		res := response.BaseResponseDTO{StatusCode: http.StatusBadRequest, Message: "Invalid Request"}
		c.JSON(http.StatusBadRequest, res)
		return
	}
	var skills []model.Skill
	if err := database.GlobalDB.Where("talent_id = ?", req.TalentID).Find(&skills).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve skills"})
		return
	}

	c.JSON(http.StatusOK, response.GetTalentSkillsResponseDTO{
		Skills: skills,
		BaseResponse: response.BaseResponseDTO{
			StatusCode: http.StatusOK,
			Message:    "Success",
		},
	})
}

func AddTalentSkills(c *gin.Context) {
	var req request.AddTalentSkillsRequestDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		res := response.BaseResponseDTO{StatusCode: http.StatusBadRequest, Message: "Invalid Request"}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	for _, skillID := range req.Skills {
		var skill model.Skill
		if err := database.GlobalDB.First(&skill, "skill_id = ?", skillID).Error; err != nil {
			c.JSON(http.StatusNotFound, response.BaseResponseDTO{
				StatusCode: http.StatusNotFound,
				Message:    "Skill not found",
			})
			return
		}

		if err := database.GlobalDB.Model(&model.Talent{TalentID: req.TalentID}).Association("Skills").Append(&skill); err != nil {
			c.JSON(http.StatusInternalServerError, response.BaseResponseDTO{
				StatusCode: http.StatusInternalServerError,
				Message:    "Failed to associate skill with talent",
			})
			return
		}
	}

	c.JSON(http.StatusCreated, response.BaseResponseDTO{
		StatusCode: http.StatusCreated,
		Message:    "Success",
	})
}

func SaveJob(c *gin.Context) {
	db := database.GlobalDB
	var req request.SaveJobRequestDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		res := response.BaseResponseDTO{StatusCode: http.StatusBadRequest, Message: "Invalid Request"}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	var talent model.Talent
	if err := db.Where("talent_id = ?", req.TalentID).First(&talent).Error; err != nil {
		c.JSON(http.StatusNotFound, response.BaseResponseDTO{
			StatusCode: http.StatusNotFound,
			Message:    "Talent not found",
		})
		return
	}

	var job model.Job
	if err := db.Where("job_id = ?", req.JobID).First(&job).Error; err != nil {
		c.JSON(http.StatusNotFound, response.BaseResponseDTO{
			StatusCode: http.StatusNotFound,
			Message:    "Job not found",
		})
		return
	}

	savedJob := model.SavedJobs{
		TalentID: req.TalentID,
		JobID:    req.JobID,
	}

	if err := db.Create(&savedJob).Error; err != nil {
		c.JSON(http.StatusInternalServerError, response.BaseResponseDTO{
			StatusCode: http.StatusInternalServerError,
			Message:    "Error in save job",
		})
		return
	}

	c.JSON(http.StatusCreated, response.BaseResponseDTO{
		StatusCode: http.StatusCreated,
		Message:    "Success",
	})
}
