package server

import (
	"fmt"
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/AJ-Brown-InTech/libre-api/config"

)

func Server(){
	config.Configuration()
	e := echo.New()
	e.GET("/", func(c echo.Context)error{
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
	fmt.Println(("server here."))
}