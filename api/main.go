package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"xxx.ru/ds-dit/api/routes"
	"xxx.ru/ds-dit/api/storage/postgres"
	"xxx.ru/ds-dit/api/storage/redis"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {

	// initialize Postgres connection
	db, err := postgres.New()
	if err != nil {
		log.Fatalf("Error connecting to Postgres...\n%v", err)
	}
	defer db.Close()
	log.Printf("Connected to Postgres successfuly!\n\n")

	if os.Getenv("ENV") == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	// setup router
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(redis.InitRedisSession())
	// set root api url (version 1)
	apiV1 := router.Group("/api/v1")
	routes.RegisterRoutes(apiV1, db)

	// setup & run server
	appPort := os.Getenv("APP_PORT")
	server := &http.Server{
		Addr:         fmt.Sprintf(":%s", appPort),
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}
	log.Printf("Listening and serving on port: %s", appPort)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
