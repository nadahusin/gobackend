package histories

import (
	"errors"

	"github.com/nadahusin/gobackend/src/database/orm/models"
	"gorm.io/gorm"
)

// repo untuk komunikasi ke database

type history_repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *history_repo {
	return &history_repo{db}
}

func (r *history_repo) GetAll() (*models.Histories, error) {
	var data models.Histories

	result := r.db.Find(&data)

	if result.Error != nil {
		return nil, errors.New("failed to get history")
	}

	return &data, nil
}

func (r *history_repo) AddHistory(data *models.History) (*models.History, error) {
	result := r.db.Create(data)

	if result.Error != nil {
		return nil, errors.New("failed to add history")
	}

	return data, nil
}

func (r *history_repo) DeleteHistory(data *models.History, vars string) (*models.History, error) {
	result := r.db.Where("id_history", vars).Delete(&data)
	if result.Error != nil {
		return nil, errors.New("failed to delete history")
	}

	return data, nil
}
