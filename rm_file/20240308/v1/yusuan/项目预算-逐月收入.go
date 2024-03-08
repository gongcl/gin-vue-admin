package yusuan

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/yusuan"
    yusuanReq "github.com/flipped-aurora/gin-vue-admin/server/model/yusuan/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type 项目预算-逐月收入Api struct {
}

var 项目预算Service = service.ServiceGroupApp.YusuanServiceGroup.项目预算-逐月收入Service


// Create项目预算-逐月收入 创建项目预算-逐月收入表
// @Tags 项目预算-逐月收入
// @Summary 创建项目预算-逐月收入表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body yusuan.项目预算-逐月收入 true "创建项目预算-逐月收入表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /项目预算/create项目预算-逐月收入 [post]
func (项目预算Api *项目预算-逐月收入Api) Create项目预算-逐月收入(c *gin.Context) {
	var 项目预算 yusuan.项目预算-逐月收入
	err := c.ShouldBindJSON(&项目预算)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := 项目预算Service.Create项目预算-逐月收入(&项目预算); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// Delete项目预算-逐月收入 删除项目预算-逐月收入表
// @Tags 项目预算-逐月收入
// @Summary 删除项目预算-逐月收入表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body yusuan.项目预算-逐月收入 true "删除项目预算-逐月收入表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /项目预算/delete项目预算-逐月收入 [delete]
func (项目预算Api *项目预算-逐月收入Api) Delete项目预算-逐月收入(c *gin.Context) {
	ID := c.Query("ID")
	if err := 项目预算Service.Delete项目预算-逐月收入(ID); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// Delete项目预算-逐月收入ByIds 批量删除项目预算-逐月收入表
// @Tags 项目预算-逐月收入
// @Summary 批量删除项目预算-逐月收入表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /项目预算/delete项目预算-逐月收入ByIds [delete]
func (项目预算Api *项目预算-逐月收入Api) Delete项目预算-逐月收入ByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := 项目预算Service.Delete项目预算-逐月收入ByIds(IDs); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// Update项目预算-逐月收入 更新项目预算-逐月收入表
// @Tags 项目预算-逐月收入
// @Summary 更新项目预算-逐月收入表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body yusuan.项目预算-逐月收入 true "更新项目预算-逐月收入表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /项目预算/update项目预算-逐月收入 [put]
func (项目预算Api *项目预算-逐月收入Api) Update项目预算-逐月收入(c *gin.Context) {
	var 项目预算 yusuan.项目预算-逐月收入
	err := c.ShouldBindJSON(&项目预算)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := 项目预算Service.Update项目预算-逐月收入(项目预算); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// Find项目预算-逐月收入 用id查询项目预算-逐月收入表
// @Tags 项目预算-逐月收入
// @Summary 用id查询项目预算-逐月收入表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query yusuan.项目预算-逐月收入 true "用id查询项目预算-逐月收入表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /项目预算/find项目预算-逐月收入 [get]
func (项目预算Api *项目预算-逐月收入Api) Find项目预算-逐月收入(c *gin.Context) {
	ID := c.Query("ID")
	if re项目预算, err := 项目预算Service.Get项目预算-逐月收入(ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"re项目预算": re项目预算}, c)
	}
}

// Get项目预算-逐月收入List 分页获取项目预算-逐月收入表列表
// @Tags 项目预算-逐月收入
// @Summary 分页获取项目预算-逐月收入表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query yusuanReq.项目预算-逐月收入Search true "分页获取项目预算-逐月收入表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /项目预算/get项目预算-逐月收入List [get]
func (项目预算Api *项目预算-逐月收入Api) Get项目预算-逐月收入List(c *gin.Context) {
	var pageInfo yusuanReq.项目预算-逐月收入Search
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := 项目预算Service.Get项目预算-逐月收入InfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}
