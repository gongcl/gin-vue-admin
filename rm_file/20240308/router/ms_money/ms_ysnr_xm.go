package ms_money

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type MsYsnrXmRouter struct {
}

// InitMsYsnrXmRouter 初始化 msYsnrXm表 路由信息
func (s *MsYsnrXmRouter) InitMsYsnrXmRouter(Router *gin.RouterGroup) {
	msYsnrXmRouter := Router.Group("msYsnrXm").Use(middleware.OperationRecord())
	msYsnrXmRouterWithoutRecord := Router.Group("msYsnrXm")
	var msYsnrXmApi = v1.ApiGroupApp.Ms_moneyApiGroup.MsYsnrXmApi
	{
		msYsnrXmRouter.POST("createMsYsnrXm", msYsnrXmApi.CreateMsYsnrXm)   // 新建msYsnrXm表
		msYsnrXmRouter.DELETE("deleteMsYsnrXm", msYsnrXmApi.DeleteMsYsnrXm) // 删除msYsnrXm表
		msYsnrXmRouter.DELETE("deleteMsYsnrXmByIds", msYsnrXmApi.DeleteMsYsnrXmByIds) // 批量删除msYsnrXm表
		msYsnrXmRouter.PUT("updateMsYsnrXm", msYsnrXmApi.UpdateMsYsnrXm)    // 更新msYsnrXm表
	}
	{
		msYsnrXmRouterWithoutRecord.GET("findMsYsnrXm", msYsnrXmApi.FindMsYsnrXm)        // 根据ID获取msYsnrXm表
		msYsnrXmRouterWithoutRecord.GET("getMsYsnrXmList", msYsnrXmApi.GetMsYsnrXmList)  // 获取msYsnrXm表列表
	}
}
