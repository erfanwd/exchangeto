package handlers

import (
	"github.com/erfanwd/exchangeto/services"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Messages(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	
	switch update.Message.Text {
		case "افزودن ریمایندر جدید" : 
			services.ExchangeList(bot, update)
		case "لیست ریمایندر ها" : 
			services.RemindersList(bot, update)	
	}
	
	state := services.GetState(update.Message.Chat.ID)

	if state == "choosing_amount" {
		services.SetAmount(bot, update)
	}
}
