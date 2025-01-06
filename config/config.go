package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// 환경 변수 초기화
func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Warning: .env file not found !!, using system environment variables")
	}
}

// 환경 변수 로드 함수
func GetEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// API URL 설정
var APIURL = GetEnv("UPBIT_API_URL", "https://api.upbit.com/v1")
