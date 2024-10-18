package service

import (
	"loom/helper"

	"github.com/gin-gonic/gin"
)

func AuthUserRoutes(public *gin.RouterGroup) {
	public.POST("/talent-login",
		helper.LoginTalent,
	)
	public.POST("/talent-register", helper.RegisterTalent)
	public.POST("/sme-login", helper.LoginSME)
	public.POST("/sme-register", helper.RegisterSME)
	public.POST("/check-user-role", helper.DetermineUserRoleByEmail)
	public.GET("/check-user-role-by-jwt", helper.DetermineUserRoleByJWT)
	public.GET("/get-talent-profile", helper.GetTalentProfile)
}

func UserRoutes(private *gin.RouterGroup) {
	private.GET("/get-all-talent", helper.GetAllTalents)
	private.POST("/get-all-talent-by-appid", helper.GetAllTalentByAppID)
	private.GET("/get-sme-detail", helper.GetSMEDetail)
	private.GET("/get-talent-detail", helper.GetTalentDetail)
}
