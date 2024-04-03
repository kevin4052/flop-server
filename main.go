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

	// router.GET("/ws", func(c *gin.Context) {
	// 	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	// 	if err != nil {
	// 		return
	// 	}
	// 	defer conn.Close()
	// 	for {
	// 		conn.WriteMessage(websocket.TextMessage, []byte("Hello, WebSocket!"))
	// 		time.Sleep(time.Second)
	// 	}
	// })

	router.Run(":8080")
}
