//入口程序

package main

import (
	"cm_platform/httpServer"
	"cm_platform/internal/config"
	"cm_platform/internal/initializers"
	"cm_platform/internal/repository"
	"cm_platform/internal/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	configPath := "/configPath/cm_platform.yaml"
	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// 初始化数据库
	db := initializers.InitDB(cfg)

	//初始化仓库和服务层
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)

	//创建路由
	r := gin.Default()
	r.Use(httpServer.CORSMiddleware())
	go httpServer.SetupRoutes(r, userService)
	err = r.Run(fmt.Sprintf("%s:%d", cfg.KubeVisionary.CmPlatform.Host, cfg.KubeVisionary.CmPlatform.Port))
	if err != nil {
		panic(err)
	}
}
