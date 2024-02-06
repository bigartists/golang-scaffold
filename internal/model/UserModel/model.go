package UserModel

import (
	"golang.org/x/crypto/bcrypt"
	"time"
)

// UserImpl 创建 Users struct
type UserImpl struct {
	Id int64 `json:"userid" gorm:"column:id; primaryKey; autoIncrement"`
	//Email 邮箱，不能为空， 必须是邮箱格式，且不能重复；
	Email string `json:"email" gorm:"column:email;unique" binding:"required,email"`
	//Username 用户名，创建自定义验证器： 长度在 6-20 之间，且不能重复，只能包含大小写字母，数字，下划线；第一个字符必须是字母；
	Username string `json:"username" gorm:"column:username;unique" binding:"usernameValid"`
	//Password 密码， 长度在 6-20 之间，只能包含字母，数字，下划线；
	Password string `json:"password" gorm:"column:password" binding:"passwordValid"`
	// admin 0 和 1
	Admin int `json:"admin" gorm:"column:admin"`
	// active 0 和 1
	Active int `json:"active" gorm:"column:active"`
	// nickname
	Nickname string `json:"nickname" gorm:"column:nickname"`
	// description
	Description string `json:"description" gorm:"column:description"`
	// avatar
	Avatar string `json:"avatar" gorm:"column:avatar"`
	// 自动维护时间

	CreateAt time.Time `json:"create_time" gorm:"column:create_time;type:datetime(0); autoCreateTime"`
	UpdateAt time.Time `json:"update_time" gorm:"column:update_time;type:datetime(0); autoUpdateTime"`
	CreateBy int64     `json:"-" gorm:"column:create_by"`
	UpdateBy int64     `json:"-" gorm:"column:update_by"`
}

func (u *UserImpl) TableName() string {
	return "user"
}

func New(attrs ...UserModelAttrFunc) *UserImpl {
	u := &UserImpl{}
	UserModelAttrFuncs(attrs).apply(u)
	return u
}

func (u *UserImpl) Mutate(attrs ...UserModelAttrFunc) *UserImpl {
	UserModelAttrFuncs(attrs).apply(u)
	return u
}

// 生成密码
func (u *UserImpl) GeneratePassword() error {
	// 使用 bcrypt 生成密码, bcrypt.DefaultCost 表示默认的加密强度，值越大加密强度越大，但是会消耗更多的资源
	pas, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(pas)
	return nil
}

// 检查密码
func (u *UserImpl) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

//func (u *UserImpl) BeforeSave() error {
//	//turn password into hash
//	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
//	if err != nil {
//		return err
//	}
//	u.Password = string(hashedPassword)
//	return nil
//}
