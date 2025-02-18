package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
	"osh.com/rps/registrar/internal/interfaces"
)

type LoggerConfig struct {
	Logger Logger `yaml:"logger" validate:"required"`
}

type Logger struct {
	Institutes    []string `yaml:"institutes"`
	LogLevel      string   `yaml:"log.level"`
	LogFile       string   `yaml:"log.file"`
	LogPath       string   `yaml:"log.path"`
	RootDir       string   `yaml:"root.directory"`
	DefaultLogDir string   `yaml:"default.log.dir"`
	MaxSize       int      `yaml:"max.size"`
	MaxBackups    int      `yaml:"max.backups"`
	MaxAge        int      `yaml:"max.age"`
	Compress      bool     `yaml:"compress"`
}

func (c *LoggerConfig) LoadConfig(conf interfaces.ConfigInterface) (err error) {
	err = conf.GetConfig(c)
	if err != nil {
		log.Fatal().Msgf("loading logger config failed : %s", err.Error())
		return
	}

	validate := validator.New()

	err = validate.Struct(c)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			log.Error().Msgf("Error: %s field is required\n", err.Field())
		}
		log.Fatal().Msg("Gateway config Validation failed")
	}
	return
}
