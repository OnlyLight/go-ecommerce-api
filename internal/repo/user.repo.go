package repo

import (
	"github.com/onlylight29/go-ecommerce-backend-api/global"
	"github.com/onlylight29/go-ecommerce-backend-api/internal/database"
	"go.uber.org/zap"
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
	// GetUserByEmail(email string) bool
	GetUserByEmailSQLC(email string) bool
}

// Like Class in Java
type userRepository struct {
	sqlc *database.Queries
}

// Like Class implement Interface in Java
func NewUserRepository() IUserRepo {
	return &userRepository{
		sqlc: database.New(global.Mdbc),
	}
}

// GetUserByEmail implements IUserRepo.
// func (ur *userRepository) GetUserByEmail(email string) bool {
// 	row := global.MDB.Table(TableNameGoCrmUser).Where("usr_email = ?", email).First(&model.GoCrmUser{}).RowsAffected

// 	return row != 0
// }

func (ur *userRepository) GetUserByEmailSQLC(email string) bool {
	user, err := ur.sqlc.GetUserByEmailSQLC(ctx, email)
	if err != nil {
		global.Logger.Info("Not found User", zap.Error(err))
		return false
	}

	return user.UsrID != 0
}
