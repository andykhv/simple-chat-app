package main

import (
    "fmt"

    static "github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
)

func main() {
	r := gin.Default()
	m := melody.New()

	r.Use(static.Serve("/", static.LocalFile("../client", true)))
	r.GET("/ws", func(c *gin.Context) {
        fmt.Println("hello world")
		m.HandleRequest(c.Writer, c.Request)
	})

	m.HandleMessage(func(s *melody.Session, msg []byte) {
        fmt.Println("hello world 2")
		m.Broadcast(msg)
	})

	r.Run(":5000")
}

