package cache

import (
	"encoding/json"

	"github.com/pericles-luz/go-base/pkg/conf"
)

type Config struct {
	file     conf.ConfigBase
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Password string `json:"password"`
	DB       int    `json:"db"`
}

func (c *Config) Load(file string) error {
	raw, err := c.file.ReadConfigurationFile(file)
	if err != nil {
		return err
	}
	err = json.Unmarshal([]byte(raw), c)
	if err != nil {
		return err
	}
	return nil
}

func (c *Config) GetConfig() map[string]interface{} {
	return map[string]interface{}{
		"DE_Host":  c.Host,
		"NU_Port":  c.Port,
		"PW_Senha": c.Password,
		"NU_Banco": c.DB,
	}
}

func NewConfig() *Config {
	return &Config{}
}
