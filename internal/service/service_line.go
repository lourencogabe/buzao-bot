package service

import (
	"github.com/lourencogabe/buzao-bot/internal/models"
	"github.com/lourencogabe/buzao-bot/internal/repository"
)

// GetLineByNumber busca uma linha pelo número
func GetLineByNumber(number int) (*models.BusLine, error) {
	return repository.FindLineByNumber(number)
}

// SearchLinesByDescription busca linhas por descrição
func SearchLinesByDescription(description string) ([]models.BusLine, error) {
	return repository.FindLinesByDescription(description)
}

// GetAllLines retorna todas as linhas
func GetAllLines() ([]models.BusLine, error) {
	return repository.FindAllLines()
}
