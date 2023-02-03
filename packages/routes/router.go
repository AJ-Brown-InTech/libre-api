package routes

import (

	"github.com/AJ-Brown-InTech/libre-ra/packages/utils"
	"github.com/gofiber/fiber/v2"
)



func RouteManager(app *fiber.App, log utils.Logger){
	//Login
	app.Server()
	//api := app.Group("/*", middleware.SessionAuth(app, log)) 
	// app.Group("/",func(c *fiber.Ctx) error {
	// 	// err := c.App()lo
	// 	 return fiber.ErrBadGateway
	// } )

}