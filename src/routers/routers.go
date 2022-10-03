package routers

import (
	"errors"

	"github.com/gorilla/mux"
	"github.com/nadahusin/gobackend/src/database/orm"
	"github.com/nadahusin/gobackend/src/modules/v1/auth"
	"github.com/nadahusin/gobackend/src/modules/v1/histories"
	"github.com/nadahusin/gobackend/src/modules/v1/users"
	"github.com/nadahusin/gobackend/src/modules/v1/vehicles"
)

func New() (*mux.Router, error) {

	mainRoute := mux.NewRouter()

	db, err := orm.New()
	if err != nil {
		return nil, errors.New("failed to init database")
	}

	vehicles.New(mainRoute, db)
	users.New(mainRoute, db)
	histories.New(mainRoute, db)
	auth.New(mainRoute, db)
	return mainRoute, nil
}
