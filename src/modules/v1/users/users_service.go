package users

import (
	"github.com/nadahusin/gorent/src/database/orm/models"
	"github.com/nadahusin/gorent/src/interfaces"
	"github.com/nadahusin/gorent/src/libs"
)

// berinteraksi dengan repo dan controller
// berisi logic bisnis

type users_service struct {
	repo interfaces.UsersRepo
}

func NewService(reps interfaces.UsersRepo) *users_service {
	return &users_service{reps}
}

func (r users_service) Gets() *libs.Response {
	data, err := r.repo.GetAllUser()
	if err != nil {
		return libs.Respone(err.Error(), 400, true)
	}
	return libs.Respone(data, 200, false)
}

func (r users_service) GetsUname(username string) *libs.Response {
	data, err := r.repo.GetByUsername(username)
	if err != nil {
		return libs.Respone(err.Error(), 400, true)
	}
	return libs.Respone(data, 200, false)
}

func (r users_service) Adds(data *models.User) *libs.Response {
	if check := r.repo.UserExsist(data.Username); check {
		return libs.Respone("username sudah terdaftar", 400, true)
	}

	hassPassword, err := libs.HashingPassword(data.Password)
	if err != nil {
		return libs.Respone(err.Error(), 400, true)
	}

	data.Password = hassPassword
	result, err := r.repo.AddUser(data)
	if err != nil {
		return libs.Respone(err.Error(), 400, true)
	}

	return libs.Respone(result, 200, false)
}

func (r users_service) Updates(data *models.User, vars string) *libs.Response {
	if check := r.repo.UserExsist(vars); !check {
		return libs.Respone("username not found", 404, true)
	}

	hashPassword, err := libs.HashingPassword(data.Password)
	if err != nil {
		return libs.Respone(err.Error(), 400, true)
	}

	data.Password = hashPassword
	result, err := r.repo.UpdateUser(data, vars)
	if err != nil {
		return libs.Respone(err.Error(), 400, true)
	}
	return libs.Respone(result, 200, false)
}

func (r users_service) Deletes(data *models.User, vars string) *libs.Response {
	if check := r.repo.UserExsist(vars); !check {
		return libs.Respone("username not found", 404, true)
	}
	data, err := r.repo.DeleteUser(data, vars)
	if err != nil {
		return libs.Respone(err.Error(), 400, true)
	}
	return libs.Respone(data, 200, false)
}
