package repo

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
	panic("unimplemented")
}
