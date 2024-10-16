package models

import (
	"log"
	"tugas-4/config"
	"tugas-4/entities"
)

func CreateUser(user *entities.User) error {
	return config.DB.Create(user).Error
}

// Get all users
func GetAllUsers() ([]entities.User, error) {
	var users []entities.User
	err := config.DB.Find(&users).Error
	return users, err
}

func FindUserByID(id uint) (*entities.User, error) {
	var user entities.User
	err := config.DB.First(&user, id).Error
	return &user, err
}

func FindUserByUsernameAndPassword(username string) (*entities.User, error) {
	var user entities.User
	log.Println(username)
	// Cari user berdasarkan username
	err := config.DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, err
}

func UpdateUser(user *entities.User) error {
	return config.DB.Save(user).Error
}

func DeleteUser(id uint) error {
	var user *entities.User
	return config.DB.Delete(&user, id).Error
}
