package api

import (
	"bitGoin/internal/util"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// JWT 생성 함수
func GenerateJWT(accessKey, secretKey string) string {
	util.Logger.Info("JWT 생성 시작")
	claims := jwt.MapClaims{
		"access_key": accessKey,
		"nonce":      uuid.New().String(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		util.Logger.WithError(err).Error("JWT 서명 실패")
		panic(err)
	}
	util.Logger.Info("JWT 생성 성공")
	return signedToken
}