package po

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UUID     uuid.UUID `json:"uuid" gorm:"column:uuid;type:varchar(255);not null;unique;"`
	Username string    `json:"username" gorm:"column:user_name;type:varchar(255);"`
	IsActive bool      `json:"is_active" gorm:"column:is_active;type:boolean;"`
	Roles    []Role    `json:"roles" gorm:"many2many:go_user_roles;"`
}

func (u *User) TableName() string {
	return "go_db_users"
}
