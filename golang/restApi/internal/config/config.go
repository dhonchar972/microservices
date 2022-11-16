package config

import (
	"github.com/spf13/viper"
	"restApi/pkg/logging"
	"sync"
)

type Config struct {
	IsDebug bool `json:"isDebug" env-required:"true"`
	Listen  struct {
		Type   string `json:"type" env-default:"port"`
		BindIp string `json:"bindIp" env-default:"127.0.0.1"`
		Port   string `json:"port" env-default:"8082"`
	} `json:"listen"`
}

var instance *Config
var once sync.Once

func GetConfig() (*Config, error) {
	var returnError error
	once.Do(func() {
		logger := logging.GetLogger()
		logger.Info("read config file")
		instance = &Config{}

		vp := viper.New()
		vp.SetConfigName("appsettings")
		vp.SetConfigType("json")
		vp.AddConfigPath(".")
		err := vp.ReadInConfig()
		if err != nil {
			logger.Info(err)
		}

		err = vp.Unmarshal(&instance)
		if err != nil {
			returnError = err
		}
	})

	if returnError != nil {
		return &Config{}, returnError
	}
	return instance, nil
}
