package services

import (
	"telegram-todolist/keyboards"
	"telegram-todolist/repositories"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var userStates = make(map[int64]string)
var userSelections = make(map[int64]map[string]interface{})

func StartChat(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	user,_ := repositories.CreateUser(update)
	userStates[update.Message.Chat.ID] = "user_started"
	text := " سلام" + user.FirstName + " از گزینه های پایین یکی رو انتخاب کن :)"
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
	msg.ReplyMarkup = keyboards.CmdKeyboard()
	if _, err := bot.Send(msg); err != nil {
		panic(err)
	}
}

func GetState(chat_id int64) string{
	return userStates[chat_id]
}

func GetStateSelections(chat_id int64) map[string]interface{}{
	return userSelections[chat_id]
}

