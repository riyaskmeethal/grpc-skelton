package pkg

import (
	"osh.com/rps/registrar/config"
	"osh.com/rps/registrar/internal/interfaces"
	"osh.com/rps/registrar/logger/zerolog"
)

type MultiLogger struct {
	loggers map[string]interfaces.LogInterface
	conf    *config.LoggerConfig
}

func (ml MultiLogger) GetLogger() interfaces.LogInterface {

	if logger, ok := ml.loggers["default"]; ok {
		return logger
	}
	return zerolog.GetNewLogger(ml.conf, "")

}

func (ml MultiLogger) GetInstituteLogger(institute string) interfaces.LogInterface {

	if logger, ok := ml.loggers[institute]; ok {
		return logger
	} else if logger, ok := ml.loggers["default"]; ok {
		return logger
	}
	return zerolog.GetNewLogger(ml.conf, "")

}

func GetMultilogger(conf interfaces.ConfigInterface) (ml interfaces.MultyLogInterface) {

	return loadMultilogger(conf)

}

func loadMultilogger(conf interfaces.ConfigInterface) (ml MultiLogger) {

	loggerConf := new(config.LoggerConfig)
	loggerConf.LoadConfig(conf)

	ml.conf = loggerConf
	ml.loggers = make(map[string]interfaces.LogInterface)

	//for logging based on institues
	for _, inst := range loggerConf.Logger.Institutes {

		ml.loggers[inst] = zerolog.GetNewLogger(loggerConf, inst)
		ml.loggers[inst].Info(inst, ": logger loaded.")
	}
	ml.loggers["default"] = zerolog.GetNewLogger(loggerConf, "")
	return ml
}
