package user_extra

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/user_extra"
	user_extraReq "github.com/flipped-aurora/gin-vue-admin/server/model/user_extra/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserExtraApi struct{}

// CreateUserExtra 创建user extra info
// @Tags UserExtra
// @Summary 创建user extra info
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body user_extra.UserExtra true "创建user extra info"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /userExtra/createUserExtra [post]
func (userExtraApi *UserExtraApi) CreateUserExtra(c *gin.Context) {
	// 创建业务用Context
	zap.L().Info("api v1 user_extra.go CreateUserExtra")
	ctx := c.Request.Context()

	var userExtra user_extra.UserExtra
	err := c.ShouldBindJSON(&userExtra)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	zap.L().Info("userExtra content", zap.Any("userExtra", userExtra))
	// Inside your handler function:
	claims, _ := utils.GetClaims(c)
	userID := claims.BaseClaims.ID
	zap.L().Info("current user id", zap.Any("userID", userID))

	err = userExtraService.CreateUserExtra(ctx, &userExtra)
	if err != nil {
		global.GVA_LOG.Error(global.Translate("general.creationFail"), zap.Error(err))
		response.FailWithMessage(global.Translate("general.creationFailErr"), c)
		return
	}
	response.OkWithMessage(global.Translate("general.createSuccess"), c)
}

// DeleteUserExtra 删除user extra info
// @Tags UserExtra
// @Summary 删除user extra info
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body user_extra.UserExtra true "删除user extra info"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /userExtra/deleteUserExtra [delete]
func (userExtraApi *UserExtraApi) DeleteUserExtra(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	id := c.Query("id")
	err := userExtraService.DeleteUserExtra(ctx, id)
	if err != nil {
		global.GVA_LOG.Error(global.Translate("general.deleteFail"), zap.Error(err))
		response.FailWithMessage(global.Translate("general.deleteFailErr"), c)
		return
	}
	response.OkWithMessage(global.Translate("general.deleteSuccess"), c)
}

// DeleteUserExtraByIds 批量删除user extra info
// @Tags UserExtra
// @Summary 批量删除user extra info
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /userExtra/deleteUserExtraByIds [delete]
func (userExtraApi *UserExtraApi) DeleteUserExtraByIds(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ids := c.QueryArray("ids[]")
	err := userExtraService.DeleteUserExtraByIds(ctx, ids)
	if err != nil {
		global.GVA_LOG.Error(global.Translate("system.sys_operation_record.batchDeleteFail"), zap.Error(err))
		response.FailWithMessage(global.Translate("system.sys_operation_record.batchDeleteFailErr"), c)
		return
	}
	response.OkWithMessage(global.Translate("system.sys_operation_record.batchDeleteSuccess"), c)
}

// UpdateUserExtra 更新user extra info
// @Tags UserExtra
// @Summary 更新user extra info
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body user_extra.UserExtra true "更新user extra info"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /userExtra/updateUserExtra [put]
func (userExtraApi *UserExtraApi) UpdateUserExtra(c *gin.Context) {
	// 从ctx获取标准context进行业务行为
	ctx := c.Request.Context()

	var userExtra user_extra.UserExtra
	err := c.ShouldBindJSON(&userExtra)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = userExtraService.UpdateUserExtra(ctx, userExtra)
	if err != nil {
		global.GVA_LOG.Error(global.Translate("general.updateFail"), zap.Error(err))
		response.FailWithMessage(global.Translate("general.updateFailErr"), c)
		return
	}
	response.OkWithMessage(global.Translate("general.updateSuccess"), c)
}

// FindUserExtra 用id查询user extra info
// @Tags UserExtra
// @Summary 用id查询user extra info
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param id query int true "用id查询user extra info"
// @Success 200 {object} response.Response{data=user_extra.UserExtra,msg=string} "查询成功"
// @Router /userExtra/findUserExtra [get]
func (userExtraApi *UserExtraApi) FindUserExtra(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	id := c.Query("id")
	reuserExtra, err := userExtraService.GetUserExtra(ctx, id)
	if err != nil {
		global.GVA_LOG.Error(global.Translate("general.queryFail"), zap.Error(err))
		response.FailWithMessage(global.Translate("general.queryFailErr"), c)
		return
	}
	response.OkWithData(reuserExtra, c)
}

// GetUserExtraList 分页获取user extra info列表
// @Tags UserExtra
// @Summary 分页获取user extra info列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query user_extraReq.UserExtraSearch true "分页获取user extra info列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /userExtra/getUserExtraList [get]
func (userExtraApi *UserExtraApi) GetUserExtraList(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var pageInfo user_extraReq.UserExtraSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := userExtraService.GetUserExtraInfoList(ctx, pageInfo)
	if err != nil {
		global.GVA_LOG.Error(global.Translate("general.getDataFail"), zap.Error(err))
		response.FailWithMessage(global.Translate("general.getDataFailErr"), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, global.Translate("general.getDataSuccess"), c)
}

// GetUserExtraPublic 不需要鉴权的user extra info接口
// @Tags UserExtra
// @Summary 不需要鉴权的user extra info接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /userExtra/getUserExtraPublic [get]
func (userExtraApi *UserExtraApi) GetUserExtraPublic(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	userExtraService.GetUserExtraPublic(ctx)
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的user extra info接口信息",
	}, "获取成功", c)
}
