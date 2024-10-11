package service

import (
	"help/helper"

	"github.com/gin-gonic/gin"
)

func UserRoutes(public *gin.RouterGroup) {
	public.POST("/talent-login",
		helper.LoginTalent,
	)
	public.POST("/talent-register", helper.RegisterTalent)
	public.POST("/sme-login", helper.LoginSME)
	public.POST("/sme-register", helper.RegisterSME)
}
