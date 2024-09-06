package repositories

import (
	"log"
	"time"

	"github.com/erfanwd/exchangeto/models"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func CreateUser(update tgbotapi.Update) (*models.User, error) {

	user := models.User{
		ChatId:    update.Message.Chat.ID,
		Username:  update.Message.Chat.UserName,
		FirstName: update.Message.Chat.FirstName,
		CreatedAt: time.Now(),
	}

	if result := DB.FirstOrCreate(&user, models.User{ChatId: user.ChatId}); result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func GetUserByChatId(chatId int64) (*models.User, error) {
	var user models.User
	result := DB.Where("chat_id = ?", chatId).First(&user)
	if result.Error != nil {
		log.Println("Error finding user by chat ID:", result.Error)
		return nil, result.Error
	}
	return &user, nil
}
