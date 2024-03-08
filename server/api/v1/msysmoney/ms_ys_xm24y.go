package msysmoney

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/msysmoney"
    msysmoneyReq "github.com/flipped-aurora/gin-vue-admin/server/model/msysmoney/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type MsYsXm24yApi struct {
}

var msYsXm24yService = service.ServiceGroupApp.MsysmoneyServiceGroup.MsYsXm24yService


// CreateMsYsXm24y 创建msYsXm24y表
// @Tags MsYsXm24y
// @Summary 创建msYsXm24y表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body msysmoney.MsYsXm24y true "创建msYsXm24y表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /msYsXm24y/createMsYsXm24y [post]
func (msYsXm24yApi *MsYsXm24yApi) CreateMsYsXm24y(c *gin.Context) {
	var msYsXm24y msysmoney.MsYsXm24y
	err := c.ShouldBindJSON(&msYsXm24y)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    msYsXm24y.CreatedBy = utils.GetUserID(c)

	if err := msYsXm24yService.CreateMsYsXm24y(&msYsXm24y); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteMsYsXm24y 删除msYsXm24y表
// @Tags MsYsXm24y
// @Summary 删除msYsXm24y表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body msysmoney.MsYsXm24y true "删除msYsXm24y表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /msYsXm24y/deleteMsYsXm24y [delete]
func (msYsXm24yApi *MsYsXm24yApi) DeleteMsYsXm24y(c *gin.Context) {
	ID := c.Query("ID")
    	userID := utils.GetUserID(c)
	if err := msYsXm24yService.DeleteMsYsXm24y(ID,userID); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteMsYsXm24yByIds 批量删除msYsXm24y表
// @Tags MsYsXm24y
// @Summary 批量删除msYsXm24y表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /msYsXm24y/deleteMsYsXm24yByIds [delete]
func (msYsXm24yApi *MsYsXm24yApi) DeleteMsYsXm24yByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
    userID := utils.GetUserID(c)
	if err := msYsXm24yService.DeleteMsYsXm24yByIds(IDs,userID); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateMsYsXm24y 更新msYsXm24y表
// @Tags MsYsXm24y
// @Summary 更新msYsXm24y表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body msysmoney.MsYsXm24y true "更新msYsXm24y表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /msYsXm24y/updateMsYsXm24y [put]
func (msYsXm24yApi *MsYsXm24yApi) UpdateMsYsXm24y(c *gin.Context) {
	var msYsXm24y msysmoney.MsYsXm24y
	err := c.ShouldBindJSON(&msYsXm24y)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    msYsXm24y.UpdatedBy = utils.GetUserID(c)

	if err := msYsXm24yService.UpdateMsYsXm24y(msYsXm24y); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindMsYsXm24y 用id查询msYsXm24y表
// @Tags MsYsXm24y
// @Summary 用id查询msYsXm24y表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query msysmoney.MsYsXm24y true "用id查询msYsXm24y表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /msYsXm24y/findMsYsXm24y [get]
func (msYsXm24yApi *MsYsXm24yApi) FindMsYsXm24y(c *gin.Context) {
	ID := c.Query("ID")
	if remsYsXm24y, err := msYsXm24yService.GetMsYsXm24y(ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"remsYsXm24y": remsYsXm24y}, c)
	}
}

// GetMsYsXm24yList 分页获取msYsXm24y表列表
// @Tags MsYsXm24y
// @Summary 分页获取msYsXm24y表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query msysmoneyReq.MsYsXm24ySearch true "分页获取msYsXm24y表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /msYsXm24y/getMsYsXm24yList [get]
func (msYsXm24yApi *MsYsXm24yApi) GetMsYsXm24yList(c *gin.Context) {
	var pageInfo msysmoneyReq.MsYsXm24ySearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := msYsXm24yService.GetMsYsXm24yInfoList(pageInfo); err != nil {
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
