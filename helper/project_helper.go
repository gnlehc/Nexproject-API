package helper

import (
	"errors"
	"net/http"
	"nexproject/database"
	"nexproject/model"
	"nexproject/model/request"
	"nexproject/model/response"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetAllProjectBySMEID(c *gin.Context) {
	smeID := c.Query("sme_id")
	if smeID == "" {
		c.JSON(http.StatusBadRequest, response.BaseResponseDTO{
			Message:    "SMEID is required",
			StatusCode: http.StatusBadRequest,
		})
		return
	}
	var projects []model.Project
	if err := database.GlobalDB.Where("sme_id = ?", smeID).Find(&projects).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, response.BaseResponseDTO{
				StatusCode: http.StatusNotFound,
				Message:    "There are no projects right now",
			})
		} else {
			c.JSON(http.StatusInternalServerError, response.BaseResponseDTO{
				StatusCode: http.StatusInternalServerError,
				Message:    "Failed to retrieve projects: " + err.Error(),
			})
		}
		return
	}
	for i := range projects {
		var jobs []model.Job
		if err := database.GlobalDB.Where("project_id = ?", projects[i].ProjectID).Find(&jobs).Error; err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusInternalServerError, response.BaseResponseDTO{
					StatusCode: http.StatusInternalServerError,
					Message:    "Failed to retrieve Jobs: " + err.Error(),
				})
				return
			}
		}
		projects[i].Jobs = jobs
	}
	c.JSON(http.StatusOK, response.GetAllProjectsResponse{
		Projects: projects,
		BaseResponse: response.BaseResponseDTO{
			StatusCode: http.StatusOK,
			Message:    "Success",
		},
	})
}

func GetProjectByID(c *gin.Context) {
	projectId := c.Query("project_id")
	if projectId == "" {
		c.JSON(http.StatusBadRequest, response.BaseResponseDTO{
			Message:    "Project ID is required",
			StatusCode: http.StatusBadRequest,
		})
		return
	}
	var project model.Project
	if err := database.GlobalDB.Where("project_id = ?", projectId).Find(&project).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, response.BaseResponseDTO{
				StatusCode: http.StatusNotFound,
				Message:    "There are no projects right now",
			})
		} else {
			c.JSON(http.StatusInternalServerError, response.BaseResponseDTO{
				StatusCode: http.StatusInternalServerError,
				Message:    "Failed to retrieve projects: " + err.Error(),
			})
		}
		return
	}
	var jobs []model.Job
	if err := database.GlobalDB.Where("project_id = ?", project.ProjectID).Find(&jobs).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusInternalServerError, response.BaseResponseDTO{
				StatusCode: http.StatusInternalServerError,
				Message:    "Failed to retrieve Jobs: " + err.Error(),
			})
			return
		}
	}
	project.Jobs = jobs
	c.JSON(http.StatusOK, response.GetProjectByID{
		Project: project,
		BaseResponse: response.BaseResponseDTO{
			StatusCode: http.StatusOK,
			Message:    "Success",
		},
	})
}

func PostProject(c *gin.Context) {
	db := database.GlobalDB
	role, err := GetUserRole(c)
	if err != nil || role != "SME" {
		c.JSON(http.StatusForbidden, response.BaseResponseDTO{
			Message:    "Unauthorized: Only SME can create project",
			StatusCode: http.StatusForbidden,
		})
		return
	}

	var req request.AddProjectRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.BaseResponseDTO{
			Message:    err.Error(),
			StatusCode: http.StatusBadRequest,
		})
		return
	}

	if len(req.Jobs) == 0 {
		c.JSON(http.StatusBadRequest, response.BaseResponseDTO{
			Message:    "At least one job must be provided",
			StatusCode: http.StatusBadRequest,
		})
		return
	}

	var sme model.SME
	if err := db.First(&sme, "sme_id = ?", req.SMEID).Error; err != nil {
		c.JSON(http.StatusBadRequest, response.BaseResponseDTO{
			Message:    "Invalid SME ID: " + err.Error(),
			StatusCode: http.StatusBadRequest,
		})
		return
	}

	tx := db.Begin()

	projectID := uuid.New()
	project := model.Project{
		ProjectID:          projectID,
		SMEID:              req.SMEID,
		ProjectName:        req.ProjectName,
		ProjectDescription: req.ProjectDescription,
	}

	if err := tx.Create(&project).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, response.BaseResponseDTO{
			Message:    "Failed to create project: " + err.Error(),
			StatusCode: http.StatusInternalServerError,
		})
		return
	}

	for _, jobReq := range req.Jobs {
		job := model.Job{
			JobID:          uuid.New(),
			ProjectID:      projectID,
			JobTitle:       jobReq.JobTitle,
			JobDescription: jobReq.JobDescription,
			JobType:        jobReq.JobType,
			Qualification:  jobReq.Qualification,
			JobArrangement: jobReq.JobArrangement,
			Wage:           jobReq.Wage,
			Active:         jobReq.Active,
			Location:       jobReq.Location,
			CreatedAt:      time.Now(),
		}
		if err := tx.Create(&job).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, response.BaseResponseDTO{
				Message:    "Failed to create job: " + err.Error(),
				StatusCode: http.StatusInternalServerError,
			})
			return
		}
	}

	if err := tx.Commit().Error; err != nil {
		c.JSON(http.StatusInternalServerError, response.BaseResponseDTO{
			Message:    "Failed to commit transaction: " + err.Error(),
			StatusCode: http.StatusInternalServerError,
		})
		return
	}

	c.JSON(http.StatusCreated, response.BaseResponseDTO{
		StatusCode: http.StatusCreated,
		Message:    "Project and associated jobs created successfully",
	})
}
