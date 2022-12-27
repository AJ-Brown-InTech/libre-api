package main

import (
	"fmt"
	"log"
	"os"

	"github.com/AJ-Brown-InTech/libre-api/config"
	"github.com/AJ-Brown-InTech/libre-api/server"
	"github.com/AJ-Brown-InTech/libre-api/utils"
	"github.com/AJ-Brown-InTech/libre-api/database"
)

func main (){
fmt.Printf("api server starting %s\n", utils.Timestamp())

configPath := utils.GetConfigPath(os.Getenv("config")) 

	cfgFile, err:= config.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("Load config file fail %v", err)
	}
	confg, err := config.ParseConfig(cfgFile)
	if err != nil {
		log.Fatalf("Parsing config file fail %v", err)
	}
	appLogger := utils.NewApiLogger(confg)

	//initalize new logger and connect to db
	appLogger.InitLogger()
	appLogger.Infof("AppVersion: %s, LogLevel: %s, SSLMode:%v, ServerMode: %s ", confg.Server.AppVersion, confg.Logger.Level, confg.Server.SSL, confg.Server.Mode)
	pgDB, err := database.NewPsqlDb(confg);
	if err != nil{
		appLogger.Errorf("Postgres Database in 2 %s", err)
	} else {
		appLogger.Infof("Postgres Connected, Status is: 3 %v", pgDB.Stats())
	}


	s := server.NewServer(confg,pgDB,appLogger)
	if err = s.Run(); err != nil {
		appLogger.Panicf("Server it not running, %v", err)
	}



}