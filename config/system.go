package config

import (
	"github.com/go-playground/validator/v10"
	"osh.com/rps/registrar/internal/interfaces"
)

type SystemConfig struct {
	System struct {
		TimeFormat string `yaml:"time.format" validate:"required"`
	} `yaml:"system" validate:"required"`
}

func (sc *SystemConfig) LoadSystemConfig(conf interfaces.ConfigInterface, log interfaces.LogInterface) (err error) {

	err = conf.GetConfig(sc)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	validate := validator.New()

	err = validate.Struct(sc)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			log.Error("Error: %s field is required\n", err.Field())
		}
		log.Fatal("System config Validation failed")
	}
	return
}
