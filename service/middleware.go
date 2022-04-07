package service

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"online-chat-server/models"
	"online-chat-server/utils"
	"strconv"
	"time"
)

func Collect() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip, _ := c.RemoteIP()
		log.Printf("request %s from %s\n", c.FullPath(), ip)
		c.Next()
	}
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET,POST,OPTIONS,HEADER,PUT,PATCH,DELETE")
		c.Header("Access-Control-Allow-Headers", "Content-Type,Authorization,AccessToken,X-CSRF-Token,Token")
		c.Header("Access-Control-Expose-Headers", "Access-Control-Allow-Origin,Access-Control-Allow-Headers,Content-Length,Content-Type")
		c.Header("Access-Control-Allow-Credentials","true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}

		c.Next()
	}
}

func BeforeRegister() gin.HandlerFunc  {
	return func(c *gin.Context) {
		nickname := c.PostForm("nickname")
		password := c.PostForm("password")
		if nickname == "" || password == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"msg": "params missing",
			})
			return
		}

		p, err := utils.DecryptOAEP(password)
		if err != nil {
			log.Println(err)
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"msg": "password error",
			})
			return
		}

		cryptPassword := utils.EncryptData(p)
		c.Set("password", cryptPassword)
		c.Next()
	}
}

func BeforeLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		account := c.PostForm("account")
		password := c.PostForm("password")
		if account == "" || password == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"msg": "params missing",
			})
			return
		}

		p, err := utils.DecryptOAEP(password)
		if err != nil {
			log.Println(err)
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"msg": "password error",
			})
			return
		}

		cryptPassword := utils.EncryptData(p)
		var user models.TUser
		models.GetDB().Where("account = ?", account).First(&user)
		if user.Password != cryptPassword {
			log.Printf("account %s password %s do not match: %s", account, user.Password, cryptPassword)
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"msg": "password error",
			})
			return
		}

		c.Set("id", user.ID)
		c.Next()
	}
}

func AuthToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if c.Request.Method == "GET" && token == "" {
			token = c.Query("token")
		}
		log.Println(token)
		data := models.GetRedis().Get(token)
		if data.Err() != nil {
			log.Println(data.Err())
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		log.Printf("user:%s request:%s \n", data.Val(), c.FullPath())
		id, _ := strconv.Atoi(data.Val())
		c.Set("id", uint32(id))

		ttl := models.GetRedis().TTL(token)
		if c.FullPath() != "/logout" && ttl.Val() <= 30 * time.Minute {
			models.GetRedis().Expire(token, 3 * time.Hour)
		}

		c.Next()
	}
}
