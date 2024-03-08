package yusuan

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type 项目预算-逐月收入Router struct {
}

// Init项目预算-逐月收入Router 初始化 项目预算-逐月收入表 路由信息
func (s *项目预算-逐月收入Router) Init项目预算-逐月收入Router(Router *gin.RouterGroup) {
	项目预算Router := Router.Group("项目预算").Use(middleware.OperationRecord())
	项目预算RouterWithoutRecord := Router.Group("项目预算")
	var 项目预算Api = v1.ApiGroupApp.YusuanApiGroup.项目预算-逐月收入Api
	{
		项目预算Router.POST("create项目预算-逐月收入", 项目预算Api.Create项目预算-逐月收入)   // 新建项目预算-逐月收入表
		项目预算Router.DELETE("delete项目预算-逐月收入", 项目预算Api.Delete项目预算-逐月收入) // 删除项目预算-逐月收入表
		项目预算Router.DELETE("delete项目预算-逐月收入ByIds", 项目预算Api.Delete项目预算-逐月收入ByIds) // 批量删除项目预算-逐月收入表
		项目预算Router.PUT("update项目预算-逐月收入", 项目预算Api.Update项目预算-逐月收入)    // 更新项目预算-逐月收入表
	}
	{
		项目预算RouterWithoutRecord.GET("find项目预算-逐月收入", 项目预算Api.Find项目预算-逐月收入)        // 根据ID获取项目预算-逐月收入表
		项目预算RouterWithoutRecord.GET("get项目预算-逐月收入List", 项目预算Api.Get项目预算-逐月收入List)  // 获取项目预算-逐月收入表列表
	}
}
