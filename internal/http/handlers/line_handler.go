package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetLine(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "line found successfully",
	})
}
