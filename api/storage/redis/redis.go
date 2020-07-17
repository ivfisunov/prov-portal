package redis

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

// InitRedisSession inits redis store for managing sessions
func InitRedisSession() gin.HandlerFunc {
	var (
		host           = os.Getenv("REDIS_HOST")
		port           = os.Getenv("REDIS_PORT")
		password       = os.Getenv("REDIS_PASSWORD")
		sessionSecret  = os.Getenv("SESSION_SECRET")
		sessionName    = os.Getenv("SESSION_NAME")
		sessionTimeout = os.Getenv("SESSION_TIMEOUT")
	)

	if runtime.GOOS != "linux" {
		host = "host.docker.internal"
	}

	store, err := redis.NewStore(
		10,
		"tcp",
		fmt.Sprintf("%s:%s", host, port),
		password,
		[]byte(sessionSecret))
	if err != nil {
		log.Fatalf("Error connecting to Redis...\n%v", err)
	}
	log.Printf("Connected to Redis successfuly!\n\n")

	sTimeout, _ := strconv.Atoi(sessionTimeout)
	store.Options(sessions.Options{
		MaxAge: sTimeout,
	})

	return sessions.Sessions(sessionName, store)
}
