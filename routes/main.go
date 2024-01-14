package routes

import (
	"net/http"

	"github.com/Pugpaprika21/go-gin/controller"
	"github.com/gin-gonic/gin"
)

func Setup(router *gin.Engine) {
	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "home.html", gin.H{})
	})

	user := controller.User{}
	userGroup := router.Group("/user")
	{
		userGroup.GET("/", user.PageHome)
		userGroup.GET("/create", user.PageCreate)
		userGroup.POST("/create", user.Create)
		userGroup.GET("/update/:id", user.PageUpdate)
		userGroup.PUT("/update/:id", user.Update)
		userGroup.DELETE("/delete/:id", user.Delete)
	}

	dashboard := controller.Dashboard{}
	dashboardGroup := router.Group("/dashboard")
	{
		dashboardGroup.GET("/", dashboard.PageHome)
	}
}
