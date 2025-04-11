package config

import (
	"log"
	"testing"
)

func TestGenerateToken(t *testing.T) {
	userID := uint(1)
	username := "testuser"
	token, err := GenerateToken(userID, username)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("生成的token: ", token)

}

func TestVerifyToken(t *testing.T) {
	userID := uint(1)
	username := "testuser"
	token, err := GenerateToken(userID, username)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("生成的token: ", token)

	verifyToken, err := VerifyToken(token)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(verifyToken)

}
