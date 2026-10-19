package config

import (
	"scaffold-demo/utils/logs"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const (
	TimeFormat = "2006-01-02 15:04:05"
)

var (
	Port string
)

func initLogConfig(logLevel string) {
	if logLevel == "debug" {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.InfoLevel)
	}

	logrus.SetReportCaller(true)
	logrus.SetFormatter(&logrus.JSONFormatter{TimestampFormat: TimeFormat})
}

func init() {
	viper.SetDefault("LOG_LEVEL", "debug")
	viper.SetDefault("PORT", ":8080")
	viper.AutomaticEnv()
	logLevel := viper.GetString("LOG_LEVEL")
	Port = viper.GetString("PORT")
	initLogConfig(logLevel)
	logs.Info(nil, "初始化")
}
