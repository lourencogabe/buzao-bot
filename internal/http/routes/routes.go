package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lourencogabe/buzao-bot/internal/http/handlers"
)

func InitializeRoutes(router *gin.Engine) {
	v1 := router.Group("buzao-bot/v1")
	{
		// Endpoints de linhas
		v1.GET("/lines", handlers.GetAllLines)
		v1.GET("/lines/:number", handlers.GetLineByNumber)
		v1.GET("/lines/search", handlers.SearchLines)
		v1.GET("/lines/:number/formatted", handlers.GetLineFormatted)

		// Endpoint de horários (será implementado)
	}
}
