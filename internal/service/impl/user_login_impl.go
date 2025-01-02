package impl

import (
	"context"

	"github.com/onlylight29/go-ecommerce-backend-api/internal/database"
)

type sUserLogin struct {
	r *database.Queries
}

func NewUserLoginImpl(r *database.Queries) *sUserLogin {
	return &sUserLogin{
		r: r,
	}
}

// Implement interface of IUserLogin
func (s *sUserLogin) Login(ctx context.Context) error {
	panic("unimplement yet")
}

func (s *sUserLogin) Register(ctx context.Context) error {
	panic("unimplement yet")
}

func (s *sUserLogin) VerifyOTP(ctx context.Context) error {
	panic("unimplement yet")
}

func (s *sUserLogin) UpdatePasswordRegister(ctx context.Context) error {
	panic("unimplement yet")
}
