package user

import (
	"context"
	"gorm.io/gorm"
	"purple-school/pkg/db"
)

type Repository struct {
	Database *db.DB
}

func NewUserRepository(db *db.DB) *Repository {
	return &Repository{
		Database: db,
	}
}

func (r *Repository) CreateUser(user *User) (*User, error) {
	err := gorm.G[User](r.Database.DB).Create(context.Background(), user)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *Repository) FindByEmail(email string) (*User, error) {
	user, err := gorm.G[*User](r.Database.DB).Where("email = ?", email).First(context.Background())

	if err != nil {
		return nil, err
	}

	return user, nil
}
