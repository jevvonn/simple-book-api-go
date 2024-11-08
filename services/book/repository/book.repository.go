package repository

import (
	"ls-go/model/domain"
	"ls-go/services/book"

	"gorm.io/gorm"
)

type BookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) book.BookRepository {
	return &BookRepository{db}
}

func (r BookRepository) GetAllBooks() ([]domain.Book, error) {
	var books []domain.Book
	err := r.db.Find(&books).Error
	return books, err
}

func (r BookRepository) Create(book *domain.Book) error {
	err := r.db.Create(&book).Error
	return err
}
