package config

import (
	"fmt"
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

func TestViperLoadConfig(t *testing.T) {
	configFile := "E:\\StudyCode\\go_code\\cm_platform\\cmd\\configPath\\cm_platform.yaml"
	config, err := ViperLoadConfig(configFile)
	if err != nil {
		t.Error("err: ", err)
	}

	f := config.KubeVisionary.KubeConfig
	fmt.Println(f)
}
