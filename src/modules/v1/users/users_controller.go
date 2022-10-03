package users

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nadahusin/rental/src/database/orm/models"
	"github.com/nadahusin/rental/src/interfaces"
	"github.com/nadahusin/rental/src/libs"
)

// berinteraksi dengan service dan router
// untuk mengahandle request yang masuk

type users_ctrl struct {
	svc interfaces.UsersService
}

func NewCtrl(reps interfaces.UsersService) *users_ctrl {
	return &users_ctrl{svc: reps}
}

func (re *users_ctrl) GetAllUser(w http.ResponseWriter, r *http.Request) {
	claims_users := r.Context().Value("username")
	result := re.svc.GetsUname(claims_users.(string))
	result.Send(w)
}

func (re *users_ctrl) AddUser(w http.ResponseWriter, r *http.Request) {
	var datas models.User
	err := json.NewDecoder(r.Body).Decode(&datas)
	if err != nil {
		libs.Respone(err, 400, true)
	}
	re.svc.Adds(&datas).Send(w)
}

func (re *users_ctrl) UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := r.Context().Value("id").(string)

	var datas *models.User
	err := json.NewDecoder(r.Body).Decode(&datas)
	if err != nil {
		libs.Respone(err.Error(), 400, true)
	}
	re.svc.Updates(datas, vars).Send(w)
}

func (re *users_ctrl) DeleteUser(w http.ResponseWriter, r *http.Request) {
	var datas *models.User
	vars := mux.Vars(r)
	re.svc.Deletes(datas, vars["id"]).Send(w)
}
