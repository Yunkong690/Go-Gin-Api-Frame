package v1

import (
	"fmt"
	"personFrame/model"
	"personFrame/pkg/common"
	"personFrame/utils"
)

type UserService struct {
}

func (userService *UserService) Login(u *model.User) (err error, userInter *model.User) {
	if common.DB == nil {
		return fmt.Errorf("db not init"), nil
	}
	var user model.User
	u.Password = utils.MD5V([]byte(u.Password))
	err = common.DB.Where("username=? and password =?", u.Username, u.Password).First(&user).Error
	return err, &user
}
