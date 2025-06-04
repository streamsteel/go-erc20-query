package config

import "os"

// Config 应用配置结构
type Config struct {
	EthRPCURL  string
	Port       string
	GinMode    string
	PrivateKey string
}

// Load 加载配置
func Load() *Config {
	return &Config{
		EthRPCURL:  getEnv("ETH_RPC_URL", "https://sepolia.infura.io/v3/YOUR_PROJECT_ID"),
		Port:       getEnv("PORT", "8080"),
		GinMode:    getEnv("GIN_MODE", "debug"),
		PrivateKey: getEnv("PRIVATE_KEY", ""),
	}
}

// getEnv 获取环境变量，如果不存在则返回默认值
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
