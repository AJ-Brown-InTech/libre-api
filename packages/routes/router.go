// Structure of the Router Pacakage
// -----> Stucts
// -----> Router
// -----> Handler(Individual Routes)
package routes

import (
	"fmt"
	"net/mail"
	"time"
	"github.com/AJ-Brown-InTech/libre-ra/packages/utils"
	"github.com/gofiber/fiber/v2"
)

//structs
type User struct {
    Name string `json:"username"`
    Email string `json:"email"` 
	Password string `json:"password"`   
	Birthdate string `json:"birthdate"`
	FirstName string `json:"firstName"`   
	LastName string `json:"lastName"`   
}


//Router
func RouteManager(app *fiber.App, log utils.Logger){
	//Login
	app.Post("/login", Login)
	//Register
	app.Post("/register", Register)

	//api := app.Group("/*", middleware.SessionAuth(app, log)) 
	// app.Group("/",func(c *fiber.Ctx) error {
	// 	// err := c.App()lo
	// 	 return fiber.ErrBadGateway
	// } )
}

// Handlers
func Login(c *fiber.Ctx) error {	
	//var response bool; //db response  either data(user) is present or not
	user := new(User)
	err := c.BodyParser(user)
	if err !=nil {
		
		fmt.Printf("EROROROROROR%v", err)
	}

	//login params
	email := user.Email
	username := user.Name
	password := user.Password

	fmt.Printf("%v", len(email))
	if password == ""{

		return c.JSON(fiber.Map{"message": "Password needed"})
	}
	if email == "" && username == ""{
		
		return c.JSON(fiber.Map{"message":"Email or Username needed"})
	}
	
	if email != "" ||  username != ""{
		// var user_check string; 
		// if len(email) > 0 {
		// 	user_check = email
		// }
		//run query to check is user exist 
		//if user present return user id and token to store on device & set response to true
		//else set response to false with error message
		
	} 
	fmt.Printf("testone1111 %v", user)
	return c.Context().Err()
}

func valid(email string) bool {
    _, err := mail.ParseAddress(email)
    return err == nil
}

func Register(c *fiber.Ctx) error {
	var (
		email, username, password, first_name, last_name, birthday string
	)

	user:= new(User)
	
	
	if err := c.BodyParser(user); err != nil{
		fmt.Printf("Register error %v", err)
	}	

	email = user.Email
	username = user.Name
	password = user.Password
	first_name = user.FirstName
	last_name = user.LastName
	birthday = user.Birthdate

	//verify email is valid
	 if  valid(email) == false{
	 	return c.JSON(fiber.Map{"message": "Email is not valid."})
	 }	
	 //nested conditionals so we verify in order from front-end input
	 //verify username is eligble
	 if len(username) > 24 {
		// also need to check if username exist in db
		return c.JSON(fiber.Map{"message": "Username is too long."})
	 }else if len(username) < 6{
		return c.JSON(fiber.Map{"message": "Username is too short."})
	 }
	//verify firstname and lastname is eligble
	if len(first_name) > 24 || len(last_name) > 24 {
		return c.JSON(fiber.Map{"message": "First Name/Last Name is too long."})
	} 
	if len(password) < 6 {
	 	return c.JSON(fiber.Map{"message": "Password too short."})
	}
	//check if date is valid
	t, err := time.Parse("01/02/2006", birthday)
	if err != nil{
		fmt.Printf("Register error with birthdate, %v", err)
	}
	//run query to create user and add other data if you want
	//then redirect to home page
	fmt.Printf("birthdate, %v", t)

return c.Context().Err()
}