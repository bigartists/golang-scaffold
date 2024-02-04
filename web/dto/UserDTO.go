package dto

import "time"

// 输入对象

type (
	SimpleUserReq struct {
		Id int `uri:"id" binding:"required,min=1"`
	}
)

// 以下是输出对象
type SimpleUserInfo struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	City string `json:"city"`
}

func (s SimpleUserInfo) Error() string {
	//TODO implement me
	panic("implement me")
}

type UserLog struct {
	Id   int       `json:"id"`
	Log  string    `json:"log"`
	Date time.Time `json:"date"`
}

type UserInfo struct {
	Id    int        `json:"id"`
	Name  string     `json:"name"`
	City  string     `json:"city"`
	Phone string     `json:"phone"`
	Logs  []*UserLog `json:"logs"`
}
