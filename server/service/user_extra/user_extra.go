package user_extra

import (
	"context"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/user_extra"
	user_extraReq "github.com/flipped-aurora/gin-vue-admin/server/model/user_extra/request"
	"go.uber.org/zap"
)

type UserExtraService struct{}

// CreateUserExtra 创建user extra info记录
// Author [yourname](https://github.com/yourname)
func (userExtraService *UserExtraService) CreateUserExtra(ctx context.Context, userExtra *user_extra.UserExtra) (err error) {
	zap.L().Info("service user_extra.go 111 CreateUserExtra")
	err = global.GVA_DB.Create(userExtra).Error
	return err
}

// DeleteUserExtra 删除user extra info记录
// Author [yourname](https://github.com/yourname)
func (userExtraService *UserExtraService) DeleteUserExtra(ctx context.Context, id string) (err error) {
	err = global.GVA_DB.Delete(&user_extra.UserExtra{}, "id = ?", id).Error
	return err
}

// DeleteUserExtraByIds 批量删除user extra info记录
// Author [yourname](https://github.com/yourname)
func (userExtraService *UserExtraService) DeleteUserExtraByIds(ctx context.Context, ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]user_extra.UserExtra{}, "id in ?", ids).Error
	return err
}

// UpdateUserExtra 更新user extra info记录
// Author [yourname](https://github.com/yourname)
func (userExtraService *UserExtraService) UpdateUserExtra(ctx context.Context, userExtra user_extra.UserExtra) (err error) {
	err = global.GVA_DB.Model(&user_extra.UserExtra{}).Where("id = ?", userExtra.Id).Updates(&userExtra).Error
	return err
}

// GetUserExtra 根据id获取user extra info记录
// Author [yourname](https://github.com/yourname)
func (userExtraService *UserExtraService) GetUserExtra(ctx context.Context, id string) (userExtra user_extra.UserExtra, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&userExtra).Error
	return
}

// GetUserExtraInfoList 分页获取user extra info记录
// Author [yourname](https://github.com/yourname)
func (userExtraService *UserExtraService) GetUserExtraInfoList(ctx context.Context, info user_extraReq.UserExtraSearch) (list []user_extra.UserExtra, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&user_extra.UserExtra{})
	var userExtras []user_extra.UserExtra
	// 如果有条件搜索 下方会自动创建搜索语句

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&userExtras).Error
	return userExtras, total, err
}
func (userExtraService *UserExtraService) GetUserExtraPublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
