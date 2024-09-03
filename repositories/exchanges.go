package repositories

import (
	"telegram-todolist/models"
)



func GetAllExchanges() ([]models.Exchange, error) {
	var exchanges []models.Exchange
	if result := DB.Find(&exchanges); result.Error != nil {
		return exchanges, result.Error
	}
	return exchanges, nil
}

func GetExchangeById(id string) (models.Exchange, error){
	var exchange models.Exchange
	if result := DB.First(&exchange, id); result.Error != nil {
		return exchange, result.Error
	}
	return exchange, nil
}
