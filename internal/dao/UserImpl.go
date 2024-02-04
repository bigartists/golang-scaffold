package dao

import (
	"com.github.goscaffold/internal/model/UserModel"
	"fmt"
)

var UserGetter IUserDao

func init() {
	UserGetter = NewIUserDaoImpl()
}

type IUserGetterImpl struct {
}

func (I IUserGetterImpl) FindUserByUsername(username string) (*UserModel.UserImpl, error) {
	var user UserModel.UserImpl
	err := Orm.Where("username=?", username).Find(&user).Error
	return &user, err
}

func (I IUserGetterImpl) FindUserByEmail(email string) (*UserModel.UserImpl, error) {
	var user UserModel.UserImpl
	err := Orm.Where("email=?", email).Find(&user).Error
	return &user, err
}

func (I IUserGetterImpl) CreateUser(user *UserModel.UserImpl) error {
	return Orm.Create(user).Error
}

func (I IUserGetterImpl) UpdateUser(id int, user *UserModel.UserImpl) error {
	//TODO implement me
	panic("implement me")
}

func (I IUserGetterImpl) DeleteUser(id int) error {
	//TODO implement me
	panic("implement me")
}

func (I IUserGetterImpl) FindUserById(id int64, user *UserModel.UserImpl) (*UserModel.UserImpl, error) {
	//TODO implement me
	db := Orm.Where("id=?", id).Find(user)
	if db.Error != nil || db.RowsAffected == 0 {
		return nil, fmt.Errorf("user not found, id=%d", id)
	}
	return user, nil
}

func NewIUserDaoImpl() *IUserGetterImpl {
	return &IUserGetterImpl{}
}

func (I IUserGetterImpl) FindUserAll() []*UserModel.UserImpl {
	var users []*UserModel.UserImpl
	Orm.Find(&users)
	return users
}
