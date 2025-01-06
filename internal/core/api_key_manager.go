package core

import (
	"bitGoin/internal/util"
	"encoding/json"
)

// Redis에서 API 키 조회
func GetAPIKey(userID string) (map[string]string, error) {
	// Redis에서 키 조회
	data, err := util.RedisClient.Get(util.Ctx, "user:"+userID+":api_key").Result()
	if err != nil {
		util.Logger.Errorf("Redis에서 API 키 조회 실패 (userID: %s): %v", userID, err)
		return nil, err
	}

	// JSON 파싱
	var apiKey map[string]string
	err = json.Unmarshal([]byte(data), &apiKey)
	if err != nil {
		util.Logger.Errorf("API 키 JSON 파싱 실패 (userID: %s): %v", userID, err)
		return nil, err
	}

	util.Logger.Infof("Redis에서 API 키 조회 성공 (userID: %s)", userID)
	return apiKey, nil
}
