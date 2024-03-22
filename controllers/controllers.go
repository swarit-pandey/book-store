package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/swarit-pandey/book-store/api"
	services "github.com/swarit-pandey/book-store/service"
)

type BookController struct {
    bookService services.BookService
}

func NewBookController(bookService services.BookService) *BookController {
    return &BookController{
        bookService: bookService,
    }
}

func (bc *BookController) GetBooks(c *gin.Context, params api.GetBooksParams) {
    books := bc.bookService.GetBooksByGenre(params.Genre)
    c.JSON(http.StatusOK, books)
}

func (bc *BookController) PostBooks(c *gin.Context) {
    var newBook api.NewBook
    if err := c.ShouldBindJSON(&newBook); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    book := services.Book{
        Title:  newBook.Title,
        Author: newBook.Author,
        Genre:  newBook.Genre,
    }
    bc.bookService.AddBook(book)

    books := bc.bookService.GetBooksByGenre(nil)
    c.JSON(http.StatusCreated, books)
}
