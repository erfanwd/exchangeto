package handlers

import (
	"github.com/erfanwd/exchangeto/services"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Commands(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	switch update.Message.Command() {
	case "start":
		services.StartChat(bot, update)
	case "set_reminder":
		services.ExchangeList(bot, update)
	}
}
