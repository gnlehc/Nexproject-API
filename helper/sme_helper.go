package helper

import (
	"errors"
	"help/database"
	"help/model"
	"help/model/request"
	"help/model/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func LoginSME(c *gin.Context) {
	var req request.SMELoginRequestDTO
	var sme model.SME

	if err := c.ShouldBindJSON(&req); err != nil {
		res := response.SMELoginResponseDTO{StatusCode: 400, Message: "Invalid Request"}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	if err := database.GlobalDB.Where("Email = ?", req.Email).First(&sme).Error; err != nil {
		res := response.SMELoginResponseDTO{StatusCode: 401, Message: "Credentials not matched"}
		c.JSON(http.StatusUnauthorized, res)
		return
	}

	if !CheckPasswordHash(req.Password, sme.Password) {
		res := response.SMELoginResponseDTO{StatusCode: 401, Message: "Invalid Password"}
		c.JSON(http.StatusUnauthorized, res)
		return
	}

	token, err := GenerateJWT(sme.SMEID, "SME")
	if err != nil {
		res := response.SMELoginResponseDTO{StatusCode: 500, Message: "Failed to generate token"}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	res := response.SMELoginResponseDTO{
		StatusCode: 200,
		Message:    "Login Success",
		Token:      token,
		Data: response.SMELoginData{
			SMEID: sme.SMEID,
			Email: sme.Email,
		},
	}
	c.JSON(http.StatusOK, res)
}

func RegisterSME(c *gin.Context) {
	var req request.SMERegisterRequestDTO
	var sme model.SME

	if err := c.ShouldBindJSON(&req); err != nil {
		res := response.BaseResponseDTO{StatusCode: 400, Message: "Invalid Request"}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	if err := database.GlobalDB.Where("Email = ?", req.Email).First(&sme).Error; err == nil {
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

	sme = model.SME{
		SMEID:        uuid.New(),
		Email:        req.Email,
		Password:     hashedPassword,
		CompanyName:  req.CompanyName,
		ActiveStatus: true,
		CEO:          req.CEO,
		SMEType:      req.SMEType,
		PhoneNumber:  req.PhoneNumber,
	}

	if err := database.GlobalDB.Create(&sme).Error; err != nil {
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
