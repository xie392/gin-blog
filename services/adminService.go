package services

import (
	"blog/configs"
	"blog/models"
	"blog/utils"
	"errors"
	"gorm.io/gorm"
)

func GetUser(username, password string) (models.Admin, error) {
	var admin models.Admin
	if err := configs.DB.Where("username = ? AND password = ?", username, password).First(&admin).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return models.Admin{}, errors.New("找不到该用户")
		}
		return models.Admin{}, err
	}
	token, err := utils.GenerateToken(admin.ID)
	if err != nil {
		return models.Admin{}, err
	}
	admin.Token = token
	return admin, nil
}

func CreateUser(user models.Admin) (models.Admin, error) {
	if err := configs.DB.Create(&user).Error; err != nil {
		return models.Admin{}, err
	}
	return user, nil
}
