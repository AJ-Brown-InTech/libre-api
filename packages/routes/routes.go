package routes

import (
	"fmt"
	"time"

	//"encoding/json"
	"github.com/AJ-Brown-InTech/libre-ra/packages/middleware"
	"github.com/AJ-Brown-InTech/libre-ra/packages/models"
	"github.com/AJ-Brown-InTech/libre-ra/packages/utils"
	sq "github.com/Masterminds/squirrel"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	//"gopkg.in/guregu/null.v3"
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
		 	return c.Status(fiber.StatusBadRequest).SendString("Bad Request, Email is not valid, try another.")
		}
		 if len(user.UserName) > 24 || len(user.UserName) < 6 { //check for spaces(no spaces allowed)
		 	log.Errorf("Register error with username length, %v", err)
			 return c.Status(fiber.StatusBadRequest).SendString("Bad Request, First Name/Last Name is too long.")
		 }
		 if len(user.Password) < 6 {
		 	log.Errorf("Register error with password length")
			 return c.Status(fiber.StatusBadRequest).SendString("Bad Request, Password too short.")
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
			 return c.Status(fiber.StatusBadRequest).SendString("Bad Request, Couldn't create an account, try different fields.")
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
		userUuid, err := middleware.SessionVerify(c)
		host :=  c.BaseURL()
		if len(userUuid) > 0 {
			//redirect to user account 
			c.Redirect(host + "/account/" + userUuid)
		}
		user :=  models.Account{}
		body := c.BodyParser(&user)
		if  body != nil {
			log.Errorf("Login error. Can't read user input, %v", body)
			return body
		}
	
		account := &models.Account{}
		query := fmt.Sprintf("SELECT * FROM accounts where username = '%s' AND password = '%s' LIMIT 1", user.UserName, user.Password)
		err = db.Get(account,query)
		if err != nil {
			log.Errorf("Error retrieving account from database")
			return  err
		}	
			return c.Redirect(host + "/account/" + account.Uuid)
	}
}

func GetAccountByID(db *sqlx.DB, log utils.Logger) func(c *fiber.Ctx) error {
 	return func(c *fiber.Ctx) error {
		var id models.ID
		 c.ParamsParser(&id) // "{"id": 111}"
		
		account := &models.Account{}
		query := fmt.Sprintf("SELECT * FROM accounts where uuid = '%s' limit 1", id.ID)
		err := db.Get(account,query)
		if err != nil {
			log.Errorf("Error retrieving account from database")
			return  err
		}
		
		if !account.Active {
			c.JSON(fiber.Map{"active": account.Active})
		   }
		return c.JSON(account)
		}		
 }

 func UpdateAccountByID(db *sqlx.DB, log utils.Logger) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
	   var id models.ID
		c.ParamsParser(&id) // "{"id": 111}"
	   
		var payload map[string]interface{}
		err := c.BodyParser(&payload);
		if  err != nil {
			log.Errorf("Login error. Can't read user input, %v", err)
			return err
		}
	   sb := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).RunWith(db)
	   updateStatement := sb.Update("accounts").SetMap(payload).
	   Set("updated_at",time.Now()).
	   Where(sq.Eq{"uuid":id.ID})
	   _, err = updateStatement.Exec()
	   if err != nil {
		   log.Errorf("Database doesn't like your input try again, %v", err)
		   return c.JSON(fiber.Map{"message": "update params didnt work"})
	   }

	   account := &models.Account{}
	   query := fmt.Sprintf("SELECT * FROM accounts where uuid = '%s' limit 1", id.ID)
	   err = db.Get(account,query)
	   if err != nil {
		   log.Errorf("Error retrieving account from database")
		   return  err
	   }
	   if !account.Active {
		c.JSON(fiber.Map{"active": account.Active})
	   }
	   log.Infof(" TEST: %v", account)
	   return c.JSON(account)
	   }		
}

func DeleteAccountByID(db *sqlx.DB, log utils.Logger) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
	   var id models.ID
		c.ParamsParser(&id) // "{"id": 111}"
	   
	   account := &models.Account{}
	   query := fmt.Sprintf("SELECT * FROM accounts where uuid = '%s' limit 1", id.ID)
	   err := db.Get(account,query)
	   if err != nil {
		   log.Errorf("Error retrieving account from database")
		   return  err
	   }

	   account.Active = false
	   sb := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).RunWith(db)
	   updateStatement := sb.Update("accounts").Set("active",account.Active ).
	   Set("updated_at",time.Now()).Where(sq.Eq{"uuid":id.ID})
	   _, err = updateStatement.Exec()
	   if err != nil {
		   log.Errorf("Database doesn't like your input try again, %v", err)
		   return c.JSON(fiber.Map{"message": "deactivate account in database error"})
	   }
	   

	   return c.JSON(fiber.Map{"active": account.Active})
	   }		
}