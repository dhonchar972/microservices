/*
* This package contains a standard set of configuration parameters and interaction methods for the correct application work.
 */
package config

import (
	"restApi/pkg/logging"
	"sync"

	"github.com/spf13/viper"
)

/*
	Stcucture Config contains set of configuration parameters which takes from JSON file(appsettings.json).

Parameters:

	Type - connection type
	BindIp - connection IP address
	Port - connection port
*/
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

// Return unmarshalled configuration structure.
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
