package vehicles

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"github.com/nadahusin/gobackend/src/database/orm/models"
	"github.com/nadahusin/gobackend/src/interfaces"
	"github.com/nadahusin/gobackend/src/libs"
)

type vehicles_ctrl struct {
	svc interfaces.VehiclesService
}

func NewCtrl(reps interfaces.VehiclesService) *vehicles_ctrl {
	return &vehicles_ctrl{svc: reps}
}

func (re *vehicles_ctrl) GetAll(w http.ResponseWriter, r *http.Request) {
	re.svc.Gets().Send(w)
}

func (re *vehicles_ctrl) Add(w http.ResponseWriter, r *http.Request) {
	var decode = schema.NewDecoder()
	var datas models.Vehicle
	var vars = fmt.Sprint(r.Context().Value("image"))

	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		libs.Respone(err.Error(), 400, true)
		return
	}
	decode.Decode(&datas, r.Form)
	re.svc.Adds(&datas, vars).Send(w)

}

func (re *vehicles_ctrl) Update(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	var datas *models.Vehicle
	err := json.NewDecoder(r.Body).Decode(&datas)
	if err != nil {
		libs.Respone(err.Error(), 400, true)
		return
	}
	re.svc.Updates(datas, vars["detail_id"]).Send(w)
}

func (re *vehicles_ctrl) Delete(w http.ResponseWriter, r *http.Request) {

	var datas *models.Vehicle
	vars := mux.Vars(r)
	re.svc.Deletes(datas, vars["detail_id"]).Send(w)
}

func (re *vehicles_ctrl) Search(w http.ResponseWriter, r *http.Request) {

	vars := strings.ToLower(r.URL.Query().Get("vars"))
	re.svc.Searching(vars).Send(w)
}

func (re *vehicles_ctrl) PopularVehicles(w http.ResponseWriter, r *http.Request) {
	re.svc.Populars().Send(w)
}
