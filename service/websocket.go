package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"online-chat-server/models"
	"strconv"
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
			goto exit
		}

		switch m.Type {
		case 0:
			if m.ReplyID.Valid && m.Operate.Valid {
				var r models.TMessage
				models.GetDB().Where("id = ?", m.ReplyID).First(&r)
				models.GetDB().Model(&r).Update("state", 2)
				if m.Operate.Int64 == 1 {
					models.GetDB().Create(&models.TRelation{
						UserID:    r.CreatedBy,
						CreatedBy: id,
						State:     0,
					})
					models.GetDB().Create(&models.TRelation{
						UserID:    id,
						CreatedBy: r.CreatedBy,
						State:     0,
					})
				}
			} else {
				message := models.TMessage{
					Receiver:  m.Receiver,
					ReplyID:   m.ReplyID,
					Operate:   m.Operate,
					Content:   m.Content,
					Type:      m.Type,
					State:     0,
					CreatedBy: id,
				}
				models.GetDB().Create(&message)
			}
			break
		case 1:
			var r models.TMessage
			models.GetDB().Where("receiver = ? and type = ? and created_by = ?", m.Receiver, m.Type, id).First(&r)
			if r.ID != 0 {
				ws.WriteJSON(gin.H{
					"msg": "repeat request",
					"status": http.StatusBadRequest,
				})
				break
			}
			message := models.TMessage{
				Receiver:  m.Receiver,
				Type:      m.Type,
				CreatedBy: id,
			}
			models.GetDB().Create(&message)
			ws.WriteJSON(gin.H{
				"status": http.StatusOK,
			})
			break
		case 2:
			break
		case 3:
			break
		case 4:
			var user models.User
			num, _ := strconv.Atoi(m.Content.String)
			models.GetDB().Table("t_user").Where("account like ?", uint64(num)).First(&user)
			if user.ID != 0 {
				err = ws.WriteJSON(user)
				if err != nil {
					log.Printf("chain %s WriteJSON error: %v\n", key, err)
					ws.Close()
					delete(conn, key)
					goto exit
				}
			} else {
				ws.WriteJSON(gin.H{
					"msg": "user not found",
					"status": http.StatusNoContent,
				})
			}
			break
		case 5:
			var friends []models.TRelation
			models.GetDB().Where("created_by = ?", id).Find(&friends)
			var data []gin.H
			for _, v := range friends {
				var u models.User
				models.GetDB().Table("t_user").Where("id = ?", v.UserID).First(&u)
				data = append(data, gin.H{
					"user": u,
					"relation": v,
				})
			}
			ws.WriteJSON(gin.H{
				"type": 5,
				"friends": data,
				"status": http.StatusOK,
			})
			break
		case 200:
			ws.WriteJSON(gin.H{
				"status": http.StatusOK,
			})
			break
		}
	}
	exit:
}

func WriteMessage(ws *websocket.Conn, id uint32, key string) {
	for {
		var messages []models.TMessage
		models.GetDB().Where("receiver = ? AND state = 0", id).Find(&messages)
		if len(messages) != 0 {
			for _, message := range messages {
				var err error
				if message.Type == 1 {
					var user models.User
					models.GetDB().Table("t_user").Where("id = ?", message.CreatedBy).First(&user)
					err = ws.WriteJSON(gin.H{
						"id": message.ID,
						"type": message.Type,
						"created_by": message.CreatedBy,
						"content": message.Content,
						"nickname": user.Nickname,
						"avatar": user.Avatar,
					})
				} else {
					err = ws.WriteJSON(message)
				}
				if err != nil {
					log.Printf("chain %s ReadJSON error: %v\n", key, err)
					ws.Close()
					delete(conn, key)
					goto exit
				}
				models.GetDB().Model(&message).Update("state", 1) // 值类型要对上,Update("state", "1")不会生效
			}
		}
	}
	exit:
}