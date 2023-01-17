package middleware

import (
	"fmt"
	"time"
	"math/rand"
	"github.com/AJ-Brown-InTech/libre-ra/packages/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
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
	store := session.New()
	fmt.Printf("%v", store)

	 app.Use(func(c *fiber.Ctx) error{
	 	sess, err := store.Get(c)
	 	if err != nil {
			log.Errorf("[ERROR] %v", err)
	 		return err
	 	}
	
		// return c.JSON(fiber.Map{
		// 	"keys": sess.Keys(),
		// 	"name": sess.Get("Name"),
		// 	"id":   sess.Get("user-session"),
		// })
			
		if sess.Get("user-session") == nil{
			return c.JSON(fiber.Map{
				"message": "unauthenticated",
				"authenticated": false,
			})
		}
		return c.Next()
	  })
}

// AuthReq middleware
// func AuthReq() func(*fiber.Ctx) {
//     cfg := basicauth.Config{
//         Users: map[string]string{
//           config.Config("USERNAME"): config.Config("PASSWORD"),
//         },
//       }
//     err := basicauth.New(cfg);
//     return err
// }

//  	defer sess.Save()

// // 	//EXEC SET

//  	sess.Set("name", "john")

//  	return c.JSON(fiber.Map{
//  		"id":   sess.Get("id"),
//  		"name": sess.Get("name"),
//  	})