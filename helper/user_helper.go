package helper

import (
	"loom/database"
	"loom/model"
	"loom/model/request"
	"loom/model/response"
	"net/http"
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

	if err := database.GlobalDB.Where("Email = ?", email).First(&talent).Error; err == nil {
		res := response.UserRoleResponseDTO{
			StatusCode: 200,
			Message:    "User found",
			Role:       "Talent",
		}
		c.JSON(http.StatusOK, res)
		return
	}

	if err := database.GlobalDB.Where("Email = ?", email).First(&sme).Error; err == nil {
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

func DetermineUserRoleByJWT(c *gin.Context) {
	role, err := GetUserRole(c)
	if err != nil {
		c.JSON(401, gin.H{"error": "Unauthorized: Invalid token"})
		return
	}

	profileResponse := gin.H{
		"Role": role,
	}

	c.JSON(200, profileResponse)
}
