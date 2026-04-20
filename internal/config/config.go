package config

var (
	logger *Logger
)

// Inicializa Logger
func GetLogger(p string) *Logger {
	logger = NewLogger(p)
	return logger
}
