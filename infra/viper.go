package infra

import (
	"github.com/spf13/viper"
	"log"
	"path/filepath"
	"sync"
)

var cv *viper.Viper
var cvLock sync.RWMutex

func NewViper(cfgPath string) *viper.Viper {
	if len(cfgPath) > 0 {
		viper.SetConfigFile(cfgPath)
		viper.SetConfigType("toml")
		viper.AddConfigPath(filepath.Base(cfgPath))
	} else {
		viper.SetConfigName("config") // name of config file (without extension)
		viper.SetConfigType("toml")
		/*viper.AddConfigPath("../etc/")
		viper.AddConfigPath("./etc/")*/
		viper.AddConfigPath(AppPath + "configs/") // TODO:路径不应该先死
	}

	ReloadViper()

	//spew.Dump(cv)

	return cv
}

func ReloadViper() *viper.Viper {
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatal(err)
		} else {
			log.Fatal(err)
		}
	}

	if err := viper.Unmarshal(&cv); err != nil {
		log.Fatal(err)
	}

	return cv
}
