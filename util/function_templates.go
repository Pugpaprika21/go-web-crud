package util

import (
	"os"
	"strings"
	"time"

	"github.com/Pugpaprika21/go-gin/db"
	"github.com/Pugpaprika21/go-gin/models"
)

func TimeNow() time.Time {
	currentTime := time.Now()
	return currentTime
}

func Rows(index int) int {
	return index + 1
}

func ConvertYMD(createAt time.Time) string {
	parts := strings.Split(createAt.String(), " ")
	if len(parts) >= 1 {
		datePart := parts[0]
		return datePart
	}
	return ""
}

func UserTotal() int64 {
	var userTotal int64
	db.Conn.Model(&models.User{}).Count(&userTotal)
	return userTotal
}

func AppURL() string {
	return os.Getenv("APP_URL")
}
