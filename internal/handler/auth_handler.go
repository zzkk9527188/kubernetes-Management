package handler

import (
	"cm_platform/internal/model"
	"cm_platform/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type AuthHandler struct {
	authService *service.UserService
}

func NewAuthHandler(authService *service.UserService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

// RegisterAPI 用户注册
func (h *AuthHandler) RegisterAPI(c *gin.Context) {
	var req model.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "无效的请求参数",
		})
		return
	}
	//获取上下文
	ctx := c.Request.Context()
	// 调用 UserService.Register 注册用户
	err := h.authService.Register(ctx, req.Username, req.Password, req.Email)
	if err != nil {
		// 检查错误类型并返回适当的响应
		switch err.Error() {
		case "用户名已存在":
			c.JSON(http.StatusConflict, gin.H{"error": "用户名已存在"}) // 409 Conflict
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "服务器内部错误"})
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "注册成功"})
}

// LoginAPI 用户登录
func (h *AuthHandler) LoginAPI(c *gin.Context) {
	var req model.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := c.Request.Context()
	// 调用 UserService 登录逻辑Login
	loginToken, err := h.authService.Login(ctx, req.Username, req.Password)
	if err != nil {
		switch {
		case strings.Contains(err.Error(), "用户名或密码错误"):
			c.JSON(http.StatusUnauthorized, gin.H{"error": "用户名或密码错误"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "服务器内部错误"})
		}
		return
	}
	// 返回成功响应
	c.JSON(http.StatusOK, gin.H{"message": "登录成功", "token": loginToken})
}

// GetAllUsersAPI 获取用户信息
func (h *AuthHandler) GetAllUsersAPI(c *gin.Context) {
	// 获取上下文
	ctx := c.Request.Context()
	// 调用服务层方法获取所有用户
	users, err := h.authService.GetAllUsers(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "服务器内部错误"})
	}
	// 返回用户列表
	c.JSON(http.StatusOK, gin.H{"users": users})
}
