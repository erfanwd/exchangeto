package services

import (
	"log"
	"strconv"

	"github.com/erfanwd/exchangeto/repositories"

	// "github.com/erfanwd/exchangeto/repositories"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func SetReminder(bot *tgbotapi.BotAPI, update tgbotapi.Update, value string) {
	chatID := update.CallbackQuery.Message.Chat.ID

	userSelections[chatID]["strategy"] = value

	userSelections := GetStateSelections(chatID)

	if _, err := repositories.CreateReminder(chatID, userSelections); err != nil {
		log.Println("Error creating reminder:", err)
		return
	}

	text := "ریماندر شما با موفقیت ثبت شد."
	msg := tgbotapi.NewMessage(chatID, text)

	if _, err := bot.Send(msg); err != nil {
		log.Println("Error sending message:", err)
		panic(err) // Consider handling this more gracefully
	}

	delete(userStates, chatID)
	delete(userSelections, strconv.FormatInt(chatID, 10))
}

func SetAmount(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	chatID := update.Message.Chat.ID
	userSelections[chatID]["amount"] = update.Message.Text

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "استراتژی مورد نظر را انتخاب کنید")
	crypto, _ := repositories.GetExchangeById(userSelections[chatID]["crypto"].(string))
	higherMsg := crypto.Name + " بالاتر از " + userSelections[chatID]["amount"].(string) + " دلار بود، بهم پیغام بده "
	lowerMsg := crypto.Name + " پایین تر از " + userSelections[chatID]["amount"].(string) + " دلار بود، بهم پیغام بده "

	var btns []tgbotapi.InlineKeyboardButton

	btn := tgbotapi.NewInlineKeyboardButtonData(higherMsg, "selected_strategy=higher")
	btns = append(btns, btn)
	btn = tgbotapi.NewInlineKeyboardButtonData(lowerMsg, "selected_strategy=lower")
	btns = append(btns, btn)

	var rows [][]tgbotapi.InlineKeyboardButton
	for i := 0; i < len(btns); i += 1 {
		row := tgbotapi.NewInlineKeyboardRow(btns[i])
		rows = append(rows, row)
	}
	var keyboard = tgbotapi.InlineKeyboardMarkup{InlineKeyboard: rows}

	msg.ReplyMarkup = keyboard

	if _, err := bot.Send(msg); err != nil {
		log.Println("Error sending message:", err)
		panic(err) // Consider handling this more gracefully
	}
	userStates[update.Message.Chat.ID] = "choosing_strategy"

}
