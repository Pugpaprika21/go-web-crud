package main

import (
	"log"
	"text/template"

	"github.com/Pugpaprika21/go-gin/db"
	"github.com/Pugpaprika21/go-gin/middleware"
	"github.com/Pugpaprika21/go-gin/routes"
	"github.com/Pugpaprika21/go-gin/util"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db.ConnectDB()
	db.Migrate()

	router := gin.Default()
	router.Use(middleware.Cors())

	router.SetFuncMap(template.FuncMap{
		"rows":       util.Rows,
		"convertYMD": util.ConvertYMD,
		"userTotal":  util.UserTotal,
		"isTime":     util.TimeNow,
		"app_url":    util.AppURL,
	})

	router.LoadHTMLGlob("resources/templates/**/*")
	router.Static("/assets", "./assets")

	routes.Setup(router)

	router.Run(":8080")
}
