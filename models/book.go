package models

type Book struct {
	ID          string  `json:"id"`
	Title       string  `json:"title" binding:"required"`
	Author      string  `json:"author" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Price       float64 `json:"price" binding:"required"`
}

type Status struct {
	Status string `json:"status"`
}
