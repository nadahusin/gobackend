package vehicles

import (
	"github.com/nadahusin/gobackend/src/database/orm/models"
	"github.com/nadahusin/gobackend/src/interfaces"
	"github.com/nadahusin/gobackend/src/libs"
)

// berinteraksi dengan repo dan controller
// berisi logic bisnis

type vehicles_service struct {
	repo interfaces.VehiclesRepo
}

func NewService(reps interfaces.VehiclesRepo) *vehicles_service {
	return &vehicles_service{reps}
}

func (w *vehicles_service) Gets() *libs.Response {
	data, err := w.repo.GetAll()
	if err != nil {
		return libs.Respone(err.Error(), 400, true)
	}
	return libs.Respone(data, 200, false)

}

func (w *vehicles_service) Adds(data *models.Vehicle, vars string) *libs.Response {
	data.Image = vars
	data, err := w.repo.Add(data)
	if err != nil {
		return libs.Respone(err.Error(), 400, true)
	}
	return libs.Respone(data, 200, false)
}

func (w *vehicles_service) Updates(data *models.Vehicle, vars string) *libs.Response {
	data, err := w.repo.Update(data, vars)

	if err != nil {
		return libs.Respone(err.Error(), 400, true)
	}
	return libs.Respone(data, 200, false)
}

func (w *vehicles_service) Deletes(data *models.Vehicle, vars string) *libs.Response {
	data, err := w.repo.Delete(data, vars)

	if err != nil {
		return libs.Respone(err.Error(), 400, true)
	}
	return libs.Respone(data, 200, false)
}

func (w *vehicles_service) Searching(vars string) *libs.Response {
	data, err := w.repo.Search(vars)

	if err != nil {
		return libs.Respone(err.Error(), 400, true)
	}
	return libs.Respone(data, 200, false)
}

func (w *vehicles_service) Populars() *libs.Response {
	data, err := w.repo.PopularVehicles()
	if err != nil {
		return libs.Respone(err.Error(), 400, true)
	}
	return libs.Respone(data, 200, false)
}
