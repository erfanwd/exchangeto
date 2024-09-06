package services

import (
	"fmt"
	"log"
	"strconv"

	"github.com/erfanwd/exchangeto/repositories"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func ExchangeList(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	data, _ := repositories.GetAllExchanges()
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
	userStates[update.Message.Chat.ID] = "choosing_crypto"
}

func SetExchange(bot *tgbotapi.BotAPI, update tgbotapi.Update, value string) {
	// Access the chat ID from the CallbackQuery's message
	chatID := update.CallbackQuery.Message.Chat.ID

	userSelections[chatID] = map[string]interface{}{
		"crypto": value,
	}

	// Send the response message
	text := "حالا لطفا اون رقمی که میخوای اگه این ارز بهش رسید برات پیغام بیاد رو بنویس (به دلار)"
	msg := tgbotapi.NewMessage(chatID, text)

	if _, err := bot.Send(msg); err != nil {
		log.Println("Error sending message:", err)
		panic(err) // Consider handling this more gracefully
	}

	userStates[chatID] = "choosing_amount"
}
