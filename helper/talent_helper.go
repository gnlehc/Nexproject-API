package helper

import (
	"errors"
	"help/database"
	"help/model"
	"help/model/request"
	"help/model/response"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func LoginTalent(c *gin.Context) {
	var req request.TalentLoginRequestDTO
	var talent model.Talent

	if err := c.ShouldBindJSON(&req); err != nil {
		res := response.TalentLoginResponseDTO{StatusCode: 400, Message: "Invalid Request"}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	if err := database.GlobalDB.Where("Email = ?", strings.ToLower(req.Email)).First(&talent).Error; err != nil {
		res := response.TalentLoginResponseDTO{StatusCode: 401, Message: "Credentials not matched"}
		c.JSON(http.StatusUnauthorized, res)
		return
	}

	if !CheckPasswordHash(req.Password, talent.Password) {
		res := response.TalentLoginResponseDTO{StatusCode: 401, Message: "Invalid Password"}
		c.JSON(http.StatusUnauthorized, res)
		return
	}

	token, err := GenerateJWT(talent.TalentID, "Talent")
	if err != nil {
		res := response.TalentLoginResponseDTO{StatusCode: 500, Message: "Failed to generate token"}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	res := response.TalentLoginResponseDTO{
		StatusCode: 200,
		Message:    "Login Success",
		Token:      token,
		Data: response.TalentLoginData{
			TalentID: talent.TalentID,
			Email:    talent.Email,
		},
	}
	c.JSON(http.StatusOK, res)
}

func RegisterTalent(c *gin.Context) {
	var req request.TalentRegisterRequestDTO
	var talent model.Talent

	if err := c.ShouldBindJSON(&req); err != nil {
		res := response.BaseResponseDTO{StatusCode: 400, Message: "Invalid Request"}
		c.JSON(http.StatusBadRequest, res)
		return
	}
	req.Email = strings.ToLower(req.Email)
	if err := database.GlobalDB.Where("Email = ?", req.Email).First(&talent).Error; err == nil {
		res := response.BaseResponseDTO{StatusCode: 409, Message: "Account already exists"}
		c.JSON(http.StatusConflict, res)
		return
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		res := response.BaseResponseDTO{StatusCode: 500, Message: "Database error"}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	hashedPassword, err := HashPassword(req.Password)
	if err != nil {
		res := response.BaseResponseDTO{StatusCode: 500, Message: "Failed to hash password"}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	talent = model.Talent{
		TalentID:     uuid.New(),
		Email:        req.Email,
		Password:     hashedPassword,
		FullName:     req.FullName,
		ActiveStatus: true,
		HireCount:    0,
		PhoneNumber:  req.PhoneNumber,
	}

	if err := database.GlobalDB.Create(&talent).Error; err != nil {
		res := response.BaseResponseDTO{StatusCode: 500, Message: "Failed to create account"}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	res := response.BaseResponseDTO{
		StatusCode: 200,
		Message:    "Register Success",
	}
	c.JSON(http.StatusCreated, res)
}

func EditProfile() {

}
