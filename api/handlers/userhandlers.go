package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"xxx.ru/ds-dit/api/forms"
	"xxx.ru/ds-dit/api/models"
	"xxx.ru/ds-dit/api/storage"
)

var sessionName string
var sessionTimeout int

func init() {
	sessionName = os.Getenv("SESSION_NAME")
	sessionTimeout, _ = strconv.Atoi(os.Getenv("SESSION_TIMEOUT"))
}

// Login logs in user & set session.
//
// route : api/v1/users/login
//
// access: public
func Login(db storage.Services) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get("user")
		if user != nil {
			c.JSON(http.StatusForbidden, gin.H{"message": "Forbidden. You have already loged in."})
			return
		}

		var loginForm forms.LoginForm
		if c.ShouldBindJSON(&loginForm) != nil {
			log.Printf("%+v", loginForm)
			c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
			return
		}

		// TODO: implement AD (LDAPS) user authentication
		if loginForm.Name == "admin" && loginForm.Pass == "123456" {
			user, _ := json.Marshal(
				models.SessionUser{
					Name:      loginForm.Name,
					LoginTime: time.Now().UTC().Local().String(),
				},
			)
			session.Set("user", string(user))
			session.Options(sessions.Options{Path: "/", MaxAge: sessionTimeout})
			session.Save()
			c.JSON(http.StatusOK, gin.H{"message": "Logged in"})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Login failed"})
		}
	}
}

// Logout clears user session / delete cookies
//
// route : api/v1/users/logout
//
// access: private
func Logout() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		session.Set("user", "")
		session.Clear()
		session.Options(sessions.Options{MaxAge: -1, Path: "/"})
		session.Save()
		c.SetCookie(sessionName, "", -1, "/", "", false, false)
		c.JSON(http.StatusOK, gin.H{"message": "You have logged out successfuly"})
	}
}
