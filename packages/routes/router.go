package routes

import (
	"github.com/AJ-Brown-InTech/libre-ra/packages/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func RouteManager(app *fiber.App, log utils.Logger, db *sqlx.DB){
	app.Get("/",API(db,log))
	app.Post("/login", Login(db, log))
	app.Post("/register", Register(db, log))
	app.Get("/account/:id", GetAccountByID(db, log))
	app.Put("/account/:id", UpdateAccountByID(db, log))
	app.Delete("/account/:id", DeleteAccountByID(db, log))
	// app.Get("/account/:id/followers", GetAccountFollowers(db, log))
	// app.Get("/account/:id/following", GetAccountFollowing(db, log))
	// app.Get("/account/:id/following/:id", GetAccountFollowingByID(db, log))
	// app.Put("/account/:id/followers/:id", AddAccountFollowingByID(db, log))
	// app.Delete("/account/:id/followers/:id", RemoveAccountFollowingByID(db, log))
	// app.Post("/account/:id/post/", CreatePost(db, log))
	// app.Get("/account/:id/post/:id", GetPostByID(db, log))
	// app.Put("/account/:id/post/:id", UpdatePostByID(db, log))
	// app.Delete("/account/:id/post/:id", DeletePostByID(db, log))

	//api := app.Group("/*", middleware.SessionAuth(app, log)) 
	// app.Group("/",func(c *fiber.Ctx) error {
	// 	// err := c.App()lo
	// 	 return fiber.ErrBadGateway
	// } )
}

