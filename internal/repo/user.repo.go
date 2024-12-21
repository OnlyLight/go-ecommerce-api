package repo

type UserRepo struct {
}

func NewUserRepo() *UserRepo {
	return &UserRepo{}
}

func (ur *UserRepo) GetInfoUserRepo() []string {
	return []string{"Duy", "Khue"}
}
