package database

import "github.com/Prameesh-P/chat-api-cleancode/pkg/models"

func SyncDatabase() {

	DB.AutoMigrate(
		models.Users{},
	)
}