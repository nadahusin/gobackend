package vehicles

import (
	"errors"
	"strings"

	"github.com/nadahusin/gobackend/src/database/orm/models"
	"gorm.io/gorm"
)

// repo untuk komunikasi ke database

type vehicles_repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *vehicles_repo {
	return &vehicles_repo{db}
}

func (r *vehicles_repo) GetAll() (*models.Vehicles, error) {
	var data models.Vehicles

	result := r.db.Find(&data)

	if result.Error != nil {
		return nil, errors.New("failed to get data")
	}

	return &data, nil
}

func (r *vehicles_repo) Add(data *models.Vehicle) (*models.Vehicle, error) {
	result := r.db.Create(data)

	if result.Error != nil {
		return nil, errors.New("failed to add data")
	}

	return data, nil
}

func (r *vehicles_repo) Update(data *models.Vehicle, vars string) (*models.Vehicle, error) {
	result := r.db.Model(&data).Where("detail_id = ?", vars).Updates(&data)

	if result.Error != nil {
		return nil, errors.New("failed to update data")
	}

	return data, nil
}

func (r *vehicles_repo) Delete(data *models.Vehicle, vars string) (*models.Vehicle, error) {
	result := r.db.Where("detail_id", vars).Delete(&data)
	if result.Error != nil {
		return nil, errors.New("failed to delete data")
	}

	return data, nil
}

func (r *vehicles_repo) Search(vars string) (*models.Vehicles, error) {
	var data models.Vehicles

	x := strings.ToLower((vars))
	result := r.db.Where("LOWER(goods) LIKE ?", "%"+x+"%").Find(&data)
	if result.Error != nil {
		return nil, errors.New("failed to search data")
	}

	return &data, nil
}

func (r *vehicles_repo) PopularVehicles() (*models.Vehicles, error) {
	var data models.Vehicles

	result := r.db.Order("order DESC").Find(&data)
	if result.Error != nil {
		return nil, errors.New("failed to get data")
	}

	return &data, nil
}
