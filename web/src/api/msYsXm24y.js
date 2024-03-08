import service from '@/utils/request'

// @Tags MsYsXm24y
// @Summary 创建msYsXm24y表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MsYsXm24y true "创建msYsXm24y表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /msYsXm24y/createMsYsXm24y [post]
export const createMsYsXm24y = (data) => {
  return service({
    url: '/msYsXm24y/createMsYsXm24y',
    method: 'post',
    data
  })
}

// @Tags MsYsXm24y
// @Summary 删除msYsXm24y表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MsYsXm24y true "删除msYsXm24y表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /msYsXm24y/deleteMsYsXm24y [delete]
export const deleteMsYsXm24y = (params) => {
  return service({
    url: '/msYsXm24y/deleteMsYsXm24y',
    method: 'delete',
    params
  })
}

// @Tags MsYsXm24y
// @Summary 批量删除msYsXm24y表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除msYsXm24y表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /msYsXm24y/deleteMsYsXm24y [delete]
export const deleteMsYsXm24yByIds = (params) => {
  return service({
    url: '/msYsXm24y/deleteMsYsXm24yByIds',
    method: 'delete',
    params
  })
}

// @Tags MsYsXm24y
// @Summary 更新msYsXm24y表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MsYsXm24y true "更新msYsXm24y表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /msYsXm24y/updateMsYsXm24y [put]
export const updateMsYsXm24y = (data) => {
  return service({
    url: '/msYsXm24y/updateMsYsXm24y',
    method: 'put',
    data
  })
}

// @Tags MsYsXm24y
// @Summary 用id查询msYsXm24y表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.MsYsXm24y true "用id查询msYsXm24y表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /msYsXm24y/findMsYsXm24y [get]
export const findMsYsXm24y = (params) => {
  return service({
    url: '/msYsXm24y/findMsYsXm24y',
    method: 'get',
    params
  })
}

// @Tags MsYsXm24y
// @Summary 分页获取msYsXm24y表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取msYsXm24y表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /msYsXm24y/getMsYsXm24yList [get]
export const getMsYsXm24yList = (params) => {
  return service({
    url: '/msYsXm24y/getMsYsXm24yList',
    method: 'get',
    params
  })
}
