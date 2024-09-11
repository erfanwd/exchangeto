package models

import "time"

type Reminder struct {
	ID         uint      `gorm:"primaryKey"`
	UserId     uint      `gorm:"column:user_id"`   
	User       User      `gorm:"foreignKey:UserId"`
	ExchangeId uint      `gorm:"column:exchange_id"` 
	Exchange   Exchange  `gorm:"foreignKey:ExchangeId"` 
	Value      int64     `gorm:"column:value"`    
	Strategy   string    `gorm:"type:varchar(50);column:strategy"` 
	CreatedAt  time.Time `gorm:"autoCreateTime"`  
}


var strategyTranslations = map[string]string{
	"higher":  "بالا تر از مقدار تعیین شده",
	"lower": "پایین تر از مقدار تعیین شده",
}

func (r *Reminder) GetPersianStrategy() string {
	if persianStrategy, ok := strategyTranslations[r.Strategy]; ok {
		return persianStrategy
	}
	return r.Strategy 
}


