package vehicles

import (
	"fmt"
	"testing"

	"github.com/nadahusin/gobackend/src/database/orm/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGets(t *testing.T) {
	repo := RepoMock{mock.Mock{}}
	service := NewService(&repo)

	var dataMock = models.Vehicles{
		{DetailId: 1, Goods: "Fixie Green"},
		{DetailId: 2, Goods: "Fixie Black"},
	}

	repo.mock.On("GetAll").Return(&dataMock, nil)
	data := service.Gets()
	fmt.Println(data)

	result := data.Data.(*models.Vehicles)
	for i, v := range *result {
		assert.Equal(t, dataMock[i].DetailId, v.DetailId, "id must be like data mock")
	}
}
