package initializers

import (
	"github.com/Bouchiba43/Auth-Go/models"
)

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}