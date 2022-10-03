package interfaces

import (
	"github.com/nadahusin/rental/src/database/orm/models"
	"github.com/nadahusin/rental/src/libs"
)

type AuthService interface {
	Login(body models.User) *libs.Response
}
