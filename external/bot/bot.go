package bot

import (
	"log"
	"os"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func StartBot() {
	token := os.Getenv("TELEGRAM_TOKEN")
	if token == "" {
		log.Println("TELEGRAM_TOKEN não definido")
		return
	}

	b, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal(err)
	}
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 30
	updates := b.GetUpdatesChan(u)

	for upd := range updates {
		if upd.Message == nil {
			continue
		}

		text := strings.TrimSpace(upd.Message.Text)
		if text == "" {
			continue
		}
		
		if num, err := strconv.Atoi(text); err == nil {
			b.Send(tgbotapi.NewMessage(upd.Message.Chat.ID, "Buscando linha: "+strconv.Itoa(num)))
			continue
		}

		// heurística: se parecer nome de terminal/ponto, buscar por nome
		lower := strings.ToLower(text)
		if strings.Contains(lower, "terminal") || strings.Contains(lower, "ponto") || len(text) > 2 {
			b.Send(tgbotapi.NewMessage(upd.Message.Chat.ID, "Buscando por terminal/ponto: "+text))
			continue
		}

		b.Send(tgbotapi.NewMessage(upd.Message.Chat.ID, "Não entendi. Envie /onibus <número> ou apenas o número/terminal."))
	}
}
