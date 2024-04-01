package main

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	router := gin.Default()
	router.GET("/ws", func(c *gin.Context) {
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			return
		}
		defer conn.Close()
		for {
			conn.WriteMessage(websocket.TextMessage, []byte("Hello, WebSocket!"))
			time.Sleep(time.Second)
		}
	})
	router.Run(":8080")
}
