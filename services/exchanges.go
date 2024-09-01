package services

import (
	"fmt"
	"strconv"
	"telegram-todolist/repositories"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func ExchangeList(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	data, _ := repositories.GetAllExchanges(update.Message.Chat.ID)
	var btns []tgbotapi.InlineKeyboardButton
	
	for i := 0; i < len(data); i++ {	
		btn := tgbotapi.NewInlineKeyboardButtonData(data[i].Name, "selected_exchange="+strconv.FormatUint(uint64(data[i].ID), 10))
		btns = append(btns, btn)
	}

	var rows [][]tgbotapi.InlineKeyboardButton
	for i := 0; i < len(btns); i += 2 {
		if i < len(btns) && i+1 < len(btns) {
			row := tgbotapi.NewInlineKeyboardRow(btns[i], btns[i+1])
			rows = append(rows, row)
		} else if i < len(btns) {
			row := tgbotapi.NewInlineKeyboardRow(btns[i])
			rows = append(rows, row)
		}
	}
	fmt.Println(len(rows))
	var keyboard = tgbotapi.InlineKeyboardMarkup{InlineKeyboard: rows}
	//keyboard.InlineKeyboard = rows

	text := "لطفا ارز مورد نظر را انتخاب کنید"
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
	msg.ReplyMarkup = keyboard
	if _, err := bot.Send(msg); err != nil {
		panic(err)
	}
}
