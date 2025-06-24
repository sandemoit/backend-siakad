package service

import (
	"errors"
	"siakad/api/models"
	"siakad/config"
	"siakad/utils"
)

func AuthenticateUser(email, password string) (*models.User, error) {
	var user models.User
	if err := config.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	if ok := utils.CheckPassword(user.Password, password); !ok {
		return nil, errors.New("invalid password")
	}

	return &user, nil
}
