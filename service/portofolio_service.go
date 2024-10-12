package service

import (
	"help/helper"

	"github.com/gin-gonic/gin"
)

func PortofolioRoute(private *gin.RouterGroup) {
	private.POST("/add-portofolio", helper.AddPortofolio)
}
