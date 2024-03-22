package main

import (
	"github.com/gin-gonic/gin"
	"github.com/swarit-pandey/book-store/api"
	"github.com/swarit-pandey/book-store/controllers"
	services "github.com/swarit-pandey/book-store/service"
)

func main() {
	router := gin.Default()

	bookService := services.NewBookService()
	bookController := controllers.NewBookController(*bookService)

	api.RegisterHandlers(router, bookController)

	_ = router.Run(":8080")
}
