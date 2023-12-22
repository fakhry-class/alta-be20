package repositories

import (
	"errors"
	"fakhry/mvc/config"
	"fakhry/mvc/entities"
	"fakhry/mvc/models"
)

func InsertUser(newUser entities.UserCore) error {
	// proses mapping dari struct entities core ke model gorm
	userInputGorm := models.User{
		Name:        newUser.Name,
		Email:       newUser.Email,
		Password:    newUser.Password,
		Address:     newUser.Address,
		PhoneNumber: newUser.PhoneNumber,
		Role:        newUser.Role,
	}
	// simpan ke DB
	tx := config.DB.Create(&userInputGorm) // proses query insert
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("insert failed, row affected = 0")
	}
	return nil
}

func SelectAllUsers() ([]entities.UserCore, error) {
	var usersDataGorm []models.User
	tx := config.DB.Find(&usersDataGorm) // select * from users;
	if tx.Error != nil {
		return nil, tx.Error
	}

	// proses mapping dari struct gorm model ke struct core
	var usersDataCore []entities.UserCore
	for _, value := range usersDataGorm {
		var userCore = entities.UserCore{
			ID:          value.ID,
			Name:        value.Name,
			Email:       value.Email,
			Password:    value.Password,
			Address:     value.Address,
			PhoneNumber: value.PhoneNumber,
			Role:        value.Role,
			CreatedAt:   value.CreatedAt,
			UpdatedAt:   value.UpdatedAt,
		}
		usersDataCore = append(usersDataCore, userCore)
	}

	return usersDataCore, nil
}

func UpdateUserById(id int, userUpdate models.User) error {
	tx := config.DB.Model(&models.User{}).Where("id = ?", id).Updates(userUpdate)
	if tx.Error != nil {
		// fmt.Println("err:", tx.Error)
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("error record not found ")
	}
	return nil
}
