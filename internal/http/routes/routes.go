package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lourencogabe/buzao-bot/internal/http/handlers"
)

func InitializeRoutes(router *gin.Engine) {
	v1 := router.Group("buzao-bot/v1")
	{
		v1.GET("/line", handlers.GetLine)
		v1.GET("/time", handlers.GetTime)
	}
}
