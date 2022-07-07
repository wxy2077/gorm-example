package model

type DepartmentUser struct {
	ID int64 `json:"id"`

	UserID int64 `json:"user_id"`

	DepID int64 `json:"dep_id"`

	Department *Department `gorm:"foreignKey:ID;references:DepID" json:"department"`
}

func (du *DepartmentUser) TableName() string {
	return "department_users"
}
