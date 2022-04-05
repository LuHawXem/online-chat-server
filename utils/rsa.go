package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"online-chat-server/config"
)

func encryptOAEP(publicKey []byte, text string) ([]byte, error) {
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("public key error")
	}

	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	pubKey := pubInterface.(*rsa.PublicKey)
	r := rand.Reader
	cipherData, err := rsa.EncryptOAEP(sha256.New(), r, pubKey, []byte(text), nil)
	if err != nil {
		return nil, err
	}

	ciphertext := base64.StdEncoding.EncodeToString(cipherData)
	return []byte(ciphertext), nil
}

func EncryptOAEP(text string) (data []byte, err error) {
	publicKey := config.Conf.Secret.PublicKey
	return encryptOAEP([]byte(publicKey), text)
}

func decryptOAEP(privateKey []byte, ciphertext string) ([]byte, error) {
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key error")
	}

	priInterface, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	priKey := priInterface.(*rsa.PrivateKey)
	r := rand.Reader
	cipherData, _ := base64.StdEncoding.DecodeString(ciphertext)
	plaintext, err := rsa.DecryptOAEP(sha256.New(), r, priKey, cipherData, nil)
	if err != nil {
		return nil, err
	}
	return plaintext, nil
}

func DecryptOAEP(ciphertext string) ([]byte, error) {
	privateKey := config.Conf.Secret.PrivateKey
	return decryptOAEP([]byte(privateKey), ciphertext)
}