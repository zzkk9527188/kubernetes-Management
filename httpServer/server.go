package httpServer

import (
	"cm_platform/internal/handler"
	"cm_platform/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRoutes(r *gin.Engine, userService *service.UserService) {
	api := r.Group("/api")
	{
		authHandler := handler.NewAuthHandler(userService)
		api.POST("/register", authHandler.RegisterAPI)
		api.POST("/login", authHandler.LoginAPI)
		api.GET("/users", authHandler.GetAllUsersAPI)

		v1 := api.Group("/v1/cluster")
		{
			v1.GET("/information", GetClusterInfoHandler)
			v1.GET("/nodes", GetNodeInfoHandler)
		}
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取请求的 Origin 头部
		origin := c.Request.Header.Get("Origin")
		// 动态设置允许的来源
		if origin != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		}
		// 设置其他响应头
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true") // 允许携带凭据
		c.Writer.Header().Set("Access-Control-Max-Age", "43200")
		// 处理预检请求
		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		//继续处理请求
		c.Next()
	}

}
