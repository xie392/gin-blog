package services

import (
	"blog/configs"
	"blog/models"
	"blog/utils"
	"errors"
	"gorm.io/gorm"
)

func GetUser(username, password string) (models.AdminResponse, error) {
	var admin models.Admin
	if err := configs.DB.Where("username = ? AND password = ?", username, password).First(&admin).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return models.AdminResponse{}, errors.New("找不到该用户")
		}
		return models.AdminResponse{}, err
	}
	token, err := utils.GenerateToken(admin.ID)
	if err != nil {
		return models.AdminResponse{}, err
	}
	result := models.AdminResponse{
		ID:       admin.ID,
		Username: admin.Username,
		Token:    token,
	}
	return result, nil
}

func CreateUser(user models.Admin) (models.Admin, error) {
	if err := configs.DB.Create(&user).Error; err != nil {
		return models.Admin{}, err
	}
	return user, nil
}
