import service from '@/utils/request'

// @Tags MsXmList
// @Summary 创建msXmList表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MsXmList true "创建msXmList表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /msXmList/createMsXmList [post]
export const createMsXmList = (data) => {
  return service({
    url: '/msXmList/createMsXmList',
    method: 'post',
    data
  })
}

// @Tags MsXmList
// @Summary 删除msXmList表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MsXmList true "删除msXmList表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /msXmList/deleteMsXmList [delete]
export const deleteMsXmList = (params) => {
  return service({
    url: '/msXmList/deleteMsXmList',
    method: 'delete',
    params
  })
}

// @Tags MsXmList
// @Summary 批量删除msXmList表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除msXmList表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /msXmList/deleteMsXmList [delete]
export const deleteMsXmListByIds = (params) => {
  return service({
    url: '/msXmList/deleteMsXmListByIds',
    method: 'delete',
    params
  })
}

// @Tags MsXmList
// @Summary 更新msXmList表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MsXmList true "更新msXmList表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /msXmList/updateMsXmList [put]
export const updateMsXmList = (data) => {
  return service({
    url: '/msXmList/updateMsXmList',
    method: 'put',
    data
  })
}

// @Tags MsXmList
// @Summary 用id查询msXmList表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.MsXmList true "用id查询msXmList表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /msXmList/findMsXmList [get]
export const findMsXmList = (params) => {
  return service({
    url: '/msXmList/findMsXmList',
    method: 'get',
    params
  })
}

// @Tags MsXmList
// @Summary 分页获取msXmList表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取msXmList表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /msXmList/getMsXmListList [get]
export const getMsXmListList = (params) => {
  return service({
    url: '/msXmList/getMsXmListList',
    method: 'get',
    params
  })
}
