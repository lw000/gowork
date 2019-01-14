// melody_test project main.go
package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/olahol/melody"
)

func main() {
	engine := gin.Default()

	m := melody.New()

	engine.GET("/", func(c *gin.Context) {
		http.ServeFile(c.Writer, c.Request, "templates/index.html")
	})

	//websocket
	engine.GET("/ws", func(c *gin.Context) {
		m.HandleRequest(c.Writer, c.Request)
	})

	m.HandleMessage(func(s *melody.Session, msg []byte) {
		log.Println(string(msg))
		m.Broadcast(msg)
	})

	m.HandleConnect(func(s *melody.Session) {

	})

	m.HandleDisconnect(func(s *melody.Session) {

	})

	m.HandleClose(func(s *melody.Session, code int, msg string) error {

		return nil
	})

	engine.Run()
}
