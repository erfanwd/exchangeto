package models

import "time"


type User struct {
	ID        uint      `gorm:"primaryKey"`
	FirstName string    `gorm:"type:varchar(100);column:first_name"`
	Username  string    `gorm:"type:varchar(50);column:username"`
	ChatId    int64     `gorm:"chat_id"`
	CreatedAt time.Time
}
