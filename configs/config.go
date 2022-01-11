package configs

import (
	"sync"

	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
)

type AppConfig struct {
	Port     int `yaml:"port"`
	Database struct {
		Driver   string `yaml:"driver"`
		Name     string `yaml:"name"`
		Address  string `yaml:"address"`
		Port     int    `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	}
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
	defaultConfig.Database.Driver = "mysql"
	defaultConfig.Database.Name = "defal"
	defaultConfig.Database.Address = "dbbe5.c7svlamncepc.ap-southeast-1.rds.amazonaws.com"
	defaultConfig.Database.Port = 3306
	defaultConfig.Database.Username = "admin"
	defaultConfig.Database.Password = "admin123"

	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath("./layered/configs/")
	if err := viper.ReadInConfig(); err != nil {
		log.Warn("failed to open file")
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
