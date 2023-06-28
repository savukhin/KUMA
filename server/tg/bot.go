package tg

import (
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

const (
	checkInCommand  = "check-in"
	checkOutCommand = "check-out"
)

type TelegramBot struct {
	botApi *tgbotapi.BotAPI
}

func NewTelegramBot(token string) (*TelegramBot, error) {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_APITOKEN"))
	if err != nil {
		return nil, err
	}

	return &TelegramBot{
		botApi: bot,
	}, nil
}

func (bot *TelegramBot) Start(db *gorm.DB, logger *zap.Logger) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.botApi.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message updates
			continue
		}

		if !update.Message.IsCommand() { // ignore any non-command Messages
			continue
		}

		// Create a new MessageConfig. We don't have text yet,
		// so we leave it empty.
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		switch update.Message.Command() {
		case checkInCommand:
			err := checkEmployee(db, update.SentFrom().UserName, true)
			if err == nil {
				msg.Text = "Sucessfully check in"
			} else {
				msg.Text = "Server error while check in: " + err.Error()
			}
		case checkOutCommand:
			err := checkEmployee(db, update.SentFrom().UserName, false)
			if err == nil {
				msg.Text = "Sucessfully check out"
			} else {
				msg.Text = "Server error while check out: " + err.Error()
			}
		default:
			msg.Text = "I don't know that command"
		}

		if _, err := bot.botApi.Send(msg); err != nil {
			logger.Error("Cannot send message " + err.Error())
		}
	}
}
