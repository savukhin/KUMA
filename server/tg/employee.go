package tg

import (
	"errors"
	"server/db/query"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func checkEmployee(db *gorm.DB, telegramUsername string, checkIn bool) error {
	emp := query.Use(db).Employee
	result, err := emp.
		Where(emp.TelegramUserName.Eq(telegramUsername)).
		Update(emp.CheckedIN, checkIn)

	if err != nil {
		return err
	}

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("no such employee")
	}

	return nil
}

func (bot *TelegramBot) BroadcastCheckedInEmployees(text string, db *gorm.DB, logger *zap.Logger) error {
	emp := query.Use(db).Employee
	employees, err := emp.Where(emp.CheckedIN.Is(true)).Find()
	if err != nil {
		return err
	}

	for _, employee := range employees {
		username := employee.TelegramUserName
		msg := tgbotapi.NewMessageToChannel(username, text)
		_, err := bot.botApi.Send(msg)

		if err != nil {
			logger.Warn(err.Error())
		}
	}

	return nil
}
