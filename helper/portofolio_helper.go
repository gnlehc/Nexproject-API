package helper

import (
	"loom/database"
	"loom/model"
	"loom/model/request"
	"loom/model/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func AddPortofolio(c *gin.Context) {
	db := database.GlobalDB
	var req request.PortofolioRequestDTO
	claims, _ := c.Get("claims")
	JWTClaims := claims.(*JWTClaims)
	talentID := JWTClaims.UserID

	var talent model.Talent
	if err := db.First(&talent, "talent_id = ?", talentID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Talent not found"})
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.BaseResponseDTO{
			StatusCode: 400,
			Message:    "All fields are required",
		})
		return
	}

	newPortofolio := model.Portofolio{
		PortofolioID: uuid.New(),
		Title:        req.Title,
		Description:  req.Description,
		ProjectLink:  req.ProjectLink,
		CoverImage:   req.CoverImage,
		TalentID:     JWTClaims.UserID,
	}

	if err := db.Create(&newPortofolio).Error; err != nil {
		c.JSON(http.StatusInternalServerError, response.BaseResponseDTO{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to add portfolio",
		})
		return
	}

	talent.Portofolio = append(talent.Portofolio, newPortofolio)

	if err := db.Save(&talent).Error; err != nil {
		c.JSON(http.StatusInternalServerError, response.BaseResponseDTO{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to update talent with portfolio",
		})
		return
	}
	c.JSON(http.StatusOK, response.BaseResponseDTO{
		StatusCode: http.StatusOK,
		Message:    "Portfolio Added Successfully",
	})
}
