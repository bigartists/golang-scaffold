package dao

import "com.github.goscaffold/internal/model/UserModel"

type IUserDao interface {
	FindUserAll() []*UserModel.UserImpl
	FindUserById(id int64, user *UserModel.UserImpl) (*UserModel.UserImpl, error)
	FindUserByUsername(username string) (*UserModel.UserImpl, error)
	FindUserByEmail(email string) (*UserModel.UserImpl, error)
	CreateUser(user *UserModel.UserImpl) error
	UpdateUser(id int, user *UserModel.UserImpl) error
	DeleteUser(id int) error
}
