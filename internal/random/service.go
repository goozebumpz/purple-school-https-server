package random

import (
	"math/rand"
	"strings"
)

type Service struct {
}

func (r *Service) Random() string {
	const chars = "123456"
	randIndex := rand.Intn(len(chars))
	return strings.Split(chars, "")[randIndex]
}
