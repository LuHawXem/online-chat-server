package utils

import (
	"encoding/json"
	"log"
	"net"
	"online-chat-server/config"
	"time"
)

type token struct {
	UserID uint32 `json:"user_id"`
	From net.IP `json:"from"`
	IssueAt time.Time `json:"issue_at"`
}

func GenerateToken(userId uint32, ip net.IP) string {
	t := token{
		UserID:  userId,
		From:    ip,
		IssueAt: time.Now(),
	}
	b, err := json.Marshal(t)
	if err != nil {
		log.Printf("marshal token error: %v \n", err)
		return ""
	}

	return EncryptHmac(config.Conf.Secret.TokenKey, b)
}
