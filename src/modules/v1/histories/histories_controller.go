package histories

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nadahusin/gorent/src/database/orm/models"
	"github.com/nadahusin/gorent/src/interfaces"
	"github.com/nadahusin/gorent/src/libs"
)

type history_ctrl struct {
	svc interfaces.HistoriesService
}

func NewCtrl(reps interfaces.HistoriesService) *history_ctrl {
	return &history_ctrl{svc: reps}
}

func (re *history_ctrl) GetAll(w http.ResponseWriter, r *http.Request) {
	re.svc.Gets().Send(w)
}

func (re *history_ctrl) AddHistory(w http.ResponseWriter, r *http.Request) {

	var datas models.History
	err := json.NewDecoder(r.Body).Decode(&datas)
	if err != nil {
		libs.Respone(err.Error(), 400, true)
		return
	}
	re.svc.Adds(&datas).Send(w)
}

func (re *history_ctrl) DeleteHistory(w http.ResponseWriter, r *http.Request) {
	var datas *models.History
	vars := mux.Vars(r)
	re.svc.Deletes(datas, vars["id_history"]).Send(w)
}
