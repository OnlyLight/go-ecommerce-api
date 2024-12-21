package service

import "github.com/onlylight29/go-ecommerce-backend-api/internal/repo"

type PongService struct {
	pongRepo *repo.PongRepo
}

func NewPongService() *PongService {
	return &PongService{
		pongRepo: repo.NewPongRepo(),
	}
}

func (ps *PongService) GetInfoPongService() string {
	return ps.pongRepo.GetInfoPong()
}
