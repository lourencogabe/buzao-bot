package data

import (
	"encoding/csv"
	"encoding/json"
	"os"
	"path/filepath"
	"strconv"

	"github.com/lourencogabe/buzao-bot/internal/config"
	"github.com/lourencogabe/buzao-bot/internal/models"
	"go.etcd.io/bbolt"
)

// IsDatabasePopulated verifica se o banco já foi populado
func IsDatabasePopulated() bool {
	var count int

	DB.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(LinesBucket))
		if b == nil {
			return nil
		}
		count = b.Stats().KeyN
		return nil
	})

	return count > 0
}

// LoadCSVToDB carrega o arquivo CSV no banco de dados (apenas se vazio)
func LoadCSVToDB() error {
	logger := config.GetLogger("CSVLoader")

	// Verificar se banco já foi populado
	if IsDatabasePopulated() {
		stats := countLines()
		logger.InfoF("Banco de dados já populado com %d linhas. Pulando carregamento de CSV.", stats)
		return nil
	}

	logger.Info("Banco vazio. Carregando dados do CSV...")

	// Caminho relativo ao arquivo CSV
	csvPath := filepath.Join("internal", "data", "ListaInicial.csv")

	file, err := os.Open(csvPath)
	if err != nil {
		logger.ErrorF("erro ao abrir arquivo CSV: %v", err)
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		logger.ErrorF("erro ao ler CSV: %v", err)
		return err
	}

	// Pula o header (primeira linha)
	if len(records) > 0 {
		records = records[1:]
	}

	count := 0
	err = DB.Batch(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(LinesBucket))

		for _, record := range records {
			if len(record) < 6 {
				continue
			}

			// Trata o número da linha (pode conter letras como "X27")
			numberStr := record[1]
			number, err := strconv.Atoi(numberStr)
			if err != nil {
				logger.WarningF("número de linha inválido: %s", numberStr)
				continue
			}

			line := models.BusLine{
				Number:      number,
				Description: record[2],
				UrlUrbs:     record[3],
				UrlGazeta:   record[4],
				UrlMoovit:   record[5],
			}

			// Converter para JSON e salvar
			data, err := json.Marshal(line)
			if err != nil {
				logger.ErrorF("erro ao serializar linha %d: %v", number, err)
				continue
			}

			key := []byte(strconv.Itoa(number))
			if err := b.Put(key, data); err != nil {
				logger.ErrorF("erro ao inserir linha %d: %v", number, err)
				continue
			}
			count++
		}
		return nil
	})

	if err != nil {
		logger.ErrorF("erro durante transação: %v", err)
		return err
	}

	logger.InfoF("CSV carregado com sucesso: %d linhas inseridas", count)
	return nil
}

// countLines conta quantas linhas existem no banco
func countLines() int {
	var count int

	DB.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(LinesBucket))
		if b == nil {
			return nil
		}
		count = b.Stats().KeyN
		return nil
	})

	return count
}
