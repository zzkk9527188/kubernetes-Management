package interfaces

import (
	"cm_platform/internal/model"
	"context"
)

// UserService 定义用户服务的接口
type UserService interface {
	Register(ctx context.Context, username, password string) error
	Login(ctx context.Context, username, password string) (string, error)
	GetProfile(ctx context.Context, userID string) (map[string]interface{}, error)
}

// UserRepository 定义用户数据存储的接口
type UserRepository interface {
	CreateUser(ctx context.Context, username, password, email string) error
	GetUserByUsername(ctx context.Context, username string) (map[string]interface{}, error)
	GetUserByID(ctx context.Context, userID uint) (map[string]interface{}, error)
	QueryAllUsers(ctx context.Context) ([]model.User, error)
}
