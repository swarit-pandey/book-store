package services

type BookService struct {
    books []Book
}

func NewBookService() *BookService {
    return &BookService{
        books: []Book{},
    }
}

func (s *BookService) GetBooksByGenre(genre *string) []Book {
    if genre == nil {
        return s.books
    }

    var result []Book
    for _, book := range s.books {
        if book.Genre == *genre {
            result = append(result, book)
        }
    }
    return result
}

func (s *BookService) AddBook(book Book) {
    s.books = append(s.books, book)
}
