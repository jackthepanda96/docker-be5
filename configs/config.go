package configs

import (
	"fmt"
	"sync"

	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
)

// type AppConfig struct {
// 	Port     int `yaml:"port"`
// 	Database struct {
// 		Driver   string `yaml:"driver"`
// 		Name     string `yaml:"name"`
// 		Address  string `yaml:"address"`
// 		Port     int    `yaml:"port"`
// 		Username string `yaml:"username"`
// 		Password string `yaml:"password"`
// 	}
// }
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
	// defaultConfig.Port = 8000
	// defaultConfig.Database.Driver = ""
	// defaultConfig.Database.Name = ""
	// defaultConfig.Database.Address = ""
	// defaultConfig.Database.Port = 3306
	// defaultConfig.Database.Username = ""
	// defaultConfig.Database.Password = ""
	defaultConfig.Port = 8000
	defaultConfig.Driver = ""
	defaultConfig.Name = ""
	defaultConfig.Address = ""
	defaultConfig.DB_Port = 3306
	defaultConfig.Username = ""
	defaultConfig.Password = ""

	// viper.SetConfigType("env")
	// viper.SetConfigName("config")
	// viper.AddConfigPath("./configs/")
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
	fmt.Println(finalConfig)
	return &finalConfig
}
