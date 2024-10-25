package service

import (
	"crud/database"
	"crud/models"
	"errors"
	"strings"
)

func CreateUser(user *models.User) error {
	if err := database.DB.Create(user).Error; err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			return errors.New("error")
		}
		return err
	}
	return nil
}

func GetAllUsers() ([]models.User, error) {
	var users []models.User
	err := database.DB.Find(&users).Error
	return users, err
}

func GetUserById(id uint) (models.User, error) {
	var user models.User
	err := database.DB.First(&user, id).Error
	return user, err
}

func UpdateUser(user *models.User) error {
	if err := database.DB.Save(user).Error; err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			return errors.New("email ou cpf ja existe")
		}
		return err
	}
	return nil
}

func DeleteUser(id uint) error {
	return database.DB.Delete(&models.User{}, id).Error
}
