package utils

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"online-chat-server/config"
)

func EncryptHmac(key string, text []byte) string {
	h := hmac.New(sha1.New, []byte(key))
	h.Write(text)
	s := base64.StdEncoding.EncodeToString(h.Sum([]byte("")))
	return s
}

func EncryptData(text []byte) string {
	key := config.Conf.Secret.PrivateKey
	return EncryptHmac(key, text)
}
