package vehicles

import (
	"github.com/gorilla/mux"
	"github.com/nadahusin/gorent/src/middleware"
	"gorm.io/gorm"
)

// akan memangil semua method
// inisialisasi endpoint

func New(rt *mux.Router, db *gorm.DB) {
	route := rt.PathPrefix("/vehicles").Subrouter()

	repo := NewRepo(db)
	svc := NewService(repo)
	ctrl := NewCtrl(svc)

	route.HandleFunc("/", middleware.CheckAuth(middleware.CheckRole(ctrl.GetAll))).Methods("GET")
	route.HandleFunc("/", middleware.CheckAuth(middleware.FileUpload(ctrl.Add))).Methods("POST")
	route.HandleFunc("/{detail_id}", ctrl.Update).Methods("PUT")
	route.HandleFunc("/{detail_id}", ctrl.Delete).Methods("DELETE")
	route.HandleFunc("{vars}", ctrl.Search).Methods("GET")

	route.HandleFunc("/", ctrl.PopularVehicles).Methods("GET")
}
