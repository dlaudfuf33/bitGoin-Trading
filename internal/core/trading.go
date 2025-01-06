package core

import (
	"fmt"
	"time"
)

// 조건 감시 및 매매 실행
func MonitorUser(user *User) {
	fmt.Printf("[%s] 조건 감시 시작\n", user.ID)
// TODO : 조건 감시 및 매매 실행
	for {
		// API 호출 (예: 계좌 조회)		
		// 10초 간격으로 실행
		time.Sleep(10 * time.Second)
	}
}

// 사용자별 고루틴 실행
func StartMonitoring() {
	for _, user := range users {
		go MonitorUser(user)
	}
}
