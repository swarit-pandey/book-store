package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cucumber/godog"
	"github.com/gin-gonic/gin"
	"github.com/swarit-pandey/book-store/api"
	"github.com/swarit-pandey/book-store/controllers"
	services "github.com/swarit-pandey/book-store/service"
)

func TestFeatures(t *testing.T) {
    suite := godog.TestSuite{
        ScenarioInitializer: InitializeScenario,
        Options: &godog.Options{
            Format:   "pretty",
            Paths:    []string{"features"},
            TestingT: t,
        },
    }

    if suite.Run() != 0 {
        t.Fatal("non-zero status returned, failed to run feature tests")
    }
}

var (
    bookService    *services.BookService
    bookController *controllers.BookController
    response       *httptest.ResponseRecorder
    responseStatus int
)

func InitializeScenario(ctx *godog.ScenarioContext) {
    ctx.Step(`^the following books exist in the system:$`, theFollowingBooksExistInTheSystem)
    ctx.Step(`^I retrieve the list of all books$`, iRetrieveTheListOfAllBooks)
    ctx.Step(`^I retrieve the list of books by genre "([^"]*)"$`, iRetrieveTheListOfBooksByGenre)
    ctx.Step(`^I should receive the following books:$`, iShouldReceiveTheFollowingBooks)
    ctx.Step(`^the response status code should be (\d+)$`, theResponseStatusCodeShouldBe)
    ctx.Step(`^I add a book with the following details:$`, iAddABookWithTheFollowingDetails)
    ctx.Step(`^the book list should include:$`, theBookListShouldInclude)

    // TODO: Update this to use Before
    ctx.BeforeScenario(func(*godog.Scenario) {
        gin.SetMode(gin.TestMode)
        bookService = services.NewBookService()
        bookController = controllers.NewBookController(*bookService)
        response = httptest.NewRecorder()
        responseStatus = 0
    })
}

func theFollowingBooksExistInTheSystem(table *godog.Table) error {
    for _, row := range table.Rows[1:] {
        book := services.Book{
            Title:  row.Cells[0].Value,
            Author: row.Cells[1].Value,
            Genre:  row.Cells[2].Value,
        }
        bookService.AddBook(book)
    }
    return nil
}

func iRetrieveTheListOfAllBooks() error {
    req, _ := http.NewRequest("GET", "/books", nil)
    router := gin.Default()
    api.RegisterHandlers(router, bookController)
    router.ServeHTTP(response, req)
    responseStatus = response.Code
    return nil
}

func iRetrieveTheListOfBooksByGenre(genre string) error {
    req, _ := http.NewRequest("GET", fmt.Sprintf("/books?genre=%s", genre), nil)
    router := gin.Default()
    api.RegisterHandlers(router, bookController)
    router.ServeHTTP(response, req)
    responseStatus = response.Code
    return nil
}

func iShouldReceiveTheFollowingBooks(table *godog.Table) error {
    var expectedBooks []services.Book
    for _, row := range table.Rows[1:] {
        book := services.Book{
            Title:  row.Cells[0].Value,
            Author: row.Cells[1].Value,
            Genre:  row.Cells[2].Value,
        }
        expectedBooks = append(expectedBooks, book)
    }

    var actualBooks []services.Book
    json.Unmarshal(response.Body.Bytes(), &actualBooks)

    if len(expectedBooks) != len(actualBooks) {
        return fmt.Errorf("expected %d books, but got %d", len(expectedBooks), len(actualBooks))
    }

    for i, expectedBook := range expectedBooks {
        if expectedBook != actualBooks[i] {
            return fmt.Errorf("expected book %v, but got %v", expectedBook, actualBooks[i])
        }
    }

    return nil
}

func theResponseStatusCodeShouldBe(statusCode int) error {
    if responseStatus != statusCode {
        return fmt.Errorf("expected status code %d, but got %d", statusCode, responseStatus)
    }
    return nil
}

func iAddABookWithTheFollowingDetails(table *godog.Table) error {
    row := table.Rows[1]
    newBook := api.NewBook{
        Title:  row.Cells[0].Value,
        Author: row.Cells[1].Value,
        Genre:  row.Cells[2].Value,
    }

    jsonBook, _ := json.Marshal(newBook)
    req, _ := http.NewRequest("POST", "/books", bytes.NewBuffer(jsonBook))
    req.Header.Set("Content-Type", "application/json")
    router := gin.Default()
    api.RegisterHandlers(router, bookController)
    router.ServeHTTP(response, req)
    responseStatus = response.Code

    return nil
}

func theBookListShouldInclude(table *godog.Table) error {
    var expectedBooks []services.Book
    for _, row := range table.Rows[1:] {
        book := services.Book{
            Title:  row.Cells[0].Value,
            Author: row.Cells[1].Value,
            Genre:  row.Cells[2].Value,
        }
        expectedBooks = append(expectedBooks, book)
    }

    var actualBooks []services.Book
    json.Unmarshal(response.Body.Bytes(), &actualBooks)

    if len(expectedBooks) != len(actualBooks) {
        return fmt.Errorf("expected %d books, but got %d", len(expectedBooks), len(actualBooks))
    }

    for i, expectedBook := range expectedBooks {
        if expectedBook != actualBooks[i] {
            return fmt.Errorf("expected book %v, but got %v", expectedBook, actualBooks[i])
        }
    }

    return nil
}
