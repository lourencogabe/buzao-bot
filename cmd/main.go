package main

import (
	"github.com/joho/godotenv"
	"github.com/lourencogabe/buzao-bot/external/bot"
	"github.com/lourencogabe/buzao-bot/internal/config"
	"github.com/lourencogabe/buzao-bot/internal/data"
	server "github.com/lourencogabe/buzao-bot/internal/http"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	if err := godotenv.Load(); err != nil {
		logger.ErrorF("env import failed", err)
		return
	}

	data.Connect()
	go bot.StartBot()
	server.StartServer()
}
