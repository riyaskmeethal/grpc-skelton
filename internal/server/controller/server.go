package controller

import (
	"osh.com/rps/registrar/config"
	"osh.com/rps/registrar/internal/interfaces"
	"osh.com/rps/registrar/internal/server/services"
)

type RegstrarServer struct {
	// someappPb.UnimplementedTerminalRegistrationServer
	config  *config.SystemConfig
	db      interfaces.DbInterface
	mlog    interfaces.MultyLogInterface
	service services.Service
}

func GetNewServer(conf interfaces.ConfigInterface, db interfaces.DbInterface, vDb interfaces.ValidationDbInterface, mlog interfaces.MultyLogInterface) RegstrarServer {

	log := mlog.GetLogger()
	config := new(config.SystemConfig)
	config.LoadSystemConfig(conf, log)

	return RegstrarServer{
		config:  config,
		db:      db,
		mlog:    mlog,
		service: services.GetNewRegstrarService(config, db, vDb, mlog),
	}
}
