package repositories

import (
	"HackOn/database"
	"HackOn/models"
)

func GetMerchantByID(merchantID string) (*models.Merchant, error) {
	var merchant models.Merchant
	if err := database.DB.Where("merchantid = ?", merchantID).First(&merchant).Error; err != nil {
		return nil, err
	}
	return &merchant, nil
}

// Add similar CRUD operations for Parameter and Pricing
