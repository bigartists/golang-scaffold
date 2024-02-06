package UserModel

type UserModelAttrFunc func(u *UserImpl)

type UserModelAttrFuncs []UserModelAttrFunc

func WithUserId(id int64) UserModelAttrFunc {
	return func(u *UserImpl) {
		u.Id = id
	}
}

func WithUsername(name string) UserModelAttrFunc {
	return func(u *UserImpl) {
		u.Username = name
	}
}

func WithPassword(pwd string) UserModelAttrFunc {
	return func(u *UserImpl) {
		u.Password = pwd
	}
}

func WithEmail(email string) UserModelAttrFunc {
	return func(u *UserImpl) {
		u.Email = email
	}
}

func (this UserModelAttrFuncs) apply(u *UserImpl) {
	for _, f := range this {
		f(u)
	}
}
