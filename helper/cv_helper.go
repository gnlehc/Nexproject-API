package helper

import (
	"loom/database"
	"loom/model"
	"loom/model/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UploadCV(c *gin.Context) {
	var cvLink string
	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, response.BaseResponseDTO{
			Message:    "Unauthorized",
			StatusCode: http.StatusUnauthorized,
		})
		return
	}

	jwtClaims, ok := claims.(*JWTClaims)
	if !ok {
		c.JSON(http.StatusUnauthorized, response.BaseResponseDTO{
			Message:    "Invalid claims",
			StatusCode: http.StatusUnauthorized,
		})
		return
	}

	db := database.GlobalDB

	var updateError error
	switch jwtClaims.Role {
	case "Talent":
		updateError = db.Model(&model.Talent{}).Where("talent_id = ?", jwtClaims.UserID).Update("cv", cvLink).Error
	default:
		c.JSON(http.StatusForbidden, response.BaseResponseDTO{
			Message:    "Invalid user role",
			StatusCode: http.StatusForbidden,
		})
		return
	}

	if err := c.ShouldBindJSON(cvLink); err != nil {
		c.JSON(400, response.BaseResponseDTO{
			StatusCode: 400,
			Message:    "CV Link is required",
		})
		return
	}
	if updateError != nil {
		c.JSON(http.StatusInternalServerError, response.BaseResponseDTO{
			Message:    "Unable to upload CV",
			StatusCode: http.StatusInternalServerError,
		})
		return
	}

	c.JSON(http.StatusOK, response.BaseResponseDTO{
		StatusCode: http.StatusOK,
		Message:    "CV Uploaded Successfully",
	})
}
