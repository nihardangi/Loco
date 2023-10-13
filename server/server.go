package server

import (
	"fmt"
	"log"
	"net/http"

	"Loco/config"
	"Loco/db/postgres"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

/*
newHTTPServer type function
	initiates new http server with GIN framework
	sets up all required middleware.
*/
func newHTTPServer() {
	router := gin.Default()
	srv := &http.Server{
		Addr:    config.Env().HTTPPort,
		Handler: router,
	}

	// ping server
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	db := postgres.GetConnection()
	sqlDB, err := db.DB()
	if err != nil {

	}
	defer sqlDB.Close()

	// Initialize the routes
	InitializeRoutes(db, router)

	// Run service server
	log.Println(fmt.Sprintf("status: initiating application on %s...\n", config.Env().HTTPPort))
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal("error in starting server", err)
	}
}

/*
Start type function
	runs new http server
*/
func Start() {
	newHTTPServer()
}
