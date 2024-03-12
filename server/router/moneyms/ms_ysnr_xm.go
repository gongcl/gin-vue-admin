package moneyms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type MsYsnrXmRouter struct {
}

// InitMsYsnrXmRouter 初始化 项目预算表 路由信息
func (s *MsYsnrXmRouter) InitMsYsnrXmRouter(Router *gin.RouterGroup) {
	msYsnrXmRouter := Router.Group("msYsnrXm").Use(middleware.OperationRecord())
	msYsnrXmRouterWithoutRecord := Router.Group("msYsnrXm")
	var msYsnrXmApi = v1.ApiGroupApp.MoneymsApiGroup.MsYsnrXmApi
	{
		msYsnrXmRouter.POST("createMsYsnrXm", msYsnrXmApi.CreateMsYsnrXm)   // 新建项目预算表
		msYsnrXmRouter.DELETE("deleteMsYsnrXm", msYsnrXmApi.DeleteMsYsnrXm) // 删除项目预算表
		msYsnrXmRouter.DELETE("deleteMsYsnrXmByIds", msYsnrXmApi.DeleteMsYsnrXmByIds) // 批量删除项目预算表
		msYsnrXmRouter.PUT("updateMsYsnrXm", msYsnrXmApi.UpdateMsYsnrXm)    // 更新项目预算表
	}
	{
		msYsnrXmRouterWithoutRecord.GET("findMsYsnrXm", msYsnrXmApi.FindMsYsnrXm)        // 根据ID获取项目预算表
		msYsnrXmRouterWithoutRecord.GET("getMsYsnrXmList", msYsnrXmApi.GetMsYsnrXmList)  // 获取项目预算表列表
	}
}
