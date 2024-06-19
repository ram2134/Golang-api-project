package main

import (
    "bookstore-api/config"
    "bookstore-api/controller"
    "bookstore-api/repository"
    "bookstore-api/services"
    "bookstore-api/routes"
    "github.com/gin-gonic/gin"
)

func main() {
    config.ConnectDB()
    
    bookRepo := repository.NewBookRepository()
    bookService := services.NewBookService(bookRepo)
    bookController := controller.NewBookController(bookService)
    
    r := gin.Default()
    routes.RegisterBookRoutes(r, bookController)
    r.Run()
}
