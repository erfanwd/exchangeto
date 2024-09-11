package services

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/erfanwd/exchangeto/repositories"
    "github.com/dustin/go-humanize"

	// "github.com/erfanwd/exchangeto/repositories"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func SetReminder(bot *tgbotapi.BotAPI, update tgbotapi.Update, value string) {
	chatID := update.CallbackQuery.Message.Chat.ID
	_, hasState := userStates[chatID]
	text := ""
	if !hasState {
		text = "لطفا از منو پایین فرایند افزودن ریماندر را مجددا شروع کنید."
	} else {
		userSelections[chatID]["strategy"] = value

		userSelections := GetStateSelections(chatID)

		if _, err := repositories.CreateReminder(chatID, userSelections); err != nil {
			log.Println("Error creating reminder:", err)
			return
		}

		text = "ریماندر شما با موفقیت ثبت شد."
		delete(userStates, chatID)
		delete(userSelections, strconv.FormatInt(chatID, 10))

	}

	msg := tgbotapi.NewMessage(chatID, text)

	if _, err := bot.Send(msg); err != nil {
		log.Println("Error sending message:", err)
		panic(err) // Consider handling this more gracefully
	}

}

func SetAmount(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	chatID := update.Message.Chat.ID
	userSelections[chatID]["amount"] = update.Message.Text

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "استراتژی مورد نظر را انتخاب کنید")
	crypto, _ := repositories.GetExchangeById(userSelections[chatID]["crypto"].(string))
	amount, _ := strconv.ParseInt(userSelections[chatID]["amount"].(string), 10, 64)
	higherMsg := crypto.Name + " بالاتر از " + humanize.Comma(amount) + " دلار بود، بهم پیغام بده "
	lowerMsg := crypto.Name + " پایین تر از " + humanize.Comma(amount) + " دلار بود، بهم پیغام بده "

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

func RemindersList(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	user, _ := repositories.GetUserByChatId(update.Message.Chat.ID)
	reminders, _ := repositories.GetRemindersByUserId(user.ID)

	var output []string
	fmt.Println(reminders)
	if len(reminders) == 0 {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "هنوز ریمایندری ست نکردی که :)")
		if _, err := bot.Send(msg); err != nil {
			log.Println("Error sending message:", err)
			return
		}
		return
	}

	for key, reminder := range reminders {

		output = append(output, fmt.Sprintf("%d \n استراتژی: %s \n ارز: %s \n مقدار تعیین شده: %s دلار \n --------------------------------------------------", key+1, reminder.GetPersianStrategy(), reminder.Exchange.Name,humanize.Comma(reminder.Value) ))
	}

	result := strings.Join(output, "\n")

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, result)

	if _, err := bot.Send(msg); err != nil {
		log.Println("Error sending message:", err)
		panic(err.Error()) 
	}
}
