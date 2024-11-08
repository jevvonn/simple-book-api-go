package book

import "ls-go/model/domain"

type BookRepository interface {
	GetAllBooks() ([]domain.Book, error)
	Create(book *domain.Book) error
}
