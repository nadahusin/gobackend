package vehicles

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/nadahusin/gobackend/src/database/orm/models"
	"github.com/nadahusin/gobackend/src/libs"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var repos = RepoMock{mock.Mock{}}
var service = NewService(&repos)
var ctrl = NewCtrl(service)

var dataMock = models.Vehicles{
	{DetailId: 1, Goods: "Fixie Green"},
	{DetailId: 2, Goods: "Fixie Black"},
}

func TestCtrlGetAll(t *testing.T) {
	w := httptest.NewRecorder()
	mux := mux.NewRouter()

	repos.mock.On("GetAll").Return(&dataMock, nil)

	mux.HandleFunc("/test/vehicles", ctrl.GetAll)

	mux.ServeHTTP(w, httptest.NewRequest("GET", "/test/vehicles", nil))

	var vehicle models.Vehicles
	respone := libs.Response{
		Data: &vehicle,
	}

	if err := json.Unmarshal(w.Body.Bytes(), &respone); err != nil {
		t.Fatal(err)
	}
	// fmt.Println(respone)

	assert.Equal(t, 200, w.Code, "status must be code 200")
	assert.False(t, respone.IsError, "isError must be false")
}
