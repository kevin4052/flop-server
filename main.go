package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	manager := NewManager()

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowHeaders: []string{"Origin"},
		MaxAge:       12 * time.Hour,
	}))

	router.GET("/ws", manager.serveWS)

	router.Run(":8080")
}
