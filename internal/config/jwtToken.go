package config

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"log"
	"time"
)

var (
	JWTSecret = []byte("123456")
	JWTExpire = 24 * time.Hour
)

func GenerateToken(id uint, username string) (string, error) {
	// 定义 Claims（JWT 载荷）
	claims := jwt.MapClaims{
		"id":       id,                               // 用户 ID
		"username": username,                         // 用户名
		"exp":      time.Now().Add(JWTExpire).Unix(), // 过期时间
	}
	//生成Token
	signedToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := signedToken.SignedString(JWTSecret)
	if err != nil {
		log.Println("签名失败", err)
		return "", err
	}
	return token, nil
}

func VerifyToken(tokenString string) (map[string]interface{}, error) {
	//解析token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 检查签名方法是否为 HMAC
		log.Println("检查签名方法是否为 HMAC")
		if hmac, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			log.Println("签名是HMAC", hmac)
			return nil, errors.New("无效的签名方法")
		}
		// 返回密钥
		return JWTSecret, nil
	})

	//检查解析过程中是否有错误
	if err != nil {
		log.Println("解析token异常: ", err)
		return nil, err
	}

	// 检查 Token 是否有效
	if !token.Valid {
		log.Println("token无效")
		return nil, errors.New("无效 Token")
	}
	// 提取 Claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		log.Println("claims is: ", claims)
		return claims, nil
	}
	log.Println("token无效")

	// 如果 Claims 类型不匹配，返回错误
	return nil, errors.New("无效 Token")
}
