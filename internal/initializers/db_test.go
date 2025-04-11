package initializers

import (
	"cm_platform/internal/config"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestInitMySQL(t *testing.T) {
	// 模拟配置
	cfg := &config.Config{
		Databases: struct {
			Mysql struct {
				Enabled  bool   `mapstructure:"enabled"`
				Host     string `mapstructure:"host"`
				Port     int    `mapstructure:"port"`
				User     string `mapstructure:"user"`
				Password string `mapstructure:"password"`
				Database string `mapstructure:"database"`
			} `mapstructure:"mysql"`
		}{
			Mysql: struct {
				Enabled  bool   `mapstructure:"enabled"`
				Host     string `mapstructure:"host"`
				Port     int    `mapstructure:"port"`
				User     string `mapstructure:"user"`
				Password string `mapstructure:"password"`
				Database string `mapstructure:"database"`
			}{
				Enabled:  true,
				Host:     "10.10.10.113",
				Port:     27878,
				User:     "root",
				Password: "zzkk.9527",
				Database: "test",
			},
		},
	}

	// 初始化 MySQL
	db := InitDB(cfg)

	// 确保数据库连接成功
	assert.NotNil(t, db)

	// 验证表是否创建成功
	var tableNames []string
	db.Raw("SHOW TABLES").Scan(&tableNames)
	expectedTables := []string{"users"}
	for _, table := range expectedTables {
		log.Println("Table:", table)
		assert.Contains(t, tableNames, table)
	}
}
