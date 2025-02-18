package config

import (
	"github.com/go-playground/validator/v10"
	"osh.com/rps/registrar/internal/interfaces"
)

type GateWayConfig struct {
	RestEnable bool    `yaml:"rest.enable"`
	RestHost   string  `yaml:"rest.host" validate:"required"`
	GrpcHost   string  `yaml:"grpc.host" validate:"required"`
	TLSConf    TLSConf `yaml:"tls" validate:"required"`
}

type TLSConf struct {
	Enable   bool   `yaml:"enable"`
	CertFile string `yaml:"certFile"`
	KeyFile  string `yaml:"keyFile"`
}

func (gc *GateWayConfig) LoadGateWayConfig(conf interfaces.ConfigInterface, log interfaces.LogInterface) (err error) {

	err = conf.GetConfig(gc)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	validate := validator.New()

	err = validate.Struct(gc)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			log.Error("Error: %s field is required\n", err.Field())
		}
		log.Fatal("Gateway config Validation failed")
	}
	return
}
