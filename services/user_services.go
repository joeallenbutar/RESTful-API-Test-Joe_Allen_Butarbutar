package database

import (
	"RESTful-API-Test-Joe_Allen_Butarbutar/models"
	"errors"
	"gorm.io/gorm"
	"fmt"
	"strconv"
)

func CreateUser(db *gorm.DB, user *models.User) error {
	if err := db.Save(&user).Error; err != nil {
		return err
	}

	return nil
}

func GetAllUsers(db *gorm.DB, pagination *models.Pagination) ([]models.User, error) {
	users := []models.User{}

	reqPage, _ := strconv.Atoi(pagination.Page)
	reqLimit, _ := strconv.Atoi(pagination.Limit)
	offset := (reqPage - 1) * reqLimit
	query := db.Select("users.*").Group("users.id")

	if reqPage > 0 || reqLimit > 0 || offset > 0{
		fmt.Println("query panjang")
		queryExt := query.Where(pagination.Filter).Limit(reqLimit).Order(pagination.Sort).Offset(offset)
		if err := queryExt.Find(&users).Error; err != nil {
			return users, err
		}

	}else if pagination.Filter != "" {
		fmt.Println("query hanya filter yang ada")
		fmt.Println("pagination.Filter")
		queryExt := query.Where(pagination.Filter)
		if err := queryExt.Find(&users).Error; err != nil {
			return users, err
		}

	}else if pagination.Sort != ""{
		fmt.Println("query hanya sort yang ada")
		queryExt := query.Order(pagination.Sort)
		if err := queryExt.Find(&users).Error; err != nil {
			return users, err
		}

	}else{
		fmt.Println("query semua kosong")
		queryExt := query.Order(pagination.Sort)
		if err := queryExt.Find(&users).Error; err != nil {
			return users, err
		}
	}

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