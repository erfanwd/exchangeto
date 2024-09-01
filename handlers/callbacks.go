package handlers

import (
	"fmt"
	"telegram-todolist/services"
	"telegram-todolist/utils"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Callbacks(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	fmt.Println(update.Message.Chat.ID)
	cmd, value := utils.GetKeyValue(update.CallbackQuery.Data)
	switch {
	case cmd == "selected_exchange":
	
		services.SetReminder(bot, update, value)
	}
}
