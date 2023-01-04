package middleware

import (
	"time"

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
	 	sess.Set("user-session", "1")
		sess.SetExpiry((1 * time.Minute))
         sess.Save()
		 log.Warningf("Cookie session created, [INFO]: %v", sess)
	 	return c.Next()
	 })
}


func CookieAuth(app *fiber.App, log utils.Logger, endpoint string){
	store := session.New()
	 app.Get(endpoint, func(c *fiber.Ctx) error {
	 	sess, err := store.Get(c)
	 	if err != nil {
	 		return err
	 	}
	 	defer sess.Save()

	// 	//EXEC SET

	 	sess.Set("name", "john")

	 	return c.JSON(fiber.Map{
	 		"id":   sess.Get("id"),
	 		"name": sess.Get("name"),
	 	})
	 })
}
