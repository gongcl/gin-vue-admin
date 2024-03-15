package htxm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type MsXmListRouter struct {
}

// InitMsXmListRouter 初始化 msXmList表 路由信息
func (s *MsXmListRouter) InitMsXmListRouter(Router *gin.RouterGroup) {
	msXmListRouter := Router.Group("msXmList").Use(middleware.OperationRecord())
	msXmListRouterWithoutRecord := Router.Group("msXmList")
	var msXmListApi = v1.ApiGroupApp.HtxmApiGroup.MsXmListApi
	{
		msXmListRouter.POST("createMsXmList", msXmListApi.CreateMsXmList)   // 新建msXmList表
		msXmListRouter.DELETE("deleteMsXmList", msXmListApi.DeleteMsXmList) // 删除msXmList表
		msXmListRouter.DELETE("deleteMsXmListByIds", msXmListApi.DeleteMsXmListByIds) // 批量删除msXmList表
		msXmListRouter.PUT("updateMsXmList", msXmListApi.UpdateMsXmList)    // 更新msXmList表
	}
	{
		msXmListRouterWithoutRecord.GET("findMsXmList", msXmListApi.FindMsXmList)        // 根据ID获取msXmList表
		msXmListRouterWithoutRecord.GET("getMsXmListList", msXmListApi.GetMsXmListList)  // 获取msXmList表列表
	}
}
