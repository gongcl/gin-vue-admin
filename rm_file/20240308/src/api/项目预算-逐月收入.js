import service from '@/utils/request'

// @Tags 项目预算-逐月收入
// @Summary 创建项目预算-逐月收入表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.项目预算-逐月收入 true "创建项目预算-逐月收入表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /项目预算/create项目预算-逐月收入 [post]
export const create项目预算-逐月收入 = (data) => {
  return service({
    url: '/项目预算/create项目预算-逐月收入',
    method: 'post',
    data
  })
}

// @Tags 项目预算-逐月收入
// @Summary 删除项目预算-逐月收入表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.项目预算-逐月收入 true "删除项目预算-逐月收入表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /项目预算/delete项目预算-逐月收入 [delete]
export const delete项目预算-逐月收入 = (params) => {
  return service({
    url: '/项目预算/delete项目预算-逐月收入',
    method: 'delete',
    params
  })
}

// @Tags 项目预算-逐月收入
// @Summary 批量删除项目预算-逐月收入表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除项目预算-逐月收入表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /项目预算/delete项目预算-逐月收入 [delete]
export const delete项目预算-逐月收入ByIds = (params) => {
  return service({
    url: '/项目预算/delete项目预算-逐月收入ByIds',
    method: 'delete',
    params
  })
}

// @Tags 项目预算-逐月收入
// @Summary 更新项目预算-逐月收入表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.项目预算-逐月收入 true "更新项目预算-逐月收入表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /项目预算/update项目预算-逐月收入 [put]
export const update项目预算-逐月收入 = (data) => {
  return service({
    url: '/项目预算/update项目预算-逐月收入',
    method: 'put',
    data
  })
}

// @Tags 项目预算-逐月收入
// @Summary 用id查询项目预算-逐月收入表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.项目预算-逐月收入 true "用id查询项目预算-逐月收入表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /项目预算/find项目预算-逐月收入 [get]
export const find项目预算-逐月收入 = (params) => {
  return service({
    url: '/项目预算/find项目预算-逐月收入',
    method: 'get',
    params
  })
}

// @Tags 项目预算-逐月收入
// @Summary 分页获取项目预算-逐月收入表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取项目预算-逐月收入表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /项目预算/get项目预算-逐月收入List [get]
export const get项目预算-逐月收入List = (params) => {
  return service({
    url: '/项目预算/get项目预算-逐月收入List',
    method: 'get',
    params
  })
}
