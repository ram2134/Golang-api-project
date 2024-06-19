package repository

import (
    "bookstore-api/config"
    "bookstore-api/entities"
)

type bookRepository struct{}

func NewBookRepository() *bookRepository {
    return &bookRepository{}
}

func (r *bookRepository) CreateBook(book entities.Book) (entities.Book, error) {
    err := config.DB.Create(&book).Error
    return book, err
}

func (r *bookRepository) GetBooks() ([]entities.Book, error) {
    var books []entities.Book
    err := config.DB.Find(&books).Error
    return books, err
}

func (r *bookRepository) GetBook(id uint) (entities.Book, error) {
    var book entities.Book
    err := config.DB.First(&book, id).Error
    return book, err
}

func (r *bookRepository) UpdateBook(book entities.Book) (entities.Book, error) {
    err := config.DB.Save(&book).Error
    return book, err
}

func (r *bookRepository) DeleteBook(id uint) error {
    err := config.DB.Delete(&entities.Book{}, id).Error
    return err
}
