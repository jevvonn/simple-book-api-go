package dto

type BookRequest struct {
	Title       string  `json:"title" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Rating      float64 `json:"rating" binding:"required,number"`
	Price       int     `json:"price" binding:"required,number,gt=1000"`
}
