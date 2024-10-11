package service

import (
	"help/helper"
	"help/model/response"

	"github.com/gin-gonic/gin"
)

func CVRoute(private *gin.RouterGroup) {
	private.POST("/upload-cv", func(c *gin.Context) {
		var requestBody struct {
			CVLink string `json:"cvLink" binding:"required"`
		}

		if err := c.ShouldBindJSON(&requestBody); err != nil {
			c.JSON(400, response.BaseResponseDTO{
				StatusCode: 400,
				Message:    "cvLink is required",
			})
			return
		}

		if err := helper.UploadCV(c, requestBody.CVLink); err != nil {
			c.JSON(500, response.BaseResponseDTO{
				StatusCode: 500,
				Message:    "Failed to upload CV",
			})
			return
		}
	})
}
