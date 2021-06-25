package main

import (
	api "drouotBack/API"
	"drouotBack/models"
	websockets "drouotBack/webSockets"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// load credentials from .env file
	/*if err := godotenv.Load("./.env"); err != nil {
		log.Fatalf("Error loading .env file")
	}*/

	// Open connection to database
	models.ConnectDataBase()
	//close database connection after server close
	defer models.DB.Close()

	// create a gin engine
	r := gin.Default()

	// initialize sessions
	key := []byte(os.Getenv("SESSION_ID"))
	store := cookie.NewStore(key)
	store.Options(sessions.Options{MaxAge: 0}) // expire on close
	r.Use(sessions.Sessions("localSession", store))

	// Create endpoints
	api.InitializeRoutes(r)

	r.GET("/ws/v1", websockets.OpenSocket)

	// Run server on default port (8080)
	r.Run(":8080")
}
