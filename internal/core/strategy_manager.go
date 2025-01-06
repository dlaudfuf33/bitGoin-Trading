package core

import (
	"bitGoin/internal/util"
	"encoding/json"
)

// Redis에서 조건 데이터 조회
func GetStrategyFromRedis(userID string) (map[string]interface{}, error) {
	// Redis에서 조건 조회
	data, err := util.RedisClient.Get(util.Ctx, "user:"+userID+":strategy").Result()
	if err != nil {
		util.Logger.Errorf("Redis에서 조건 데이터 조회 실패 (userID: %s): %v", userID, err)
		return nil, err
	}

	// JSON 파싱
	var strategy map[string]interface{}
	err = json.Unmarshal([]byte(data), &strategy)
	if err != nil {
		util.Logger.Errorf("조건 데이터 JSON 파싱 실패 (userID: %s): %v", userID, err)
		return nil, err
	}

	util.Logger.Infof("Redis에서 조건 데이터 조회 성공 (userID: %s)", userID)
	return strategy, nil
}
