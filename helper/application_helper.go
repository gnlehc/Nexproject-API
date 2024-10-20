package helper

import (
	"loom/database"
	"loom/model"
	"loom/model/request"
	"loom/model/response"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetAllAppIDByJobID(c *gin.Context) {
	db := database.GlobalDB
	var req request.GetAllAppIDByJobIDRequestDTO

	role, err := GetUserRole(c)
	if err != nil || strings.ToLower(role) != "sme" {
		c.JSON(http.StatusForbidden, response.BaseResponseDTO{
			Message:    "Unauthorized",
			StatusCode: http.StatusForbidden,
		})
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.BaseResponseDTO{
			Message:    "Invalid request body",
			StatusCode: http.StatusBadRequest,
		})
		return
	}

	var applications []model.TrApplication
	if err := db.Where("job_id = ?", req.JobID).Find(&applications).Error; err != nil {
		c.JSON(http.StatusNotFound, response.BaseResponseDTO{
			Message:    "Applications not found",
			StatusCode: http.StatusNotFound,
		})
		return
	}

	var appIDs []uuid.UUID
	for _, app := range applications {
		appIDs = append(appIDs, app.AppID)
	}

	c.JSON(http.StatusOK, response.GetAllApplicationByJobIDResponseDTO{
		ListAppID: appIDs,
		BaseResponse: response.BaseResponseDTO{
			Message:    "Success",
			StatusCode: http.StatusOK,
		},
	})
}

func UpdateApplicationStatus(c *gin.Context) {
	db := database.GlobalDB
	var req request.UpdateApplicationStatusRequestDTO
	role, err := GetUserRole(c)
	if err != nil || strings.ToLower(role) != "sme" {
		c.JSON(http.StatusForbidden, response.BaseResponseDTO{
			Message:    "Unauthorized to update application status",
			StatusCode: http.StatusForbidden,
		})
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.BaseResponseDTO{
			Message:    "Invalid request body",
			StatusCode: http.StatusBadRequest,
		})
		return
	}

	var application model.TrApplication
	if err := db.Where("app_id = ? AND job_id = ? AND talent_id = ?", req.AppID, req.JobID, req.TalentID).First(&application).Error; err != nil {
		c.JSON(http.StatusNotFound, response.BaseResponseDTO{
			Message:    "Application not found",
			StatusCode: http.StatusNotFound,
		})
		return
	}

	application.StatusID = req.StatusID
	if err := db.Save(&application).Error; err != nil {
		c.JSON(http.StatusInternalServerError, response.BaseResponseDTO{
			Message:    "Failed to update application status",
			StatusCode: http.StatusInternalServerError,
		})
		return
	}

	c.JSON(http.StatusOK, response.BaseResponseDTO{
		Message:    "Application status updated successfully",
		StatusCode: http.StatusOK,
	})
}
