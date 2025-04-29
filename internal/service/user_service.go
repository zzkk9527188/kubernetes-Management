package service

import (
	"cm_platform/interfaces"
	"cm_platform/internal/config"
	"cm_platform/internal/model"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

type UserService struct {
	repo        interfaces.UserRepository
	us          interfaces.UserService
	redisClient *redis.Client
}

func NewUserService(repo interfaces.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (u *UserService) Register(ctx context.Context, username, password, email string) error {
	existingUser, err := u.repo.GetUserByUsername(ctx, username)
	if err != nil {
		log.Println("用户名不存在,开始创建用户: ", err)
	}
	if existingUser != nil {
		return errors.New("用户名已存在")
	}
	err = u.repo.CreateUser(ctx, username, password, email)
	if err != nil {
		log.Println("用户创建失败", err)
	}
	return nil
}

// GetProfile 获取用户信息
func (u *UserService) GetProfile(ctx context.Context, userID uint) (map[string]interface{}, error) {
	// 定义 Redis 缓存的 Key
	cacheKey := fmt.Sprintf("userID: %d", userID)
	//// 尝试从 Redis 获取缓存
	cacheData, err := u.redisClient.Get(ctx, cacheKey).Result()
	if err == nil {
		// 如果缓存命中，解析缓存数据
		var user model.User
		if err := json.Unmarshal([]byte(cacheData), &user); err == nil {
			return map[string]interface{}{
				"id":       user.ID,
				"username": user.Username,
				"createAt": user.CreateAt,
				"updateAt": user.UpdateAt,
				"email":    user.Email,
				"password": user.Password,
			}, nil
		}
		log.Printf("Failed to unmarshal Redis cache data: %v", err)
	} else if err != redis.Nil {
		// 如果 Redis 出现错误（非缓存未命中）
		log.Printf("Redis error: %v", err)
	}
	// 如果 Redis 中没有缓存，从数据库查询
	user, err := u.repo.GetUserByID(ctx, userID)
	if err != nil {
		log.Printf("Failed to marshal user data for Redis: %v", err)
	}
	if user == nil {
		return nil, errors.New("用户不存在")
	}
	// 将查询结果写入 Redis 缓存
	userJson, err := json.Marshal(user)
	if err != nil {
		log.Printf("Failed to marshal user data for Redis: %v", err)
	} else {
		err := u.redisClient.Set(ctx, cacheKey, userJson, 5*time.Minute).Err()
		if err != nil {
			log.Printf("Failed to set Redis cache: %v", err)
		}
	}

	return user, nil
}

func (u *UserService) Login(ctx context.Context, username, password string) (string, error) {
	// 查询用户是否存在
	user, err := u.repo.GetUserByUsername(ctx, username)
	if err != nil {
		return "", errors.New("用户名不存在,请先注册")
	}
	if user == nil {
		return "", errors.New("用户名或密码错误")
	}
	//验证密码
	passwordHash := user["password"].(string)
	err = bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return "", errors.New("用户名或密码错误")
		}
		return "", errors.New("用户名或密码错误")
	}
	// 生成 JWT Token
	token, err := config.GenerateToken(user["id"].(uint), username)
	if err != nil {
		return "", errors.New("生成Token失败")
	}
	return token, nil
}

// GetAllUsers 获取所有用户
func (u *UserService) GetAllUsers(ctx context.Context) ([]map[string]interface{}, error) {
	// 调用仓库层获取所有用户
	users, err := u.repo.QueryAllUsers(ctx)
	if err != nil {
		return nil, err
	}

	// 将数据库中的用户数据转换为 map 格式
	var userMaps []map[string]interface{}
	for _, user := range users {
		userMap := map[string]interface{}{
			"id":       user.ID,
			"username": user.Username,
			"email":    user.Email,
			"password": user.Password,
			"createAt": user.CreateAt,
			"updateAt": user.UpdateAt,
		}
		userMaps = append(userMaps, userMap)
	}

	return userMaps, nil
}
