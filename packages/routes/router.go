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
	var response bool; //db response  either data(user) is present or not
	user := new(User)
	err := c.BodyParser(user)
	if err !=nil {
		
		fmt.Printf("EROROROROROR%v", err)
	}

	//login params
	email := user.Email
	username := user.Name
	password := user.Password

	if password == ""{

		return c.JSON(fiber.Map{"message": "Password needed"})
	}
	if email == "" && username == ""{
		
		return c.JSON(fiber.Map{"message":"Email or Username needed"})
	}
	
	if(email != "" || username != ""){
		//run query to check is user exist 
	
		//if user present return user id and token to store on device & set response to true
		//else set response to false with error message
		response = false
	} 
			
	
	fmt.Printf("testone1111 %v", user)
	return c.Context().Err()
}