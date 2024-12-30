package random

import (
	"time"

	"math/rand"
)

func GenerateSixDigitOTP() int {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	return 100000 + rng.Intn(900000)
}
