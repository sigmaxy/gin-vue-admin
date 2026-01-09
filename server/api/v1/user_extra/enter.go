package user_extra

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct{ UserExtraApi }

var userExtraService = service.ServiceGroupApp.User_extraServiceGroup.UserExtraService
