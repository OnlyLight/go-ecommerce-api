package repo

import (
	"github.com/onlylight29/go-ecommerce-backend-api/global"
	"github.com/onlylight29/go-ecommerce-backend-api/internal/model"
)

// type UserRepo struct {
// }

// func NewUserRepo() *UserRepo {
// 	return &UserRepo{}
// }

// func (ur *UserRepo) GetInfoUserRepo() []string {
// 	return []string{"Duy", "Khue"}
// }

type IUserRepo interface {
	GetUserByEmail(email string) bool
}

// Like Class in Java
type userRepository struct {
}

// Like Class implement Interface in Java
func NewUserRepository() IUserRepo {
	return &userRepository{}
}

// GetUserByEmail implements IUserRepo.
func (us *userRepository) GetUserByEmail(email string) bool {
	row := global.MDB.Table(TableNameGoCrmUser).Where("usr_email = ?", email).First(&model.GoCrmUser{}).RowsAffected

	return row != 0
}
