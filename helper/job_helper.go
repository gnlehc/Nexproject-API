package helper

import (
	"loom/database"
	"loom/model"
	"loom/model/request"
	"loom/model/response"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	})
}

func PostJob(c *gin.Context) {
	role, err := GetUserRole(c)
	if err != nil || role != "SME" {
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
