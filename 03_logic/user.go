package logic

import (
	"errors"
	model "gorm-example/01_model"
	"gorm.io/gorm"
)

type UserLogic interface {

	// 登录
	Login(db *gorm.DB, account, password string) (token string, err error)

	// 其他业务逻辑方法....
}

type userLogic struct {
}

// 接口controller层直接调用 NewUserLogic
// 完成各种业务操作
func NewUserLogic() UserLogic {
	return &userLogic{}
}

func (u *userLogic) Login(db *gorm.DB, account, password string) (token string, err error) {

	user := new(model.User)
	if err := user.First(db, &model.FilterUser{
		Account: account,
	}, "id,account,password"); err != nil {
		// 错误返回可以再次进行封装
		return "", errors.New(err.Error())
	}

	// TODO
	// 处理密码散列校验
	// 生成token

	return "", err
}
