package users

import (
	"errors"

	"github.com/nadahusin/gobackend/src/database/orm/models"
	"gorm.io/gorm"
)

// repo untuk komunikasi ke database

type users_repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *users_repo {
	return &users_repo{db}
}

func (r *users_repo) GetAllUser() (*models.Users, error) {
	var data models.Users

	result := r.db.Find(&data)

	if result.Error != nil {
		return nil, errors.New("failed to get data")
	}

	return &data, nil
}

func (r *users_repo) GetByUsername(username string) (*models.User, error) {
	var data models.User

	result := r.db.First(&data, "username = ?", username)

	if result.Error != nil {
		return nil, errors.New("failed to get data")
	}
	return &data, nil
}

func (r *users_repo) UserExsist(username string) bool {
	var data models.User
	result := r.db.First(&data, "username = ?", username)

	if result.Error != nil {
		return false
	}
	return true
}

func (r *users_repo) AddUser(data *models.User) (*models.User, error) {
	result := r.db.Create(data)

	if result.Error != nil {
		return nil, errors.New("failed to add data")
	}

	return data, nil
}

func (r *users_repo) UpdateUser(data *models.User, vars string) (*models.User, error) {
	result := r.db.Model(&data).Where("id = ?", vars).Updates(&data)

	if result.Error != nil {
		return nil, errors.New("failed to update data")
	}

	return data, nil
}

func (r *users_repo) DeleteUser(data *models.User, vars string) (*models.User, error) {
	result := r.db.Where("id", vars).Delete(&data)
	if result.Error != nil {
		return nil, errors.New("failed to delete data")
	}

	return data, nil
}
