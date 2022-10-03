package interfaces

import (
	"github.com/nadahusin/gobackend/src/database/orm/models"
	"github.com/nadahusin/gobackend/src/libs"
)

type AuthService interface {
	Login(body models.User) *libs.Response
}
