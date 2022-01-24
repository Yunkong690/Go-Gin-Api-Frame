package model

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"os"
	model "personFrame/model/common"
	"personFrame/pkg/common"
)

func RegisterTables(db *gorm.DB) {
	err := db.AutoMigrate(
		model.JwtBlackList{},
		User{},
	)
	if err != nil {
		common.GlobalLog.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}
	common.GlobalLog.Info("register table success")

}
