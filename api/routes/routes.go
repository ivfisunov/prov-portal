package routes

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"xxx.ru/ds-dit/api/storage"

	"xxx.ru/ds-dit/api/handlers"
)

// RegisterRoutes registers User routes
func RegisterRoutes(route *gin.RouterGroup, db storage.Services) {

	usersGroup := route.Group("/users")
	{
		usersGroup.POST("/login", handlers.Login(db))
		usersGroup.POST("/logout", isAuthenticated(), handlers.Logout())
	}

}

// Middleware checks if the user is logged in
func isAuthenticated() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get("user")
		if user == nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			c.Abort()
			return
		}
		// c.Next()
	}

}
