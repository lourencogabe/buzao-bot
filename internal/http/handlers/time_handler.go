package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTime(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "time found successfully",
	})
}
