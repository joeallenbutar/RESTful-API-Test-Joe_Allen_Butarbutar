package database

import (
	"RESTful-API-Test-Joe_Allen_Butarbutar/models"
	"errors"
	"gorm.io/gorm"
	_"fmt"
)

func CreateUser(db *gorm.DB, user *models.User) error {
	if err := db.Save(&user).Error; err != nil {
		return err
	}

	return nil
}

func GetAllUsers(db *gorm.DB, pagination *models.Pagination) ([]models.User, error) {
	users := []models.User{}
	offset := (pagination.Page - 1) * pagination.Limit
	// query := db.Select("users.*").Group("users.id").Limit(pagination.Limit).Order(pagination.Sort).Offset(offset)

	if pagination.Page > 0 || pagination.Limit > 0{
		query := db.Select("users.*").Group("users.id").Where(pagination.Filter).Limit(pagination.Limit).Order(pagination.Sort).Offset(offset)
		if err := query.Find(&users).Error; err != nil {
			return users, err
		}
	}else{
		query := db.Select("users.*").Group("users.id").Order(pagination.Sort)
		if err := query.Find(&users).Error; err != nil {
			return users, err
		}
	}
	// fmt.Println(query)
	// if err := query.Find(&users).Error; err != nil {
	// 	return users, err
	// }

	return users, nil
}

func GetUserByID(id string, db *gorm.DB) (models.User, bool, error) {
	user := models.User{}

	query := db.Select("users.*")
	query = query.Group("users.id")
	err := query.Where("users.id = ?", id).First(&user).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return user, false, err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return user, false, nil
	}

	return user, true, nil
}

func UpdateUser(db *gorm.DB, user *models.User) error {
	if err := db.Save(&user).Error; err != nil {
		return err
	}

	return nil
}

func DeleteUser(id string, db *gorm.DB) error {
	var user models.User
	if err := db.Where("id = ? ", id).Delete(&user).Error; err != nil {
		return err
	}

	return nil
}