package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/guregu/null.v4"
	"log"
	"net/http"
	"online-chat-server/models"
	"online-chat-server/utils"
	"time"
)

func Register(c *gin.Context)  {
	nickname := c.PostForm("nickname")
	password, _ := c.Get("password")

	err := models.GetDB().Create(&models.TUser{
		Password:  password.(string),
		Nickname:  null.StringFrom(nickname),
	}).Error
	if err != nil {
		log.Printf("create user error: %v", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	var user models.TUser
	models.GetDB().Where("nickname = ? AND password = ?", nickname, password.(string)).First(&user)
	user.Account = uint64(1000000000 + user.ID)
	err = models.GetDB().Save(&user).Error
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"account": user.Account,
	})
}

func Login(c *gin.Context) {
	id, _ := c.Get("id")
	var user models.User
	err := models.GetDB().Table("t_user").Where("id = ?", id.(uint32)).First(&user).Error
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ip, _ := c.RemoteIP()
	token := utils.GenerateToken(id.(uint32), ip)
	fmt.Printf("user: %d get token: %s\n", id.(uint32), token)
	models.GetRedis().Set(token, id, 3 * time.Hour)
	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"avatar": user.Avatar,
		"nickname": user.Nickname,
	})
}

func Logout(c *gin.Context)  {
	token := c.GetHeader("token")
	err := models.GetRedis().Del(token).Err()
	if err != nil {
		log.Println(err)
	}

	c.JSON(http.StatusOK, gin.H{})
}

