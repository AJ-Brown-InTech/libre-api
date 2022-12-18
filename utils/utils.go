package utils

import (
	"time"

	"github.com/AJ-Brown-InTech/libre-api/config"
	"github.com/sirupsen/logrus"
)

//current time timestamp
func Timestamp ()string{
	currentTime := time.Now()
	var timestamp string = currentTime.Format("01-02-2006 15:04:05")
	return timestamp
}

// Get config path for local or docker
func GetConfigPath(path string ) string  {
	if path == "docker"{
		return "./config/config-docker"
	}
	 return "./config/config-local"
}

//Logger
type apiLogger struct {
	cfg  *config.Config
	base *logrus.Logger
}

//Logger constructor
func NewApiLogger(cfg *config.Config) *apiLogger{
	return &apiLogger{cfg:cfg}
}

//mapper for logrus log levels

var logLevelMapper = map[string]logrus.Level{
	"Debug": logrus.DebugLevel,
	"Info": logrus.InfoLevel,
	"Warning": logrus.WarnLevel,
	"Error": logrus.ErrorLevel,
	"Panic": logrus.PanicLevel,
}

func getLoggerLevel(cfg *config.Config)logrus.Level{
	level, exist := logLevelMapper[cfg.Logger.Level]
	if !exist {
		return logrus.DebugLevel
	}
	return level
}