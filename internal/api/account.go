package api

import (
	"bitGoin/internal/util"
)

// 계좌 정보 조회
func GetAccounts(jwt string) []byte {
	util.Logger.Info("계좌 정보 조회 시작")
	response := CallAPI("/accounts", jwt)
	util.Logger.Info("계좌 정보 조회 성공")
	util.Logger.Debug("응답 데이터:", string(response))
	return response
}
