package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"online-chat-server/models"
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
	id, _ := c.Get("id")
	log.Printf("user: %d connect\n", id.(uint32))
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

	go WriteMessage(ws, id.(uint32), key)
	go ReadMessage(ws, id.(uint32), key)
}

func ReadMessage(ws *websocket.Conn, id uint32, key string) {
	for {
		var m models.Message
		err := ws.ReadJSON(&m)
		if err != nil {
			log.Printf("chain %s ReadJSON error: %v\n", key, err)
			ws.Close()
			delete(conn, key)
			break
		}

		switch m.Type {
		case 0:
			break
		case 1:
			message := models.TMessage{
				Receiver:  m.Receiver,
				Content:   m.Content,
				Type:      m.Type,
				CreatedBy: id,
			}
			models.GetDB().Create(&message)
			break
		case 2:
			break
		case 3:
			break
		}
	}
}

func WriteMessage(ws *websocket.Conn, id uint32, key string) {
	for {
		var messages []models.TMessage
		models.GetDB().Where("receiver = ? AND state = 0", id).Find(&messages)
		if len(messages) != 0 {
			for _, message := range messages {
				err := ws.WriteJSON(message)
				if err != nil {
					log.Printf("chain %s ReadJSON error: %v\n", key, err)
					ws.Close()
					delete(conn, key)
					goto exit
				}
				models.GetDB().Model(&message).Update("state", "1")
			}
		}
	}
	exit:
}