package middleware

import (
	"time"
	"github.com/gofiber/fiber/v2"
)


func CreateSession(c *fiber.Ctx, name string, value string){
	cookie := new(fiber.Cookie)
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.Name  = name
	cookie.Value = value
	c.Cookie(cookie)
}

func SessionVerify(c *fiber.Ctx)(string, error){
	userUuid := c.Cookies("aj")
	 if len(userUuid) > 0{
	 	return userUuid, nil
	 }
	return "", c.Context().Err()	
}