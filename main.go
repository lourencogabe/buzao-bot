package main

import (
	"github.com/lourencogabe/buzao-bot/internal/data"
	"github.com/lourencogabe/buzao-bot/internal/server"
)

func main() {
	server.StartServer()
	data.Connect()
}
