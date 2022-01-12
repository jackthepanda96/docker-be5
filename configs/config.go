package configs

import (
	"sync"

	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
)

type AppConfig struct {
	Port     int
	Driver   string
	Name     string
	Address  string
	DB_Port  int
	Username string
	Password string
}

var lock = &sync.Mutex{}
var appConfig *AppConfig

func GetConfig() *AppConfig {
	lock.Lock()
	defer lock.Unlock()

	if appConfig == nil {
		appConfig = initConfig()
	}

	return appConfig
}

func initConfig() *AppConfig {
	var defaultConfig AppConfig
	defaultConfig.Port = 8000
	defaultConfig.Driver = ""
	defaultConfig.Name = ""
	defaultConfig.Address = ""
	defaultConfig.DB_Port = 3306
	defaultConfig.Username = ""
	defaultConfig.Password = ""

	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		log.Warn(err)
		return &defaultConfig
	}

	var finalConfig AppConfig
	err := viper.Unmarshal(&finalConfig)
	if err != nil {
		log.Warn("failed to extract external config, use default value")
		return &defaultConfig
	}

	return &finalConfig
}
