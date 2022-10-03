package interfaces

import (
	"github.com/nadahusin/rental/src/database/orm/models"
	"github.com/nadahusin/rental/src/libs"
)

type HistoriesRepo interface {
	GetAll() (*models.Histories, error)
	AddHistory(data *models.History) (*models.History, error)
	DeleteHistory(data *models.History, vars string) (*models.History, error)
}

type HistoriesService interface {
	Gets() *libs.Response
	Adds(data *models.History) *libs.Response
	Deletes(data *models.History, vars string) *libs.Response
}
