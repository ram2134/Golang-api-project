package services

import (
    "bookstore-api/entities"
    "bookstore-api/interface"
)

type bookService struct {
    repository _interface.BookRepository
}

func NewBookService(repo _interface.BookRepository) _interface.BookService {
    return &bookService{
        repository: repo,
    }
}

func (s *bookService) CreateBook(book entities.Book) (entities.Book, error) {
    return s.repository.CreateBook(book)
}

func (s *bookService) GetBooks() ([]entities.Book, error) {
    return s.repository.GetBooks()
}

func (s *bookService) GetBook(id uint) (entities.Book, error) {
    return s.repository.GetBook(id)
}

func (s *bookService) UpdateBook(id uint, book entities.Book) (entities.Book, error) {
    existingBook, err := s.repository.GetBook(id)
    if err != nil {
        return entities.Book{}, err
    }

    existingBook.Title = book.Title
    existingBook.Author = book.Author
    existingBook.Description = book.Description
    existingBook.PublishedYear = book.PublishedYear
    existingBook.Price = book.Price

    return s.repository.UpdateBook(existingBook)
}

func (s *bookService) DeleteBook(id uint) error {
    return s.repository.DeleteBook(id)
}
