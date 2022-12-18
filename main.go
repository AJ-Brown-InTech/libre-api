package main

import (
	"fmt"
	"log"
	"os"

	"github.com/AJ-Brown-InTech/libre-api/config"
	"github.com/AJ-Brown-InTech/libre-api/server"
	"github.com/AJ-Brown-InTech/libre-api/utils"
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
	fmt.Println("pigs")
	



	server.Server()
}