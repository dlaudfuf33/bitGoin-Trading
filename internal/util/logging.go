package util

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

var Logger = logrus.New()

func InitLogger() {
	Logger.SetLevel(logrus.InfoLevel)
	Logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	// 로그 파일 출력
	logFile, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		Logger.SetOutput(logFile)
	} else {
		Logger.Warn("로그 파일 생성 실패. 기본 stderr 사용")
		return
	}
	// 파일과 콘솔에 동시에 출력
	multiWriter := io.MultiWriter(logFile, os.Stdout)
	Logger.SetOutput(multiWriter)
}
