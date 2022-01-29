package config

import (
	"log"
	"original-covid-app-japan-prefecture-backend-echo/utils"

	"gopkg.in/go-ini/ini.v1"
)

type ConfigList struct {
	Port    string
	LogFile string
}

func init() {
	LoadConfig()
	utils.LoggingSettings(Config.LogFile)
}

var Config ConfigList

func LoadConfig() ConfigList {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Fatalln(err)
	}

	Config = ConfigList{
		Port:    cfg.Section("web").Key("port").String(),
		LogFile: cfg.Section("web").Key("logfile").String(),
	}

	return Config
}
