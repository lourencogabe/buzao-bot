package bot

import (
	"log"
	"os"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/lourencogabe/buzao-bot/internal/formatter"
	"github.com/lourencogabe/buzao-bot/internal/service"
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

	log.Println("🤖 Bot iniciado e aguardando mensagens...")

	for upd := range updates {
		if upd.Message == nil {
			continue
		}

		text := strings.TrimSpace(upd.Message.Text)
		if text == "" {
			continue
		}

		// Comando /start
		if text == "/start" {
			HandleStartCommand(b, upd.Message.Chat.ID)
			continue
		}

		// Comando /help
		if text == "/help" {
			HandleHelpCommand(b, upd.Message.Chat.ID)
			continue
		}

		// Tenta buscar por número
		if num, err := strconv.Atoi(text); err == nil {
			HandleLineSearch(b, upd.Message.Chat.ID, num)
			continue
		}

		// Tenta buscar por descrição/nome
		HandleDescriptionSearch(b, upd.Message.Chat.ID, text)
	}
}

// HandleLineSearch processa a busca de uma linha por número
func HandleLineSearch(b *tgbotapi.BotAPI, chatID int64, number int) {
	line, err := service.GetLineByNumber(number)
	if err != nil || line == nil {
		msg := tgbotapi.NewMessage(chatID, "❌ Linha "+strconv.Itoa(number)+" não encontrada.")
		b.Send(msg)
		return
	}

	response := formatter.FormatLineMessage(line)
	msg := tgbotapi.NewMessage(chatID, response)
	msg.ParseMode = "HTML"
	b.Send(msg)
}

// HandleDescriptionSearch processa a busca de linhas por descrição
func HandleDescriptionSearch(b *tgbotapi.BotAPI, chatID int64, description string) {
	lines, err := service.SearchLinesByDescription(description)
	if err != nil || len(lines) == 0 {
		msg := tgbotapi.NewMessage(
			chatID,
			"❌ Nenhuma linha encontrada para \""+description+"\".\n"+
				"💡 Dica: Use /help para ver como buscar linhas.",
		)
		b.Send(msg)
		return
	}

	// Se encontrou apenas uma linha, mostra os detalhes
	if len(lines) == 1 {
		response := formatter.FormatLineMessage(&lines[0])
		msg := tgbotapi.NewMessage(chatID, response)
		msg.ParseMode = "HTML"
		b.Send(msg)
		return
	}

	// Se encontrou múltiplas linhas, mostra a lista
	response := formatter.FormatMultipleLines(lines)
	msg := tgbotapi.NewMessage(chatID, response)
	b.Send(msg)
}

// HandleStartCommand responde ao comando /start
func HandleStartCommand(b *tgbotapi.BotAPI, chatID int64) {
	msg := tgbotapi.NewMessage(
		chatID,
		"🚌 Bem-vindo ao Búzão Bot!\n\n"+
			"Como usar:\n"+
			"📱 Envie o número da linha (ex: 464)\n"+
			"📍 Ou o nome da linha/ponto de ônibus (ex: Cabral)\n"+
			"/help - Ver mais opções",
	)
	b.Send(msg)
}

// HandleHelpCommand responde ao comando /help
func HandleHelpCommand(b *tgbotapi.BotAPI, chatID int64) {
	msg := tgbotapi.NewMessage(
		chatID,
		"📖 Comandos disponíveis:\n\n"+
			"/start - Mensagem inicial\n"+
			"/help - Esta mensagem\n\n"+
			"💡 Buscar linhas:\n"+
			"• Envie um número: 464\n"+
			"• Envie um nome: Cabral, Centro, etc\n\n"+
			"Cada linha mostra links para:\n"+
			"🔗 URBS (horários oficial)\n"+
			"📰 Gazeta do Povo\n"+
			"🗺️ Moovit",
	)
	b.Send(msg)
}
