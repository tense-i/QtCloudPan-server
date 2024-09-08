package config

import (
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

// Config 结构体定义配置参数
type Config struct {
	ServerPort  string `yaml:"server_port"`
	DatabaseDSN string `yaml:"database_dsn"`
}

// AppConfig 全局配置变量
var AppConfig Config

// LoadConfig 加载配置文件
func LoadConfig() {
	// 打印当前工作目录
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get current working directory: %v", err)
	}
	log.Printf("Current working directory: %s", dir)

	configFile, err := os.Open("config/config.yaml")
	if err != nil {
		log.Fatalf("Failed to open config file: %v", err)
	}
	defer configFile.Close()

	decoder := yaml.NewDecoder(configFile)
	if err := decoder.Decode(&AppConfig); err != nil {
		log.Fatalf("Failed to parse config file: %v", err)
	}
}
