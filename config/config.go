package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
	"osh.com/rps/registrar/internal/utils"
)

type Config struct {
	data []byte
}

func (c Config) GetConfig(value any) (err error) {
	if err := yaml.Unmarshal(c.data, value); err != nil {
		return err
	}

	return err
}

func (c Config) GetConfigByteData() []byte {
	return c.data
}

func InitConfig(confFile string, initConf bool) (config Config, err error) {

	if initConf && !FileExists(confFile) {
		err = CreateYamlConfFile(confFile)
		if err != nil {
			return config, err
		}
	}
	config = Config{}
	config.data, err = os.ReadFile(confFile)
	if err != nil {
		return
	}

	return

}

// pass key as "StructFieldName.FieldName" for nested struct.
func (config Config) GetConfigValue(key string) (interface{}, error) {
	// Get the field by name

	val, err := utils.GetValueByField(config, key)

	// Check if the field exists and is valid
	if err != nil {
		return nil, fmt.Errorf("no such field: %s", key)
	}
	return val, nil
}

func FileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func CreateYamlConfFile(fileName string) error {

	defualtConfigs := new(Config)

	data, err := yaml.Marshal(defualtConfigs)
	if err != nil {
		return err
	}

	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}
