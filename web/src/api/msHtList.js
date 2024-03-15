import service from '@/utils/request'

// @Tags MsHtList
// @Summary 创建msHtList表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MsHtList true "创建msHtList表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /msHtList/createMsHtList [post]
export const createMsHtList = (data) => {
  return service({
    url: '/msHtList/createMsHtList',
    method: 'post',
    data
  })
}

// @Tags MsHtList
// @Summary 删除msHtList表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MsHtList true "删除msHtList表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /msHtList/deleteMsHtList [delete]
export const deleteMsHtList = (params) => {
  return service({
    url: '/msHtList/deleteMsHtList',
    method: 'delete',
    params
  })
}

// @Tags MsHtList
// @Summary 批量删除msHtList表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除msHtList表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /msHtList/deleteMsHtList [delete]
export const deleteMsHtListByIds = (params) => {
  return service({
    url: '/msHtList/deleteMsHtListByIds',
    method: 'delete',
    params
  })
}

// @Tags MsHtList
// @Summary 更新msHtList表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MsHtList true "更新msHtList表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /msHtList/updateMsHtList [put]
export const updateMsHtList = (data) => {
  return service({
    url: '/msHtList/updateMsHtList',
    method: 'put',
    data
  })
}

// @Tags MsHtList
// @Summary 用id查询msHtList表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.MsHtList true "用id查询msHtList表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /msHtList/findMsHtList [get]
export const findMsHtList = (params) => {
  return service({
    url: '/msHtList/findMsHtList',
    method: 'get',
    params
  })
}

// @Tags MsHtList
// @Summary 分页获取msHtList表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取msHtList表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /msHtList/getMsHtListList [get]
export const getMsHtListList = (params) => {
  return service({
    url: '/msHtList/getMsHtListList',
    method: 'get',
    params
  })
}
