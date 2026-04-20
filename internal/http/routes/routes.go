package routes

import (
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {
	v1 := router.Group("buzao-bot/v1")
	{
		v1.GET("/line")
		v1.GET("/time")
		v1.GET("/schedules")
	}

}
