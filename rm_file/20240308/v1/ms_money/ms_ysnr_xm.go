package ms_money

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/ms_money"
    ms_moneyReq "github.com/flipped-aurora/gin-vue-admin/server/model/ms_money/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type MsYsnrXmApi struct {
}

var msYsnrXmService = service.ServiceGroupApp.Ms_moneyServiceGroup.MsYsnrXmService


// CreateMsYsnrXm 创建msYsnrXm表
// @Tags MsYsnrXm
// @Summary 创建msYsnrXm表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body ms_money.MsYsnrXm true "创建msYsnrXm表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /msYsnrXm/createMsYsnrXm [post]
func (msYsnrXmApi *MsYsnrXmApi) CreateMsYsnrXm(c *gin.Context) {
	var msYsnrXm ms_money.MsYsnrXm
	err := c.ShouldBindJSON(&msYsnrXm)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := msYsnrXmService.CreateMsYsnrXm(&msYsnrXm); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteMsYsnrXm 删除msYsnrXm表
// @Tags MsYsnrXm
// @Summary 删除msYsnrXm表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body ms_money.MsYsnrXm true "删除msYsnrXm表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /msYsnrXm/deleteMsYsnrXm [delete]
func (msYsnrXmApi *MsYsnrXmApi) DeleteMsYsnrXm(c *gin.Context) {
	projectName := c.Query("projectName")
	if err := msYsnrXmService.DeleteMsYsnrXm(projectName); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteMsYsnrXmByIds 批量删除msYsnrXm表
// @Tags MsYsnrXm
// @Summary 批量删除msYsnrXm表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /msYsnrXm/deleteMsYsnrXmByIds [delete]
func (msYsnrXmApi *MsYsnrXmApi) DeleteMsYsnrXmByIds(c *gin.Context) {
	projectNames := c.QueryArray("projectNames[]")
	if err := msYsnrXmService.DeleteMsYsnrXmByIds(projectNames); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateMsYsnrXm 更新msYsnrXm表
// @Tags MsYsnrXm
// @Summary 更新msYsnrXm表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body ms_money.MsYsnrXm true "更新msYsnrXm表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /msYsnrXm/updateMsYsnrXm [put]
func (msYsnrXmApi *MsYsnrXmApi) UpdateMsYsnrXm(c *gin.Context) {
	var msYsnrXm ms_money.MsYsnrXm
	err := c.ShouldBindJSON(&msYsnrXm)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := msYsnrXmService.UpdateMsYsnrXm(msYsnrXm); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindMsYsnrXm 用id查询msYsnrXm表
// @Tags MsYsnrXm
// @Summary 用id查询msYsnrXm表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query ms_money.MsYsnrXm true "用id查询msYsnrXm表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /msYsnrXm/findMsYsnrXm [get]
func (msYsnrXmApi *MsYsnrXmApi) FindMsYsnrXm(c *gin.Context) {
	projectName := c.Query("projectName")
	if remsYsnrXm, err := msYsnrXmService.GetMsYsnrXm(projectName); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"remsYsnrXm": remsYsnrXm}, c)
	}
}

// GetMsYsnrXmList 分页获取msYsnrXm表列表
// @Tags MsYsnrXm
// @Summary 分页获取msYsnrXm表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query ms_moneyReq.MsYsnrXmSearch true "分页获取msYsnrXm表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /msYsnrXm/getMsYsnrXmList [get]
func (msYsnrXmApi *MsYsnrXmApi) GetMsYsnrXmList(c *gin.Context) {
	var pageInfo ms_moneyReq.MsYsnrXmSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := msYsnrXmService.GetMsYsnrXmInfoList(pageInfo); err != nil {
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
