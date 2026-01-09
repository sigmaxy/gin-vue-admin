package user_extra

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type UserExtraRouter struct {}

// InitUserExtraRouter 初始化 user extra info 路由信息
func (s *UserExtraRouter) InitUserExtraRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	userExtraRouter := Router.Group("userExtra").Use(middleware.OperationRecord())
	userExtraRouterWithoutRecord := Router.Group("userExtra")
	userExtraRouterWithoutAuth := PublicRouter.Group("userExtra")
	{
		userExtraRouter.POST("createUserExtra", userExtraApi.CreateUserExtra)   // 新建user extra info
		userExtraRouter.DELETE("deleteUserExtra", userExtraApi.DeleteUserExtra) // 删除user extra info
		userExtraRouter.DELETE("deleteUserExtraByIds", userExtraApi.DeleteUserExtraByIds) // 批量删除user extra info
		userExtraRouter.PUT("updateUserExtra", userExtraApi.UpdateUserExtra)    // 更新user extra info
	}
	{
		userExtraRouterWithoutRecord.GET("findUserExtra", userExtraApi.FindUserExtra)        // 根据ID获取user extra info
		userExtraRouterWithoutRecord.GET("getUserExtraList", userExtraApi.GetUserExtraList)  // 获取user extra info列表
	}
	{
	    userExtraRouterWithoutAuth.GET("getUserExtraPublic", userExtraApi.GetUserExtraPublic)  // user extra info开放接口
	}
}
