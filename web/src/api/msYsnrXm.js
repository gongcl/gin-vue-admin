import service from '@/utils/request'

// @Tags MsYsnrXm
// @Summary 创建项目预算表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MsYsnrXm true "创建项目预算表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /msYsnrXm/createMsYsnrXm [post]
export const createMsYsnrXm = (data) => {
  return service({
    url: '/msYsnrXm/createMsYsnrXm',
    method: 'post',
    data
  })
}

// @Tags MsYsnrXm
// @Summary 删除项目预算表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MsYsnrXm true "删除项目预算表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /msYsnrXm/deleteMsYsnrXm [delete]
export const deleteMsYsnrXm = (params) => {
  return service({
    url: '/msYsnrXm/deleteMsYsnrXm',
    method: 'delete',
    params
  })
}

// @Tags MsYsnrXm
// @Summary 批量删除项目预算表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除项目预算表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /msYsnrXm/deleteMsYsnrXm [delete]
export const deleteMsYsnrXmByIds = (params) => {
  return service({
    url: '/msYsnrXm/deleteMsYsnrXmByIds',
    method: 'delete',
    params
  })
}

// @Tags MsYsnrXm
// @Summary 更新项目预算表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MsYsnrXm true "更新项目预算表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /msYsnrXm/updateMsYsnrXm [put]
export const updateMsYsnrXm = (data) => {
  return service({
    url: '/msYsnrXm/updateMsYsnrXm',
    method: 'put',
    data
  })
}

// @Tags MsYsnrXm
// @Summary 用id查询项目预算表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.MsYsnrXm true "用id查询项目预算表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /msYsnrXm/findMsYsnrXm [get]
export const findMsYsnrXm = (params) => {
  return service({
    url: '/msYsnrXm/findMsYsnrXm',
    method: 'get',
    params
  })
}

// @Tags MsYsnrXm
// @Summary 分页获取项目预算表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取项目预算表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /msYsnrXm/getMsYsnrXmList [get]
export const getMsYsnrXmList = (params) => {
  return service({
    url: '/msYsnrXm/getMsYsnrXmList',
    method: 'get',
    params
  })
}
