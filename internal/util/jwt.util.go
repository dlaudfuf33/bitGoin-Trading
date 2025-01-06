package util

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// JWT 검증 함수
func ValidateServiceJWT(tokenString, secret string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 서명 방법 검증
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return "", err
	}

	// 클레임에서 유저 ID 추출
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID := claims["sub"].(string) // JWT의 `sub` 클레임을 유저 ID로 사용
		exp := int64(claims["exp"].(float64)) // 만료 시간

		// 만료 확인
		if time.Now().Unix() > exp {
			return "", fmt.Errorf("JWT가 만료되었습니다")
		}

		return userID, nil
	}
	return "", fmt.Errorf("유효하지 않은 JWT")
}
