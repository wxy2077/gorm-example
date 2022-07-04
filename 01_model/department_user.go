package model

type DepartmentUser struct {
	ID int64 `json:"id"`

	UserID int64 `json:"user_id"`

	DepartmentID int64 `json:"department_id"`

	Department *Department `gorm:"foreignKey:ID;references:DepartmentID" json:"department"`
}

func (du *DepartmentUser) TableName() string {
	return "department_users"
}
