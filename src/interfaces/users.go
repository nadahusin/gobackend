package interfaces

import (
	"github.com/nadahusin/gorent/src/database/orm/models"
	"github.com/nadahusin/gorent/src/libs"
)

type UsersRepo interface {
	GetAllUser() (*models.Users, error)
	GetByUsername(username string) (*models.User, error)
	UserExsist(username string) bool
	AddUser(data *models.User) (*models.User, error)
	UpdateUser(data *models.User, vars string) (*models.User, error)
	DeleteUser(data *models.User, vars string) (*models.User, error)
}

type UsersService interface {
	Gets() *libs.Response
	GetsUname(username string) *libs.Response
	Adds(data *models.User) *libs.Response
	Updates(data *models.User, vars string) *libs.Response
	Deletes(data *models.User, vars string) *libs.Response
}
