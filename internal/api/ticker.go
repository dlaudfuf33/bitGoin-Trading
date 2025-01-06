package api

import (
	"bitGoin/internal/util"
	"strings"
)

// 특정 코인의 시세 조회
func GetTickerInfo(jwt string, markets []string) []byte {
	util.Logger.Info("시세 조회 시작")

	// markets를 콤마로 연결
	marketParam := strings.Join(markets, ",")
	response := CallAPI("/ticker?markets="+marketParam, jwt)

	util.Logger.Info("시세 조회 성공")
	util.Logger.Debug("응답 데이터:", string(response))
	return response
}
