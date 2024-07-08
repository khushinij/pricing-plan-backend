package repositories

import (
	"HackOn/database"
	"HackOn/models"
)

func CreateChatMessage(msg *models.ChatMessage) error {
	if err := database.DB.Create(msg).Error; err != nil {
		return err
	}
	return nil
}

func GetChatMessagesSorted() ([]models.ChatMessage, error) {
	var messages []models.ChatMessage
	if err := database.DB.Order("timestamp DESC").Find(&messages).Error; err != nil {
		return nil, err
	}
	return messages, nil
}
