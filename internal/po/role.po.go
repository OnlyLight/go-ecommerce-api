package po

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	ID       int64  `json:"id" gorm:"column:id;type:int;not null;primaryKey;autoIncrement;"`
	RoleName string `json:"role_name" gorm:"column:role_name;"`
	RoleNote string `json:"role_note" gorm:"column:role_note;type:text;"`
}

func (r *Role) TableName() string {
	return "go_db_roles"
}
