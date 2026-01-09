import service from '@/utils/request'
// @Tags UserExtra
// @Summary 创建user extra info
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.UserExtra true "创建user extra info"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /userExtra/createUserExtra [post]
export const createUserExtra = (data) => {
  return service({
    url: '/userExtra/createUserExtra',
    method: 'post',
    data
  })
}

// @Tags UserExtra
// @Summary 删除user extra info
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.UserExtra true "删除user extra info"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /userExtra/deleteUserExtra [delete]
export const deleteUserExtra = (params) => {
  return service({
    url: '/userExtra/deleteUserExtra',
    method: 'delete',
    params
  })
}

// @Tags UserExtra
// @Summary 批量删除user extra info
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除user extra info"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /userExtra/deleteUserExtra [delete]
export const deleteUserExtraByIds = (params) => {
  return service({
    url: '/userExtra/deleteUserExtraByIds',
    method: 'delete',
    params
  })
}

// @Tags UserExtra
// @Summary 更新user extra info
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.UserExtra true "更新user extra info"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /userExtra/updateUserExtra [put]
export const updateUserExtra = (data) => {
  return service({
    url: '/userExtra/updateUserExtra',
    method: 'put',
    data
  })
}

// @Tags UserExtra
// @Summary 用id查询user extra info
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.UserExtra true "用id查询user extra info"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /userExtra/findUserExtra [get]
export const findUserExtra = (params) => {
  return service({
    url: '/userExtra/findUserExtra',
    method: 'get',
    params
  })
}

// @Tags UserExtra
// @Summary 分页获取user extra info列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取user extra info列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /userExtra/getUserExtraList [get]
export const getUserExtraList = (params) => {
  return service({
    url: '/userExtra/getUserExtraList',
    method: 'get',
    params
  })
}

// @Tags UserExtra
// @Summary 不需要鉴权的user extra info接口
// @Accept application/json
// @Produce application/json
// @Param data query user_extraReq.UserExtraSearch true "分页获取user extra info列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /userExtra/getUserExtraPublic [get]
export const getUserExtraPublic = () => {
  return service({
    url: '/userExtra/getUserExtraPublic',
    method: 'get',
  })
}
