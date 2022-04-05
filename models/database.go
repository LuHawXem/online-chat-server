package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"online-chat-server/config"
	"time"
)

var db *gorm.DB

func GetDB() *gorm.DB {
	return db
}

func init()  {
	d := config.Conf.Db
	var dbArgs = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		d.UserName, d.PassWord, d.Host, d.Port, d.DbName, d.Charset)

	var err error
	db, err = gorm.Open(d.EngineName, dbArgs)
	if err != nil {
		log.Fatalf("open database failed: %v \n", err)
		return
	}

	log.Printf("database: %s@%s:%s", d.DbName, d.Host, d.Port)

	db.DB().SetConnMaxLifetime(1 * time.Hour)
	db.SingularTable(true)

	err = db.AutoMigrate(&TUser{}).Error
	if err != nil {
		log.Printf("auto migrate failed: %v \n", err)
		return
	}
}