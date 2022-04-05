package router

import (
	"github.com/gin-gonic/gin"
	"online-chat-server/service"
)

func init()  {
	r := gin.Default()
	r.Use(service.Cors())
	r.POST("/register", service.BeforeRegister(), service.Register)
	r.POST("/login", service.BeforeLogin(), service.Login)
	r.POST("/logout", service.AuthToken(), service.Logout)
	r.Any("/ws", service.AuthToken(), service.Websocket)
	r.Run(":9000")
}