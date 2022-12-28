// utilities == lib for reusable code
package utils

import (
	"crypto/rand"
	//"encoding/base32"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/securecookie"

	//"github.com/klauspost/compress/s2"
	"github.com/AJ-Brown-InTech/libre-api/config"
//	"github.com/AJ-Brown-InTech/libre-api/packages/utils"
	"github.com/labstack/echo/v4"
)

//Cookie interface  == just crud operations for cookies
type Cookies interface {
	NewCookieSession() *securecookie.SecureCookie
	CreateSessionCookie() *http.Cookie
	DeleteSessionCookie()
	
}

//generate a cookie instance
func NewCookieSession() *securecookie.SecureCookie{
	randomBytes := make([]byte,32)
	if _, err := rand.Read(randomBytes); err != nil {
		fmt.Println(err)
		panic(err)
	}
	var hashKey = []byte("user-session")
	securecookie.GenerateRandomKey(6)
	var secret = securecookie.New(hashKey, randomBytes)
	//stamp :=  base32.StdEncoding.EncodeToString(randomBytes)[:length] // makes a string
	//store := sessions.NewCookieStore(randomBytes)
	return secret
}

// Create a new session cookie 
func CreateSessionCookie(cfg *config.Config) *http.Cookie {
	//random cookie value 
	newCookieVal := securecookie.GenerateRandomKey(10)

	return &http.Cookie{
		Name:  cfg.Session.Name,
		Value: string(newCookieVal),
		Path:  "/",
		// Domain: "/",
		Expires:    time.Now().Add(1 * time.Minute),
		RawExpires: "",
		MaxAge:     cfg.Session.Expire,
		Secure:     false,
		HttpOnly:   true,
		SameSite:   0,
	}
}

// Delete session cookie
func DeleteSessionCookie(c echo.Context, sessionName string) {
	c.SetCookie(&http.Cookie{
		Name:   sessionName,
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	})
}

// Get user ip address
func GetIPAddress(c echo.Context) string {
	return c.Request().RemoteAddr
}

//read my fucking cookie
func ReadCookie(c *config.Config, w http.ResponseWriter, r *http.Request) bool {
	
	sessionKey := c.Session.Name //specific key we are looking for
	_, err := r.Cookie(string(sessionKey)) //specific cookie id for our api
	
	if err != nil {
		//l.Errorf("Cookie not available %v", ourCookie)
		return false
	} else {
		//l.Infof("Cookie is available %v", ourCookie)
		return true
	}
	
}


 
 