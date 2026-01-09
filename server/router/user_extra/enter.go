package user_extra

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct{ UserExtraRouter }

var userExtraApi = api.ApiGroupApp.User_extraApiGroup.UserExtraApi
