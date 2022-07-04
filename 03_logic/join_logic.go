package logic

import (
	model "gorm-example/01_model"
	"gorm-example/utils"
	"gorm.io/gorm"
)

type JoinLogic interface {

	// UserDepList 连接查询用户所在部门
	UserDepList(db *gorm.DB, page int64) (list []*UserList, pagination *utils.Pagination)

	// 其他业务逻辑方法....
}

type joinLogic struct {
}

// NewJoinLogic 接口controller层直接调用
// 完成各种业务操作
func NewJoinLogic() JoinLogic {
	return &joinLogic{}
}

type UserList struct {
	*model.User
	Title string `json:"title"`
}

func (jl *joinLogic) UserDepList(db *gorm.DB, page int64) (list []*UserList, pagination *utils.Pagination) {

	db.Model(&model.User{}).Select("d.title,users.*").
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
