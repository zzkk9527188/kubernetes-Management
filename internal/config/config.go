package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	//数据库默认mysql
	Databases struct {
		Mysql struct {
			Enabled  bool   `mapstructure:"enabled"`
			Host     string `mapstructure:"host"`
			Port     int    `mapstructure:"port"`
			User     string `mapstructure:"user"`
			Password string `mapstructure:"password"`
			Database string `mapstructure:"database"`
		} `mapstructure:"mysql"`
	} `mapstructure:"databases"`
	//mq
	Rabbitmq struct {
		Enabled     bool   `mapstructure:"enabled"`
		Host        string `mapstructure:"host"`
		Port        int    `mapstructure:"port"`
		Username    string `mapstructure:"username"`
		Password    string `mapstructure:"password"`
		VirtualHost string `mapstructure:"virtualHost"`
		Exchange    string `mapstructure:"exchange"`
		Queue       string `mapstructure:"queue"`
		RoutingKey  string `mapstructure:"routingKey"`
		ConsumerTag string `mapstructure:"consumerTag"`
	} `mapstructure:"rabbitmq"`
	//redis默认单点
	Redis struct {
		Standalone struct {
			Enabled  bool   `mapstructure:"enabled"`
			Host     string `mapstructure:"host"`
			Port     int    `mapstructure:"port"`
			Password string `mapstructure:"password"`
		} `mapstructure:"standalone"`
	} `mapstructure:"redis"`
	//平台配置
	KubeVisionary struct {
		CmPlatform struct {
			Port int    `mapstructure:"port"`
			Host string `mapstructure:"host"`
		} `mapstructure:"cm_platform"`
		KubeConfig string `mapstructure:"kubeConfig"`
	} `mapstructure:"kube-visionary"`
}

// LoadConfig 使用 Viper 加载配置文件并返回 Config 结构体
func ViperLoadConfig(configFile string) (*Config, error) {
	v := viper.New()
	v.SetConfigFile(configFile)
	v.SetConfigType("yaml")

	// 读取配置文件
	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}
	// 将配置映射到结构体
	var cfg Config
	err := v.Unmarshal(&cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}
	return &cfg, nil
}
