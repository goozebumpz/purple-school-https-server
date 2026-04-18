package link

import (
	"gorm.io/gorm"
	"math/rand"
)

type Link struct {
	gorm.Model
	Url  string `json:"url"`
	Hash string `json:"hash" gorm:"uniqueIndex"`
}

func NewLink(url string) *Link {
	return &Link{
		Url:  url,
		Hash: createHash(10),
	}
}

func (l *Link) GenerateNewHash() {
	l.Hash = createHash(10)
}

func createHash(n int) string {
	var letters = []byte("abcdefghijklmnoprstuvwxyzABCDEFGHIJKLMNOPRSTUVWXYZ")
	r := make([]byte, n)

	for i := range r {
		r[i] = letters[rand.Intn(len(letters))]
	}

	return string(r)
}
