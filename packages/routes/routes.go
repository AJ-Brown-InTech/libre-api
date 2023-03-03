package routes

import (
	"time"
	//"encoding/json"
	"github.com/AJ-Brown-InTech/libre-ra/packages/middleware"
	"github.com/AJ-Brown-InTech/libre-ra/packages/models"
	"github.com/AJ-Brown-InTech/libre-ra/packages/utils"
	sq "github.com/Masterminds/squirrel"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

//API Home Page
func API(db *sqlx.DB, log utils.Logger) func(c *fiber.Ctx) error{
	return func(c *fiber.Ctx) error {
	
		middleware.SessionVerify(c)
		return 	c.JSON(fiber.Map{
				"message": "Welcome to libre.",
				"instruction": "Login or register to utilize api"})}
}

// user registeration for an account
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
		 //create cookie session
		 middleware.CreateSession(c,user.UserName,user.Uuid)
		 //redirect to home
		// return a url to home
		return c.Context().Err()
	}
}

// User log into account
func Login(db *sqlx.DB, log utils.Logger) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		userUuid, _ := middleware.SessionVerify(c)
		if len(userUuid) > 0 {
			//redirect to user account
		}
		user := new(models.Account) // read in datat
		err := c.BodyParser(&user);
		if  err != nil {
			log.Errorf("Login error. Can't read user input, %v", err)
			return err
		}

		var account models.Account
		err = db.Get(&account,"SELECT * FROM accounts WHERE username = $1 AND password = $2", user.UserName, user.Password )
		if err != nil {
			log.Errorf("Error retrieving account from dataase")
			return  err
		}

		if account.UserName == user.UserName && account.Password == user.Password{
			// get user by id and return their account
		}
		
		return c.Context().Err()
	}
}

func GetAccountByID(db *sqlx.DB, log utils.Logger) func(c *fiber.Ctx) error {
 	return func(c *fiber.Ctx) error {
		var id models.ID
		 c.ParamsParser(&id) // "{"id": 111}"
		log.Infof(" TEST: %v", id)
		
		var user models.Account
		err := db.Select(&user, "SELECT * FROM accounts WHERE uuid = $1", id)
		if err !=nil{
			log.Errorf("Error fetching user ")
		}
		return c.JSON(user)
		}		
 }