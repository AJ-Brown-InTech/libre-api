package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/AJ-Brown-InTech/libre-ra/config"
	"github.com/AJ-Brown-InTech/libre-ra/packages/database"
	"github.com/AJ-Brown-InTech/libre-ra/packages/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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
	database.NewPsqlDb(confg, appLogger)
	appLogger.Infof("Libre API spects: AppVersion: %s, LogLevel: %s, SSLMode:%v ", confg.Server.AppVersion, confg.Logger.Level, confg.Server.SSL)
	// Connect with database

	// pgDB, err := database.NewPsqlDb(confg);
	// if err != nil{
	// 	appLogger.Errorf("Postgres Database in 2 %s", err)
	// } else {
	// 	appLogger.Infof("Postgres Connected, Status is: 3 %v", pgDB.Stats())
	// }
	 app := fiber.New(fiber.Config{
	 	ServerHeader:         "Libre",
	 	StrictRouting:        true,
	 	Concurrency:          256,
	 	ViewsLayout:          "",
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

	// Custom Timer middleware
	app.Use(Timer())
	
	// Start server
	log.Fatal(app.Listen(":8080"))

}

// Timer will measure how long it takes before a response is returned
func Timer() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// start timer
		start := time.Now()
		// next routes
		err := c.Next()
		// stop timer
		stop := time.Now()
		// Do something with response
		c.Append("Server-Timing", fmt.Sprintf("app;dur=%v", stop.Sub(start).String()))
		// return stack error if exist
		return err
	}
}