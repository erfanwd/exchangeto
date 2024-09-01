package services

import (
	"fmt"
	"telegram-todolist/repositories"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func SetReminder(bot *tgbotapi.BotAPI, update tgbotapi.Update, value string) {
	fmt.Println(update.Message.Chat.ID)
	repositories.CreateReminder(update, value)
	text := "please set amount"
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
	if _, err := bot.Send(msg); err != nil {
		panic(err)
	}
}
