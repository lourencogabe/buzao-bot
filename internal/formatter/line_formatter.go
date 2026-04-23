package formatter

import (
	"fmt"

	"github.com/lourencogabe/buzao-bot/internal/models"
)

// FormatLineMessage formata uma linha para exibição
func FormatLineMessage(line *models.BusLine) string {
	message := fmt.Sprintf("🚌 Linha %d: %s\n", line.Number, line.Description)

	if line.UrlUrbs != "" {
		message += fmt.Sprintf("🔗 URBS: %s\n", line.UrlUrbs)
	}
	if line.UrlGazeta != "" {
		message += fmt.Sprintf("📰 Gazeta: %s\n", line.UrlGazeta)
	}
	if line.UrlMoovit != "" {
		message += fmt.Sprintf("🗺️ Moovit: %s\n", line.UrlMoovit)
	}

	return message
}

// FormatMultipleLines formata múltiplas linhas para exibição
func FormatMultipleLines(lines []models.BusLine) string {
	if len(lines) == 0 {
		return "❌ Nenhuma linha encontrada."
	}

	message := fmt.Sprintf("🚌 Encontradas %d linha(s):\n\n", len(lines))
	for _, line := range lines {
		message += fmt.Sprintf("• Linha %d: %s\n", line.Number, line.Description)
	}
	message += "\n📍 Use o número da linha para mais detalhes."
	return message
}
