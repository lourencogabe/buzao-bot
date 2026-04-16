package bot

import (
    "log"
    tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"os"
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
        if upd.Message.IsCommand() {
            switch upd.Message.Command() {
            case "start":
                b.Send(tgbotapi.NewMessage(upd.Message.Chat.ID, "Olá! Use /onibus <número>"))
            case "onibus":
                arg := upd.Message.CommandArguments()
                if arg == "" {
                    b.Send(tgbotapi.NewMessage(upd.Message.Chat.ID, "Uso: /onibus <número>"))
                    continue
                }
                // Aqui você chamaria sua lógica DB/handler para buscar horários
                b.Send(tgbotapi.NewMessage(upd.Message.Chat.ID, "Buscando linha: "+arg))
            default:
                b.Send(tgbotapi.NewMessage(upd.Message.Chat.ID, "Comando não reconhecido"))
            }
        }
    }
}
