package model

import (
	"github.com/guregu/null"
	"gorm-example/utils"
	"gorm.io/gorm"
)

// 使用sql-migrate时,不需要关心gorm tag的写法
// 以下可以适用gorm自带迁移工具执行迁移
// 生成的建表SQL和 02_migrate/mysql/20220626_init.sql下的`users`表SQL一致。
type User struct {
	ID       int64       `gorm:"primaryKey" json:"id,omitempty"`
	Account  string      `gorm:"column:account;type:varchar(191);default'';comment:账号;" json:"account,omitempty"`
	Password string      `gorm:"column:password;type:varchar(191);comment:密码;" json:"-"`
	Username null.String `gorm:"column:username;type:varchar(191);comment:昵称;" json:"username,omitempty"`
	Phone    string      `gorm:"column:phone;type:varchar(16);comment:手机号;" json:"phone,omitempty"`
	Avatar   null.String `gorm:"column:avatar;type:varchar(191);comment:头像;" json:"avatar,omitempty"`
	Email    null.String `gorm:"column:email;type:varchar(191);comment:头像;" json:"email,omitempty"`

	CreatedAt *utils.LocalTimeX `json:"created_at,omitempty"` // 自定义时间JSON序列化
	UpdatedAt *utils.LocalTimeX `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt    `gorm:"index" json:"deleted_at,omitempty"`
}

// TableName
func (u *User) TableName() string {
	return "users"
}

func (u *User) AfterFind(tx *gorm.DB) (err error) {

	// 没有设置头像时，设置默认头像
	if u.Avatar.IsZero() {
		u.Avatar.SetValid("http://default.avatar")
	}

	return nil
}

// 筛选条件
type FilterUser struct {
	ID       int64
	IDs      []int64
	Account  string
	Username null.String
	Phone    string
	Page     int64
	PageSize int64
}

// 获取一条数据
// 尽可能封装更多的条件场景
func (u *User) First(db *gorm.DB, filter *FilterUser, columns ...string) error {
	db = db.Model(&User{})
	if len(columns) > 0 {
		db = db.Select(columns)
	}
	if filter.ID > 0 {
		db = db.Where("id", filter.ID)
	}
	if filter.Account != "" {
		db = db.Where("account", filter.Account)
	}
	if filter.Username.Valid {
		db = db.Where("username", filter.Username.String)
	}
	if filter.Phone != "" {
		db = db.Where("phone", filter.Phone)
	}
	return db.First(u).Error
}

// 分页
// model 层只写不包含任何业务逻辑的纯SQL操作
// 然后在其他层组合基本操作逻辑成业务
func (u *User) Find(db *gorm.DB, filter *FilterUser, columns ...string) (list []*User, pagination *utils.Pagination) {
	db = db.Model(&User{})

	if len(columns) > 0 {
		db = db.Select(columns)
	}
	if len(filter.IDs) > 0 {
		db = db.Where("id in (?)", filter.IDs)
	}
	if filter.Account != "" {
		db = db.Where("account", filter.Account)
	}
	if filter.Username.Valid {
		db = db.Where("username", filter.Username.String)
	}
	if filter.Phone != "" {
		db = db.Where("phone", filter.Phone)
	}
	pagination = utils.Paginate(&utils.Param{
		DB:      db,
		Page:    filter.Page,
		Limit:   filter.PageSize,
		OrderBy: []string{"created_at desc"},
	}, &list)
	return list, pagination
}
