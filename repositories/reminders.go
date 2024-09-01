package repositories

import (
	"fmt"
	"log"
	"strconv"
	"telegram-todolist/models"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func CreateReminder(update tgbotapi.Update, value string) (*models.Reminder, error) {
	fmt.Println(update.Message.Chat)
	panic("fgdfg")
	user, err := getUserByChatId(update.Message.Chat.ID)
	if err != nil {
        log.Println("Error fetching user:", err)
        
    }
	intvalue, _ := strconv.Atoi(value)
	uintValue := uint(intvalue)
	reminder := models.Reminder{
		ExchangeId:    uintValue,
		UserId:  user.ID,
	}

	if result := DB.Create(&reminder); result.Error != nil {
		return nil, result.Error
	}
	return &reminder, nil
}
