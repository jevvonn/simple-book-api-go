package book

import (
	"ls-go/model/domain"
	"ls-go/model/dto"

	"github.com/gin-gonic/gin"
)

type BookUsecase interface {
	GetBooks(ctx *gin.Context) ([]domain.Book, error)
	CreateBook(ctx *gin.Context, bookRequest *dto.BookRequest) (domain.Book, error)
}
