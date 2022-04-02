package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET,POST,OPTIONS,HEADER,PUT,PATCH,DELETE")
		c.Header("Access-Control-Allow-Headers", "Content-Type,Authorization")
		c.Header("Access-Control-Expose-Headers", "Access-Control-Allow-Origin,Access-Control-Allow-Headers,Content-Length,Content-Type")
		c.Header("Access-Control-Allow-Credentials","true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}

		c.Next()
	}
}

func BeforeLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")
		if username == "" || password == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": "params missing",
			})
			c.Abort()
		}
		c.Next()
	}
}
