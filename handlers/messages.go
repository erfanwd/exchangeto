package handlers

import (
	"telegram-todolist/services"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Messages(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	state := services.GetState(update.Message.Chat.ID)

	if state == "choosing_amount" {
		services.SetAmount(bot, update)
	}
}
