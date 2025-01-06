package core

import "sync"

// 사용자 정보 저장 구조체
type User struct {
	ID        string
	AccessKey string
	SecretKey string
	Strategy map[string]interface{}
}

// 사용자 관리 맵 및 동기화
var users = make(map[string]*User)
var usersLock sync.Mutex

// 사용자 추가
func AddUser(id, accessKey, secretKey string) {
	usersLock.Lock()
	defer usersLock.Unlock()
	users[id] = &User{
		ID:        id,
		AccessKey: accessKey,
		SecretKey: secretKey,
		Strategy:  make(map[string]interface{}), // 전략 초기화
	}
}

// 사용자 조회
func GetUser(id string) (*User, bool) {
	usersLock.Lock()
	defer usersLock.Unlock()
	user, exists := users[id]
	return user, exists
}

// 사용자 전략 조회
func GetUserStrategy(id string) (map[string]interface{}, bool) {
	usersLock.Lock()
	defer usersLock.Unlock()
	if user, exists := users[id]; exists {
		return user.Strategy, true
	}
	return nil, false
}