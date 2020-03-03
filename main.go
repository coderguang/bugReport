package main

import (
	"bugReport/config"
	"bugReport/db"
	"bugReport/httpHandle"
	"log"

	"github.com/coderguang/GameEngine_go/sgcfg"
	"github.com/coderguang/GameEngine_go/sgcmd"
	"github.com/coderguang/GameEngine_go/sglog"
	"github.com/coderguang/GameEngine_go/sgserver"
)

func main() {
	sgcfg.SetServerCfgDir("./../globalConfig/bugReport/")
	sgserver.StartServer(sgserver.ServerTypeLog, "debug", "./log/", log.LstdFlags, true)
	sgserver.StartServer(sgserver.ServerTypeMail)

	db.InitDb()
	config.InitUtilCfg()

	go httpHandle.NewWebServer(config.GetUtilCfg().Port)

	sglog.Info("start bug report server")

	sgcmd.StartCmdWaitInputLoop()
}