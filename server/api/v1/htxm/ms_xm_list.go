package htxm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/htxm"
	htxmReq "github.com/flipped-aurora/gin-vue-admin/server/model/htxm/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type MsXmListApi struct {
}

var msXmListService = service.ServiceGroupApp.HtxmServiceGroup.MsXmListService

// CreateMsXmList 创建msXmList表
// @Tags MsXmList
// @Summary 创建msXmList表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body htxm.MsXmList true "创建msXmList表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /msXmList/createMsXmList [post]
func (msXmListApi *MsXmListApi) CreateMsXmList(c *gin.Context) {
	var msXmList htxm.MsXmList
	err := c.ShouldBindJSON(&msXmList)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	msXmList.CreatedBy = utils.GetUserID(c)

	if err := msXmListService.CreateMsXmList(&msXmList); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteMsXmList 删除msXmList表
// @Tags MsXmList
// @Summary 删除msXmList表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body htxm.MsXmList true "删除msXmList表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /msXmList/deleteMsXmList [delete]
func (msXmListApi *MsXmListApi) DeleteMsXmList(c *gin.Context) {
	ID := c.Query("ID")
	userID := utils.GetUserID(c)
	if err := msXmListService.DeleteMsXmList(ID, userID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteMsXmListByIds 批量删除msXmList表
// @Tags MsXmList
// @Summary 批量删除msXmList表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /msXmList/deleteMsXmListByIds [delete]
func (msXmListApi *MsXmListApi) DeleteMsXmListByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	userID := utils.GetUserID(c)
	if err := msXmListService.DeleteMsXmListByIds(IDs, userID); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateMsXmList 更新msXmList表
// @Tags MsXmList
// @Summary 更新msXmList表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body htxm.MsXmList true "更新msXmList表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /msXmList/updateMsXmList [put]
func (msXmListApi *MsXmListApi) UpdateMsXmList(c *gin.Context) {
	var msXmList htxm.MsXmList
	err := c.ShouldBindJSON(&msXmList)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	msXmList.UpdatedBy = utils.GetUserID(c)

	if err := msXmListService.UpdateMsXmList(msXmList); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindMsXmList 用id查询msXmList表
// @Tags MsXmList
// @Summary 用id查询msXmList表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query htxm.MsXmList true "用id查询msXmList表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /msXmList/findMsXmList [get]
func (msXmListApi *MsXmListApi) FindMsXmList(c *gin.Context) {
	ID := c.Query("ID")
	if remsXmList, err := msXmListService.GetMsXmList(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"remsXmList": remsXmList}, c)
	}
}

// GetMsXmListList 分页获取msXmList表列表
// @Tags MsXmList
// @Summary 分页获取msXmList表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query htxmReq.MsXmListSearch true "分页获取msXmList表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /msXmList/getMsXmListList [get]
func (msXmListApi *MsXmListApi) GetMsXmListList(c *gin.Context) {
	var pageInfo htxmReq.MsXmListSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := msXmListService.GetMsXmListInfoList(pageInfo); err != nil {
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
