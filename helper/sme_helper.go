package helper

import (
	"errors"
	"log"
	"net/http"
	"nexproject/database"
	"nexproject/model"
	"nexproject/model/request"
	"nexproject/model/response"
	"strings"

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
		SMETypeID:    req.SMETypeID,
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

func GetSMEDetail(c *gin.Context) {
	db := database.GlobalDB
	smeID := c.Query("sme_id")
	if smeID == "" {
		c.JSON(http.StatusBadRequest, response.BaseResponseDTO{
			Message:    "SMEID is required",
			StatusCode: http.StatusBadRequest,
		})
		return
	}
	var smeDetail model.SME

	if err := db.Where("sme_id = ?", smeID).Find(&smeDetail).Error; err != nil {
		c.JSON(http.StatusInternalServerError, response.BaseResponseDTO{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to get SME Detail",
		})
		return
	}

	c.JSON(http.StatusOK, response.SMEDetailResponseDTO{
		Data: smeDetail,
		BaseResponse: response.BaseResponseDTO{
			StatusCode: http.StatusOK,
			Message:    "Success",
		},
	})
}

func EditSMEDetails(c *gin.Context) {
	db := database.GlobalDB
	var req request.EditSMERequestDTO

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

	var sme model.SME
	if err := db.First(&sme, req.SMEID).Error; err != nil {
		c.JSON(http.StatusNotFound, response.BaseResponseDTO{
			Message:    "SME not found",
			StatusCode: http.StatusNotFound,
		})
		return
	}
	hashedPassword, err := HashPassword(req.Password)
	if err != nil {
		log.Printf("Error hashing password for %s: %v\n", req.Password, err)
	}
	sme.Email = req.Email
	if req.Password != "" {
		sme.Password = hashedPassword
	}

	sme.CompanyName = req.CompanyName
	sme.CompanyDescription = req.CompanyDescription
	sme.CEO = req.CEO
	sme.Social = req.Social
	sme.PhoneNumber = req.PhoneNumber
	sme.ActiveStatus = req.ActiveStatus
	sme.SMETypeID = req.SMETypeID

	if err := db.Save(&sme).Error; err != nil {
		c.JSON(http.StatusInternalServerError, response.BaseResponseDTO{
			Message:    "Failed to update SME: " + err.Error(),
			StatusCode: http.StatusInternalServerError,
		})
		return
	}

	c.JSON(http.StatusOK, response.BaseResponseDTO{
		Message:    "SME updated successfully",
		StatusCode: http.StatusOK,
	})
}
