package book

import (
	"project/playground/be5/docker-play/entities"

	"gorm.io/gorm"
)

type BookRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *BookRepository {
	return &BookRepository{
		db: db,
	}
}

func (bk BookRepository) Insert(newBook entities.Book) (entities.Book, error) {
	if err := bk.db.Create(&newBook).Error; err != nil {
		return newBook, err
	}

	return newBook, nil
}
