package repo

type PongRepo struct {
}

func NewPongRepo() *PongRepo {
	return &PongRepo{}
}

func (pr *PongRepo) GetInfoPong() string {
	return "Pong..."
}
