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
