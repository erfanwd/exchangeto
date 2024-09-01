package keyboards

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func CmdKeyboard() tgbotapi.ReplyKeyboardMarkup {
	var cmdKeyboard = tgbotapi.NewOneTimeReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("/set_reminder"),
		),
	)
	return cmdKeyboard
}



