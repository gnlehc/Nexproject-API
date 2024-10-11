package helper

import (
	"errors"
	"help/database"
	"help/model"
	"help/model/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UploadCV(c *gin.Context, cvLink string) error {
	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, response.BaseResponseDTO{
			Message:    "Unauthorized",
			StatusCode: http.StatusUnauthorized,
		})
		return errors.New("unauthorized")
	}

	jwtClaims, ok := claims.(*JWTClaims)
	if !ok {
		c.JSON(http.StatusUnauthorized, response.BaseResponseDTO{
			Message:    "Invalid claims",
			StatusCode: http.StatusUnauthorized,
		})
		return errors.New("invalid claims")
	}

	db := database.GlobalDB

	var updateError error
	switch jwtClaims.Role {
	case "Talent":
		// Update the CV field in the Talent struct
		updateError = db.Model(&model.Talent{}).Where("talent_id = ?", jwtClaims.UserID).Update("cv", cvLink).Error
	default:
		c.JSON(http.StatusForbidden, response.BaseResponseDTO{
			Message:    "Invalid user role",
			StatusCode: http.StatusForbidden,
		})
		return errors.New("invalid user role")
	}

	if updateError != nil {
		c.JSON(http.StatusInternalServerError, response.BaseResponseDTO{
			Message:    "Unable to upload CV",
			StatusCode: http.StatusInternalServerError,
		})
		return updateError
	}

	c.JSON(http.StatusOK, response.BaseResponseDTO{
		StatusCode: http.StatusOK,
		Message:    "CV Uploaded Successfully",
	})
	return nil
}
