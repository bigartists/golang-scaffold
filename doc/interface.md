# interface struct instance 三者之间的关系

```go

package repos

type IUserRepo interface {
    FindById(*models.UserModel) error
    FindByName(name string) *models.UserModel
    SaveUser(*models.UserModel) error
    UpdateUser(*models.UserModel) error
    DeleteUser(*models.UserModel) error
}

type IUserLogRepo interface {
    FindByName(name string) *models.UserLogsModel
    SaveLog(model *models.UserLogsModel) error
}

package services


type UserLoginService struct {
    userRepo repos.IUserRepo
}

func (this *UserLoginService) UserLogin(userName string, userPwd string) (string, error) {
    user := this.userRepo.FindByName(userName)
    if user.UserId > 0 { // 有这个用户
        // md5
    if user.UserPwd == utils.MD5(userPwd) {
        // 用户登录安全判断，比如IP校验
        // 根据用户登录次数，给予奖励。也都写在这里的；
        // todo 思考题 当用户登录成功之后，登录日志的动作要怎么做，要写在哪里？ 答案 调用聚合就成；
        return "1000200", nil
    } else {
    return "1000400", fmt.Errorf("密码不正确")
    }
    } else {
        return "1000404", fmt.Errorf("用户不存在")
    }
}

package repos

type UserRepo struct {
	// 添加仓储所需的任何依赖项或配置
}

func (r *UserRepo) FindById(user *models.UserModel) error {
	// FindById 方法的具体实现
	return nil
}

func (r *UserRepo) FindByName(name string) *models.UserModel {
	// FindByName 方法的具体实现
	return nil
}

func (r *UserRepo) SaveUser(user *models.UserModel) error {
	// SaveUser 方法的具体实现
	return nil
}

func (r *UserRepo) UpdateUser(user *models.UserModel) error {
	// UpdateUser 方法的具体实现
	return nil
}

func (r *UserRepo) DeleteUser(user *models.UserModel) error {
	// DeleteUser 方法的具体实现
	return nil
}


package services

type UserLoginService struct {
	userRepo repos.IUserRepo
}

func (s *UserLoginService) UserLogin(userName string, userPwd string) (string, error) {
	user := s.userRepo.FindByName(userName)
	// 其余代码...
}

// 示例用法:
func main() {
	repo := &repos.UserRepo{} // 某个具体实现， 这里可以有多个实现； 通过这种方式就实现了解耦，实现了依赖倒置， 同时也区分了不同的实现
	service := &services.UserLoginService{
		userRepo: repo, // todo 在真正需要调用的时候，将某个具体的实现赋值给 构造函数；
	}
	// 使用 UserLoginService...
}
```