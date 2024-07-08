package repositories

import (
	"HackOn/database"
	"HackOn/models"
)

// CreateParameter saves the parameter in the database
func CreateParameter(param *models.Parameter) error {
	return database.DB.Create(param).Error
}

func GetParamerterByRoundCompAndVolume(round string, volume string, competition string) (*models.Parameter, error) {
	var parameter models.Parameter
	round = "Round " + round
	err := database.DB.Where("round = ? AND volume = ? AND competition = ?", round, volume, competition).First(&parameter).Error
	if err != nil {
		return nil, err
	}

	return &parameter, nil
}
