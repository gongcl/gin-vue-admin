import service from '@/utils/request'

// @Tags HrEmployee
// @Summary 创建hrEmployee表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HrEmployee true "创建hrEmployee表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /hrEmployee/createHrEmployee [post]
export const createHrEmployee = (data) => {
  return service({
    url: '/hrEmployee/createHrEmployee',
    method: 'post',
    data
  })
}

// @Tags HrEmployee
// @Summary 删除hrEmployee表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HrEmployee true "删除hrEmployee表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hrEmployee/deleteHrEmployee [delete]
export const deleteHrEmployee = (params) => {
  return service({
    url: '/hrEmployee/deleteHrEmployee',
    method: 'delete',
    params
  })
}

// @Tags HrEmployee
// @Summary 批量删除hrEmployee表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除hrEmployee表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hrEmployee/deleteHrEmployee [delete]
export const deleteHrEmployeeByIds = (params) => {
  return service({
    url: '/hrEmployee/deleteHrEmployeeByIds',
    method: 'delete',
    params
  })
}

// @Tags HrEmployee
// @Summary 更新hrEmployee表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HrEmployee true "更新hrEmployee表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hrEmployee/updateHrEmployee [put]
export const updateHrEmployee = (data) => {
  return service({
    url: '/hrEmployee/updateHrEmployee',
    method: 'put',
    data
  })
}

// @Tags HrEmployee
// @Summary 用id查询hrEmployee表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.HrEmployee true "用id查询hrEmployee表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hrEmployee/findHrEmployee [get]
export const findHrEmployee = (params) => {
  return service({
    url: '/hrEmployee/findHrEmployee',
    method: 'get',
    params
  })
}

// @Tags HrEmployee
// @Summary 分页获取hrEmployee表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取hrEmployee表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hrEmployee/getHrEmployeeList [get]
export const getHrEmployeeList = (params) => {
  return service({
    url: '/hrEmployee/getHrEmployeeList',
    method: 'get',
    params
  })
}
