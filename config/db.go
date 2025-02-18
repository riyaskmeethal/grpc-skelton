package config

import (
	"osh.com/rps/registrar/internal/interfaces"
)

type MongoConf map[string]MongoDBconf

type SqlConf map[string]SqlDBconf

type DatabaseConfig struct {
	DBSystem string `yaml:"database.system" validate:"required"`
}

type MongoDatabaseConfig struct {
	SystemConf  System    `yaml:"system" validate:"required"`
	MongoDBconf MongoConf `yaml:"mongodbs" validate:"required"`
}

type SqlDatabaseConfig struct {
	SystemConf System  `yaml:"system" validate:"required"`
	SqlDBconf  SqlConf `yaml:"sqldbs" validate:"required"`
}

type SqlDBconf struct {
	ServerIp string `yaml:"hostip" validate:"required"`
	Port     string `yaml:"port" validate:"required"`
	User     string `yaml:"user" validate:"required"`
	Password string `yaml:"password" validate:"required"`
	DBName   string `yaml:"dbname" validate:"required"`
}

type MongoDBconf struct {
	Host       string `yaml:"host" validate:"required"`
	Name       string `yaml:"namePref" validate:"required"`
	Collection string `yaml:"collection" validate:"required"`
	MaxConn    uint64 `yaml:"maxConn"`
}

type System struct {
	TimeFormat string `yaml:"time.format" validate:"required"`
}

func (dc *DatabaseConfig) LoadDbConfig(conf interfaces.ConfigInterface, log interfaces.LogInterface) (err error) {

	err = conf.GetConfig(dc)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	return ValidateConfig(dc)

}

func (dc *MongoDatabaseConfig) LoadDbConfig(conf interfaces.ConfigInterface, log interfaces.LogInterface) (err error) {

	err = conf.GetConfig(dc)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	if len(dc.MongoDBconf) <= 0 {
		log.Fatal("Database configuration not found.")
	}
	return ValidateConfig(dc)
}

func (dc *SqlDatabaseConfig) LoadDbConfig(conf interfaces.ConfigInterface, log interfaces.LogInterface) (err error) {

	err = conf.GetConfig(dc)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	if len(dc.SqlDBconf) <= 0 {
		log.Fatal("Database configuration not found.")
	}
	return ValidateConfig(dc)
}
