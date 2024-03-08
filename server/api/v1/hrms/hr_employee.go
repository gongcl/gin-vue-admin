package hrms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/hrms"
    hrmsReq "github.com/flipped-aurora/gin-vue-admin/server/model/hrms/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type HrEmployeeApi struct {
}

var hrEmployeeService = service.ServiceGroupApp.HrmsServiceGroup.HrEmployeeService


// CreateHrEmployee 创建hrEmployee表
// @Tags HrEmployee
// @Summary 创建hrEmployee表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body hrms.HrEmployee true "创建hrEmployee表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /hrEmployee/createHrEmployee [post]
func (hrEmployeeApi *HrEmployeeApi) CreateHrEmployee(c *gin.Context) {
	var hrEmployee hrms.HrEmployee
	err := c.ShouldBindJSON(&hrEmployee)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    hrEmployee.CreatedBy = utils.GetUserID(c)

	if err := hrEmployeeService.CreateHrEmployee(&hrEmployee); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteHrEmployee 删除hrEmployee表
// @Tags HrEmployee
// @Summary 删除hrEmployee表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body hrms.HrEmployee true "删除hrEmployee表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hrEmployee/deleteHrEmployee [delete]
func (hrEmployeeApi *HrEmployeeApi) DeleteHrEmployee(c *gin.Context) {
	ID := c.Query("ID")
    	userID := utils.GetUserID(c)
	if err := hrEmployeeService.DeleteHrEmployee(ID,userID); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteHrEmployeeByIds 批量删除hrEmployee表
// @Tags HrEmployee
// @Summary 批量删除hrEmployee表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /hrEmployee/deleteHrEmployeeByIds [delete]
func (hrEmployeeApi *HrEmployeeApi) DeleteHrEmployeeByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
    userID := utils.GetUserID(c)
	if err := hrEmployeeService.DeleteHrEmployeeByIds(IDs,userID); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateHrEmployee 更新hrEmployee表
// @Tags HrEmployee
// @Summary 更新hrEmployee表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body hrms.HrEmployee true "更新hrEmployee表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hrEmployee/updateHrEmployee [put]
func (hrEmployeeApi *HrEmployeeApi) UpdateHrEmployee(c *gin.Context) {
	var hrEmployee hrms.HrEmployee
	err := c.ShouldBindJSON(&hrEmployee)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    hrEmployee.UpdatedBy = utils.GetUserID(c)

	if err := hrEmployeeService.UpdateHrEmployee(hrEmployee); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindHrEmployee 用id查询hrEmployee表
// @Tags HrEmployee
// @Summary 用id查询hrEmployee表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query hrms.HrEmployee true "用id查询hrEmployee表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hrEmployee/findHrEmployee [get]
func (hrEmployeeApi *HrEmployeeApi) FindHrEmployee(c *gin.Context) {
	ID := c.Query("ID")
	if rehrEmployee, err := hrEmployeeService.GetHrEmployee(ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rehrEmployee": rehrEmployee}, c)
	}
}

// GetHrEmployeeList 分页获取hrEmployee表列表
// @Tags HrEmployee
// @Summary 分页获取hrEmployee表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query hrmsReq.HrEmployeeSearch true "分页获取hrEmployee表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hrEmployee/getHrEmployeeList [get]
func (hrEmployeeApi *HrEmployeeApi) GetHrEmployeeList(c *gin.Context) {
	var pageInfo hrmsReq.HrEmployeeSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := hrEmployeeService.GetHrEmployeeInfoList(pageInfo); err != nil {
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
