package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type DbConfig struct {
	Driver          string `yaml:"driver"`
	DSN             string `yaml:"dsn"`
	MaxOpenConns    int    `yaml:"max_open_conns"`
	MaxIdleConns    int    `yaml:"max_idle_conns"`    // 对应 database.max_idle_conns
	ConnMaxLifetime int    `yaml:"conn_max_lifetime"` // 对应 database.conn_max_lifetime
}

// AppConfig 其他业务配置结构体
type AppConfig struct {
	Name     string `yaml:"name"`      // 对应 app.name
	LogLevel string `yaml:"log_level"` // 对应 app.log_level
	Timeout  int    `yaml:"timeout"`   // 对应 app.timeout
	Port     int    `yaml:"port"`      // 对应 app.port
}

// Config 总配置结构体（包含所有配置）
type Config struct {
	Database DbConfig  `yaml:"database"` // 对应配置文件的 database 节点
	App      AppConfig `yaml:"app"`      // 对应配置文件的 app 节点
}

// 2. 读取配置文件的工具函数
func LoadConfig(filePath string) (*Config, error) {
	// 读取配置文件内容
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("读取配置文件失败: %w", err)
	}

	// 解析 YAML 到结构体
	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, fmt.Errorf("解析配置文件失败: %w", err)
	}

	return &config, nil
}
