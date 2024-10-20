package helper

import (
	"errors"
	"loom/database"
	"loom/model"
	"loom/model/request"
	"loom/model/response"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetAllJobs(c *gin.Context) {
	var jobs []model.Job

	if err := database.GlobalDB.Find(&jobs).Error; err != nil {
		c.JSON(http.StatusInternalServerError, response.BaseResponseDTO{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to retrieve Jobs",
		})
		return
	}

	if len(jobs) == 0 {
		c.JSON(http.StatusNotFound, response.BaseResponseDTO{
			StatusCode: http.StatusNotFound,
			Message:    "No jobs found",
		})
		return
	}

	c.JSON(http.StatusOK, response.GetAllJobResponseDTO{
		Jobs: jobs,
		BaseResponse: response.BaseResponseDTO{
			StatusCode: http.StatusOK,
			Message:    "Success",
		},
	})
}

func PostJob(c *gin.Context) {
	role, err := GetUserRole(c)
	if err != nil || strings.TrimSpace(strings.ToLower(role)) != "sme" {
		c.JSON(http.StatusForbidden, response.BaseResponseDTO{
			Message:    "You are not authorized to post a job",
			StatusCode: http.StatusForbidden,
		})
		return
	}
	var jobRequest request.JobRequestDTO

	if err := c.ShouldBindJSON(&jobRequest); err != nil {
		c.JSON(http.StatusBadRequest, response.BaseResponseDTO{
			Message:    err.Error(),
			StatusCode: http.StatusBadRequest,
		})
		return
	}
	// Check if SME exists
	var sme model.SME
	if err := database.GlobalDB.First(&sme, jobRequest.SMEID).Error; err != nil {
		c.JSON(http.StatusBadRequest, response.BaseResponseDTO{
			Message:    "Invalid SME ID: " + err.Error(),
			StatusCode: http.StatusBadRequest,
		})
		return
	}

	job := model.Job{
		JobID:          uuid.New(),
		SMEID:          jobRequest.SMEID,
		JobTitle:       jobRequest.JobTitle,
		JobDescription: jobRequest.JobDescription,
		JobType:        jobRequest.JobType,
		Qualification:  jobRequest.Qualification,
		JobArrangement: jobRequest.JobArrangement,
		Wage:           jobRequest.Wage,
		Active:         jobRequest.Active,
		CreatedAt:      time.Now(),
		Location:       jobRequest.Location,
	}

	db := database.GlobalDB

	if err := db.Create(&job).Error; err != nil {
		c.JSON(http.StatusInternalServerError, response.BaseResponseDTO{
			Message:    "Unable to create job: " + err.Error(),
			StatusCode: http.StatusInternalServerError,
		})
		return
	}

	job.Skills = append(job.Skills, jobRequest.Skills...)

	if err := db.Save(&job).Error; err != nil {
		c.JSON(http.StatusInternalServerError, response.BaseResponseDTO{
			Message:    "Unable to associate skills with job: " + err.Error(),
			StatusCode: http.StatusInternalServerError,
		})
		return
	}

	c.JSON(http.StatusCreated, response.BaseResponseDTO{
		StatusCode: http.StatusCreated,
		Message:    "Job created successfully",
	})
}

func ApplyJob(c *gin.Context) {
	db := database.GlobalDB
	role, err := GetUserRole(c)
	if err != nil || role != "Talent" {
		c.JSON(http.StatusForbidden, response.BaseResponseDTO{
			Message:    "Unauthorized to apply for a job",
			StatusCode: http.StatusForbidden,
		})
		return
	}
	var req request.ApplyJobRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.BaseResponseDTO{
			Message:    err.Error(),
			StatusCode: http.StatusBadRequest,
		})
		return
	}

	var existingApplication model.TrApplication
	if err := db.Where("job_id = ? AND talent_id = ?", req.JobID, req.TalentID).First(&existingApplication).Error; err == nil {
		c.JSON(http.StatusBadRequest, response.BaseResponseDTO{
			StatusCode: http.StatusBadRequest,
			Message:    "You have already applied for this job",
		})
		return
	}

	application := model.TrApplication{
		AppID:    uuid.New(),
		TalentID: req.TalentID,
		JobID:    req.JobID,
		StatusID: 1,
	}

	if err := db.Create(&application).Error; err != nil {
		c.JSON(http.StatusInternalServerError, response.BaseResponseDTO{
			Message:    "Unable to apply job: " + err.Error(),
			StatusCode: http.StatusInternalServerError,
		})
		return
	}
	c.JSON(http.StatusCreated, response.BaseResponseDTO{
		StatusCode: http.StatusCreated,
		Message:    "Job applied successfully",
	})
}

func GetAllJobsPostedBySME(c *gin.Context) {
	smeID := c.Query("sme_id")
	if smeID == "" {
		c.JSON(http.StatusBadRequest, response.BaseResponseDTO{
			Message:    "SMEID is required",
			StatusCode: http.StatusBadRequest,
		})
		return
	}

	var jobs []model.Job
	if err := database.GlobalDB.Where("sme_id = ?", smeID).Find(&jobs).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, response.BaseResponseDTO{
				StatusCode: http.StatusNotFound,
				Message:    "There are no jobs right now",
			})
		} else {
			c.JSON(http.StatusInternalServerError, response.BaseResponseDTO{
				StatusCode: http.StatusInternalServerError,
				Message:    "Failed to retrieve jobs: " + err.Error(),
			})
		}
		return
	}

	c.JSON(http.StatusOK, response.GetAllJobResponseDTO{
		Jobs: jobs,
		BaseResponse: response.BaseResponseDTO{
			StatusCode: http.StatusOK,
			Message:    "Success",
		},
	})
}

