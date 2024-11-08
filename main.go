package main

import (
	"log"
	"ls-go/database"
	"ls-go/services/book/delivery"
	"ls-go/services/book/repository"
	"ls-go/services/book/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	db, err := database.ConnectDB()

	if err != nil {
		log.Fatal(err)
		return
	}

	bookRepository := repository.NewBookRepository(db)
	bookUsecase := usecase.NewBookUsecase(bookRepository)
	delivery.NewBookDelivery(bookUsecase, router)

	router.Run("localhost:4001")
}
