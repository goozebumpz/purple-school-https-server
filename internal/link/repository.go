package link

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"purple-school/pkg/db"
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
	_, err := gorm.G[Link](repo.Database.DB).Where("hash = ?", hash).First(context.Background())

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return true
	}

	return false
}

func (repo *Repository) UpdateLink(link *Link) (*Link, error) {
	db := gorm.G[*Link](repo.Database.DB, clause.Returning{}).Where("id = ?", link.ID)
	rowsAffected, err := db.Updates(context.Background(), link)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("link not found")
		}
		return nil, err
	}

	fmt.Println(rowsAffected)

	if rowsAffected == 0 {
		return nil, fmt.Errorf("rows affected zero")
	}

	return link, nil
}

func (repo *Repository) DeleteLink(id uint) error {
	ctx := context.Background()
	_, err := gorm.G[Link](repo.Database.DB).Where("id = ?", id).Delete(ctx)

	if err != nil {
		return err
	}

	return nil
}

func (repo *Repository) ExistLink(id uint) (bool, error) {
	link, err := gorm.G[*Link](repo.Database.DB).Where("id = ?", id).First(context.Background())

	if err != nil {
		return false, err
	}

	return link != nil, nil
}
