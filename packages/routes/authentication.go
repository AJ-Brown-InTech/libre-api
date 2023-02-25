package routes

import (
	"time"
	//"encoding/json"
	"github.com/AJ-Brown-InTech/libre-ra/packages/models"
	"github.com/AJ-Brown-InTech/libre-ra/packages/utils"
	sq "github.com/Masterminds/squirrel"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

// User log into account
func Login(db *sqlx.DB, log utils.Logger) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		//var response bool; //db response  either data(user) is present or not

		// err := c.BodyParser(user)
		// if err !=nil {

		// 	fmt.Printf("EROROROROROR%v", err)
		// }

		// //login params
		// email := user.Email
		// username := user.Name
		// password := user.Password

		// fmt.Printf("%v", len(email))
		// if password == ""{

		// 	return c.JSON(fiber.Map{"message": "Password needed"})
		// }
		// if email == "" && username == ""{

		// 	return c.JSON(fiber.Map{"message":"Email or Username needed"})
		// }

		// if email != "" ||  username != ""{
		// 	// var user_check string;
		// 	// if len(email) > 0 {
		// 	// 	user_check = email
		// 	// }
		// 	//run query to check is user exist
		// 	//if user present return user id and token to store on device & set response to true
		// 	//else set response to false with error message

		// }
		// fmt.Printf("testone1111 %v", user)
		return c.Context().Err()
	}
}

// user register for an account
func Register(db *sqlx.DB, log utils.Logger) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		//read in user data
		user := new(models.Account)
		err := c.BodyParser(user);
		if  err != nil {
			log.Errorf("ERROR: Register %v", err)
			return err
		}
		 if !utils.Valid(user.Email) {
		 	log.Errorf("Email error %v", err)
		 	return c.JSON(fiber.Map{"message": "Email is not valid, try another."})
		}
		 if len(user.UserName) > 24 || len(user.UserName) < 6 { //check for spaces(no spaces allowed)
		 	log.Errorf("Register error with username length, %v", err)
		 	return c.JSON(fiber.Map{"message": "First Name/Last Name is too long."})
		 }
		 if len(user.Password) < 6 {
		 	log.Errorf("Register error with password length")
		 	return c.JSON(fiber.Map{"message": "Password too short."})
		 }
		 log.Infof("%v",user.Dob)
		 //data cleanup before the db query
		 user.Uuid = uuid.New().String()
		 user.Rating = 0
		 user.Verfied = false
		 user.Active = true
		 user.CreadtedAt = time.Now()
		 user.UpdatedAt = time.Now()

		//create an account for user
		 query := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).RunWith(db)
		 insertStatement := query.Insert("Accounts").
		 	Columns("username", "uuid", "dob", "password", "email", "rating", "active", "verified", "created_at", "updated_at").
		 	Values(user.UserName, user.Uuid, user.Dob, user.Password, user.Email, user.Rating, user.Active, user.Verfied, user.CreadtedAt, user.UpdatedAt)
		 _, err = insertStatement.Exec()
		 if err != nil {
		 	log.Errorf("Database doesn't like your input try again, %v", err)
		 	return c.JSON(fiber.Map{"message": "Couldn't create an account, try different fields."})
		 }
		 
		 log.Infof("TEST!!!:%v")
		// return a url to home
		return c.Context().Err()
	}
}
