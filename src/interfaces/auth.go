package interfaces

import (
	"github.com/nadahusin/gorent/src/database/orm/models"
	"github.com/nadahusin/gorent/src/libs"
)

type AuthService interface {
	Login(body models.User) *libs.Response
}
