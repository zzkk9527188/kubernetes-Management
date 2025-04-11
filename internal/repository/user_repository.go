package repository

import (
	"cm_platform/internal/model"
	"context"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

//

// CreateUser创建用户
func (r *UserRepository) CreateUser(ctx context.Context, username, password, email string) error {
	log.Println("创建用户: ", username)
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("密码加密失败")
		return err
	}
	user := model.User{
		Username: username,
		Password: string(hashPassword),
		Email:    email,
	}
	result := r.db.WithContext(ctx).Create(&user)
	return result.Error
}

// GetUserByUsername 查询用户
func (r *UserRepository) GetUserByUsername(ctx context.Context, username string) (map[string]interface{}, error) {
	var user model.User
	err := r.db.WithContext(ctx).Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return map[string]interface{}{
		"id":       user.ID,
		"username": user.Username,
		"email":    user.Email,
		"password": user.Password,
		"createAt": user.CreateAt,
		"updateAt": user.UpdateAt,
	}, nil
}

// GetUserByID 根据ID查询用户
func (r *UserRepository) GetUserByID(ctx context.Context, userID uint) (map[string]interface{}, error) {
	var user model.User
	err := r.db.WithContext(ctx).Where("id=?", userID).First(&user).Error
	if err != nil {
		return nil, err
	}
	return map[string]interface{}{
		"id":       user.ID,
		"username": user.Username,
		"email":    user.Email,
		"password": user.Password,
		"createAt": user.CreateAt,
		"updateAt": user.UpdateAt,
	}, nil
}

// QueryAllUsers 查询所有用户，用户web默认展示
func (r *UserRepository) QueryAllUsers(ctx context.Context) ([]model.User, error) {
	var users []model.User
	err := r.db.WithContext(ctx).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}
