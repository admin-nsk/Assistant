package config

import (
	"note_service/app/pkg/logging"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	IsDebug *bool `ysml:"is_debug"`
	JWT struct {
		Secret string `json:"secret" env-required:"true"`
	}
	Listen  struct {
		Type   string `yaml:"type" env-default="port"`
		BindIP string `yaml:"bind_ip" env-default="localhost"`
		Port   string `yaml:"port" env-default="8080"`
	}
}

var instance *Config
var once sync.Once


func GetConfig() *Config {
	once.Do(func() {
		logger := logging.Getlogger()
		logger.Info("read application config...")
		instance = &Config{}
		if err := cleanenv.ReadConfig("config.yml", instance); err != nil {
			help, _ := cleanenv.GetDescription(&instance, nil)
			logger.Info(help)
			logger.Fatal(err)
		}
	})

	return instance
}
