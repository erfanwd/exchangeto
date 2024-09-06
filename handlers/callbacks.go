package handlers

import (
	"github.com/erfanwd/exchangeto/services"
	"github.com/erfanwd/exchangeto/utils"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Callbacks(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	cmd, value := utils.GetKeyValue(update.CallbackQuery.Data)
	switch {
	case cmd == "selected_exchange":
		services.SetExchange(bot, update, value)
	case cmd == "selected_strategy":
		services.SetReminder(bot, update, value)
	}

}
