package gateway

import (
	"github.com/go-playground/validator/v10"
	"osh.com/rps/registrar/internal/interfaces"
)

type GateWayConfig struct {
	RestEnable bool    `yaml:"rest.enable" validate:"required"`
	RestHost   string  `yaml:"rest.host" validate:"required"`
	GrpcHost   string  `yaml:"grpc.host" validate:"required"`
	TLSConf    TLSConf `yaml:"tls" validate:"required"`
}

type TLSConf struct {
	Enable   bool   `yaml:"enable" validate:"required"`
	CertFile string `yaml:"certFile"`
	KeyFile  string `yaml:"keyFile"`
}

func (gc *GateWayConfig) LoadGateWayConfig(conf interfaces.ConfigInterface, log interfaces.LogInterface) (Config *GateWayConfig, err error) {
	Config = new(GateWayConfig)
	err = conf.GetConfig(Config)
	if err != nil {
		log.Error(err.Error())
	}

	validate := validator.New()

	err = validate.Struct(Config)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			log.Error("Error: %s field is required\n", err.Field())
		}
		log.Error("Gateway config Validation failed")
	}
	return
}
