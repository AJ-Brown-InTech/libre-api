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

