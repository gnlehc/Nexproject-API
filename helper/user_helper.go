package helper

import (
	"net/http"
	"nexproject/database"
	"nexproject/model"
	"nexproject/model/request"
	"nexproject/model/response"
	"strings"

	"github.com/gin-gonic/gin"
)

func DetermineUserRoleByEmail(c *gin.Context) {
	var req request.UserRoleRequestDTO
	var talent model.Talent
	var sme model.SME

	if err := c.ShouldBindJSON(&req); err != nil {
		res := response.BaseResponseDTO{StatusCode: 400, Message: "Invalid Request"}
		c.JSON(http.StatusBadRequest, res)
		return
	}
	email := strings.ToLower(req.Email)

	if err := database.GlobalDB.Where("LOWER(Email) = ?", strings.ToLower(email)).First(&talent).Error; err == nil {
		res := response.UserRoleResponseDTO{
			StatusCode: 200,
			Message:    "User found",
			Role:       "Talent",
		}
		c.JSON(http.StatusOK, res)
		return
	}

	if err := database.GlobalDB.Where("LOWER(Email) = ?", strings.ToLower(email)).First(&sme).Error; err == nil {
		res := response.UserRoleResponseDTO{
			StatusCode: 200,
			Message:    "User found",
			Role:       "SME",
		}
		c.JSON(http.StatusOK, res)
		return
	}

	res := response.UserRoleResponseDTO{
		StatusCode: 404,
		Message:    "User not found",
		Role:       "Unknown",
	}
	c.JSON(http.StatusNotFound, res)
}
