package data

import (
	"encoding/json"
	"github.com/lourencogabe/buzao-bot/internal/config"
	"github.com/lourencogabe/buzao-bot/internal/models"
	"go.etcd.io/bbolt"
)

var DB *bbolt.DB
var logger *config.Logger

// Nomes dos buckets
const (
	LinesBucket = "lines"
	TimesBucket = "times"
)

func Connect() {
	logger = config.GetLogger("BoltDB")

	var err error
	DB, err = bbolt.Open("bus.db", 0600, nil)
	if err != nil {
		logger.ErrorF("failed to connect to database: %v", err)
		return
	}

	// Criar buckets se não existirem
	err = DB.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(LinesBucket))
		if err != nil {
			return err
		}
		_, err = tx.CreateBucketIfNotExists([]byte(TimesBucket))
		return err
	})

	if err != nil {
		logger.ErrorF("failed to create buckets: %v", err)
		return
	}

	logger.Info("Database connected successfully!")
}

// SaveLine salva uma linha no banco
func SaveLine(line *models.BusLine) error {
	return DB.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(LinesBucket))
		data, err := json.Marshal(line)
		if err != nil {
			return err
		}
		key := []byte(toKey(line.Number))
		return b.Put(key, data)
	})
}

// SaveTime salva um horário no banco
func SaveTime(time *models.BusTime) error {
	return DB.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(TimesBucket))
		data, err := json.Marshal(time)
		if err != nil {
			return err
		}
		key := []byte(toKey(int(time.ID)))
		return b.Put(key, data)
	})
}

// toKey converte número para string key
func toKey(num int) string {
	return string(rune(num))
}
