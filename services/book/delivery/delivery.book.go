package delivery

import (
	"fmt"
	"log"
	"ls-go/model/dto"
	"ls-go/services/book"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type BookDelivery struct {
	usecase book.BookUsecase
	router  *gin.Engine
}

func NewBookDelivery(usecase book.BookUsecase, router *gin.Engine) {
	handler := &BookDelivery{
		usecase,
		router,
	}

	router.GET("/books", handler.GetAllBooks)
	router.POST("/books", handler.CreateBooks)
}

func (d *BookDelivery) GetAllBooks(ctx *gin.Context) {
	books, err := d.usecase.GetBooks(ctx)

	if err != nil {
		log.Fatal(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message":    "Erorr encounterd",
			"errMessage": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"data": gin.H{
			"books": books,
		},
	})
}

func (d *BookDelivery) CreateBooks(ctx *gin.Context) {
	var bookRequest dto.BookRequest

	err := ctx.Bind(&bookRequest)
	if err != nil {
		errorMassages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			message := fmt.Sprintf("Error in field %s, condition: %s", e.Field(), e.ActualTag())
			errorMassages = append(errorMassages, message)
		}

		ctx.JSON(http.StatusBadRequest, gin.H{
			"messages": errorMassages,
		})
		return
	}

	book, err := d.usecase.CreateBook(ctx, &bookRequest)
	if err != nil {
		log.Fatal(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message":    "Erorr encounterd",
			"errMessage": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Success Created Book!",
		"data": gin.H{
			"book": book,
		},
	})
}
