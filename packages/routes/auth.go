package routes

import (
	"github.com/gofiber/fiber/v2"
)


func Authorized(c *fiber.Ctx){
		  c.Status(201).JSON(fiber.Map{
			"message": "authorized",
			"authenticated": true,
		})
   return
}

func Unauthorized(c *fiber.Ctx){
	c.Status(401).JSON(fiber.Map{
				"message": "unauthorized, create a session",
				"authenticated": false,
			})
				return
}