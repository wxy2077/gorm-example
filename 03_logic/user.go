package logic

import (
	"errors"
	model "gorm-example/01_model"
	"gorm-example/utils"
	"gorm.io/gorm"
)

type UserLogic interface {

	// Login 登录
	Login(db *gorm.DB, account, password string) (token string, err error)

	// UserDepList 连接查询用户所在部门
	UserDepList(db *gorm.DB, page int64) (list []*UserList, pagination *utils.Pagination)

	// PreloadUserDep 预加载查询出所有的部门
	PreloadUserDep(db *gorm.DB, page int64) (list []*model.User, pagination *utils.Pagination)

	// 同上
	PreloadUserDeps(db *gorm.DB, page int64) (list []*model.User, pagination *utils.Pagination)

	// 原生SQL查询

	// 其他业务逻辑方法....
}

type userLogic struct {
}

// NewUserLogic 接口controller层直接调用
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

type UserList struct {
	*model.User
	Title string `json:"title"`
}

func (u *userLogic) UserDepList(db *gorm.DB, page int64) (list []*UserList, pagination *utils.Pagination) {

	db = db.Debug().Model(&model.User{}).Select("d.title,users.*").
		Joins("LEFT JOIN department_users du ON du.user_id = users.id").
		Joins("LEFT JOIN departments d ON d.dep_id = du.dep_id")

	pagination = utils.Paginate(&utils.Param{
		DB:      db,
		Page:    page,
		Limit:   15,
		OrderBy: []string{"id desc"},
	}, &list)

	return list, pagination
}

func (u *userLogic) PreloadUserDep(db *gorm.DB, page int64) (list []*model.User, pagination *utils.Pagination) {
	// Debug() 显示SQL语句方便调试
	db = db.Debug().Model(&model.User{}).
		Preload("DepartmentUser", func(tx *gorm.DB) *gorm.DB {
			return tx.Preload("Department", func(tx *gorm.DB) *gorm.DB {
				return tx.Select("id,title")
			})
		})

	pagination = utils.Paginate(&utils.Param{
		DB:      db,
		Page:    page,
		Limit:   15,
		OrderBy: []string{"id desc"},
	}, &list)

	return list, pagination
}

func (u *userLogic) PreloadUserDeps(db *gorm.DB, page int64) (list []*model.User, pagination *utils.Pagination) {

	db = db.Debug().Model(&model.User{}).Preload("Department")

	pagination = utils.Paginate(&utils.Param{
		DB:      db,
		Page:    page,
		Limit:   15,
		OrderBy: []string{"id desc"},
	}, &list)

	return list, pagination
}
