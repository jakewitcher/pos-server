package users

import (
	"github.com/jakewitcher/pos-server/graph/model"
	"strconv"
)

type UserEntity struct {
	Id         int64  `json:"id"`
	EmployeeId int64  `json:"employee_id"`
	Username   string `json:"username"`
	Password   string `json:"password"`
}

func (u *UserEntity) ToDTO() *model.User {
	return &model.User{
		ID:         strconv.FormatInt(u.Id, 10),
		EmployeeID: strconv.FormatInt(u.EmployeeId, 10),
		Username:   u.Username,
	}
}
