package htxm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/htxm"
    htxmReq "github.com/flipped-aurora/gin-vue-admin/server/model/htxm/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type MsHtListApi struct {
}

var msHtListService = service.ServiceGroupApp.HtxmServiceGroup.MsHtListService


// CreateMsHtList 创建msHtList表
// @Tags MsHtList
// @Summary 创建msHtList表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body htxm.MsHtList true "创建msHtList表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /msHtList/createMsHtList [post]
func (msHtListApi *MsHtListApi) CreateMsHtList(c *gin.Context) {
	var msHtList htxm.MsHtList
	err := c.ShouldBindJSON(&msHtList)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    msHtList.CreatedBy = utils.GetUserID(c)

	if err := msHtListService.CreateMsHtList(&msHtList); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteMsHtList 删除msHtList表
// @Tags MsHtList
// @Summary 删除msHtList表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body htxm.MsHtList true "删除msHtList表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /msHtList/deleteMsHtList [delete]
func (msHtListApi *MsHtListApi) DeleteMsHtList(c *gin.Context) {
	htmc := c.Query("htmc")
    	userID := utils.GetUserID(c)
	if err := msHtListService.DeleteMsHtList(htmc,userID); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteMsHtListByIds 批量删除msHtList表
// @Tags MsHtList
// @Summary 批量删除msHtList表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /msHtList/deleteMsHtListByIds [delete]
func (msHtListApi *MsHtListApi) DeleteMsHtListByIds(c *gin.Context) {
	htmcs := c.QueryArray("htmcs[]")
    userID := utils.GetUserID(c)
	if err := msHtListService.DeleteMsHtListByIds(htmcs,userID); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateMsHtList 更新msHtList表
// @Tags MsHtList
// @Summary 更新msHtList表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body htxm.MsHtList true "更新msHtList表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /msHtList/updateMsHtList [put]
func (msHtListApi *MsHtListApi) UpdateMsHtList(c *gin.Context) {
	var msHtList htxm.MsHtList
	err := c.ShouldBindJSON(&msHtList)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    msHtList.UpdatedBy = utils.GetUserID(c)

	if err := msHtListService.UpdateMsHtList(msHtList); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindMsHtList 用id查询msHtList表
// @Tags MsHtList
// @Summary 用id查询msHtList表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query htxm.MsHtList true "用id查询msHtList表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /msHtList/findMsHtList [get]
func (msHtListApi *MsHtListApi) FindMsHtList(c *gin.Context) {
	htmc := c.Query("htmc")
	if remsHtList, err := msHtListService.GetMsHtList(htmc); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"remsHtList": remsHtList}, c)
	}
}

// GetMsHtListList 分页获取msHtList表列表
// @Tags MsHtList
// @Summary 分页获取msHtList表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query htxmReq.MsHtListSearch true "分页获取msHtList表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /msHtList/getMsHtListList [get]
func (msHtListApi *MsHtListApi) GetMsHtListList(c *gin.Context) {
	var pageInfo htxmReq.MsHtListSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := msHtListService.GetMsHtListInfoList(pageInfo); err != nil {
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
