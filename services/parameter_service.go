package services

import (
	"HackOn/models"
	"HackOn/repositories"
)

// CreateParameter calls the repository to create a new parameter
func CreateParameter(param *models.Parameter) error {
	return repositories.CreateParameter(param)
}
