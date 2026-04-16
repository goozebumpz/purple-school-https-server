package link

import (
	"context"
	"errors"
	"fmt"
	"purple-school/pkg/db"

	"gorm.io/gorm"
)

type Repository struct {
	Database *db.DB
}

func NewRepository(database *db.DB) *Repository {
	return &Repository{
		Database: database,
	}
}

func (repo *Repository) CreateLink(link *Link) (*Link, error) {
	result := repo.Database.Create(link)

	if result.Error != nil {
		return nil, result.Error
	}

	return link, nil
}

func (repo *Repository) GetByHash(hash string) (*Link, error) {
	ctx := context.Background()
	link, err := gorm.G[Link](repo.Database.DB).Where("hash = ?", hash).Last(ctx)

	if err != nil {
		return nil, err
	}

	return &link, nil
}

func (repo *Repository) CheckUniqueHash(hash string) bool {
	ctx := context.Background()
	_, err := gorm.G[Link](repo.Database.DB).Where("hash = ?", hash).Last(ctx)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return true
	}

	return false
}

func (repo *Repository) UpdateLink(hash string, body *UpdateRequest) (*Link, error) {
	_, err := gorm.G[Link](repo.Database.DB).Where("hash = ?", hash).Update(context.Background(), "url", body.Url)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("link not found")
		}
		return nil, err
	}

	link, err := gorm.G[Link](repo.Database.DB).Where("hash = ?", hash).Last(context.Background())

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("link not found")
		}
		return nil, err
	}

	return &link, nil
}

func (repo *Repository) DeleteLink(hash string) (uint, error) {
	ctx := context.Background()
	link, err := gorm.G[Link](repo.Database.DB).Where("hash = ?").Last(ctx)
	_, err = gorm.G[Link](repo.Database.DB).Where("hash = ?", hash).Delete(ctx)

	if err != nil {
		return 0, err
	}

	return link.ID, nil
}
