package services

import (
	"osh.com/rps/registrar/config"
	"osh.com/rps/registrar/internal/interfaces"
)

type Service struct {
	Config *config.SystemConfig
	Db     interfaces.DbInterface
	mlog   interfaces.MultyLogInterface
	// validationService validation_services.ValidationService
}

func GetNewRegstrarService(conf *config.SystemConfig, db interfaces.DbInterface, vDb interfaces.ValidationDbInterface, mlog interfaces.MultyLogInterface) Service {
	return Service{
		Config: conf,
		Db:     db,
		mlog:   mlog,
	}
}
