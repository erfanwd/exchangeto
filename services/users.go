package services

import (

	"telegram-todolist/keyboards"
	"telegram-todolist/repositories"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func StartChat(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	user,_ := repositories.CreateUser(update)
	text := " سلام" + user.FirstName + " از گزینه های پایین یکی رو انتخاب کن :)"
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
	msg.ReplyMarkup = keyboards.CmdKeyboard()
	if _, err := bot.Send(msg); err != nil {
		panic(err)
	}
}

