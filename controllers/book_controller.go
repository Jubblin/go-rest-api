package controllers

import (
	"go-rest-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @title           Book Store API
// @version         1.0
// @description     A simple book store API
// @host            localhost:8080
// @BasePath        /api/v1

// In-memory storage for books
var books = make(map[string]models.Book)

// Initialize sample data
func init() {
	// Add some sample books
	books[uuid.New().String()] = models.Book{
		ID:     uuid.New().String(),
		Title:  "The Go Programming Language",
		Author: "Alan A. A. Donovan & Brian W. Kernighan",
		Price:  49.99,
	}

	books[uuid.New().String()] = models.Book{
		ID:     uuid.New().String(),
		Title:  "Clean Code",
		Author: "Robert C. Martin",
		Price:  39.99,
	}

	books[uuid.New().String()] = models.Book{
		ID:     uuid.New().String(),
		Title:  "Design Patterns",
		Author: "Erich Gamma, Richard Helm, Ralph Johnson, John Vlissides",
		Price:  54.99,
	}
}

// GetBooks godoc
// @Summary      Get all books
// @Description  Returns a list of all books
// @Tags         books
// @Accept       json
// @Produce      json
// @Success      200  {array}   models.Book
// @Router       /books [get]
func GetBooks(c *gin.Context) {
	booksList := make([]models.Book, 0)
	for _, book := range books {
		booksList = append(booksList, book)
	}
	c.JSON(http.StatusOK, booksList)
}

// GetBook godoc
// @Summary      Get a book by ID
// @Description  Returns a single book by its ID
// @Tags         books
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Book ID"
// @Success      200  {object}  models.Book
// @Failure      404  {object}  string
// @Router       /books/{id} [get]
func GetBook(c *gin.Context) {
	id := c.Param("id")
	if book, exists := books[id]; exists {
		c.JSON(http.StatusOK, book)
		return
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
}

// CreateBook godoc
// @Summary      Create a new book
// @Description  Creates a new book with the provided details
// @Tags         books
// @Accept       json
// @Produce      json
// @Param        book  body      models.Book  true  "Book object"
// @Success      201   {object}  models.Book
// @Failure      400   {object}  string
// @Router       /books [post]
func CreateBook(c *gin.Context) {
	var newBook models.Book
	if err := c.ShouldBindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Generate UUID for new book
	newBook.ID = uuid.New().String()
	books[newBook.ID] = newBook

	c.JSON(http.StatusCreated, newBook)
}

// UpdateBook godoc
// @Summary      Update a book
// @Description  Updates a book's details by its ID
// @Tags         books
// @Accept       json
// @Produce      json
// @Param        id    path      string      true  "Book ID"
// @Param        book  body      models.Book  true  "Book object"
// @Success      200   {object}  models.Book
// @Failure      404   {object}  string
// @Router       /books/{id} [put]
func UpdateBook(c *gin.Context) {
	id := c.Param("id")
	if _, exists := books[id]; !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	var updatedBook models.Book
	if err := c.ShouldBindJSON(&updatedBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedBook.ID = id
	books[id] = updatedBook

	c.JSON(http.StatusOK, updatedBook)
}

// DeleteBook godoc
// @Summary      Delete a book
// @Description  Deletes a book by its ID
// @Tags         books
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Book ID"
// @Success      204  {object}  nil
// @Failure      404  {object}  string
// @Router       /books/{id} [delete]
func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	if _, exists := books[id]; !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	delete(books, id)
	c.Status(http.StatusNoContent)
} 