package usecase

import (
	"ls-go/model/domain"
	"ls-go/model/dto"
	"ls-go/services/book"

	"github.com/gin-gonic/gin"
)

type BookUsecase struct {
	repo book.BookRepository
}

func NewBookUsecase(repo book.BookRepository) book.BookUsecase {
	return &BookUsecase{repo}
}

func (u *BookUsecase) GetBooks(ctx *gin.Context) ([]domain.Book, error) {
	books, err := u.repo.GetAllBooks()
	return books, err
}

func (u *BookUsecase) CreateBook(ctx *gin.Context, bookRequest *dto.BookRequest) (domain.Book, error) {
	book := domain.Book{
		Title:       bookRequest.Title,
		Description: bookRequest.Description,
		Price:       bookRequest.Price,
		Rating:      bookRequest.Rating,
	}

	err := u.repo.Create(&book)
	return book, err
}
