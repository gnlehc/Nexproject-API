package service

import (
	"nexproject/helper"

	"github.com/gin-gonic/gin"
)

func PortofolioRoute(private *gin.RouterGroup) {
	private.POST("/add-portofolio", helper.AddPortofolio)
	private.GET("/get-portofolio", helper.GetTalentPortofolio)
	private.GET("/get-portofolio-by-talent-id", helper.GetPortfolioByTalentID)
}
