package helper

import (
	"errors"
	"loom/database"
	"loom/model"
	"loom/model/request"
	"loom/model/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
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

func GetTalentPortofolio(c *gin.Context) {
	db := database.GlobalDB
	claims, _ := c.Get("claims")
	JWTClaims := claims.(*JWTClaims)
	talentID := JWTClaims.UserID

	var talent model.Talent
	if err := db.First(&talent, "talent_id = ?", talentID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Talent not found"})
		return
	}

	var portfolios []model.Portofolio
	if err := db.Where("talent_id = ?", talentID).First(&talent).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, response.BaseResponseDTO{
				StatusCode: http.StatusNotFound,
				Message:    "Talent not found",
			})
		} else {
			c.JSON(http.StatusInternalServerError, response.BaseResponseDTO{
				StatusCode: http.StatusInternalServerError,
				Message:    "Failed to check Talent: " + err.Error(),
			})
		}
		return
	}
	if err := db.Where("talent_id = ?", talentID).Find(&portfolios).Error; err != nil {
		c.JSON(http.StatusInternalServerError, response.BaseResponseDTO{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to get Talent Portfolios",
		})
		return
	}

	c.JSON(http.StatusOK, response.TalentPortfolioResponseDTO{
		Data: portfolios,
		BaseResponse: response.BaseResponseDTO{
			StatusCode: http.StatusOK,
			Message:    "Success",
		},
	})
}

func GetPortfolioByTalentID(c *gin.Context) {
    db := database.GlobalDB
    talentID := c.Query("talent_id")  

    if talentID == "" {
        c.JSON(http.StatusBadRequest, response.BaseResponseDTO{
            StatusCode: http.StatusBadRequest,
            Message:    "Talent ID is required",
        })
        return
    }

    var talent model.Talent
    if err := db.First(&talent, "talent_id = ?", talentID).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            c.JSON(http.StatusNotFound, response.BaseResponseDTO{
                StatusCode: http.StatusNotFound,
                Message:    "Talent not found",
            })
        } else {
            c.JSON(http.StatusInternalServerError, response.BaseResponseDTO{
                StatusCode: http.StatusInternalServerError,
                Message:    "Failed to retrieve talent: " + err.Error(),
            })
        }
        return
    }

    var portfolios []model.Portofolio
    if err := db.Where("talent_id = ?", talentID).Find(&portfolios).Error; err != nil {
        c.JSON(http.StatusInternalServerError, response.BaseResponseDTO{
            StatusCode: http.StatusInternalServerError,
            Message:    "Failed to retrieve portfolios",
        })
        return
    }

    if len(portfolios) == 0 {
        c.JSON(http.StatusNotFound, response.BaseResponseDTO{
            StatusCode: http.StatusNotFound,
            Message:    "No portfolios found for this talent",
        })
        return
    }

    c.JSON(http.StatusOK, response.TalentPortfolioResponseDTO{
        Data: portfolios,
        BaseResponse: response.BaseResponseDTO{
            StatusCode: http.StatusOK,
            Message:    "Success",
        },
    })
}