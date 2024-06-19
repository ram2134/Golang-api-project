package controller

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "bookstore-api/entities"
    "bookstore-api/interface"
)

type BookController struct {
    Service _interface.BookService
}

func NewBookController(service _interface.BookService) *BookController {
    return &BookController{
        Service: service,
    }
}

func (ctrl *BookController) CreateBook(c *gin.Context) {
    var book entities.Book
    if err := c.ShouldBindJSON(&book); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    createdBook, err := ctrl.Service.CreateBook(book)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, createdBook)
}

func (ctrl *BookController) GetBooks(c *gin.Context) {
    books, err := ctrl.Service.GetBooks()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, books)
}

func (ctrl *BookController) GetBook(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    book, err := ctrl.Service.GetBook(uint(id))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
        return
    }
    c.JSON(http.StatusOK, book)
}

func (ctrl *BookController) UpdateBook(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    var book entities.Book
    if err := c.ShouldBindJSON(&book); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    updatedBook, err := ctrl.Service.UpdateBook(uint(id), book)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, updatedBook)
}

func (ctrl *BookController) DeleteBook(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    if err := ctrl.Service.DeleteBook(uint(id)); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
}
