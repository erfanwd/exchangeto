package repositories

import (
	"telegram-todolist/models"
)



func GetAllExchanges(chatId int64) ([]models.Exchange, error) {
	var exchanges []models.Exchange
	if result := DB.Find(&exchanges); result.Error != nil {
		return exchanges, result.Error
	}
	return exchanges, nil
}
