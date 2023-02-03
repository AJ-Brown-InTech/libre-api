// Structure of the Router Pacakage
// -----> Stucts
// -----> Router
// -----> Handler(Individual Routes)
package routes
import (
	"fmt"
	"github.com/AJ-Brown-InTech/libre-ra/packages/utils"
	"github.com/gofiber/fiber/v2"
	//"encoding/json"  
)

//structs 
type User struct {
    Name string `json:"name"`
    Email string `json:"email"` 
	Password string `json:"password"`   
}


//Router
func RouteManager(app *fiber.App, log utils.Logger){
	//Login
	app.Post("/login", Login)
	//api := app.Group("/*", middleware.SessionAuth(app, log)) 
	// app.Group("/",func(c *fiber.Ctx) error {
	// 	// err := c.App()lo
	// 	 return fiber.ErrBadGateway
	// } )
}

// Handlers
func Login(c *fiber.Ctx) error {		
	user := new(User)
	if err := c.BodyParser(user); err !=nil {
		fmt.Printf("%v", err)
	}
	//fmt.Printf("n\test %v", c.Request())
	data := c.BodyParser(user)
	// fmt.Printf("test %v", user.Name)
	 fmt.Printf("test %v", data)
	return c.Context().Err()
}