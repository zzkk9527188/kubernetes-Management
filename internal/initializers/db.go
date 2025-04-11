package initializers

import (
	"cm_platform/internal/config"
	"cm_platform/internal/model"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

// InitMySQL 初始化 MySQL 数据库连接

func InitDB(cfg *config.Config) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Databases.Mysql.User,
		cfg.Databases.Mysql.Password,
		cfg.Databases.Mysql.Host,
		cfg.Databases.Mysql.Port,
		cfg.Databases.Mysql.Database)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to MySQL: %v", err)
	}
	err = db.AutoMigrate(&model.User{})
	if err != nil {
		fmt.Errorf("mysql 连接失败")
	}
	return db
}
