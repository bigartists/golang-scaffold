package service

import (
	"com.github.goscaffold/config"
	"com.github.goscaffold/internal/dao"
	"com.github.goscaffold/internal/model/UserModel"
	"com.github.goscaffold/pkg/utils"
	"com.github.goscaffold/web/result"
	"fmt"
	"time"
)

var UserServiceGetter IUser

func init() {
	UserServiceGetter = NewIUserGetterImpl()
}

type IUserServiceGetterImpl struct {
}

func (this *IUserServiceGetterImpl) SignIn(username string, password string) *result.ErrorResult {
	user, err := dao.UserGetter.FindUserByUsername(username)
	if err != nil {
		return result.Result(nil, err)
	}
	if user.Password != password {
		return result.Result(nil, fmt.Sprintf("用户密码错误, username=%s", username))
	}
	// 生成 token
	prikey := []byte(config.SysYamlconfig.Jwt.PrivateKey)
	curTime := time.Now().Add(time.Second * 60 * 60 * 24)
	token, _ := utils.GenerateToken(user.Id, prikey, curTime)

	target := map[string]interface{}{
		"user":  user,
		"token": token,
	}

	return result.Result(target, nil)
}

func (this *IUserServiceGetterImpl) SignUp(user *UserModel.UserImpl) error {
	return dao.UserGetter.CreateUser(user)
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
