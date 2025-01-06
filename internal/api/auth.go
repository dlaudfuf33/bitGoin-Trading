package api

import (
	"bitGoin/internal/util"
	"net/http"
)

func HandleAuthRequest(w http.ResponseWriter, r *http.Request) {
	// JWT 검증
	token := r.Header.Get("Authorization")[7:] // Bearer 제거
	userID, err := util.ValidateServiceJWT(token, "your_service_jwt_secret")
	if err != nil {
		http.Error(w, "JWT 검증 실패: "+err.Error(), http.StatusUnauthorized)
		return
	}

	// 유저 정보 반환
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("유저 ID: " + userID))
}