package service

import (
	"com.github.goscaffold/internal/dao"
	"com.github.goscaffold/internal/model/UserModel"
	"com.github.goscaffold/web/result"
	"fmt"
)

var UserServiceGetter IUser

func init() {
	UserServiceGetter = NewIUserGetterImpl()
}

type IUserServiceGetterImpl struct {
}

func (this *IUserServiceGetterImpl) SignIn(username string, password string) (*UserModel.UserImpl, error) {
	user, err := dao.UserGetter.FindUserByUsername(username)
	if err != nil {
		return nil, err
	}
	//if user.Password != password {
	//	err = fmt.Errorf("用户名%s或密码错误", username)
	//	return nil, err
	//}
	// 校验密码
	if !user.CheckPassword(password) {
		err = fmt.Errorf("用户名%s或密码错误", username)
		return nil, err
	}
	return user, nil
}

func (this *IUserServiceGetterImpl) SignUp(email string, username string, password string) error {
	//return dao.UserGetter.CreateUser(user)

	if _, err := dao.UserGetter.FindUserByUsername(username); err != nil {
		return fmt.Errorf("用户名%s已存在", username)
	}
	if _, err := dao.UserGetter.FindUserByEmail(email); err != nil {
		return fmt.Errorf("邮箱%s已存在", email)
	}

	user := UserModel.New(UserModel.WithEmail(email), UserModel.WithUsername(username), UserModel.WithPassword(password))
	err := user.GeneratePassword()
	if err != nil {
		return fmt.Errorf("密码加密失败")
	}
	err = dao.UserGetter.CreateUser(user)

	if err != nil {
		return fmt.Errorf("用户注册失败")
	}

	return nil
}

func (this *IUserServiceGetterImpl) GetUserList() []*UserModel.UserImpl {
	users := dao.UserGetter.FindUserAll()
	return users
}

func (this *IUserServiceGetterImpl) GetUserDetail(id int64) *result.ErrorResult {
	//TODO implement me
	user := UserModel.New()
	_, err := dao.UserGetter.FindUserById(id, user)
	if err != nil {
		return result.Result(nil, err)
	}
	return result.Result(user, nil)
}

func (this *IUserServiceGetterImpl) CreateUser(user *UserModel.UserImpl) *result.ErrorResult {
	//TODO implement me
	panic("implement me")
}

func (this *IUserServiceGetterImpl) UpdateUser(id int, user *UserModel.UserImpl) *result.ErrorResult {
	//TODO implement me
	panic("implement me")
}

func (this *IUserServiceGetterImpl) DeleteUser(id int) *result.ErrorResult {
	//TODO implement me
	panic("implement me")
}

func NewIUserGetterImpl() *IUserServiceGetterImpl {
	return &IUserServiceGetterImpl{}
}

//
//// 创建用户
//func (this *IUserGetterImpl) CreateUser(user *UserModel.UserModelImpl) *result.ErrorResult {
//	db := dbs.Orm.Create(user)
//	if db.Error != nil {
//		return result.Result(nil, db.Error)
//	}
//	return result.Result(user, nil)
//}
//

//
//// 更新用户
//func (this *IUserGetterImpl) UpdateUser(id int, user *UserModel.UserModelImpl) *result.ErrorResult {
//	db := dbs.Orm.Where("id=?", id).Updates(user)
//	if db.Error != nil {
//		return result.Result(nil, db.Error)
//	}
//	return result.Result(user, nil)
//}
//
//// 删除用户
//func (this *IUserGetterImpl) DeleteUser(id int) *result.ErrorResult {
//	user := UserModel.New()
//	db := dbs.Orm.Where("id=?", id).Delete(user)
//	if db.Error != nil {
//		return result.Result(nil, db.Error)
//	}
//	return result.Result(user, nil)
//}
