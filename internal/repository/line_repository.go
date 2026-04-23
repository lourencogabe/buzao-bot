package repository

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/lourencogabe/buzao-bot/internal/data"
	"github.com/lourencogabe/buzao-bot/internal/models"
	"go.etcd.io/bbolt"
)

// FindLineByNumber busca uma linha pelo número
func FindLineByNumber(number int) (*models.BusLine, error) {
	var line *models.BusLine

	err := data.DB.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(data.LinesBucket))
		if b == nil {
			return errors.New("bucket não existe")
		}

		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			var l models.BusLine
			if err := json.Unmarshal(v, &l); err != nil {
				continue
			}
			if l.Number == number {
				line = &l
				return nil
			}
		}
		return errors.New("linha não encontrada")
	})

	return line, err
}

// FindLinesByDescription busca linhas por descrição (partial match)
func FindLinesByDescription(description string) ([]models.BusLine, error) {
	var lines []models.BusLine
	description = strings.ToUpper(description)

	err := data.DB.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(data.LinesBucket))
		if b == nil {
			return errors.New("bucket não existe")
		}

		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			var line models.BusLine
			if err := json.Unmarshal(v, &line); err != nil {
				continue
			}
			if strings.Contains(strings.ToUpper(line.Description), description) {
				lines = append(lines, line)
			}
		}
		return nil
	})

	return lines, err
}

// FindAllLines retorna todas as linhas
func FindAllLines() ([]models.BusLine, error) {
	var lines []models.BusLine

	err := data.DB.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(data.LinesBucket))
		if b == nil {
			return errors.New("bucket não existe")
		}

		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			var line models.BusLine
			if err := json.Unmarshal(v, &line); err != nil {
				continue
			}
			lines = append(lines, line)
		}
		return nil
	})

	return lines, err
}
