// utilities == lib for reusable code
package utils

import (
	"crypto/rand"
	"encoding/base32"
	"fmt"
	"net/http"
	"time"
	"github.com/gorilla/sessions"
)

type Cookie struct{
Name *http.Cookie
Value *http.Cookie
RawExpires *http.Cookie
MaxAge *http.Cookie

}

type Cookie interface{
	GenerateCookie()//returns string

}

//generate a cookie hash/token
func(c *Cookie)GenerateToken(length int)string{
	randomBytes := make([]byte,32)
	if _, err := rand.Read(randomBytes); err != nil {
		fmt.Println(err)
		panic(err)
	}
	stamp :=  base32.StdEncoding.EncodeToString(randomBytes)[:length]
	return  stamp
}
// Configure jwt cookie
func CreateSessionCookie(cfg *config.Config, session string) *http.Cookie {
	return &http.Cookie{
		Name:  cfg.Session.Name,
		Value: session,
		Path:  "/",
		// Domain: "/",
		Expires:    time.Now().Add(1 * time.Minute),
		RawExpires: "",
		MaxAge:     cfg.Session.Expire,
		HttpOnly:   cfg.Cookie.HTTPOnly,
		SameSite:   0,
	}

	// Delete session
func DeleteSessionCookie(c echo.Context, sessionName string) {
	c.SetCookie(&http.Cookie{
		Name:   sessionName,
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	})
}
}
