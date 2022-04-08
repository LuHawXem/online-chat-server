package main

import (
	"online-chat-server/models"
	_ "online-chat-server/router"
)

func main()  {
	defer models.GetDB().Close()
}