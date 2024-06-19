package _interface

import (
    "bookstore-api/entities"
)

type BookRepository interface {
    CreateBook(book entities.Book) (entities.Book, error)
    GetBooks() ([]entities.Book, error)
    GetBook(id uint) (entities.Book, error)
    UpdateBook(book entities.Book) (entities.Book, error)
    DeleteBook(id uint) error
}

type BookService interface {
    CreateBook(book entities.Book) (entities.Book, error)
    GetBooks() ([]entities.Book, error)
    GetBook(id uint) (entities.Book, error)
    UpdateBook(id uint, book entities.Book) (entities.Book, error)
    DeleteBook(id uint) error
}
