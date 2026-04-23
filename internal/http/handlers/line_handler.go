package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lourencogabe/buzao-bot/internal/formatter"
	"github.com/lourencogabe/buzao-bot/internal/service"
)

// GetLineByNumber busca uma linha pelo número
// GET /lines/:number
func GetLineByNumber(ctx *gin.Context) {
	numberStr := ctx.Param("number")
	number, err := strconv.Atoi(numberStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "número inválido",
		})
		return
	}

	line, err := service.GetLineByNumber(number)
	if err != nil || line == nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "linha não encontrada",
		})
		return
	}

	ctx.JSON(http.StatusOK, line)
}

// SearchLines busca linhas por descrição
// GET /lines/search?q=termo
func SearchLines(ctx *gin.Context) {
	query := ctx.Query("q")
	if query == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "parâmetro 'q' é obrigatório",
		})
		return
	}

	lines, err := service.SearchLinesByDescription(query)
	if err != nil || len(lines) == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "nenhuma linha encontrada",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"count": len(lines),
		"lines": lines,
	})
}

// GetAllLines retorna todas as linhas
// GET /lines
func GetAllLines(ctx *gin.Context) {
	lines, err := service.GetAllLines()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "erro ao buscar linhas",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"count": len(lines),
		"lines": lines,
	})
}

// GetLineFormatted busca uma linha e retorna formatada para exibição
// GET /lines/:number/formatted
func GetLineFormatted(ctx *gin.Context) {
	numberStr := ctx.Param("number")
	number, err := strconv.Atoi(numberStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "número inválido",
		})
		return
	}

	line, err := service.GetLineByNumber(number)
	if err != nil || line == nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "linha não encontrada",
		})
		return
	}

	formatted := formatter.FormatLineMessage(line)
	ctx.JSON(http.StatusOK, gin.H{
		"formatted": formatted,
		"line":      line,
	})
}
