package utils

import (
	"io"
	"os"
	"time"

	"github.com/AJ-Brown-InTech/libre-api/config"
	"github.com/sirupsen/logrus"
)


type Logger interface{
	InitLogger()
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warningf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Panicf(format string, args ...interface{})
}


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

func (x *apiLogger) getLoggerLevel(cfg *config.Config) logrus.Level{
	level, exist := logLevelMapper[cfg.Logger.Level]
	if !exist {
		return logrus.DebugLevel
	}
	return level
}


func (x *apiLogger) InitLogger()  {
	Formatter := new(logrus.TextFormatter)
    Formatter.TimestampFormat = "02-01-2006 15:04:05"
    Formatter.FullTimestamp = true
	Formatter.DisableColors = false
    logrus.SetFormatter(Formatter)

	logWriter := x.base.Writer()
	//logLevel := x.getLoggerLevel(x.cfg)

	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
    if err != nil {
        x.base.Fatal(err)
    }
	 io.MultiWriter(logWriter, file )
}


//Logger methods dont wanna add more becasue its redundant
func (x *apiLogger) Debugf(format string, args ...interface{}) {
	x.base.Debugf(format, args)
}

func (x *apiLogger) Infof(format string, args ...interface{}) {
	x.base.Infof(format, args)
}

func (x *apiLogger) Warningf(format string, args ...interface{}) {
	x.base.Warningf(format, args)
}

func (x *apiLogger) Errorf(format string, args ...interface{}) {
	x.base.Errorf(format, args)
}

func (x *apiLogger) Panicf(format string, args ...interface{}) {
	x.base.Panicf(format, args)
}

