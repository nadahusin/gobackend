package interfaces

import (
	"github.com/nadahusin/rental/src/database/orm/models"
	"github.com/nadahusin/rental/src/libs"
)

type VehiclesRepo interface {
	GetAll() (*models.Vehicles, error)
	Add(data *models.Vehicle) (*models.Vehicle, error)
	Delete(data *models.Vehicle, vars string) (*models.Vehicle, error)
	Update(data *models.Vehicle, vars string) (*models.Vehicle, error)
	Search(vars string) (*models.Vehicles, error)
	PopularVehicles() (*models.Vehicles, error)
}

type VehiclesService interface {
	Gets() *libs.Response
	Adds(data *models.Vehicle, vars string) *libs.Response
	Deletes(data *models.Vehicle, vars string) *libs.Response
	Updates(data *models.Vehicle, vars string) *libs.Response
	Searching(vars string) *libs.Response
	Populars() *libs.Response
}
