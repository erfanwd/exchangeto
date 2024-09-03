package repositories

import (
	"log"
	"strconv"
	"telegram-todolist/models"
)


func CreateReminder(chat_id int64, userSelections map[string]interface{}) (*models.Reminder, error) {
	// Get the user by chat ID
	user, err := getUserByChatId(chat_id)
	if err != nil {
		log.Println("Error fetching user:", err)
		return nil, err
	}

	crypto := userSelections["crypto"].(string)
	strategy := userSelections["strategy"].(string)
	amount := userSelections["amount"].(string)
	// Convert the value to an integer
	intcrypto, err := strconv.Atoi(crypto)
	if err != nil {
		log.Println("Error converting value to integer:", err)
		return nil, err
	}
	int64Value, err := strconv.ParseInt(amount, 10, 64)
	if err != nil {
		log.Println("Error converting value to integer:", err)
		return nil, err
	}
	uintCrypto := uint(intcrypto)

	// Create and save the reminder
	reminder := models.Reminder{
		ExchangeId: uintCrypto,
		UserId:     user.ID,
		Strategy:   strategy,
		Value:      int64Value,
	}

	if result := DB.Create(&reminder); result.Error != nil {
		return nil, result.Error
	}

	return &reminder, nil
}

func GetAllReminders() ([]models.Reminder, error) {
    var reminders []models.Reminder
    if err := DB.Preload("User").Preload("Exchange").Find(&reminders).Error; err != nil {
        return nil, err
    }
    return reminders, nil
}

func DeleteReminder(reminder models.Reminder){
	DB.Delete(reminder)
}


