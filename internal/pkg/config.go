package pkg

import (
	"osh.com/rps/registrar/config"
	"osh.com/rps/registrar/internal/interfaces"
)

func GetConfig(confFile string, initConf bool) (conf interfaces.ConfigInterface, err error) {

	return config.InitConfig(confFile, initConf)
}
