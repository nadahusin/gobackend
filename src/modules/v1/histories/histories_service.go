package histories

import (
	"github.com/nadahusin/rental/src/database/orm/models"
	"github.com/nadahusin/rental/src/interfaces"
	"github.com/nadahusin/rental/src/libs"
)

// berinteraksi dengan repo dan controller
// berisi logic bisnis

type history_service struct {
	repo interfaces.HistoriesRepo
}

func NewService(reps interfaces.HistoriesRepo) *history_service {
	return &history_service{reps}
}

func (w history_service) Gets() *libs.Response {
	_, err := w.repo.GetAll()
	if err != nil {
		return libs.Respone(err.Error(), 400, true)
	}
	return libs.Respone("success to get history", 200, false)
}

func (w *history_service) Adds(data *models.History) *libs.Response {
	_, err := w.repo.AddHistory(data)
	if err != nil {
		return libs.Respone(err.Error(), 400, true)
	}
	return libs.Respone("success to add history", 200, false)
}

func (w *history_service) Deletes(data *models.History, vars string) *libs.Response {
	_, err := w.repo.DeleteHistory(data, vars)
	if err != nil {
		return libs.Respone(err.Error(), 400, true)
	}
	return libs.Respone("success to delete history", 200, false)
}
