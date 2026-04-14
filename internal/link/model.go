package link

import (
	"math/rand"

	"gorm.io/gorm"
)

type Link struct {
	gorm.Model
	Url  string `json:"url"`
	Hash string `json:"hash" gorm:"uniqueIndex"`
}

func createHash(n int) string {
	var letters = []byte("abcdefghijklmnoprstuvwxyzABCDEFGHIJKLMNOPRSTUVWXYZ")
	r := make([]byte, n)

	for i := range r {
		r[i] = letters[rand.Intn(len(letters))]
	}

	return string(r)
}

func NewLink(url string) *Link {
	return &Link{
		Url:  url,
		Hash: createHash(10),
	}
}
