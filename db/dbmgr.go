package db

import (
	"bugReport/define"

	"github.com/coderguang/GameEngine_go/sgcfg"
	"github.com/coderguang/GameEngine_go/sgdb/sgmysql"
	"github.com/coderguang/GameEngine_go/sglog"
	"github.com/coderguang/GameEngine_go/sgthread"
	"github.com/jinzhu/gorm"
)

var (
	globalDb *gorm.DB
)

func InitDb() {

	cfg, err := sgmysql.ReadCfg(sgcfg.MySQLCfgFile)
	if err != nil {
		sglog.Error("init mysql db error,err:", err)
		sgthread.DelayExit(2)
	}

	globalDb, err = sgmysql.Open(cfg)
	if err != nil {
		sglog.Error("open mysql db error,err:", err)
		sgthread.DelayExit(2)
	}

	sglog.Info("init db connection ok")

	err = globalDb.AutoMigrate(define.SReportData{}).Error
	if err != nil {
		sglog.Error("init report table ok", err)
	}
}
