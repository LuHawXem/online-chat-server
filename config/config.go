package config

import (
	"github.com/spf13/viper"
	"log"
)

type Configure struct {
	Db struct {
		EngineName string
		UserName string
		PassWord string
		Host string
		Port string
		DbName string
		Charset string
	}
	
	Redis struct {
		PassWord string
		Host string
		Port string
	}
	
	Secret struct {
		PrivateKey string
		PublicKey string
		TokenKey string
	}
}

var Conf = &Configure{}

func init()  {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("read config failed: %v \n", err)
		return
	}

	err = viper.Unmarshal(Conf)
	if err != nil {
		log.Printf("unmarshal config failed: %v \n", err)
	}
}