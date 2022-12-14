package conf

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const (
	linuxCfg   = "/var/config/config.yaml"
	devCfg     = "./config.yaml"
	projectCfg = "github.com/doge-verse/easy-upgrade-backend/config.yaml"
)

// Init
// linux default path > windows default path > local
func Init() {
	cfgFile := linuxCfg
	if _, err := os.Stat(cfgFile); err == nil {
		initViper(cfgFile)
		logrus.Infoln("Using linux configuration")
		return
	}

	cfgFile = devCfg
	if _, err := os.Stat(cfgFile); err == nil {
		initViper(cfgFile)
		logrus.Infoln("Using cfgFile configuration")
		return
	}

	cfgFile = os.Getenv("GOPATH") + projectCfg
	if _, err := os.Stat(cfgFile); err != nil {
		panic("Find config error: " + err.Error())
	}
	initViper(cfgFile)
	logrus.Infoln("Using projectCfg configuration")
}

func initViper(file string) {
	viper.SetConfigFile(file)
	if err := viper.ReadInConfig(); err != nil {
		panic("Read config error: " + err.Error())
	}
	fmt.Println("Using config file:", viper.ConfigFileUsed())
}
