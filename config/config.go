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
	Port            string
	JwtSignKey      string
	JWT_EXPIRE_TIME int64
	Username		string
	Password		string
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    map[string]any `json:"data,omitempty"`
}

func NewResponse(code int, message string, data map[string]any) *Response {
	return &Response{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

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
	viper.SetDefault("JWT_EXPIRE_TIME", 120)
	viper.SetDefault("JWT_SIGN_KEY", "abc")
	viper.SetDefault("USERNAME", "admin")
	viper.SetDefault("PASSWORD", "password")
	viper.AutomaticEnv()

	logLevel := viper.GetString("LOG_LEVEL")
	Port = viper.GetString("PORT")
	JWT_EXPIRE_TIME = viper.GetInt64("JWT_EXPIRE_TIME")
	initLogConfig(logLevel)
	logs.Info(nil, "初始化")
	JwtSignKey = viper.GetString("JWT_SIGN_KEY")
	Username = viper.GetString("USERNAME")
	Password = viper.GetString("PASSWORD")
	logs.Info(map[string]any{"JWT_EXPIRE_TIME": JWT_EXPIRE_TIME, "JwtSignKey": JwtSignKey, "Username": Username, "Password": Password}, "配置文件加载完成")	

}
