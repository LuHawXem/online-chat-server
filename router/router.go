package router

import (
	"github.com/gin-gonic/gin"
	"online-chat-server/service"
)

func init()  {
	r := gin.Default()
	r.Use(service.Cors())
	r.POST("/login", service.BeforeLogin(), service.Login)
}