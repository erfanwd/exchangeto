package repositories

import (
	"github.com/erfanwd/exchangeto/database"

	"gorm.io/gorm"
)

var DB *gorm.DB = database.Init()
