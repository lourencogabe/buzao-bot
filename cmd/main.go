package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/lourencogabe/buzao-bot/external/bot"
	"github.com/lourencogabe/buzao-bot/internal/server"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println(".env not loaded:", err)
	}

	//data.Connect()
	go bot.StartBot()
	server.StartServer()
}
