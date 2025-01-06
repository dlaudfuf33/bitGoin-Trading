package util

import (
	"context"

	"github.com/go-redis/redis/v8"
)

var RedisClient *redis.Client
var Ctx = context.Background()

// Redis 초기화
func InitRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis 주소
		Password: "", // 비밀번호
		DB:       0,  // Redis DB 인덱스
	})
	// 연결 테스트
	_, err := RedisClient.Ping(Ctx).Result()
	if err != nil {
		panic("Redis 연결 실패: " + err.Error())
	}
}
