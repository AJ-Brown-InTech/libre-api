package routes

import (
	//"fmt"
	//-Brown-InTech/libre-ra/packages/middleware"
	//"github.com/AJ-Brown-InTech/libre-ra/packages/middleware"
	"github.com/AJ-Brown-InTech/libre-ra/packages/middleware"
	"github.com/AJ-Brown-InTech/libre-ra/packages/utils"
	"github.com/gofiber/fiber/v2"
)



func RouteManager(app *fiber.App, log utils.Logger){

	api := app.Group("/*", middleware.SessionAuth(app, log)) 
	// app.Group("/",func(c *fiber.Ctx) error {
	// 	// err := c.App()lo
	// 	 return fiber.ErrBadGateway
	// } )
//fmt.Printf("%v", app)
// fmt.Println("")
// fmt.Printf("%v", log)
// fmt.Printf("%v", app)
//fmt.Printf("%v", auth)
//api := app.Group("/*",  )




}


