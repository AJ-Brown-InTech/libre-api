package middleware

import (

	"math/rand"
	"time"
	 "strconv"
	"github.com/AJ-Brown-InTech/libre-ra/packages/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

const  KEY = "libre-session"

func rando() string{
	first := rand.Int()
	second := rand.Int()
	third := (first * second / 10) / 2
	final := "libre" + strconv.Itoa(third) + "erbil"
	
	return final

}
func CreateCookieSession(app *fiber.App, log utils.Logger){
	// storage instance for session cookie
	store := session.New()
	app.Use(func(c *fiber.Ctx) error {
		//create session
		store.CookieHTTPOnly =  true
		store.CookieSecure = true
		store.KeyLookup = KEY
		store.Expiration = 30 * time.Second
		store.KeyGenerator = rando 
		
		// store session	
		session, err := store.Get(c)
		if err != nil{
			log.Errorf("[ERROR] %v,", err)
			return err
		}
		log.Infof("Success, %s", "session created. You're authorized now!")
		return session.Save()
  })
}


func SessionAuth(app *fiber.App, log utils.Logger){
	//access storage instance
	store := session.New()

	//obtain the session
	app.Use(func(c *fiber.Ctx) string {
	session, err:= store.Get(c);

	//check if its present => if present send to home else send to unauthorized page
	if session.Get(KEY) == nil {
		log.Errorf("%v", err)
		return "no auth" 
	} else {
		session.Regenerate()
		//route = c.OriginalURL()
		log.Infof("%v",  c.Route().Path) 
	
		 c.Next()
		 return c.OriginalURL()
	}
	
		})
}

