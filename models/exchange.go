package models

type Exchange struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"type:varchar(50);column:name"`
	Code      string `gorm:"type:varchar(50);column:code"`
}
