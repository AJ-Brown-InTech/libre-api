package main

import (
	"log"
	"os"
	"time"
	"github.com/AJ-Brown-InTech/libre-ra/config"
	"github.com/AJ-Brown-InTech/libre-ra/packages/database"
	"github.com/AJ-Brown-InTech/libre-ra/packages/routes"
	"github.com/AJ-Brown-InTech/libre-ra/packages/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/cache"
)

func main(){
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

	//database.NewPsqlDb(confg, appLogger)
	appLogger.Infof("Libre API Spects: AppVersion: %s, LogLevel: %s, SSLMode:%v ", confg.Server.AppVersion, confg.Logger.Level, confg.Server.SSL)
	
	// Connect with database
	  pgDB, err := database.NewPsqlDb(confg, appLogger);
	  if err != nil{
	  	appLogger.Errorf("Postgres Database connection error, [ERROR]: %s", err)
	  } else {
	  	appLogger.Infof("Postgres Connected, [INFO]: %v", pgDB.Stats())
	  }

	  err = pgDB.Ping()
	  if err != nil {
		appLogger.Errorf("Failed to ping, %v", err)
	  }

	 app := fiber.New(fiber.Config{
	 	ServerHeader:         "Libre",
	 	StrictRouting:        true,
	 	Concurrency:          256,
	 	ReadTimeout:          time.Second * 10,
	 	WriteTimeout:         time.Second * 10,
	 	IdleTimeout:          time.Second * (10 * 2),
	 	AppName:              "Libre Api v1.0.0",
	 	EnablePrintRoutes:    true,
	 	ColorScheme:          fiber.Colors{},
	 	RequestMethods:       []string{},
	 })
	 
	// CORS for external resources
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowHeaders:     "Cache-Control",
		AllowCredentials: true,
	}))

	//middleware
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(cache.New())
	
	//Routes/RouteManager
	routes.RouteManager(app, appLogger, pgDB)
	// Start server
	appLogger.Panicf("%v",app.Listen(":8080") )
}


