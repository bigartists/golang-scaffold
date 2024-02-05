package service

import (
	"com.github.goscaffold/internal/model/UserModel"
	"com.github.goscaffold/web/result"
)

type IUser interface {
	GetUserList() []*UserModel.UserImpl
	GetUserDetail(id int64) *result.ErrorResult
	CreateUser(user *UserModel.UserImpl) *result.ErrorResult
	UpdateUser(id int, user *UserModel.UserImpl) *result.ErrorResult
	DeleteUser(id int) *result.ErrorResult
	SignIn(username string, password string) (*UserModel.UserImpl, error)
	SignUp(impl *UserModel.UserImpl) error
}
