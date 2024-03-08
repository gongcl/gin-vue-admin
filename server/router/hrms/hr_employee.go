package hrms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type HrEmployeeRouter struct {
}

// InitHrEmployeeRouter 初始化 hrEmployee表 路由信息
func (s *HrEmployeeRouter) InitHrEmployeeRouter(Router *gin.RouterGroup) {
	hrEmployeeRouter := Router.Group("hrEmployee").Use(middleware.OperationRecord())
	hrEmployeeRouterWithoutRecord := Router.Group("hrEmployee")
	var hrEmployeeApi = v1.ApiGroupApp.HrmsApiGroup.HrEmployeeApi
	{
		hrEmployeeRouter.POST("createHrEmployee", hrEmployeeApi.CreateHrEmployee)   // 新建hrEmployee表
		hrEmployeeRouter.DELETE("deleteHrEmployee", hrEmployeeApi.DeleteHrEmployee) // 删除hrEmployee表
		hrEmployeeRouter.DELETE("deleteHrEmployeeByIds", hrEmployeeApi.DeleteHrEmployeeByIds) // 批量删除hrEmployee表
		hrEmployeeRouter.PUT("updateHrEmployee", hrEmployeeApi.UpdateHrEmployee)    // 更新hrEmployee表
	}
	{
		hrEmployeeRouterWithoutRecord.GET("findHrEmployee", hrEmployeeApi.FindHrEmployee)        // 根据ID获取hrEmployee表
		hrEmployeeRouterWithoutRecord.GET("getHrEmployeeList", hrEmployeeApi.GetHrEmployeeList)  // 获取hrEmployee表列表
	}
}
