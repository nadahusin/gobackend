package vehicles

import (
	"github.com/nadahusin/rental/src/database/orm/models"
	"github.com/stretchr/testify/mock"
)

type RepoMock struct {
	mock mock.Mock
}

func (m *RepoMock) GetAll() (*models.Vehicles, error) {
	args := m.mock.Called()
	return args.Get(0).(*models.Vehicles), nil
}

func (m *RepoMock) Add(data *models.Vehicle) (*models.Vehicle, error) {
	args := m.mock.Called(data)
	return args.Get(0).(*models.Vehicle), nil
}

func (m *RepoMock) Delete(data *models.Vehicle, vars string) (*models.Vehicle, error) {
	args := m.mock.Called(data, vars)
	return args.Get(0).(*models.Vehicle), nil
}

func (m *RepoMock) Update(data *models.Vehicle, vars string) (*models.Vehicle, error) {
	args := m.mock.Called(data, vars)
	return args.Get(0).(*models.Vehicle), nil
}

func (m *RepoMock) Search(vars string) (*models.Vehicles, error) {
	args := m.mock.Called()
	return args.Get(0).(*models.Vehicles), nil
}

func (m *RepoMock) PopularVehicles() (*models.Vehicles, error) {
	args := m.mock.Called()
	return args.Get(0).(*models.Vehicles), nil
}
