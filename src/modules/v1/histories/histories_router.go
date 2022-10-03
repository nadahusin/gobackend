package histories

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// akan memangil semua method
// inisialisasi endpoint

func New(rt *mux.Router, db *gorm.DB) {
	route := rt.PathPrefix("/history").Subrouter()

	repo := NewRepo(db)
	svc := NewService(repo)
	ctrl := NewCtrl(svc)

	route.HandleFunc("/", ctrl.GetAll).Methods("GET")
	route.HandleFunc("/", ctrl.AddHistory).Methods("POST")
	route.HandleFunc("/{id_history}", ctrl.DeleteHistory).Methods("DELETE")

}
