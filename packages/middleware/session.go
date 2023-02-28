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

func SessionVerify(c *fiber.Ctx)error{
	user := c.Cookies("name")
	if len(user) > 0{
		return nil
	}
	return fiber.ErrLocked
}