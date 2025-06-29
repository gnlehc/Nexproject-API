package service

import (
	"nexproject/helper"

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

}

func UserRoutes(private *gin.RouterGroup) {
	private.GET("/get-all-talent", helper.GetAllTalents)
	private.POST("/get-all-talent-by-appid", helper.GetAllTalentByAppID)
	private.GET("/get-sme-detail", helper.GetSMEDetail)
	private.GET("/get-talent-detail", helper.GetTalentDetail)
	private.GET("/get-user-role", helper.DetermineUserRoleByEmail)
	private.POST("/get-all-applicants-on-a-job", helper.GetAllApplicantsByJobID)
	private.POST("/get-talent-skills", helper.GetTalentSkills)
	private.POST("/add-talent-skills", helper.AddTalentSkills)
	private.POST("/save-job", helper.SaveJob)

	private.POST("/edit-talent", helper.EditTalentDetail)
	private.POST("/edit-sme", helper.EditSMEDetails)
	private.POST("/get-all-my-applied-jobs", helper.GetAllAppliedJobsByTalentID)
}
