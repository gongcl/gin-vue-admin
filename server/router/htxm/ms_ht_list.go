package htxm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type MsHtListRouter struct {
}

// InitMsHtListRouter 初始化 msHtList表 路由信息
func (s *MsHtListRouter) InitMsHtListRouter(Router *gin.RouterGroup) {
	msHtListRouter := Router.Group("msHtList").Use(middleware.OperationRecord())
	msHtListRouterWithoutRecord := Router.Group("msHtList")
	var msHtListApi = v1.ApiGroupApp.HtxmApiGroup.MsHtListApi
	{
		msHtListRouter.POST("createMsHtList", msHtListApi.CreateMsHtList)   // 新建msHtList表
		msHtListRouter.DELETE("deleteMsHtList", msHtListApi.DeleteMsHtList) // 删除msHtList表
		msHtListRouter.DELETE("deleteMsHtListByIds", msHtListApi.DeleteMsHtListByIds) // 批量删除msHtList表
		msHtListRouter.PUT("updateMsHtList", msHtListApi.UpdateMsHtList)    // 更新msHtList表
	}
	{
		msHtListRouterWithoutRecord.GET("findMsHtList", msHtListApi.FindMsHtList)        // 根据ID获取msHtList表
		msHtListRouterWithoutRecord.GET("getMsHtListList", msHtListApi.GetMsHtListList)  // 获取msHtList表列表
	}
}
