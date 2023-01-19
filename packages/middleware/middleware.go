package middleware

import (
	//"fmt"
	"time"
	"math/rand"
	"github.com/AJ-Brown-InTech/libre-ra/packages/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
 	"github.com/gofiber/redirect/v2"
)

func CreateCookieSession(app *fiber.App, log utils.Logger){
	store := session.New()

	app.Use(func(c *fiber.Ctx) error {
	 	sess, err := store.Get(c)
	 	if err != nil {
			log.Warningf("Cookie session not available, [WARNING]: %v", err)
	 		return err
	 	}
	 	sess.Set("user-session", rand.Int())
		sess.SetExpiry((30 * time.Second))
        sess.ID() 
		log.Infof("Cookie session created, [INFO]: %s", "Succesful Session Created!")

		   c.JSON(fiber.Map{
		 	"message": "Welcome to libre",
		 	"Token":   sess.Get("user-session"),
		 	"authenticated": true,
		 })
		
		return sess.Save()
	 })
}


func MiddlwareAuth(app *fiber.App, log utils.Logger){
	// store := session.New()
	// fmt.Printf("%v", store)

	//  app.Use(func(c *fiber.Ctx) error{
	//  	sess, err := store.Get(c)
	//  	if err != nil {
	// 		log.Errorf("[ERROR]Session Authentication error: %v", err)
	//  		return err
	//  	}
		
	// if sess.Get("user-session") == nil{
	// 		return c.JSON(fiber.Map{
	// 			"message": "unauthenticated try logging in",
	// 			"authenticated": false,
	// 		})
	// 	}

	// 	return c.Next()
	//   })

	app.Use(redirect.New(redirect.Config{
		Rules: map[string]string{
		  "/":   "/home",
		},
		StatusCode: 301,
	  }))

	  log.Infof("log, %v ", app )

	  app.Get("/new", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	  })
}
