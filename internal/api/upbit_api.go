package api

import (
	"bitGoin/config"
	"bitGoin/internal/util"
	"io"
	"net/http"
)

// 공통 API 호출 함수
func CallAPI(endpoint, jwtToken string) []byte {
	client := &http.Client{}
	req, err := http.NewRequest("GET", config.APIURL+endpoint, nil)
	if err != nil {
		util.Logger.WithError(err).Error("API 요청 생성 실패")
		panic(err)
	}
	if jwtToken != "" {
		req.Header.Set("Authorization", "Bearer "+jwtToken)
	}

	resp, err := client.Do(req)
	if err != nil {
		util.Logger.WithError(err).Error("API 호출 실패")
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		util.Logger.WithError(err).Error("응답 본문 읽기 실패")
		panic(err)
	}

	return body
}
