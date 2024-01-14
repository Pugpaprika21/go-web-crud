package controller

import (
	"net/http"

	"github.com/Pugpaprika21/go-gin/db"
	"github.com/Pugpaprika21/go-gin/models"
	"github.com/gin-gonic/gin"
)

type Dashboard struct{}

var usersNotNull []struct {
	UserId    int    `gorm:"column:id"`
	Username  string `gorm:"column:username"`
	Password  string `gorm:"column:password"`
	DateStart string `gorm:"column:date_start"`
}

var usersDataDashboard struct {
	IsNotNull int64 `gorm:"column:is_not_null"`
	IsNull    int64 `gorm:"column:is_null"`
}

func (d Dashboard) PageHome(ctx *gin.Context) {
	dateStart := ctx.Query("date_start")
	dateEnd := ctx.Query("date_end")

	query := db.Conn.Model(&models.User{}).
		Select("id, username, password, DATE_FORMAT(created_at, '%Y-%m-%d') as date_start")

	if dateStart != "" && dateEnd != "" {
		query = query.Where("created_at BETWEEN ? AND ?", dateStart+" 00:00:00", dateEnd+" 23:59:59")
	}

	query.Where("deleted_at IS NULL").Find(&usersNotNull)

	db.Conn.Raw(`SELECT COALESCE(SUM(CASE WHEN deleted_at IS NOT NULL THEN 1 END), 0) as is_not_null,
		COALESCE(SUM(CASE WHEN deleted_at IS NULL THEN 1 END), 0) as is_null
		FROM users
	`).Scan(&usersDataDashboard)

	ctx.HTML(http.StatusOK, "home-dashboard.html", gin.H{
		"Data": gin.H{
			"User": usersDataDashboard,
		},
	})
}
