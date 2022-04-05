package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

var conn = make(map[string]*websocket.Conn)
var up = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func Websocket(c *gin.Context) {
	log.Printf("request websocket with method:%s", c.Request.Method)
	id, _ := c.Get("id")
	token := c.GetHeader("token")
	if c.Request.Method == "GET" && token == "" {
		token = c.Query("token")
	}

	ws, err := up.Upgrade(c.Writer, c.Request, c.Writer.Header())
	if err != nil {
		log.Println(err)
		return
	}

	key := fmt.Sprintf("chat_user:%v:%s", id, token)
	conn[key] = ws

	go func() {
		defer func() {
			ws.Close()
			log.Printf("chaining:%s disconnect\n", key)
			delete(conn, key)
		}()

		for {
			log.Println("websocket online")
			type js struct {
				Message string
				Code int
			}

			ms := &js{
				Message: "hello chat",
				Code: http.StatusOK,
			}
			err = ws.WriteJSON(ms)
			if err != nil {
				log.Println(err)
				return
			}

			time.Sleep(5 * time.Second)
		}
	}()
}