package main

import (
	"bitGoin/internal/util"
	"fmt"
)

func main() {
	// 로깅 초기화
	util.InitLogger()
	// Redis 초기화
	util.InitRedis()

		val, err := util.RedisClient.Get(util.Ctx, "test-key").Result()
	if err != nil {panic(err)}
	fmt.Println("Redis 값:", val)
	
	util.Logger.Info(val)
}
