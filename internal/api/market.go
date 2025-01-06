package api

import (
	"bitGoin/internal/util"
)

// 시장 정보 조회
func GetMarketInfo() []byte {
	util.Logger.Info("시장 정보 조회 시작")
	response := CallAPI("/market/all", "")
	util.Logger.Info("시장 정보 조회 성공")
	util.Logger.Debug("응답 데이터:", string(response))
	return response
}
