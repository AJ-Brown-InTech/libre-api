package routes

import (

	"github.com/AJ-Brown-InTech/libre-ra/packages/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func RouteManager(app *fiber.App, log utils.Logger, db *sqlx.DB){
	//Login
	app.Post("/login", Login(db, log))
	//Register
	app.Post("/register", Register(db, log))

	//api := app.Group("/*", middleware.SessionAuth(app, log)) 
	// app.Group("/",func(c *fiber.Ctx) error {
	// 	// err := c.App()lo
	// 	 return fiber.ErrBadGateway
	// } )
}

