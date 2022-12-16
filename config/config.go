package config

import (
	"fmt"
	"time"

)

// App config struct
type Config struct {
	Server   ServerConfig
	Postgres PostgresConfig
	MongoDB  MongoDB
	Cookie   Cookie
	Session  Session
	Metrics  Metrics
	Logger   Logger
}


// Server config struct
type ServerConfig struct {
	AppVersion        string
	Port              string
	PprofPort         string
	Mode              string
	JwtSecretKey      string
	CookieName        string
	ReadTimeout       time.Duration
	WriteTimeout      time.Duration
	SSL               bool
	CtxDefaultTimeout time.Duration
	CSRF              bool
	Debug             bool
}

// Logger config
type Logger struct {
	Development       bool
	DisableCaller     bool
	DisableStacktrace bool
	Encoding          string
	Level             string
}

// Postgresql config
type PostgresConfig struct {
	PostgresqlHost     string
	PostgresqlPort     string
	PostgresqlUser     string
	PostgresqlPassword string
	PostgresqlDbname   string
	PostgresqlSSLMode  bool
	PgDriver           string
}

// MongoDB config
type MongoDB struct {
	MongoURI string
}

type Metrics struct {
	URL         string
	ServiceName string
}

// Cookie config
type Cookie struct {
	Name     string
	MaxAge   int
	Secure   bool
	HTTPOnly bool
}


// Session config
type Session struct {
	Prefix string
	Name   string
	Expire int
}

func Configuration (){
	fmt.Println("config here")
}