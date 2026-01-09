package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/user_extra"
)

func bizModel() error {
	db := global.GVA_DB
	err := db.AutoMigrate(user_extra.UserExtra{})
	if err != nil {
		return err
	}
	return nil
}
