package msysmoney

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type MsYsXm24yRouter struct {
}

// InitMsYsXm24yRouter 初始化 msYsXm24y表 路由信息
func (s *MsYsXm24yRouter) InitMsYsXm24yRouter(Router *gin.RouterGroup) {
	msYsXm24yRouter := Router.Group("msYsXm24y").Use(middleware.OperationRecord())
	msYsXm24yRouterWithoutRecord := Router.Group("msYsXm24y")
	var msYsXm24yApi = v1.ApiGroupApp.MsysmoneyApiGroup.MsYsXm24yApi
	{
		msYsXm24yRouter.POST("createMsYsXm24y", msYsXm24yApi.CreateMsYsXm24y)   // 新建msYsXm24y表
		msYsXm24yRouter.DELETE("deleteMsYsXm24y", msYsXm24yApi.DeleteMsYsXm24y) // 删除msYsXm24y表
		msYsXm24yRouter.DELETE("deleteMsYsXm24yByIds", msYsXm24yApi.DeleteMsYsXm24yByIds) // 批量删除msYsXm24y表
		msYsXm24yRouter.PUT("updateMsYsXm24y", msYsXm24yApi.UpdateMsYsXm24y)    // 更新msYsXm24y表
	}
	{
		msYsXm24yRouterWithoutRecord.GET("findMsYsXm24y", msYsXm24yApi.FindMsYsXm24y)        // 根据ID获取msYsXm24y表
		msYsXm24yRouterWithoutRecord.GET("getMsYsXm24yList", msYsXm24yApi.GetMsYsXm24yList)  // 获取msYsXm24y表列表
	}
}
